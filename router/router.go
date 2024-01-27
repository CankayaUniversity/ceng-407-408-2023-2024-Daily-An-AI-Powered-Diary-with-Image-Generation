package router

import (
	"os"

	"github.com/Final-Projectors/daily-server/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := New()
	router.Run(os.Getenv("PORT"))
}

func New() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	api.POST("/users", handler.CreateUser)
	api.GET("/users", handler.GetUsers)
	api.POST("/dailies", handler.CreateDaily)
	api.GET("/dailies", handler.GetDailies)
	return router
}
