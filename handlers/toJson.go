package handlers

import (
	"encoding/json"
	"log"
)

func ConvertToJSONAndSendToAll(content any, c *Client){
	// Convert to JSON to convert to []byte
	contentJSON, err := json.Marshal(content)

	if err != nil {
		log.Fatal("Error marshalling: ", contentJSON)
		return
	}

	// Send the message to the clients
	for wsClient := range c.manager.clients {
		wsClient.dataTransfer <- []byte(contentJSON) // Send the message json Data to client
	}
}