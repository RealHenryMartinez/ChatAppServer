package handlers

import "github.com/gorilla/websocket"


type Client struct {
	conn *websocket.Conn
	manager *Manager
	// We can convert this to any data, or a string
	dataTransfer chan []byte
}

// A map of clients and check if they are connected to the web socket
type ClientList map[*Client]bool

// Update the list of clients with a new client
func createClient(c *websocket.Conn, m *Manager) *Client {
	return &Client{
		conn: c, // Connection to the server
		manager: m, // Be able to send messages to this connection's client manager
		dataTransfer: make(chan []byte), // Transfered data 
	}
}

func (c *Client) removeClient() {
	c.conn.Close()
	delete(c.manager.clients, c)
}