package main

import (
	"sync"
	"time"
)

// Todo represents a single todo item
type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}

// Store defines the interface for todo storage
type Store interface {
	Create(title string) (*Todo, error)
	GetAll() ([]*Todo, error)
	GetByID(id int) (*Todo, error)
	Update(id int, title string, completed bool) (*Todo, error)
	Delete(id int) error
	ToggleComplete(id int) (*Todo, error)
}

// MemoryStore is an in-memory implementation of the Store interface
type MemoryStore struct {
	mu     sync.RWMutex
	todos  map[int]*Todo
	nextID int
}

// NewMemoryStore creates a new in-memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		todos:  make(map[int]*Todo),
		nextID: 1,
	}
}

// Create adds a new todo to the store
func (s *MemoryStore) Create(title string) (*Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := &Todo{
		ID:        s.nextID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	s.todos[s.nextID] = todo
	s.nextID++

	return todo, nil
}

// GetAll returns all todos
func (s *MemoryStore) GetAll() ([]*Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}

	return todos, nil
}

// GetByID returns a specific todo
func (s *MemoryStore) GetByID(id int) (*Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, nil
	}

	return todo, nil
}

// Update modifies an existing todo
func (s *MemoryStore) Update(id int, title string, completed bool) (*Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, nil
	}

	todo.Title = title
	todo.Completed = completed

	return todo, nil
}

// Delete removes a todo from the store
func (s *MemoryStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.todos, id)
	return nil
}

// ToggleComplete toggles the completed status of a todo
func (s *MemoryStore) ToggleComplete(id int) (*Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, nil
	}

	todo.Completed = !todo.Completed
	return todo, nil
}
