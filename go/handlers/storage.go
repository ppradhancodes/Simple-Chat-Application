package handlers

import (
    "chat-app/models"
    "fmt"
    "strings"
    "sync"

    "github.com/google/uuid"
)

type Storage struct {
    users    map[uuid.UUID]models.User
    messages []models.Message
    mu       sync.RWMutex
}

func NewStorage() *Storage {
    return &Storage{
        users:    make(map[uuid.UUID]models.User),
        messages: make([]models.Message, 0),
    }
}

func (s *Storage) AddUser(user models.User) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    for _, existingUser := range s.users {
        if existingUser.Username == user.Username {
            return fmt.Errorf("username already exists")
        }
    }

    s.users[user.ID] = user
    return nil
}

func (s *Storage) GetUser(id uuid.UUID) (models.User, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    user, exists := s.users[id]
    return user, exists
}

func (s *Storage) GetUserByUsername(username string) (models.User, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    for _, user := range s.users {
        if user.Username == username {
            return user, true
        }
    }
    return models.User{}, false
}

func (s *Storage) AddMessage(message models.Message) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.messages = append(s.messages, message)
}

func (s *Storage) GetMessagesForUser(userID uuid.UUID) []models.Message {
    s.mu.RLock()
    defer s.mu.RUnlock()

    var userMessages []models.Message
    for _, msg := range s.messages {
        if msg.SenderID == userID || msg.ReceiverID == userID {
            userMessages = append(userMessages, msg)
        }
    }
    return userMessages
}

func (s *Storage) SearchMessages(keyword string) []models.Message {
    s.mu.RLock()
    defer s.mu.RUnlock()

    var matchingMessages []models.Message
    keyword = strings.ToLower(keyword)
    for _, msg := range s.messages {
        if strings.Contains(strings.ToLower(msg.Content), keyword) {
            matchingMessages = append(matchingMessages, msg)
        }
    }
    return matchingMessages
}

func (s *Storage) ListUsers() []models.User {
    s.mu.RLock()
    defer s.mu.RUnlock()

    users := make([]models.User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }
    return users
} 