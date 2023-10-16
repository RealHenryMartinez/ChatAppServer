package handlers

import (
	"context"
	//"encoding/json"
	"fmt"
	//"net/http"
	"sync"

	"github.com/RealHenryMartinez/ChatApp.git/database"
	"go.mongodb.org/mongo-driver/bson"
)

var muter = sync.RWMutex{}

// Read incoming client messages from frontend to server to clients
func (c *Client) ReadMessages() {
	muter.RLock() // Allows multiple goroutine readers to read messages

	defer func() {
		// Remove the client when message not sent
		c.removeClient()
		muter.RUnlock() // Cleanup the client
	}()

	for {
		// Return message as bytes
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("READING MESSAGE", string(message))

		// Creating a message for the database

		// err = json.NewDecoder(r.Body).Decode(&chatMessage)
		// if err != nil {
		// 	//http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// Specify the room number for the chat message
		roomNumber := "42" // Set to the desired room number

		// Create a filter to identify the document to update
		filter := bson.M{"room_number": roomNumber}

		// Create the chat message to be added
		chatMessage := &database.ChatDocument{
			Message: string(message),
		}

		// Create an update with the $push operator
		update := bson.M{
			"$push": bson.M{
				"messages": chatMessage,
			},
		}

		// Perform the update operation
		updateResult, err := database.ChatCollection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			// Handle the error
			// You can choose to return an error response, log the error, or perform other error handling as needed.
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// The update was successful
		fmt.Printf("Matched %v document(s) and modified %v document(s)\n", updateResult.MatchedCount, updateResult.ModifiedCount)

		// Send the message to the clients
		for wsClient := range c.manager.clients {
			fmt.Println("SENDING MESSAGE", string(message))
			wsClient.dataTransfer <- message // Send the message UTF-8 Data to client
		}
	}

}
