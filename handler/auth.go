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

func Register(c *gin.Context) {
	var user model.User
	var userRequest model.UserRegisterRequest
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
	user.Role = userRequest.Role
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(hashedPassword)
	_, err = database.Users.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func Login(c *gin.Context) {
	var userRequest model.UserLoginRequest
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
