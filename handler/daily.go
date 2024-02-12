package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/Final-Projectors/daily-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateDaily accepts a body request to POST a daily
// @Summary returns the created daily
// @Description creates a new daily resource
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.CreateDailyDTO true "CreateDailyDTO"
// @Success 200 {object} model.Daily
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 502 {object} object "Bad Gateway {"message': "Couldn't fetch the image"}"
// @Router /api/CreateDaily [post]
func CreateDaily(c *gin.Context) {
	var daily model.Daily
	var createDailyDTO model.CreateDailyDTO
	err := c.ShouldBindJSON(&createDailyDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	daily.ID = primitive.NewObjectID()
	daily.Text = createDailyDTO.Text
	daily.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	if createDailyDTO.Image == "" {
		response, err := http.Get("https://d2opxh93rbxzdn.cloudfront.net/original/2X/4/40cfa8ca1f24ac29cfebcb1460b5cafb213b6105.png")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		defer response.Body.Close()
		image, _ := io.ReadAll(response.Body)
		daily.Image = utils.ImageToBase64(image)
	} else {
		//assuming that createDailyDTO.Image is a propper base64 data
		daily.Image = createDailyDTO.Image
	}
	// getting the user_id from context and running checks
	author, _ := c.Get("user_id")
	if auth, ok := author.(primitive.ObjectID); ok {
		daily.Author = auth
	} else {
		fmt.Println("author is not a primitive.ObjectID")
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	daily.IsShared = *createDailyDTO.IsShared
	_, err = database.Dailies.InsertOne(c, daily)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, daily)
}

// GetDaily return a specific daily via daily.ID
// @Summary return a daily
// @Description return a specific daily via daily.ID
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.DailyRequestDTO true "DailyRequestDTO"
// @Success 200 {object} model.Daily
// @Failure 400 {object} object "Bad Request {"message": "Invalid JSON data"}"
// @Failure 502 {object} object "Bad Gateway {"message': "mongo: no documents in result"}"
// @Router /api/getDaily [get]
func GetDaily(c *gin.Context) {
	var daily model.DailyRequestDTO
	var getDaily model.Daily
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}
	err := database.Dailies.FindOne(c, bson.M{"_id": daily.ID}).Decode(&getDaily)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, getDaily)
}

// GetDailies returns a list of dailies of the user based on user_id
// @Summary returns a list of dailies
// @Description returns a list of dailies
// @Tags Daily
// @Accept json
// @Produce json
// @Success 200 {array} model.Daily
// @Failure 500 {object} object "Bad Gateway {"message': "Couldn't fetch the image"}"
// @Router /api/GetDailies [get]
func GetDailies(c *gin.Context) {
	var dailies []model.Daily
	cursor, err := database.Dailies.Find(c, bson.M{"author": c.Keys["user_id"]})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = cursor.All(c, &dailies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
		reportedDaily.ReportedAt = primitive.NewDateTimeFromTime(time.Now())
		reportedDaily.Reports = 1
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

func DeleteDaily(c *gin.Context) {
	var dailies []model.Daily
	user, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user was not found"})
		return
	}
	// get all dailies of the user
	cursor, err := database.Dailies.Find(c, bson.M{"author": user})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = cursor.All(c, &dailies) // all dailies of the user are now in the "dailies" slice
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get the daily id
	var deleteDailyDTO model.DeleteDailyDTO
	err = c.ShouldBindJSON(&deleteDailyDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// check if user has a daily with given id
	var seen bool
	for _, daily := range dailies {
		if daily.ID == *deleteDailyDTO.ID {
			seen = true
		}
	}
	if !seen {
		c.JSON(http.StatusNotFound, gin.H{"message": "user does not have a daily with this id"})
		return
	}

	// if seen, delete the daily
	res, err := database.Dailies.DeleteOne(c, bson.M{"_id": *deleteDailyDTO.ID})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no document was deleted"})
	}
	c.JSON(http.StatusOK, res)
}

func EditDailyImage(c *gin.Context) {
	var daily model.EditDailyImageDTO
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}
	getDaily := bson.M{"_id": daily.ID}
	dailyOperation := bson.M{"$set": bson.M{"image": daily.Image}}
	if _, err := database.Dailies.UpdateOne(c, getDaily, dailyOperation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update daily", "error": err.Error()})
		return
	}
}
