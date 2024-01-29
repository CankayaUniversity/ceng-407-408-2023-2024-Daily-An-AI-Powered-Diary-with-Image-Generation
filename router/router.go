package router

import (
	"os"

	"github.com/Final-Projectors/daily-server/handler"
	"github.com/Final-Projectors/daily-server/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := New()
	router.Run(os.Getenv("PORT"))
}

func New() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)
	api.Use(middleware.JwtAuthMiddleware())
	api.POST("/dailies", handler.CreateDaily)
	api.GET("/dailies", handler.GetDailies)
	return router
}
