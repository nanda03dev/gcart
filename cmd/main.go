package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go2ms/config"
	"github.com/nanda03dev/go2ms/routes"
)

func main() {
	config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)

	router := routes.SetupRouter()

	PORT := os.Getenv("PORT")
	log.Println("Server running at http://localhost:" + PORT)
	log.Fatal(router.Run(":" + PORT))

}
