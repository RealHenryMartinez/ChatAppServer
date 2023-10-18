package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create a document for the given collection and return a struct of the created item
func CreateOne(collection *mongo.Collection, data interface{}) (*mongo.InsertOneResult, error) {
	// Create (Insert) a document
	result, err := collection.InsertOne(context.Background(), data)

	if err != nil { // Check if created successfully
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

func DeleteOne(collection *mongo.Collection, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
    result, err := collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GetByFilter[T ChatRoomLog | interface{}](collection *mongo.Collection, filter bson.M) (T, error) {

	// Define a filter to query for chat messages based on the chat room number.

	// Query the database and retrieve the chat messages.
	var result T
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Result: ", result, " Error: " , err)
		return result, err
	}

	return result, nil
}
