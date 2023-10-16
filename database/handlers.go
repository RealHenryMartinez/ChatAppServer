package database

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create a document for the given collection and return a struct of the created item
func CreateOne(collection *mongo.Collection, data struct{}) (*mongo.InsertOneResult, error) {
	// Create (Insert) a document
	result, err := collection.InsertOne(context.Background(), data)

	if err != nil { // Check if created successfully
		fmt.Println(err)
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

func RetrieveChatMessagesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	chatRoom := params["room_number"]
	data, err := RetrieveChatMessagesHelper(chatRoom)

	fmt.Println("data: ", data)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the data as JSON and write it to the response
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RetrieveChatMessages retrieves chat messages for a specific chat room and returns them as a string.
func RetrieveChatMessagesHelper(chatRoom string) (ChatRoomLog, error) {

	// Access the desired database and collection.

	result, err := GetByFilter[ChatRoomLog](ChatCollection, bson.M{"room_number": string(chatRoom)})

	if err != nil {
		fmt.Println("Room not found: " + chatRoom)
	}
	// Create a string containing the chat messages.
	// var chatLog string
	// for _, message := range result.Messages {
	// 	chatLog += message.Message + "\n"
	// }

	return result, nil
}
