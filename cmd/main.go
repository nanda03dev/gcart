package main

import (
	"log"
	"os"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/config"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/routes"
	"github.com/nanda03dev/go2ms/services"
	"github.com/nanda03dev/go2ms/workers"
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
