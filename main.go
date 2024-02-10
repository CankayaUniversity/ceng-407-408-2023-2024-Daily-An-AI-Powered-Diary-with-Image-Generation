package main

import (
	"log"

	docs "github.com/Final-Projectors/daily-server/docs"
	"github.com/gin-gonic/gin"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/router"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//http://localhost:9090/docs/index.html

// @title Daily API
// @version 1.0
// @host localhost:9090
// @BasePath /api

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.Init()
	router.Init()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":9090")
}
