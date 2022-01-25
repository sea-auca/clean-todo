package user

import (
	"context"
	"errors"
)

type Service interface {
	//CRUD
	Register(ctx context.Context, email, password, fullname string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error

	//Getteres
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
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

	if err := u.repo.Create(user); err != nil {
		return user, err
	}
	return user, nil
	return nil, errors.New("not implemented")
}

func (u userService) Update(ctx context.Context, user *User) error {
	return errors.New("not implemneted")
}

func (u userService) Delete(ctx context.Context, user *User) error {
	return errors.New("not implemneted")
}

func (u userService) GetByID(id int64) (*User, error) {
	return nil, errors.New("not implemneted")
}

func (u userService) GetByEmail(email string) (*User, error) {
	return nil, errors.New("not implemneted")
}
