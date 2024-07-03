package main

import (
	"log"

	"github.com/nanda03dev/oms/config"
	"github.com/nanda03dev/oms/routes"
)

func main() {
	config.LoadConfig()
	router := routes.SetupRouter()

	log.Println("Server running at http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
