package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/Final-Projectors/daily-server/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticsController struct {
	DailyRepository *repository.DailyRepository
	logger          *logrus.Logger
}

func NewStatisticsController(_repository *repository.DailyRepository) *StatisticsController {
	return &StatisticsController{
		DailyRepository: _repository,
		logger:          logrus.New(),
	}
}

// Statistics returns user statistics
// @Summary      Get user statistics
// @Description  provides statistical data about a user's activity including likes, views, number of dailies written, current mood, streak, and a predefined topic
// @Tags         Statistics
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.StatisticsDTO  "An object of statistics including likes, views, dailies written, mood, streak, and topic"
// @Failure      400  {string}  string  "bad request - error message"
// @Failure      401  {string}  string  "unauthorized - error message"
// @Router       /api/daily/statistics [get]
// @Security ApiKeyAuth
func (controller *StatisticsController) Statistics(c *gin.Context) {
	userId, exists := c.Get("user_id")

	statistics := model.StatisticsDTO{}

	controller.logger.Infof("%v, %v", exists, userId)

	if !exists {
		controller.logger.Infof("%v, %v", exists, userId)
		c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("user_id: %v, exists: %v", userId, exists)})
	}
	author := userId.(primitive.ObjectID)
	dailies, err := controller.DailyRepository.List(author, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "List failed"})
	}
	// call other functions passing by the list of "dailies"
	statistics.Likes = controller.Likes(dailies)
	statistics.Views = controller.Views(dailies)
	statistics.DailiesWritten = controller.DailiesWritten(dailies)
	statistics.Mood, err = controller.UserMood(author, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		statistics.Mood = "None"
	}
	statistics.Streak = controller.Streak(dailies)
	statistics.Topics = controller.Topics(dailies)
	statistics.Dates = controller.GetDates(author, c)
	c.JSON(http.StatusOK, statistics)
}

func (controller *StatisticsController) GetDates(userId primitive.ObjectID, c *gin.Context) []string {
	dates := bson.A{
		bson.D{{"$match", bson.D{{"author", userId}, {"createdAt", bson.D{{"$exists", true}}}}}},
		bson.D{{"$sort", bson.D{{"createdAt", 1}}}},
		bson.D{{"$project", bson.D{{"createdAt", 1}, {"_id", 0}}}},
	}

	cursor, err := database.Dailies.Aggregate(c.Request.Context(), dates)
	if err != nil {
		controller.logger.Warnf("Error on calendar aggregation")
		return []string{}
	}

	var formattedDates []string
	for cursor.Next(c.Request.Context()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			// Handle cursor decoding error
			fmt.Println("Error decoding cursor:", err)
		}

		// Extract and format the createdAt field
		if createdAt, ok := result["createdAt"].(primitive.DateTime); ok {
			goTime := createdAt.Time() // Convert MongoDB Datetime to Go's time.Time
			controller.logger.Infof("Date: %v", goTime)
			formattedDate := goTime.Format("2006-01-02") // Format the time to "YYYY-MM-DD"
			controller.logger.Infof("Date: %v", formattedDate)
			formattedDates = append(formattedDates, formattedDate)
		}
	}
	return formattedDates
}

func (controller *StatisticsController) Views(dailies []model.Daily) int {
	count := 0
	if dailies == nil {
		return count
	}
	for _, daily := range dailies {
		if daily.Viewers != nil {
			count += len(daily.Viewers)
		}
	}
	return count
}

func (controller *StatisticsController) Streak(dailies []model.Daily) int {
	// gotta check how to work with dates and date conversions
	// between mongo datetime and go datetime
	current := 0
	prev := time.Now()
	controller.logger.Infof("prev start: %v", prev)
	for _, daily := range dailies {
		controller.logger.Infof("current: %v", current)
		currentTime := daily.CreatedAt.Time()
		controller.logger.Infof("date: %v", currentTime)
		if current == 0 {
			prev = daily.CreatedAt.Time()
			controller.logger.Infof("prev time: %v, cur time: %v", prev, currentTime)
			if time.Now().Sub(currentTime).Hours() > 36.0 {
				controller.logger.Infof("ciktim diff: %v", time.Now().Sub(currentTime))
				return 0
			}
		}

		diff := prev.Sub(currentTime)
		controller.logger.Infof("Hour diff: %v", diff.Hours())
		if diff.Hours() < 36.0 {
			current += 1
			prev = daily.CreatedAt.Time()
		} else {
			return current
		}
	}
	return current
}

