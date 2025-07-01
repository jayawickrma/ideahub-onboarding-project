package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitMongoDB() {
	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DATABASE_NAME")

	if mongoURI == "" || dbName == "" {
		panic("MONGO_URI or DATABASE_NAME not set in .env")
	}

	// Mongo connection options
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to Mongo
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// Ping the DB
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	DB = client.Database(dbName)
	fmt.Println("✅ Connected to MongoDB:", dbName)
}
