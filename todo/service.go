package todo

import (
	"context"
	"fmt"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	// TODO: check for userID
	newTodo, err := s.repo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

func (s *service) ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error) {
	if userID < 0 {
		return nil, ErrNegativeUserID
	}
	if userID == 0 {
		return nil, ErrNullUserID
	}

	return nil, nil
}

func (s *service) SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error) {
	if userID < 0 {
		return nil, ErrNegativeUserID
	}
	if userID == 0 {
		return nil, ErrNullUserID
	}

	// we check for less than -1
	// because -1 stands for negating
	// limit or offset
	if offset < -1 || limit < -1 {
		return nil, fmt.Errorf("limit and offset cannot be less than -1")
	}
	todos, err := s.repo.SearchByText(ctx, text, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *service) Update(ctx context.Context, todo *Todo) (*Todo, error) {
	if todo.ID < 0 {
		return nil, ErrNegativeID
	}
	if todo.ID == 0 {
		return nil, ErrNullID
	}
	newtodo, err := s.repo.Update(ctx, todo)
	if err != nil {
		return nil, err
	}
	return newtodo, nil
}

func (s *service) Delete(ctx context.Context, todo *Todo) error {
	if todo.ID < 0 {
		return ErrNegativeID
	}
	if todo.ID == 0 {
		return ErrNullID
	}
	return s.repo.Delete(ctx, todo)
}
