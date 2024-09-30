package store

import (
	"errors"
	"sync"

	"github.com/desmomndsanctity/twilio-go-verify/internal/models"
)

type InMemoryStore struct {
	users map[string]*models.User
	mutex sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		users: make(map[string]*models.User),
	}
}

func (s *InMemoryStore) CreateUser(user *models.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.users[user.Email]; exists {
		return errors.New("user already exists")
	}

	s.users[user.Email] = user
	return nil
}

func (s *InMemoryStore) GetUserByEmail(email string) (*models.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *InMemoryStore) UpdateUser(user *models.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.users[user.Email]; !exists {
		return errors.New("user not found")
	}

	s.users[user.Email] = user
	return nil
}
