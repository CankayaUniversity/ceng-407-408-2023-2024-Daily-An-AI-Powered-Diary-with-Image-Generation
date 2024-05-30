package handler

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/Final-Projectors/daily-server/repository"
	"github.com/Final-Projectors/daily-server/utils"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DailyController struct {
	DailyRepository  *repository.DailyRepository
	ReportRepository *repository.ReportedDailyRepository
	UserRepository   *repository.UserRepository
	logger           *logrus.Logger
}

func NewDailyController(_userRepository *repository.UserRepository, _repository *repository.DailyRepository, _reports *repository.ReportedDailyRepository) *DailyController {
	return &DailyController{
		UserRepository: _userRepository, DailyRepository: _repository,
		ReportRepository: _reports,
		logger:           logrus.New(),
	}
}

// DownloadImage returns a specific daily via daily.ID
// @Summary returns a daily
// @Description returns a specific daily via daily.ID
// @Tags Daily
// @Accept json
// @Produce json
// @Param id path string true "Daily ID"
// @Success 200 {object} model.Daily
// @Failure 400 {object} object "Bad Request {"message": "Invalid JSON data"}"
// @Failure 500 {object} object "Internal Server Error {"message': "mongo: no documents in result"}"
// @Router /api/daily/image/{id} [get]
// @Security ApiKeyAuth
func (d *DailyController) DownloadImage(c *gin.Context) {
	// Define the location of the music file on the server
	id := c.Param("id")

	filename := id // Change this path to where your music file is located
	filename = "./image/" + filename

	// Check if the file exists
	_, err := os.Stat(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "File not found",
		})
		return
	}

	// Set headers for download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "image/jpeg") // Adjust the MIME type according to your music file format

	// Send the file as a response
	c.File(filename)
}

