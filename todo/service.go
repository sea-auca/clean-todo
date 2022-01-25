package todo

import (
	"context"
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
	// newTodo, err := s.repo.Create(ctx, todo)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not create a todo due to error: %v", err)
	// }
	return nil, nil
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

func (s *service) SearchByText(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error) {
	if userID < 0 {
		return nil, ErrNegativeUserID
	}
	if userID == 0 {
		return nil, ErrNullUserID
	}
	return nil, nil
}

func (s *service) Update(ctx context.Context, todo *Todo) (*Todo, error) {
	return nil, nil
}

func (s *service) Delete(ctx context.Context, todo *Todo) (*Todo, error) {
	return nil, nil
}
