package handlers

import (
	"encoding/json"
	"fmt"
	"sync"
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

		// Message received changed to json to the writer
		if err := json.Unmarshal(message, &RequestType); err != nil {
			fmt.Printf("Error unmarshaling message: %v\n", err)
			break // Continue processing other messages
		}

		fmt.Println("READING MESSAGE", string(message))

		fmt.Println("Message requested: ", RequestType.Content)

		// Call the request handler and get the response to server and user(s)
		HandlerCreators[RequestType.Type](c)
	}

}
