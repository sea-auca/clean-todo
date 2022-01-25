package todo

import (
	"context"
	"errors"
	"time"
)

// Errors for package todo
var (
	ErrNegativeID = errors.New("Todo.ID of int64 is negative")
	ErrNullID     = errors.New("Todo.ID of int64 is 0")
	ErrUpdatedAt  = errors.New("Todo.UpdatedAt of int64 is less than Todo.CreatedAt")

	ErrNegativeUserID = errors.New("UserID is negative")
	ErrNullUserID     = errors.New("UserID is 0")
)

type Todo struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	IsDueTo     time.Time `json:"IsDueTo"`
	// TODO: Todo needs a field for User
	// User
}

type Repository interface {
	Create(ctx context.Context, todo *Todo) (*Todo, error)
	ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error)
	SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error)
	Update(ctx context.Context, todo *Todo) (*Todo, error)
	Delete(ctx context.Context, todo *Todo) error
}

type Service interface {
	Create(ctx context.Context, todo *Todo) (*Todo, error)
	ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error)
	SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error)
	Update(ctx context.Context, todo *Todo) (*Todo, error)
	Delete(ctx context.Context, todo *Todo) error
}

func NewTodo() Todo {
	return Todo{}
}

func (t *Todo) isValid() bool {
	return true
}
