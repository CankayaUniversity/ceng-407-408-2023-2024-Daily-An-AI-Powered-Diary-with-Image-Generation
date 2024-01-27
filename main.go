package main

import (
	"log"

	"github.com/Final-Projectors/daily-server/database"
	"github.com/Final-Projectors/daily-server/router"
	"github.com/joho/godotenv"
)

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
