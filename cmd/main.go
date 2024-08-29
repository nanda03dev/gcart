package main

import (
	"log"
	"os"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/config"
	"github.com/nanda03dev/gcart/src/consumers"
	"github.com/nanda03dev/gcart/src/message_queue"
	"github.com/nanda03dev/gcart/src/repositories"
	"github.com/nanda03dev/gcart/src/routes"
	"github.com/nanda03dev/gcart/src/services"
	"github.com/nanda03dev/gcart/src/workers"
)

func main() {
	config.LoadConfig()

	common.InitializeChannels()

	config.SetupDatabase()

	repositories.InitializeRepositories()

	services.InitializeServices()

	message_queue.InitializeGque()

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
