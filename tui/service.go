package main

import (
	"errors"
	"sort"
)

// TodoService handles business logic for todos
type TodoService struct {
	store Store
}

// NewTodoService creates a new todo service
func NewTodoService(store Store) *TodoService {
	return &TodoService{
		store: store,
	}
}

// CreateTodo creates a new todo with validation
func (s *TodoService) CreateTodo(title string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	return s.store.Create(title)
}

// ListTodos returns all todos sorted by creation date (newest first)
func (s *TodoService) ListTodos() ([]*Todo, error) {
	todos, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}

	// Sort by ID (which correlates with creation time)
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].ID > todos[j].ID
	})

	return todos, nil
}

// ToggleTodo toggles the completion status of a todo
func (s *TodoService) ToggleTodo(id int) (*Todo, error) {
	return s.store.ToggleComplete(id)
}

// DeleteTodo removes a todo
func (s *TodoService) DeleteTodo(id int) error {
	return s.store.Delete(id)
}

// GetStats returns statistics about todos
func (s *TodoService) GetStats() (total, completed, pending int, err error) {
	todos, err := s.store.GetAll()
	if err != nil {
		return 0, 0, 0, err
	}

	total = len(todos)
	for _, todo := range todos {
		if todo.Completed {
			completed++
		}
	}
	pending = total - completed

	return total, completed, pending, nil
}