func (d *DailyController) GetImage(prompt string, dailyId primitive.ObjectID) error {
	var image repository.TextToImageImage
	image, err := repository.GenerateImage(prompt)
	if err != nil {
		return err
	}

	// Decode the base64 string
	data, err := base64.StdEncoding.DecodeString(image.Base64)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return err
	}

	// Define the directory and file path
	dir := "image"
	fileName := fmt.Sprintf("%v.jpg", dailyId.Hex())
	filePath := filepath.Join(dir, fileName)

	// Write the decoded bytes to the JPEG file
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	fmt.Println("File saved successfully to", filePath)
	return err
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
	client := openai.NewClient(os.Getenv("OPEN_API_KEY"))
	var daily model.Daily
	var createDailyDTO model.CreateDailyDTO
	err := c.ShouldBindJSON(&createDailyDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data"})
		return
	}

	daily.Favourites = 0                   // assuming Favourites is an integer.
	daily.Viewers = []primitive.ObjectID{} // assuming Viewers is an integer.

	dailyID := primitive.NewObjectID()
	daily.ID = dailyID
	daily.Text = createDailyDTO.Text
	daily.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	daily.IsShared = *createDailyDTO.IsShared

	prediction, err := utils.GetAIPrediction(daily.Text)
	d.logger.Infof("prediction: %v", prediction)
	if err != nil {
		d.logger.Infof(err.Error())
	}

	daily.Keywords = prediction.Keywords
	daily.Topics = prediction.Topics
	daily.Emotions = prediction.Emotions
	keyword_string := strings.Join(prediction.Keywords, ", ")

	if createDailyDTO.Prompt != "" {
		keyword_string = "prompt style: " + createDailyDTO.Prompt + ", prompt keywords:" + keyword_string
	}
	d.logger.Infof("keywords: %v", keyword_string)

	err = d.GetImage(keyword_string, dailyID)

	image := fmt.Sprintf("http://localhost:9090/api/daily/image/%v.jpg", dailyID.Hex())
	daily.Image = image

	// EMBEDDINGS
	queryReq := openai.EmbeddingRequest{
		Input: daily.Text,
		Model: openai.LargeEmbedding3,
	}
	targetResponse, err := client.CreateEmbeddings(c, queryReq)
	if err != nil {
		logrus.Fatalf("embedding creation failed")
	}
	embedding := targetResponse.Data[0].Embedding
	daily.Embedding = embedding

	// getting the user_id from context and running checks
	author, _ := c.Get("user_id")
	if auth, ok := author.(primitive.ObjectID); ok {
		daily.Author = auth
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
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

// GetExploreVS returns a list of shared dailies utilizing Vector Search
// @Summary returns 5 shared dailies
// @Description returns 5 shared dailies
// @Tags Daily
// @Accept json
// @Produce json
// @Success 200 {array} model.Daily
// @Failure 500 {object} object "Internal Server Error {"message': "Failed to fetch Dailies"}"
// @Failure 502 {object} object "Bad Gateway {"message': "No user"}"
// @Router /api/daily/explorevs [get]
// @Security ApiKeyAuth
func (d *DailyController) GetExploreVS(c *gin.Context) {
	author, _ := c.Get("user_id")

	dailies, err := d.DailyRepository.GetSimilarDailiesUnviewed(author.(primitive.ObjectID))
	if err != nil {
		// Assuming err would not be nil if there's an error, you can expose the error message.
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to fetch Dailies: %s", err.Error())})
		return
	}
	if len(dailies) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No Dailies found"})
		return
	}
	c.JSON(http.StatusOK, dailies)
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
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid id"))
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error() + ": " + objectID.String()})
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
// @Param        limit    query     int  false  "limit by q"  default(all)
// @Success 200 {array} model.Daily
// @Failure 500 {object} object "Bad Gateway {"message': "Couldn't fetch daily list"}"
// @Failure 502 {object} object "Bad Gateway {"message': "No user"}"
// @Router /api/daily/list [get]
// @Security ApiKeyAuth
func (d *DailyController) GetDailies(c *gin.Context) {
	author, _ := c.Get("user_id")
	limitStr := c.Query("limit") // Returns "" if "limit" is not provided
	var limit int

	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			limit = 0
		}
	} else {
		limit = 0
	}

	if _, ok := author.(primitive.ObjectID); !ok {
		fmt.Println("author is not a primitive.ObjectID")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	dailies, err := d.DailyRepository.List(author.(primitive.ObjectID), limit)
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

// GetExplore returns a list of shared dailies
// @Summary returns a list of shared dailies
// @Description returns a list of shared dailies
// @Tags Daily
// @Accept json
// @Produce json
// @Success 200 {array} model.Daily
// @Failure 500 {object} object "Bad Gateway {"message': "Failed to fetch Dailies"}"
// @Failure 502 {object} object "Bad Gateway {"message': "No user"}"
// @Router /api/daily/explore [get]
// @Security ApiKeyAuth
func (d *DailyController) GetExplore(c *gin.Context) {
	dailies, err := d.DailyRepository.GetExplore()
	if err != nil {
		// Assuming err would not be nil if there's an error, you can expose the error message.
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to fetch Dailies: %s", err.Error())})
		return
	}
	if len(dailies) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No Dailies found"})
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
// @Param id path string true "Daily ID"
// @Success 200 {object} object "Success {"message": "Favourite Success"}"
// @Failure 400 {object} object "Bad Request {"message": "Invalid JSON data"}"
// @Failure 401 {object} object "Bad Gateway {"message": "Unauthorized"}"
// @Failure 500 {object} object "Internal Server Error {"message": "Database error"}"
// @Router /api/daily/fav/{id} [put]
// @Security ApiKeyAuth
func (d *DailyController) FavDaily(c *gin.Context) {
	d.logger.Infof("FavDaily function executing")

	id := c.Param("id")                            // Extract the id from the URL.
	objectID, err := primitive.ObjectIDFromHex(id) // Convert string id to MongoDB ObjectID
	if err != nil {
		d.logger.Infof("couldn't convert daily id string to primitive.ObjectID")
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error() + objectID.String()})
		return
	}

	user, _ := c.Get("user_id")
	err = d.DailyRepository.FavouriteDaily(objectID, user.(primitive.ObjectID))
	if err != nil {
		d.logger.Infof("error in FavouriteDaily function")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	daily, err := d.DailyRepository.FindById(objectID)
	if err != nil {
		d.logger.Infof("no daily found with the given ObjectID")
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var keywords []string
	if daily.Keywords != nil && len(daily.Keywords) > 0 {
		keywords = daily.Keywords
	}
	topics := daily.Topics

	err = d.DailyRepository.UpdateUserPreferences(keywords, topics, user.(primitive.ObjectID))
	if err != nil {
		d.logger.Infof("Error in UpdateUserPreferences function")
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
// @Success 200 {object} object "Success {"message": "Viewed Successfully"}"
// @Failure 400 {object} object "Bad Request {"message": "Invalid JSON data"}"
// @Failure 401 {object} object "Bad Gateway {"message": "Wrong user id"}"
// @Failure 500 {object} object "Bad Gateway {"message": "Database error"}"
// @Router /api/daily/view [put]
// @Security ApiKeyAuth
func (d *DailyController) ViewDaily(c *gin.Context) {
	id := c.Param("id")                            // Extract the id from the URL.
	objectID, err := primitive.ObjectIDFromHex(id) // Convert string id to MongoDB ObjectID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error() + objectID.String()})
		return
	}

	user, _ := c.Get("user_id")
	err = d.DailyRepository.View(objectID, user.(primitive.ObjectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Viewed Successfully"})
}

// ReportDaily accepts a body request to update a daily
// @Summary update daily to apply report feature
// @Description report a daily
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.ReportedDaily true "ReportedDaily"
// @Success 200 {object} object "Success {"message": "Created Successfully"}"
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 502 {object} object "Bad Gateway {"message": "Failed to update daily"}"
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
		c.JSON(http.StatusOK, gin.H{"message": "Created Successfully"})
		return
	}
	if err := result.Decode(&dbReportedDaily); err == nil {
		reportOperation := bson.M{"$inc": bson.M{"reports": 1}}
		if _, err := database.ReportedDailies.UpdateOne(c, dbReportedDaily, reportOperation); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update reported daily", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Incremented Successfully"})
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
// @Success 200 {object} object "Success {"message": "Deleted Successfully"}"
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
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}

// EditDailyImage accepts a body request to update a daily's image
// @Summary update daily image
// @Description edit a daily's image
// @Tags Daily
// @Accept json
// @Produce json
// @Param daily body model.EditDailyImageDTO true "EditDailyImageDTO"
// @Success 200 {object} object "Success {"message": "Image Edited Successfully"}"
// @Failure 400 {object} object "Bad Request {"message': "Invalid JSON data"}"
// @Failure 500 {object} object "Bad Gateway {"message": "Database Error"}"
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image Edited Successfully"})
}
