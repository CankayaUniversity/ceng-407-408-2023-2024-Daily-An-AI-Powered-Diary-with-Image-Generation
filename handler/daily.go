package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/Final-Projectors/daily-server/repository"
	"github.com/Final-Projectors/daily-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DailyController struct {
	DailyRepository  *repository.DailyRepository
	ReportRepository *repository.ReportedDailyRepository
	UserRepository   *repository.UserRepository
}

func NewDailyController(_userRepository *repository.UserRepository, _repository *repository.DailyRepository, _reports *repository.ReportedDailyRepository) *DailyController {
	return &DailyController{
		UserRepository:   _userRepository,
		DailyRepository:  _repository,
		ReportRepository: _reports,
	}
}

// Create accepts a body request to POST a daily
// @Summary returns the created daily
// @Description creates a new daily resource
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.CreateDailyDTO true "CreateDailyDTO"
// @Success 200 {object} model.Daily
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 500 {object} object "Internal Server Error {"message': "Couldn't fetch the image"}"
// @Failure 502 {object} object "Bad Gateway {"message': "Couldn't fetch the image / DB error"}"
// @Router /api/daily [post]
// @Security ApiKeyAuth
func (d *DailyController) CreateDaily(c *gin.Context) {
	var daily model.Daily
	var createDailyDTO model.CreateDailyDTO
	err := c.ShouldBindJSON(&createDailyDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}
	daily.ID = primitive.NewObjectID()
	daily.Text = createDailyDTO.Text
	daily.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	if createDailyDTO.Image == "" {
		response, err := http.Get("https://d2opxh93rbxzdn.cloudfront.net/original/2X/4/40cfa8ca1f24ac29cfebcb1460b5cafb213b6105.png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch image"})
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
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unauthorized"})
		return
	}
	daily.IsShared = *createDailyDTO.IsShared

	err = d.DailyRepository.Create(&daily)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, daily)
}

// GetDaily returns a specific daily via daily.ID
// @Summary returns a daily
// @Description returns a specific daily via daily.ID
// @Tags Daily
// @Accept json
// @Produce json
// @Param id path string true "Daily ID"
// @Success 200 {object} model.Daily
// @Failure 400 {object} object "Bad Request {"message": "Invalid JSON data"}"
// @Failure 500 {object} object "Internal Server Error {"message': "mongo: no documents in result"}"
// @Router /api/daily/{id} [get]
// @Security ApiKeyAuth
func (d *DailyController) GetDaily(c *gin.Context) {
	id := c.Param("id")                            // Extract the id from the URL.
	objectID, err := primitive.ObjectIDFromHex(id) // Convert string id to MongoDB ObjectID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error() + objectID.String()})
		return
	}
	daily, err := d.DailyRepository.FindById(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() + objectID.String()})
		return
	}
	c.JSON(http.StatusOK, daily)
}

// GetDailies returns a list of dailies of the user based on user_id
// @Summary returns a list of dailies
// @Description returns a list of dailies
// @Tags Daily
// @Accept json
// @Produce json
// @Success 200 {array} model.Daily
// @Failure 500 {object} object "Bad Gateway {"message': "Couldn't fetch daily list"}"
// @Failure 502 {object} object "Bad Gateway {"message': "No user"}"
// @Router /api/daily/list [get]
// @Security ApiKeyAuth
func (d *DailyController) GetDailies(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Keys["user_id"].(string))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	dailies, err := d.DailyRepository.List(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dailies)
}

// FavDaily accepts a body request to update a daily & user
// @Summary update daily & user to apply fav feature
// @Description fav a daily
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.DailyRequestDTO true "DailyRequestDTO"
// @Success 200 {object} object
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 401 {object} object "Bad Gateway {"message': "message": "Wrong user"}"
// @Failure 500 {object} object "Bad Gateway {"message': "message": "Database error"}"
// @Router /api/daily/fav [put]
// @Security ApiKeyAuth
func (d *DailyController) Favourite(c *gin.Context) {
	var daily model.DailyRequestDTO
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}
	userId, err := primitive.ObjectIDFromHex(c.Keys["user_id"].(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong user id"})
		return
	}

	err = d.DailyRepository.FavouriteDaily(daily.ID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favourite success"})
}

// ViewDaily accepts a body request to update a daily & user
// @Summary update daily & user to apply view feature
// @Description view a daily
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.DailyRequestDTO true "DailyRequestDTO"
// @Success 200 {object} object
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 401 {object} object "Bad Gateway {"message': "message": "Wrong user"}"
// @Failure 500 {object} object "Bad Gateway {"message': "message": "Database error"}"
// @Router /api/daily/view [put]
// @Security ApiKeyAuth
func (d *DailyController) ViewDaily(c *gin.Context) {
	var daily model.DailyRequestDTO
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}
	userId, err := primitive.ObjectIDFromHex(c.Keys["user_id"].(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong user id"})
		return
	}

	err = d.DailyRepository.View(daily.ID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// ReportDaily accepts a body request to update a daily
// @Summary update daily to apply report feature
// @Description report a daily
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.ReportedDaily true "ReportedDaily"
// @Success 200 {object} object
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 502 {object} object "Bad Gateway {"message': "message": "Failed to update daily"}"
// @Router /api/daily/report [post]
// @Security ApiKeyAuth
func (d *DailyController) ReportDaily(c *gin.Context) {
	var reportedDaily model.ReportedDaily
	if err := c.ShouldBindJSON(&reportedDaily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
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

// DeleteDaily accepts a body request to update a daily
// @Summary delete the given daily
// @Description report a daily
// @Tags Daily
// @Accept json
// @Produce json
// @Param id path string true "Daily ID"
// @Success 200 {object} object
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 502 {object} object "Bad Gateway {"message': "message": "Failed to update daily"}"
// @Router /api/daily/{id} [delete]
// @Security ApiKeyAuth
func (d *DailyController) DeleteDaily(c *gin.Context) {
	user, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user was not found"})
		return
	}
	var daily2delete model.Daily
	id := c.Param("id")                          // Extract the id from the URL.
	objectID, _ := primitive.ObjectIDFromHex(id) // Convert string id to MongoDB ObjectID

	daily2delete, err := d.DailyRepository.FindById(objectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "daily could not be fetched"})
		return
	}
	// check if user has a daily with given id
	if daily2delete.Author != user {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unatuhorized to delete"})
		return
	}

	err = d.DailyRepository.Delete(daily2delete.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// EditDailyImage accepts a body request to update a daily's image
// @Summary update daily image
// @Description edit a daily's image
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.EditDailyImageDTO true "EditDailyImageDTO"
// @Success 200 {object} object
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 500 {object} object "Bad Gateway {"message': "message": "Database Error"}"
// @Router /api/daily/image [put]
// @Security ApiKeyAuth
func (d *DailyController) EditDailyImage(c *gin.Context) {
	var daily model.EditDailyImageDTO
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}
	err := d.DailyRepository.EditDailyImage(daily.ID, daily.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
