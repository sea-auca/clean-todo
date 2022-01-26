package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sea-auca/clean-todo/user"
	"github.com/sea-auca/clean-todo/user/mocks"
	"github.com/stretchr/testify/assert"
)

func setupService(t *testing.T) (m *mocks.MockRepository, us user.Service) {
	ctrl := gomock.NewController(t)
	m = mocks.NewMockRepository(ctrl)
	us = user.NewUserService(m)
	return
}

var bg = context.Background()

func setupUser(us user.Service, m *mocks.MockRepository) (usr *user.User) {
	m.EXPECT().Create(bg, gomock.Any()).Times(1).DoAndReturn(func(ctx context.Context, user *user.User) (*user.User, error) {
		user.ID = 1
		return user, nil
	})
	usr, _ = us.Register(context.Background(), "valid@mail.com", "Password_1", "User User")
	usr.ID = 1
	return
}

func TestCreateUserValid(t *testing.T) {
	m, us := setupService(t)
	m.EXPECT().Create(bg, gomock.Any()).Times(1).DoAndReturn(func(ctx context.Context, user *user.User) (*user.User, error) {
		user.ID = 1
		return user, nil
	})

	usr, err := us.Register(bg, "valid@mail.com", "Password_1", "User User")

	assert.Nilf(t, err, "Expected error to be nil")
	assert.Truef(t, user.IsValidUser(usr), "Expected to create a valid user")
	assert.NotEqualf(t, 0, usr.ID, "Expected id to be non-zero")
}

func TestCreateUserInValid(t *testing.T) {
	m, us := setupService(t)
	m.EXPECT().Create(bg, gomock.Any()).Times(0)

	usr, err := us.Register(context.Background(), "valid", "Password_1", "User User")

	assert.NotNilf(t, err, "Expected error to not be nil")
	assert.Falsef(t, user.IsValidUser(usr), "Expected to create a invalid user from invalid data")
}

func TestUpdateUserValid(t *testing.T) {
	m, us := setupService(t)

	usr := setupUser(us, m)
	params := make(map[string]interface{})

	m.EXPECT().Update(bg, gomock.AssignableToTypeOf(usr), gomock.AssignableToTypeOf(params)).Times(1)
	params["Email"] = "new_valid_email@mail.com"
	params["Fullname"] = "New eligible fullname"
	m.EXPECT().GetByID(bg, gomock.Any()).MinTimes(0).MaxTimes(1).DoAndReturn(func(ctx context.Context, id int64) (*user.User, error) {
		u := usr
		u.Email = params["Email"].(string)
		u.Fullname = params["Fullname"].(string)
		return u, nil
	})

	user, err := us.Update(context.Background(), usr, params)

	assert.Nilf(t, err, "Expected error to be nil")
	assert.Equalf(t, params["Email"], user.Email, "expected email to be %v, got %v", params["Email"], user.Email)
	assert.Equalf(t, params["Fullname"], user.Fullname, "expected fullname to be %v, got %v", params["Fullname"], user.Fullname)
}

func TestUpdateUserInvalidEmail(t *testing.T) {
	m, us := setupService(t)

	user := setupUser(us, m)
	params := make(map[string]interface{})

	m.EXPECT().Update(bg, gomock.AssignableToTypeOf(user), gomock.AssignableToTypeOf(params)).Times(1).Return(errors.New("new error"))
	m.EXPECT().GetByID(bg, gomock.Any()).MinTimes(0).MaxTimes(1)

	params["Email"] = "new_invalid_email@mail.c"
	params["Fullname"] = "New eligible fullname"

	new_user, err := us.Update(context.Background(), user, params)

	assert.NotNilf(t, err, "Expected error to not be nil")
	assert.Equalf(t, new_user, user, "Expected user to stay same. Expected %v, got %v", user, new_user)
}

func TestUserGetByID(t *testing.T) {
	m, us := setupService(t)

	usr := setupUser(us, m)

	m.EXPECT().GetByID(gomock.AssignableToTypeOf(bg), gomock.Any()).Times(1).DoAndReturn(func(ctx context.Context, id int64) (*user.User, error) {
		return usr, nil
	})

	new_user, err := us.GetByID(bg, usr.ID)
	assert.Nilf(t, err, "Expected error to be nil")
	assert.Equalf(t, new_user, usr, "Expected users to be same")
}

func TestUserEmail(t *testing.T) {
	m, us := setupService(t)

	usr := setupUser(us, m)

	m.EXPECT().GetByEmail(gomock.AssignableToTypeOf(bg), gomock.Any()).Times(1).DoAndReturn(func(ctx context.Context, _ string) (*user.User, error) {
		return usr, nil
	})

	new_user, err := us.GetByEmail(bg, usr.Email)
	assert.Nilf(t, err, "Expected error to be nil")
	assert.Equalf(t, new_user, usr, "Expected users to be same")
}
