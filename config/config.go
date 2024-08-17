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
	"github.com/nanda03dev/gque_client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var GnoSQLDB *gnosql_client.Database
var GqueClient *gque_client.Client

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

var (
	MONGO_URI     = "mongodb://localhost:27017"
	GNOSQL_SERVER = "localhost:5455"
	GQUE_SERVER   = "localhost:5456"
	DATABASE_NAME = "gcart"
)

func SetupDatabase() {
	if value := os.Getenv("MONGO_URI"); value != "" {
		MONGO_URI = value
	}

	if value := os.Getenv("GNOSQL_SERVER"); value != "" {
		GNOSQL_SERVER = value
	}

	if value := os.Getenv("GQUE_SERVER"); value != "" {
		GQUE_SERVER = value
	}

	if value := os.Getenv("DATABASE_NAME"); value != "" {
		DATABASE_NAME = value
	}

	// GQUE client connection
	GqueClient = gque_client.Connect(GQUE_SERVER, DATABASE_NAME)

	// GnoSQL DB client connection
	GnoSQLDB = gnosql_client.Connect(GNOSQL_SERVER, DATABASE_NAME, true)
	GnoSQLDB.CreateCollections(models.GetAllGnosqlCollections())

	// MONGO DB client connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_URI).SetConnectTimeout(1 * time.Second).SetServerAPIOptions(serverAPI)

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

	DB = client.Database(DATABASE_NAME)
}
