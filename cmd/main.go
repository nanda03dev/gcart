package main

import (
	"log"
	"os"

	"github.com/nanda03dev/gcart/common"
	"github.com/nanda03dev/gcart/config"
	"github.com/nanda03dev/gcart/repositories"
	"github.com/nanda03dev/gcart/routes"
	"github.com/nanda03dev/gcart/services"
	"github.com/nanda03dev/gcart/workers"
)

func main() {
	config.LoadConfig()

	common.InitializeChannels()

	database := config.SetupDatabase()

	repositories.InitializeRepositories(database)

	services.InitializeServices(database)

	go workers.InitiateWorker()

	InitiateServer()
}

func InitiateServer() {
	PORT := os.Getenv("PORT")
	log.Println("Server running at http://localhost:" + PORT)
	router := routes.InitializeRouter()
	router.Run(":" + PORT)
}
