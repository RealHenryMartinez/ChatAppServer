package database

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RetrieveChatMessagesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	chatRoom := params["room_number"]
	data, err := RetrieveChatMessagesHelper(chatRoom)

	fmt.Println("data: ", data)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the data as JSON and write it to the response
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RetrieveChatMessages retrieves chat messages for a specific chat room and returns them as a string.
func RetrieveChatMessagesHelper(chatRoom string) (ChatRoomLog, error) {

	// Access the desired database and collection.

	result, err := GetByFilter[ChatRoomLog](ChatCollection, bson.M{"room_number": string(chatRoom)})

	if err != nil {
		fmt.Println("Room not found: " + chatRoom)
	}
	// Create a string containing the chat messages.
	// var chatLog string
	// for _, message := range result.Messages {
	// 	chatLog += message.Message + "\n"
	// }

	return result, nil
}

func DeleteChatMessageByMessage(room uint64, contentToRemove string) (*mongo.UpdateResult, error) {
    filter := bson.M{"room_number": room}

    // Create an update with the $pull operator
    update := bson.M{
        "$pull": bson.M{
            "messages": bson.M{"message": contentToRemove},
        },
    }

    fmt.Println("Deleting Data: ", update)

    result, err := DeleteOne(ChatCollection, filter, update)

	fmt.Println("RESULT: ", result, "\nERROR: ", err)
    if err != nil {

        return nil, err
    }
    return result, nil
}