package user

import "context"

type Repository interface {
	Create(ctx context.Context, us *User) (*User, error)
	Update(ctx context.Context, us *User, params map[string]interface{}) error
	Delete(ctx context.Context, us *User) error

	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
