package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	//"go.mongodb.org/mongo-driver/mongo"
)

type Manager struct {
	clients ClientList
}

// Setting up configurations for websocket limits
var upgrader = websocket.Upgrader{

	// How much our channels is allowed to store in memory | Temporarily
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// Point to the client and add it to the manager
func (m *Manager) addClient(client *Client) {
	fmt.Println("Client added to manager: ", client)
	m.clients[client] = true
	fmt.Println("CLIENTLIST: ", m.clients)
}

func (m *Manager) handleNewClient(conn *websocket.Conn) {
	// Just create a new connection
	client := createClient(conn, m)
	m.addClient(client)

	// Read for the clients connection
	go client.ReadMessages()
	go client.WriteMessage() 
	
}

func (m *Manager) HandleConnections(w http.ResponseWriter, r *http.Request){

	// Handle Upgrade to websocket Errors
	conn, err := upgrader.Upgrade(w, r, nil)
	
	if err != nil {
		fmt.Print("upgrade failed: ", err)
		return
	}

	// Add to the manager
	m.handleNewClient(conn)
}

// Create a new manager with an empty client list
func NewManager() *Manager {
	manager := &Manager{
		clients: make(ClientList),
	}

	return manager
}

