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
	admin := router.Group("/admin")
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)

	// user and admin rights here
	api.Use(middleware.JwtAuthMiddleware())
	api.POST("/createDaily", handler.CreateDaily)
	api.GET("/getDailies", handler.GetDailies)
	api.PUT("/favDaily", handler.FavDaily)
	api.PUT("/viewDaily", handler.ViewDaily)
	api.POST("/reportDaily", handler.ReportDaily)
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
