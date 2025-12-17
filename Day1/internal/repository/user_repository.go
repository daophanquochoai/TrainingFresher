package repository

import (
	"errors"
	"go-rest-api/internal/model"
	"sync"
)

var ErrNotFound = errors.New("user not found")

type UserRepository interface {
	Create(u *model.User) (*model.User, error)
	GetByID(id string) (*model.User, error)
}

type InMemoryUserRepository struct {
	mu   sync.Mutex
	data map[string]*model.User
}

// NewInMemoryUserRepository initializes a new in-memory user repository.
func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		data: make(map[string]*model.User),
	}
}

// Create adds a new user to the repository.
func (r *InMemoryUserRepository) Create(u *model.User) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[u.ID] = u
	return u, nil
}

// GetByID retrieves a user by their ID.
func (r *InMemoryUserRepository) GetByID(id string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, exists := r.data[id]
	if !exists {
		return nil, ErrNotFound
	}
	return u, nil
}
