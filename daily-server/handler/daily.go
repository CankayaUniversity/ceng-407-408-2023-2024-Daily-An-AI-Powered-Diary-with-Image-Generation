package handler

import (
	"context"
	"fmt"
	"net/http"
	"sync"
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
// @Failure 401 {object} object "Unauthorized {"message': "Unauthorized"}"
// @Failure 500 {object} object "Internal Server Error {"message': "Couldn't fetch the image"}"
// @Failure 502 {object} object "Bad Gateway {"message': "Couldn't fetch the image / DB error"}"
// @Router /api/daily [post]
// @Security ApiKeyAuth
func (d *DailyController) CreateDaily(c *gin.Context) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var daily model.Daily
	var createDailyDTO model.CreateDailyDTO
	err := c.ShouldBindJSON(&createDailyDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	daily.Keywords = []string{} // assuming Keywords is a string slice.
	daily.Emotions = model.Emotion{}
	daily.Favourites = 0                   // assuming Favourites is an integer.
	daily.Viewers = []primitive.ObjectID{} // assuming Viewers is an integer.

	dailyID := primitive.NewObjectID()
	daily.ID = dailyID
	daily.Text = createDailyDTO.Text
	daily.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	daily.IsShared = *createDailyDTO.IsShared

	wg.Add(1)
	go func() {
		defer wg.Done()
		flaskData, err := utils.GetDataFromFlask(daily.Text)
		if err != nil {
			mu.Lock()
			defer mu.Unlock()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		mu.Lock()
		defer mu.Unlock()
		if emotionsMap, ok := flaskData["emotions"].(map[string]interface{}); ok {
			intEmotions := make(map[string]int)
			for key, value := range emotionsMap {
				switch v := value.(type) {
				case int:
					intEmotions[key] = v
				case float64: // Handle float64 as well
					intEmotions[key] = int(v)
				default:
					fmt.Printf("Error: Value for key %s is not an integer: %v\n", key, value)
				}
			}
			daily.Emotions.Sadness = intEmotions["Sadness"]
			daily.Emotions.Joy = intEmotions["Joy"]
			daily.Emotions.Love = intEmotions["Love"]
			daily.Emotions.Anger = intEmotions["Anger"]
			daily.Emotions.Fear = intEmotions["Fear"]
			daily.Emotions.Surprise = intEmotions["Surprise"]
		} else {
			fmt.Println("Error: Value in flaskData['emotions'] is not a map[string]interface{}")
		}
		daily.Image = flaskData["image"].(string)
	}()

	// getting the user_id from context and running checks
	author, _ := c.Get("user_id")
	if auth, ok := author.(primitive.ObjectID); ok {
		daily.Author = auth
	} else {
		fmt.Println("author is not a primitive.ObjectID")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	err = d.DailyRepository.Create(&daily)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Daily created successfuly without Image and Emotions"})

	//preparation to ai response
	var fetchedDaily model.Daily
	result := database.Dailies.FindOne(c, bson.M{"_id": dailyID})
	if result.Err() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Daily could not be fetched"})
		return
	}
	if err := result.Decode(&fetchedDaily); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode fetched daily document"})
		return
	}

	// Wait for the async operation to finish
	wg.Wait()

	fetchedDaily.Emotions = daily.Emotions
	if createDailyDTO.Image == "" {
		fetchedDaily.Image = daily.Image
	} else {
		//assuming that createDailyDTO.Image is a propper base64 data
		fetchedDaily.Image = createDailyDTO.Image
	}

	updatedBSON, err := bson.Marshal(fetchedDaily)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal updated daily document"})
		return
	}
	filter := bson.M{"_id": dailyID}
	updateResult, err := database.Dailies.ReplaceOne(context.TODO(), filter, updatedBSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update daily document"})
		return
	}
	if updateResult.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found or not updated"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Daily document updated with Image and Emotions successfully"})
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
	author, _ := c.Get("user_id")
	if _, ok := author.(primitive.ObjectID); !ok {
		fmt.Println("author is not a primitive.ObjectID")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	dailies, err := d.DailyRepository.List(author.(primitive.ObjectID))
	if err != nil {
		// Assuming err would not be nil if there's an error, you can expose the error message.
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to fetch Dailies: %s", err.Error())})
		return
	}
	if len(dailies) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No Dailies found for User"})
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
// @Failure 401 {object} object "Bad Gateway {"message': "message": "Unauthorized"}"
// @Failure 500 {object} object "Bad Gateway {"message': "message": "Database error"}"
// @Router /api/daily/fav [put]
// @Security ApiKeyAuth
func (d *DailyController) Favourite(c *gin.Context) {
	user, _ := c.Get("user_id")
	if _, ok := user.(primitive.ObjectID); !ok {
		fmt.Println("author is not a primitive.ObjectID")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var daily model.DailyRequestDTO
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	err := d.DailyRepository.FavouriteDaily(daily.ID, user.(primitive.ObjectID))
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
// @Failure 401 {object} object "Bad Gateway {"message': "message": "Wrong user id"}"
// @Failure 500 {object} object "Bad Gateway {"message': "message": "Database error"}"
// @Router /api/daily/view [put]
// @Security ApiKeyAuth
func (d *DailyController) ViewDaily(c *gin.Context) {
	user, _ := c.Get("user_id")
	if _, ok := user.(primitive.ObjectID); !ok {
		fmt.Println("author is not a primitive.ObjectID")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var daily model.DailyRequestDTO
	if err := c.ShouldBindJSON(&daily); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	err := d.DailyRepository.View(daily.ID, user.(primitive.ObjectID))
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
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data | user was not found"}"
// @Failure 400 {object} object "Unauthorized {"message': "Unauthorized"}"
// @Failure 502 {object} object "Internal Server Error"
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
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
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
