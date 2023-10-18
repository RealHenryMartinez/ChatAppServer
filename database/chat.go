package database

type Request struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

// Database types
type ChatDocument struct {
	RoomNumber uint32 `json:"room_number"`
	Message    string `bson:"message"`
	Uid        string `bson:"uid"`
}

type ChatRoomLog struct {
	Messages   []ChatDocument `bson:"messages"`
	RoomNumber uint32         `bson:"room_number"`
}
