package handler

import (
	"net/http"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := database.Users.FindOne(c, bson.D{{"email", user.Email}}).Err()
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "exist"})
		return
	}
	user.ID = primitive.NewObjectID()
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, err = database.Users.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetUsers(c *gin.Context) {
	var users []model.User
	cursor, err := database.Users.Find(c, bson.D{}) //test içindir
	// cursor, err := database.Users.Find(c, bson.D{{"favoutireDailies", "a"}}) get users with "a" in favoutireDailies array
	//_, err := database.Users.UpdateMany(c, bson.D{{"favoutireDailies", "a"}}, bson.M{"$push": bson.M{"favoutireDailies": "b"}}) //update users with "a" in favoutireDailies array, adding b
	//_, err := database.Users.UpdateMany(c, bson.D{{"favoutireDailies", "a"}}, bson.M{"$set": bson.M{"favoutireDailies.$": "b"}}) //update users with "a" in favoutireDailies array, change a as b
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	err = cursor.All(c, &users)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateDaily(c *gin.Context) {
	var daily model.Daily
	err := c.ShouldBindJSON(&daily)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	daily.ID = primitive.NewObjectID()
	daily.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, err = database.Dailies.InsertOne(c, daily)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetDailies(c *gin.Context) {
	var dailies []model.Daily
	cursor, err := database.Dailies.Find(c, bson.D{})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	err = cursor.All(c, &dailies)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dailies)
}
