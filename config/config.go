package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nanda03dev/gcart/models"
	"github.com/nanda03dev/gnosql_client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var GnoSQLDB *gnosql_client.Database

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	if os.Getenv("PORT") != "" {
		os.Setenv("PORT", "8080")
	}
	if os.Getenv("GIN_MODE") != "" {
		os.Setenv("GIN_MODE", gin.ReleaseMode)
	}
	gin.SetMode(os.Getenv("GIN_MODE"))

}

func SetupDatabase() *mongo.Database {
	DatabaseName := "gcart"

	mongoURI := "mongodb://localhost:27017"

	collections := models.GetAllGnosqlCollections()

	GnoSQLDB = gnosql_client.Connect("localhost:5455", DatabaseName, true)

	GnoSQLDB.CreateCollections(collections)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetConnectTimeout(1 * time.Second).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		println("Connection timed out")
		panic(err)
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	DB = client.Database("gcart")
	return DB
}
