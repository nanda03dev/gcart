package main

import (
	"log"
	"os"

	"github.com/nanda03dev/gcart/common"
	"github.com/nanda03dev/gcart/config"
	"github.com/nanda03dev/gcart/consumers"
	"github.com/nanda03dev/gcart/message"
	"github.com/nanda03dev/gcart/repositories"
	"github.com/nanda03dev/gcart/routes"
	"github.com/nanda03dev/gcart/services"
	"github.com/nanda03dev/gcart/workers"
)

func main() {
	config.LoadConfig()

	common.InitializeChannels()

	config.SetupDatabase()

	repositories.InitializeRepositories()

	services.InitializeServices()

	message.InitializeGque()

	go workers.InitializeWorker()

	go consumers.InitializeConsumer()

	InitializeServer()
}

var GCART_PORT = "5457"

func InitializeServer() {

	if value := os.Getenv("GCART_PORT"); value != "" {
		GCART_PORT = value
	}
	log.Println("Server running at http://localhost:" + GCART_PORT)
	router := routes.InitializeRouter()
	router.Run(":" + GCART_PORT)
}
