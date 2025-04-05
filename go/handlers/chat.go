package handlers

import (
	"fmt"

	"github.com/google/uuid"

	"chat-app/models"
)

type IChatHandler interface {
	// RegisterOrLogin will check if the input username exists or not. If exists, return the user
	// Otherwise, allow the user to register and if the registered username is there, return error
	RegisterOrLogin(username string) (models.User, error)
	// SendMessage will send the message to the user with the message
	SendMessage(senderID uuid.UUID, receiverUsername, content string) error
	// GetMessages will get all the messages that the current user has
	GetMessages(userID uuid.UUID) []models.Message
	// SearchMessages will search the messages that the current user has based on keyword
	SearchMessages(keyword string) []models.Message
	// DeleteMessage will delete the messages that the current user has based on keyword
	DeleteMessage(userID uuid.UUID, keyword string) bool
	// ListUsers will list all the users that are registered to the system
	ListUsers() []models.User
	// GetUser will get the user inforamtion based on the id
	GetUser(id uuid.UUID) (models.User, bool)
}

type ChatHandler struct {
	storage IStorage
}

func NewChatHandler() IChatHandler {
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

func (h *ChatHandler) DeleteMessage(userID uuid.UUID, keyword string) bool {
	return h.storage.DeleteMessage(userID, keyword)
}

func (h *ChatHandler) ListUsers() []models.User {
	return h.storage.ListUsers()
}

func (h *ChatHandler) GetUser(id uuid.UUID) (models.User, bool) {
	return h.storage.GetUser(id)
}
