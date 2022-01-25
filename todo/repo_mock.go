package todo

import "context"

type MockRepository struct {
}

func NewMockRepository() Repository {
	return &MockRepository{}
}

func (s *MockRepository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	return nil, nil
}

func (s *MockRepository) ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error) {
	return nil, nil
}

func (s *MockRepository) SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error) {
	return nil, nil
}

func (s *MockRepository) Update(ctx context.Context, todo *Todo) (*Todo, error) {
	return nil, nil
}

func (s *MockRepository) Delete(ctx context.Context, todo *Todo) error {
	return nil
}
