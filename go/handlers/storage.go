package handlers

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"

	"chat-app/models"
)

type IStorage interface {
	// AddUser will add the user to the system
	AddUser(user models.User) error
	// GetUser will get the user inforamtion based on the id
	GetUser(id uuid.UUID) (models.User, bool)
	// GetUser will get the user inforamtion based on the username
	GetUserByUsername(username string) (models.User, bool)
	// AddMessage will add the message to the storage
	AddMessage(message models.Message)
	// GetMessagesForUser will get all the messages that the current user has
	GetMessagesForUser(userID uuid.UUID) []models.Message
	// SearchMessages will search the messages that the current user has based on keyword
	SearchMessages(keyword string) []models.Message
	// DeleteMessage will delete the messages that the current user has based on keyword
	DeleteMessage(userID uuid.UUID, keyword string) bool
	// ListUsers will list all the users that are registered to the system
	ListUsers() []models.User
}
type Storage struct {
	users    map[uuid.UUID]models.User
	messages []models.Message
}

func NewStorage() IStorage {
	return &Storage{
		users:    make(map[uuid.UUID]models.User),
		messages: make([]models.Message, 0),
	}
}

func (s *Storage) AddUser(user models.User) error {
	for _, existingUser := range s.users {
		if existingUser.Username == user.Username {
			return fmt.Errorf("username already exists")
		}
	}

	s.users[user.ID] = user
	return nil
}

func (s *Storage) GetUser(id uuid.UUID) (models.User, bool) {
	user, exists := s.users[id]
	return user, exists
}

func (s *Storage) GetUserByUsername(username string) (models.User, bool) {
	for _, user := range s.users {
		if user.Username == username {
			return user, true
		}
	}
	return models.User{}, false
}

func (s *Storage) AddMessage(message models.Message) {
	s.messages = append(s.messages, message)
}

func (s *Storage) GetMessagesForUser(userID uuid.UUID) []models.Message {
	var userMessages []models.Message
	for _, msg := range s.messages {
		if msg.SenderID == userID || msg.ReceiverID == userID {
			userMessages = append(userMessages, msg)
		}
	}
	return userMessages
}

func (s *Storage) SearchMessages(keyword string, userID uuid.UUID) []models.Message {
	var matchingMessages []models.Message
	keyword = strings.ToLower(keyword)
	for _, msg := range s.messages {
		if msg.SenderID != userID && msg.ReceiverID != userID {
			continue
		}
		if strings.Contains(strings.ToLower(msg.Content), keyword) {
			matchingMessages = append(matchingMessages, msg)
		}
	}
	return matchingMessages
}

func (s *Storage) DeleteMessage(userID uuid.UUID, keyword string) bool {
	var nonDeleteMessage []models.Message
	keyword = strings.ToLower(keyword)
	for _, msg := range s.messages {
		if msg.SenderID != userID && msg.ReceiverID != userID || !strings.Contains(strings.ToLower(msg.Content), keyword) {
			nonDeleteMessage = append(nonDeleteMessage, msg)
		}
	}

	isDelete := len(nonDeleteMessage) != len(s.messages)
	s.messages = nonDeleteMessage
	return isDelete

}

func (s *Storage) ListUsers() []models.User {
	return maps.Values(s.users)
}
