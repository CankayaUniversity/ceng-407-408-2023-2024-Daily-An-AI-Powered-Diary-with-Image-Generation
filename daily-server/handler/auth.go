package handler

import (
	"net/http"
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
	user.ID = primitive.NewObjectID()
	user.Email = userRequest.Email
	user.Role = "user"
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
