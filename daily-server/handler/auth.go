package handler

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/Final-Projectors/daily-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Register handles the user registration process.
// @Summary Register a new user
// @Description Create a new user with the given email and password, if they don't exist already.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.UserRegisterDTO true "User Registration"
// @Success 200 {object} model.User
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 502 {object} object "Bad Gateway {"message': "Couldn't fetch the image"}"
// @Router /api/register [post]
func Register(c *gin.Context) {
	var user model.User
	var userRequest model.UserRegisterDTO
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := database.Users.FindOne(c, bson.M{"email": userRequest.Email}).Err()
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "exist"})
		return
	}
	mailErr := SendVerificationEmail(userRequest.Email)
	if mailErr != nil {
		fmt.Println("Error sending verification email:", mailErr)
		return
	}
	fmt.Println("Verification email sent successfully to", userRequest.Email)
	user.ID = primitive.NewObjectID()
	user.Email = userRequest.Email
	user.Role = "user"
	user.IsVerified = false
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(hashedPassword)
	user.FavouriteDailies = []primitive.ObjectID{} // initialize as empty slice
	user.ViewedDailies = []primitive.ObjectID{}
	_, err = database.Users.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Login authenticates a user and provides a token.
// @Summary User login
// @Description Authenticate a user using the provided email and password, and return a token on successful authentication.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.UserLoginDTO true "User login details"
// @Success 200 {object} map[string]string "Token"
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Router /api/login [post]
func Login(c *gin.Context) {
	var userRequest model.UserLoginDTO
	var result model.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := database.Users.FindOne(c, bson.M{"email": userRequest.Email}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if !result.IsVerified {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email address is not verified"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userRequest.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message1": err.Error()})
		return
	}
	token, err := utils.GenerateToken(result.ID.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message2": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func SendVerificationEmail(toEmail string) error {
	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // Port for SMTP submission (587 for TLS)
	smtpUsername := "daily2024ai@gmail.com"
	smtpPassword := "wveb wnxa zwks sycy"

	// Sender and recipient email addresses
	from := "daily2024ai@gmail.com"
	to := []string{toEmail}

	token := generateRandomToken()
	// Email content
	subject := "Verification Email"
	verificationURL := fmt.Sprintf("http://localhost:9090/api/verify/%s/&token=%s", toEmail, token)
	body := fmt.Sprintf("Please click the link below to verify your email address:\n\n %s", verificationURL)

	// Create authentication credentials
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Compose the email message
	msg := []byte("To: " + toEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}

func generateRandomToken() string {
	// Generate a random byte slice
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Encode the byte slice to base64
	token := base64.URLEncoding.EncodeToString(randomBytes)

	return token
}

func VerifyEmail(c *gin.Context) {
	var result model.User
	email := c.Param("email")
	token := c.Param("token")
	fmt.Println(email, token)
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email address is required"})
		return
	}

	err := database.Users.FindOne(c, bson.M{"email": email}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"isVerified": true}} // Assuming you have a field "isVerified" in your user document
	_, err = database.Users.UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
