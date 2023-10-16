package handlers

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Write a message to the connection | Send the message to the client
func (c *Client)WriteMessage() {
	defer c.removeClient() // Client is closed when it does not receive data
	for {
		//muter.Lock() <- don't need to block the writing to all channels
	
		//mt, message, _ := c.conn.ReadMessage()
		message := <-c.dataTransfer // Pop received message | Only receives the message bytes | 

		// Write a message to send to clients
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
		}
		fmt.Println("WRITING MESSAGE", string(message))
		//defer muter.Unlock()

	}
	

}

