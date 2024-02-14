package router

import (
	"os"

	docs "github.com/Final-Projectors/daily-server/docs"
	"github.com/Final-Projectors/daily-server/handler"
	"github.com/Final-Projectors/daily-server/middleware"
	"github.com/Final-Projectors/daily-server/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	router := New()
	router.Run(os.Getenv("PORT"))
}

func New() *gin.Engine {
	router := gin.New()

	dailyHandler := handler.NewDailyController(*repository.NewDailyRepository(), *repository.NewReportedDailyRepository())

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
	api.PUT("/daily/fav", dailyHandler.FavDaily)
	api.PUT("/daily/view", dailyHandler.ViewDaily)
	api.POST("/daily/report", dailyHandler.ReportDaily)
	api.PUT("/daily/image", dailyHandler.EditDailyImage)
	api.DELETE("/daily/:id", dailyHandler.DeleteDaily)

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
