package database

type ChatDocument struct {
    Message string `bson:"message"`
}

type ChatRoomLog struct {
    Messages   []ChatDocument `bson:"messages"`
    RoomNumber string        `bson:"room_number"`
}
