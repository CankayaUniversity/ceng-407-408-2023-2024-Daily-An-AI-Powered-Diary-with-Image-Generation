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
	api.POST("/createDaily", handler.CreateDaily)
	api.GET("/getDailies", handler.GetDailies)
	api.PUT("/favDaily", handler.FavDaily)
	api.PUT("/viewDaily", handler.ViewDaily)
	return router
}
