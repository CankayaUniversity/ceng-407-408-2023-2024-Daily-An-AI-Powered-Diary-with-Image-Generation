package handler

import (
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
	err := c.ShouldBindJSON(&daily)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	daily.ID = primitive.NewObjectID()
	daily.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	daily.Author = c.Keys["user_id"].(primitive.ObjectID)
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

func FavDaily(c *gin.Context) {
	var daily model.DailyRequestDTO
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	getDaily := bson.M{"_id": daily.ID}
	dailyOperation := bson.M{"$inc": bson.M{"favourites": 1}}
	if _, err := database.Dailies.UpdateOne(c, getDaily, dailyOperation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update daily", "error": err.Error()})
		return
	}

	getUser := bson.M{"_id": c.Keys["user_id"]}
	userOperation := bson.M{"$push": bson.M{"favouriteDailies": daily.ID}}
	if _, err := database.Users.UpdateOne(c, getUser, userOperation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Favourites updated successfully"})
}

func ViewDaily(c *gin.Context) {
	var daily model.DailyRequestDTO
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	getDaily := bson.M{"_id": daily.ID}
	dailyOperation := bson.M{"$push": bson.M{"viewers": c.Keys["user_id"]}}
	if _, err := database.Dailies.UpdateOne(c, getDaily, dailyOperation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update daily", "error": err.Error()})
		return
	}

	getUser := bson.M{"_id": c.Keys["user_id"]}
	userOperation := bson.M{"$push": bson.M{"viewedDailies": daily.ID}}
	if _, err := database.Users.UpdateOne(c, getUser, userOperation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Views updated successfully"})
}

func ReportDaily(c *gin.Context) {
	var reportedDaily model.ReportedDaily
	if err := c.ShouldBindJSON(&reportedDaily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var dbReportedDaily model.ReportedDaily
	result := database.ReportedDailies.FindOne(c, bson.M{"dailyId": reportedDaily.DailyID})
	if result.Err() != nil {
		reportedDaily.ID = primitive.NewObjectID()
		reportedDaily.DailyID = reportedDaily.DailyID
		reportedDaily.ReportedAt = primitive.NewDateTimeFromTime(time.Now())
		reportedDaily.Reports = 1
		reportedDaily.Title = reportedDaily.Title
		reportedDaily.Content = reportedDaily.Content

		var err error
		_, err = database.ReportedDailies.InsertOne(c, reportedDaily)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "create success"})
		return
	}
	if err := result.Decode(&dbReportedDaily); err == nil {
		reportOperation := bson.M{"$inc": bson.M{"reports": 1}}
		if _, err := database.ReportedDailies.UpdateOne(c, dbReportedDaily, reportOperation); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update reported daily", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "increment success"})
		return
	}
}
