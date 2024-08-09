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

const (
	MONGO_URI     = "mongodb://localhost:27017"
	GSQL_URI      = "localhost:5455"
	GQUE_URI      = "localhost:5456"
	DATABASE_NAME = "gcart"
)

func SetupDatabase() {

	collections := models.GetAllGnosqlCollections()

	GqueClient = gque_client.Connect(GQUE_URI, DATABASE_NAME)
	GnoSQLDB = gnosql_client.Connect(GSQL_URI, DATABASE_NAME, true)

	GnoSQLDB.CreateCollections(collections)

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

	DB = client.Database("gcart")
}
