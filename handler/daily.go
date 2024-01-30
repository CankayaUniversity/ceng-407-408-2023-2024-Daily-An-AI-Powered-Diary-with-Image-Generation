package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDaily(c *gin.Context) {
	var daily model.Daily
	var dailyDTO model.CreateDailyDTO
	err := c.ShouldBindJSON(&dailyDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	daily.ID = primitive.NewObjectID()
	daily.Text = dailyDTO.Text
	daily.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	// getting the user_id from context and running checks
	author, _ := c.Get("user_id")

	if auth, ok := author.(primitive.ObjectID); ok {
		daily.Author = auth
	} else {
		fmt.Println("author is not a primitive.ObjectID")
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	daily.IsShared = *dailyDTO.IsShared
	_, err = database.Dailies.InsertOne(c, daily)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetDailies(c *gin.Context) {
	var dailies []model.Daily
	cursor, err := database.Dailies.Find(c, bson.M{"author": c.Keys["user_id"]})
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
