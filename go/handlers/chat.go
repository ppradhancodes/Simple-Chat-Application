package handlers

import (
    "chat-app/models"
    "fmt"

    "github.com/google/uuid"
)

type ChatHandler struct {
    storage *Storage
}

func NewChatHandler() *ChatHandler {
    return &ChatHandler{
        storage: NewStorage(),
    }
}

func (h *ChatHandler) RegisterOrLogin(username string) (models.User, error) {
    // First check if user exists
    if existingUser, exists := h.storage.GetUserByUsername(username); exists {
        // User exists, return the existing user (login)
        return existingUser, nil
    }

    // User doesn't exist, create new user (register)
    user := models.NewUser(username)
    err := h.storage.AddUser(user)
    if err != nil {
        return models.User{}, err
    }
    return user, nil
}

func (h *ChatHandler) SendMessage(senderID uuid.UUID, receiverUsername, content string) error {
    receiver, exists := h.storage.GetUserByUsername(receiverUsername)
    if !exists {
        return fmt.Errorf("receiver not found")
    }

    message := models.NewMessage(senderID, receiver.ID, content)
    h.storage.AddMessage(message)
    return nil
}

func (h *ChatHandler) GetMessages(userID uuid.UUID) []models.Message {
    return h.storage.GetMessagesForUser(userID)
}

func (h *ChatHandler) SearchMessages(keyword string) []models.Message {
    return h.storage.SearchMessages(keyword)
}

func (h *ChatHandler) ListUsers() []models.User {
    return h.storage.ListUsers()
}

func (h *ChatHandler) GetUser(id uuid.UUID) (models.User, bool) {
    return h.storage.GetUser(id)
} 