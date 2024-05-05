package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

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
		statistics.Mood = "None"
	}
	statistics.Streak = controller.Streak(dailies)
	statistics.Topic = controller.Topics(dailies)

	c.JSON(http.StatusOK, statistics)
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
	return 0
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
			if emotion != "_id" { // Ignoring the _id field
				totalConverted, ok := total.(float64)
				if !ok {
					continue
				}

				if totalConverted > highestTotal {
					highestTotal = totalConverted
					mostProminentEmotion = emotion
				}
			}
		}
	}
	return mostProminentEmotion, nil
}

func (controller *StatisticsController) Topics(dailies []model.Daily) string {
	return "Friends"
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
