package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/RealHenryMartinez/ChatApp.git/database"
	"github.com/RealHenryMartinez/ChatApp.git/handlers"
	"github.com/gorilla/mux"
	
)

func main() {
	r := mux.NewRouter() // Router shared among all routes

	// Create a new manager
	manager := handlers.NewManager() // Globally for all clients
	fmt.Println("OIEJFOI")
	url := handlers.ViperEnvVariable("MONGO_URI")
	database.MakeDatabase(url)
	if database.DATABASE == nil {
		log.Fatal("Database is not properly initialized")
	}
	defer database.DATABASE.Client().Disconnect(context.Background())
	r.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
        manager.HandleConnections(w, r)
    })

    // HTTP route
    r.HandleFunc("/database", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Hello world")
    })

    // Subrouter for /database/get-chat-messages
    getChatRoomLog := r.PathPrefix("/database/get-chat-messages/{room_number}").Subrouter()
    getChatRoomLog.HandleFunc("", database.RetrieveChatMessagesHandler)

    // Serve static files
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./website")))

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
