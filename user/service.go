package user

import (
	"context"
)

type Service interface {
	//CRUD
	Register(ctx context.Context, email, password, fullname string) (*User, error)
	Update(ctx context.Context, user *User, params map[string]interface{}) (*User, error)
	Delete(ctx context.Context, user *User) error

	//Getteres
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type userService struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return userService{repo}
}

func (u userService) Register(ctx context.Context, email, password, fullname string) (*User, error) {
	user, err := NewUser(email, fullname, password)
	if err != nil {
		return user, err
	}

	user, err = u.repo.Create(ctx, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u userService) Update(ctx context.Context, user *User, params map[string]interface{}) (*User, error) {
	err := u.repo.Update(ctx, user, params)
	if err != nil {
		return user, err
	}
	user, _ = u.GetByID(ctx, user.ID)
	return user, nil
}

func (u userService) Delete(ctx context.Context, user *User) error {
	return u.repo.Delete(ctx, user)
}

func (u userService) GetByID(ctx context.Context, id int64) (*User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u userService) GetByEmail(ctx context.Context, email string) (*User, error) {
	return u.repo.GetByEmail(ctx, email)
}
