package todo

import (
	"context"
	"errors"
	"time"

	"github.com/sea-auca/clean-todo/user"
)

// error bellow are specific only to
// the methods in this file
var (
	ErrNegativeLimit  = errors.New("limit is less than -1")
	ErrNegativeOffset = errors.New("offset is less than -1")
	ErrEmptyText      = errors.New("text is empty")
)

type Service interface {
	Create(ctx context.Context, name, description string, dueTo time.Time, author user.User) (*Todo, error)
	Update(ctx context.Context, todo *Todo) (*Todo, error)
	Delete(ctx context.Context, todo *Todo) error

	ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error)
	SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, name, description string, dueTo time.Time, author user.User) (*Todo, error) {
	todo := NewTodo(name, description, dueTo, author)
	if err := todo.IsValid(); err != nil {
		return nil, err
	}
	if todo.Author == nil {
		return nil, ErrInvalidAuthor
	}
	if !user.IsValidUser(todo.Author) {
		return nil, ErrInvalidAuthor
	}
	return s.repo.Create(ctx, &todo)
}

func (s *service) ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error) {
	if userID < 1 {
		return nil, ErrNegativeUserID
	}
	if limit < 0 {
		return nil, ErrNegativeLimit
	}
	if offset < 0 {
		return nil, ErrNegativeOffset
	}
	return s.repo.ListByUserID(ctx, userID, limit, offset)
}

func (s *service) SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error) {
	if text == "" {
		return nil, ErrEmptyText
	}
	if userID < 1 {
		return nil, ErrNegativeUserID
	}
	if limit < 0 {
		return nil, ErrNegativeLimit
	}
	if offset < 0 {
		return nil, ErrNegativeOffset
	}
	return s.repo.SearchByText(ctx, text, userID, limit, offset)
}

func (s *service) Update(ctx context.Context, todo *Todo) (*Todo, error) {
	if todo.ID < 1 {
		return nil, ErrNegativeID
	}
	if todo.Name == "" {
		return nil, ErrEmptyName
	}
	if todo.Description == "" {
		return nil, ErrEmptyDescription
	}
	return s.repo.Update(ctx, todo)
}

func (s *service) Delete(ctx context.Context, todo *Todo) error {
	if todo.ID < 1 {
		return ErrNegativeID
	}
	return s.repo.Delete(ctx, todo)
}
