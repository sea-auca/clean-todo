package todo

import (
	"errors"
	"time"
)

// Errors for package todo
var (
	ErrNegativeID = errors.New("todo.ID of int64 is negative")
	ErrNullID     = errors.New("todo.ID of int64 is 0")
	ErrUpdatedAt  = errors.New("todo.UpdatedAt of int64 is less than Todo.CreatedAt")

	ErrNegativeUserID = errors.New("userID is negative")
	ErrNullUserID     = errors.New("userID is 0")

	ErrEmptyName        = errors.New("todo.name is empty")
	ErrEmptyDescription = errors.New("todo.description is empty")

	ErrIncorrectUpdatedAt = errors.New("updatedAt is before createdAt")
	ErrIncorrectDueTo     = errors.New("dueTo is before createdAt")
)

type Todo struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DueTo       time.Time `json:"DueTo"`
	IsDone      bool      `json:"IsDone"`
	// TODO: Todo needs a field for User
	// User
}

func NewTodo() Todo {
	return Todo{}
}

func (t *Todo) IsValid() error {
	if t.UpdatedAt.Before(t.CreatedAt) {
		return ErrIncorrectUpdatedAt
	}
	if t.Name == "" {
		return ErrEmptyName
	}
	if t.DueTo.Before(t.CreatedAt) {
		return ErrIncorrectDueTo
	}
	return nil
}
