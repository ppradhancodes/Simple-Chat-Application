package models

import (
    "time"

    "github.com/google/uuid"
)

type User struct {
    ID       uuid.UUID `json:"id"`
    Username string    `json:"username"`
}

type Message struct {
    ID         uuid.UUID `json:"id"`
    SenderID   uuid.UUID `json:"sender_id"`
    ReceiverID uuid.UUID `json:"receiver_id"`
    Content    string    `json:"content"`
    Timestamp  time.Time `json:"timestamp"`
}

func NewUser(username string) User {
    return User{
        ID:       uuid.New(),
        Username: username,
    }
}

func NewMessage(senderID, receiverID uuid.UUID, content string) Message {
    return Message{
        ID:         uuid.New(),
        SenderID:   senderID,
        ReceiverID: receiverID,
        Content:    content,
        Timestamp:  time.Now(),
    }
} 