func (controller *StatisticsController) DailiesWritten(dailies []model.Daily) int {
	if dailies == nil {
		return 0
	}
	return len(dailies)
}

func (controller *StatisticsController) UserMood(userId primitive.ObjectID, c *gin.Context) (string, error) {
	emotions := bson.A{
		bson.D{{"$match", bson.D{{"author", userId}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", primitive.Null{}},
					{"Sadness", bson.D{{"$sum", "$emotions.sadness"}}},
					{"Joy", bson.D{{"$sum", "$emotions.joy"}}},
					{"Love", bson.D{{"$sum", "$emotions.love"}}},
					{"Anger", bson.D{{"$sum", "$emotions.anger"}}},
					{"Fear", bson.D{{"$sum", "$emotions.fear"}}},
					{"Surprise", bson.D{{"$sum", "$emotions.surprise"}}},
				},
			},
		},
	}

	cursor, err := database.Dailies.Aggregate(c.Request.Context(), emotions)
	if err != nil {
		return "", err
	}
	var results []bson.M
	if err = cursor.All(c.Request.Context(), &results); err != nil {
		// Handle cursor processing error
		return "", err
	}

	mostProminentEmotion := "None"
	highestTotal := 0.0
	if len(results) > 0 {
		summary := results[0]
		for emotion, total := range summary {
			controller.logger.Infof("Emotion: %v, total: %v,  typeof total: %v", emotion, total, reflect.TypeOf(total))
			if emotion != "_id" { // Ignoring the _id field
				switch v := total.(type) {
				case int32:
					total = float64(v)
				}
				totalConverted, ok := total.(float64)

				if !ok {
					continue
				}
				controller.logger.Infof("Highest: %v, cur: %v, emotion: %v", highestTotal, totalConverted, emotion)
				if totalConverted > highestTotal {
					highestTotal = totalConverted
					mostProminentEmotion = emotion
				}
			}
		}
	}
	return mostProminentEmotion, nil
}

func (controller *StatisticsController) Topics(dailies []model.Daily) []string {
	/*
		var topicList []string

		for _, daily := range dailies {
			topicList = append(topicList, daily.Topics[0])
		}
	*/

	var a []string
	a = append(a, "Hello")
	a = append(a, "World")
	return a
}

func (controller *StatisticsController) Likes(dailies []model.Daily) int {
	count := 0
	if dailies == nil {
		return 0
	}
	for _, daily := range dailies {
		count += daily.Favourites
	}
	return count
}

func (controller *StatisticsController) GetCalendar(c *gin.Context) int {
	return 0
}

func (controller *StatisticsController) CheckForBadges(c *gin.Context) {
	badges := []string{}

	userId, exists := c.Get("user_id")

	statistics := model.StatisticsDTO{}

	controller.logger.Infof("%v, %v", exists, userId)

	if !exists {
		controller.logger.Infof("%v, %v", exists, userId)
		c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("user_id: %v, exists: %v", userId, exists)})
	}
	author := userId.(primitive.ObjectID)
	dailies, err := controller.DailyRepository.List(author, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "List failed"})
	}

	dailiesWritten := controller.DailiesWritten(dailies)
	streak := controller.Streak(dailies)
	likes := controller.Likes(dailies)
	views := controller.Views(dailies)

	if dailiesWritten >= 100 {
		badges = append(badges, "Master Writer")
	} else if dailiesWritten >= 10 {
		badges = append(badges, "Prolific Writer")
	} else if statistics.DailiesWritten >= 1 {
		badges = append(badges, "Beginner Writer")
	}
	if streak >= 7 {
		badges = append(badges, "One Week Obsessed")
	}
	if likes >= 1000 {
		badges = append(badges, "Influence")
	} else if likes >= 100 {
		badges = append(badges, "Liked by Many")
	} else if likes >= 1 {
		badges = append(badges, "Admired")
	}
	if views >= 1000 {
		badges = append(badges, "Popular Author")
	} else if views >= 1 {
		badges = append(badges, "They Look Here")
	}

	c.JSON(http.StatusOK, badges)
}
