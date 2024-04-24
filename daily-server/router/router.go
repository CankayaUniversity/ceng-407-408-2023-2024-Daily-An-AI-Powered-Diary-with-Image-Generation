package router

import (
	"fmt"
	"os"
	"time"

	docs "github.com/Final-Projectors/daily-server/docs"
	"github.com/Final-Projectors/daily-server/handler"
	"github.com/Final-Projectors/daily-server/middleware"
	"github.com/Final-Projectors/daily-server/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	router := New()
	url := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	router.Run(url)
}

func New() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorLogger())

	dailyHandler := handler.NewDailyController(repository.NewUserRepository(), repository.NewDailyRepository(repository.NewUserRepository()), repository.NewReportedDailyRepository())

	// Cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//http://localhost:9090/docs/index.html
	docs.SwaggerInfo.BasePath = ""
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	admin := router.Group("/admin")
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)

	// user and admin rights here
	api.Use(middleware.JwtAuthMiddleware())
	api.POST("/daily", dailyHandler.CreateDaily)
	api.GET("/daily/:id", dailyHandler.GetDaily)
	api.GET("/daily/list", dailyHandler.GetDailies)
	api.POST("/daily/fav/:id", dailyHandler.FavDaily)
	api.POST("/daily/view/:id", dailyHandler.ViewDaily)
	api.POST("/daily/report", dailyHandler.ReportDaily)
	api.PUT("/daily/image", dailyHandler.EditDailyImage)
	api.DELETE("/daily/:id", dailyHandler.DeleteDaily)
	api.GET("/daily/explore", dailyHandler.GetExplore)
	// moderator rights here
	admin.Use(middleware.JwtAuthMiddlewareRole("moderator"))
	admin.DELETE("/deleteUser", handler.DeleteUserAdmin)

	// admin rights here
	admin.Use(middleware.JwtAuthMiddlewareRole("admin"))
	admin.POST("/grantMod", handler.GrantModRights)
	admin.PUT("/takeMod", handler.TakeModRights)

	return router
}

/*
	func JwtAuthMiddlewareAdmin() function can be used to perform admin checks on endpoints
*/
