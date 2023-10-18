package handlers

import (
	"encoding/json"
	"strings"

	//"encoding/json"
	"fmt"
	//"net/http"
	"sync"

	"github.com/RealHenryMartinez/ChatApp.git/database"
)

var muter = sync.RWMutex{} // Be able to read the messages sent to the users

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

		_, message, _ := c.conn.ReadMessage()
		var requestType database.Request

		// Message received changed to json to the writer
		if err := json.Unmarshal(message, &requestType); err != nil {
			return
		}

		fmt.Println("READING MESSAGE", string(message))

		if strings.Compare(requestType.Type, "delete") == 0 {

		} else if strings.Compare(requestType.Type, "create_chat") == 0 {
			// Our Chat App
			var content database.ChatRoomLog

			// Out Chat room map containing keys for fields and any value accepted
			if contentMap, ok := requestType.Content.(map[string]interface{}); ok {

				// Our field that content-map accepts
				messagesInterface, _ := contentMap["messages"].([]interface{})

				// Making our field messages field of a slice of chat documents
				messages := make([]database.ChatDocument, len(messagesInterface))

				// Type assertion succeeded for a map
				content = database.ChatRoomLog{
					Messages:   messages,
					RoomNumber: uint32(contentMap["room_number"].(float64)),
				}
			}

			// Create a document in our database of a chat room
			database.CreateOne(database.ChatCollection, content)

			ConvertToJSONAndSendToAll(content, c)
		} else {
			var content database.ChatDocument

			/*
					1. Converting the Content to a map of strings for fields from chat document
				 	2. We convert the fields to the appropriate field types from chat document
			*/
			if contentMap, ok := requestType.Content.(map[string]interface{}); ok {
				// Type assertion succeeded for a map
				content = database.ChatDocument{
					Message:    contentMap["message"].(string),              // Type assertion for string
					RoomNumber: uint32(contentMap["room_number"].(float64)), // Type assertion for uint32
					Uid:        contentMap["uid"].(string),                  // Type assertion for string
				}
			}
			// Turn the request content of any type into a database.ChatDocument type
			handleAddObjectToSliceOnDB(c, content, content.RoomNumber, "room_number", "messages", database.ChatCollection)

			ConvertToJSONAndSendToAll(content, c)
		}
	}

}
