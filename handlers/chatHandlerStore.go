package handlers

import "github.com/RealHenryMartinez/ChatApp.git/database"

type FunctionStore map[string]func(*Client)

func (fs FunctionStore) AddFunction(key string, f func(*Client)) {
	fs[key] = f
}

var RequestType database.Request

var HandlerCreators = make(FunctionStore)

func CreateHandlers() {
	HandlerCreators.AddFunction("create_chat", CreateChatStore)
	HandlerCreators.AddFunction("create_message", CreateMessageStore)

}
func CreateDeleteMessage() {

}

func CreateChatStore(c *Client) {
	// Our Chat App
	var content database.ChatRoomLog

	// Out Chat room map containing keys for fields and any value accepted
	if contentMap, ok := RequestType.Content.(map[string]interface{}); ok {

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
}

func CreateMessageStore(c *Client) {
	var response database.Request
	var requestBody database.ChatDocument

	/*
			1. Converting the Content to a map of strings for fields from chat document
		 	2. We convert the fields to the appropriate field types from chat document
	*/
	if contentMap, ok := RequestType.Content.(map[string]interface{}); ok {
		// Type assertion succeeded for a map
		requestBody = database.ChatDocument{
			Message:    contentMap["message"].(string),              // Type assertion for string
			RoomNumber: uint32(contentMap["room_number"].(float64)), // Type assertion for uint32
			Uid:        contentMap["uid"].(string),                  // Type assertion for string
		}

		// Preparing the response to send to the server and user(s)
		response = database.Request{
			Type:    "sent_message",
			Content: requestBody,
		}
	}

	// Send the content to the database
	handleAddObjectToSliceOnDB(c, requestBody, requestBody.RoomNumber, "room_number", "messages", database.ChatCollection)

	// Send the response to the users channel
	ConvertToJSONAndSendToAll(response, c)
}
