package handlers

import (
	"context"

	"fmt"
	//"strconv"

	"github.com/RealHenryMartinez/ChatApp.git/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func handleDelete(c *Client, content string, roomNumber uint64) {
	database.DeleteChatMessageByMessage(roomNumber, content)
	// Send the message to the clients
	for wsClient := range c.manager.clients {
		wsClient.dataTransfer <- []byte(content) // Send the message UTF-8 Data to client
	}
}

func handleAddObjectToSliceOnDB[T database.ChatDocument | interface{}, G uint32 | string](c *Client, content T, filterBy G, filterKey string, contentKey string, collection *mongo.Collection) {
	// Create a filter to identify the document to update

	// Create the chat message to be added
	filter := bson.M{filterKey: filterBy}

	// Create an update with the $push operator
	update := bson.M{
		"$push": bson.M{
			contentKey: &content,
		},
	}
	// Update in database
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return
	}

	// The update was successful
	fmt.Printf("Matched %v document(s) and modified %v document(s)\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	
}
