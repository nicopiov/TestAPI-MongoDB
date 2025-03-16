package main

import (
	"context"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log/slog"
	_ "net/http"
	"os"
)

type charStats struct {
	Level        int `json:"level"`
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
}

func main() {
	//Retrieving mongodb string connection through env file
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file: ", err)
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		slog.Error("MONGODB_URI environment variable not set")
	}

	//Connecting to the MongoDB instance
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		slog.Error("Error connecting to MongoDB:", err)
	} else {
		slog.Info("Connected to MongoDB successfully")
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			slog.Error("An error occurred while disconnecting from mongo:", err)
		}
	}()

	//Retrieving the name of the database created with the Mongo express interface
	database := mongoClient.Database("test-data")
	fmt.Println(database.Name())
}
