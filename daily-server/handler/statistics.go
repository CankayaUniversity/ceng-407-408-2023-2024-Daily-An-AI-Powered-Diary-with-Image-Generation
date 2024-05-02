package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/Final-Projectors/daily-server/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticsController struct {
	DailyRepository  *repository.DailyRepository
	ReportRepository *repository.ReportedDailyRepository
	UserRepository   *repository.UserRepository
}

func NewStatisticsController(_userRepository *repository.UserRepository, _repository *repository.DailyRepository, _reports *repository.ReportedDailyRepository) *DailyController {
	return &DailyController{
		UserRepository:  _userRepository,
		DailyRepository: _repository, ReportRepository: _reports,
	}
}

func (controller *StatisticsController) Likes(daily model.Daily) int {

}

func (controller *StatisticsController) Streak(daily []model.Daily) int {

}

func (controller *StatisticsController) DailiesWritten(user primitive.ObjectID) int {

}

func (controller *StatisticsController) DailyMood(daily model.Daily) int {

}

func (controller *StatisticsController) UserMood(dailies []model.Daily) string {

}

func (controller *StatisticsController) Topics(dailies []model.Daily) string {

}

func (controller *StatisticsController) Views([]model.Daily) int {

}

func (controller *StatisticsController) GetCalendar(c *gin.Context) int {

}

func (controller *StatisticsController) GetStatistics(c *gin.Context) int {

}
