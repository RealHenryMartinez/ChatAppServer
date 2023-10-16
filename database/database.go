package database

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DATABASE *mongo.Database
var ChatCollection *mongo.Collection

func MakeDatabase(uri string) {
	
	// Setting the database
    clientOptions := options.Client().ApplyURI(uri)
    mongoClient, err := mongo.Connect(context.Background(), clientOptions)

    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

	// Check for the database is active
    err = mongoClient.Ping(context.Background(), readpref.Primary())
    if err != nil {
        log.Fatalf("Error pinging MongoDB: %v", err)
    }

	// Initialize the mongodb connection with the collections
    DATABASE = mongoClient.Database("MongoWithGo")
	ChatCollection = DATABASE.Collection("chat-rooms")

	if DATABASE == nil {
		log.Fatal("Database is not properly initialized")
	}
}
