package router

import (
	"os"

	docs "github.com/Final-Projectors/daily-server/docs"
	"github.com/Final-Projectors/daily-server/handler"
	"github.com/Final-Projectors/daily-server/middleware"
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

	//http://localhost:9090/docs/index.html
	docs.SwaggerInfo.BasePath = ""
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	admin := router.Group("/admin")
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)

	// user and admin rights here
	api.Use(middleware.JwtAuthMiddleware())
	api.POST("/daily", handler.CreateDaily)
	api.GET("/daily/:id", handler.GetDaily)
	api.GET("/daily/list", handler.GetDailies)
	api.PUT("/daily/fav", handler.FavDaily)
	api.PUT("/daily/view", handler.ViewDaily)
	api.POST("/daily/report", handler.ReportDaily)
	api.PUT("/daily/image", handler.EditDailyImage)

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
