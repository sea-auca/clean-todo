package user_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sea-auca/clean-todo/user"
	"github.com/sea-auca/clean-todo/user/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserValid(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mocks.NewMockRepository(ctrl)
	m.EXPECT().Create(gomock.Any()).Times(1)

	us := user.NewUserService(m)
	usr, err := us.Register(context.Background(), "valid@mail.com", "Password_1", "User User")

	assert.Nilf(t, err, "Expected error to be nil")
	assert.Truef(t, user.IsValidUser(usr), "Expected to create a valid user")
}

func TestCreateUserInValid(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mocks.NewMockRepository(ctrl)
	m.EXPECT().Create(gomock.Any()).Times(0)

	us := user.NewUserService(m)
	usr, err := us.Register(context.Background(), "valid", "Password_1", "User User")

	assert.NotNilf(t, err, "Expected error to not be nil")
	assert.Falsef(t, user.IsValidUser(usr), "Expected to create a invalid user from invalid data")
}
