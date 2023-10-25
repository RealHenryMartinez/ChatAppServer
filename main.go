package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/RealHenryMartinez/ChatApp.git/database"
    "github.com/RealHenryMartinez/ChatApp.git/handlers"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    r := mux.NewRouter() // Router shared among all routes

    // Create a new manager
    manager := handlers.NewManager() // Globally for all clients
    url := handlers.ViperEnvVariable("MONGO_URI")

    database.MakeDatabase(url)

    if database.DATABASE == nil {
        log.Fatal("Database is not properly initialized")
    }
    defer database.DATABASE.Client().Disconnect(context.Background()) // Disconnect the db connection

    // WebSocket connection
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

    // Create a CORS handler with the desired CORS options
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5173"}, // Replace with your allowed origins
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    })

    // Use the CORS handler to wrap your router
    handler := c.Handler(r)

    http.Handle("/", handler)
    http.ListenAndServe(":8080", nil)
}
