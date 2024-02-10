package main

import (
	"log"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/router"
	"github.com/joho/godotenv"
)

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
}
