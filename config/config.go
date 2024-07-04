package config

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	} else {
		log.Println("MongoDB connected successfully")
	}

	DB = client.Database("order_management_system")
}
