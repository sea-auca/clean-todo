package user_test

import (
	"testing"
	"time"

	"github.com/sea-auca/clean-todo/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		desc     string
		email    string
		password string
		fullname string
		err      error
	}{
		{
			desc:     "Create user with valid valid email and password",
			email:    "valid@mail.com",
			fullname: "User user",
			password: "Password_1",
			err:      nil,
		},
		{
			desc:     "Invalid email",
			email:    "invalid",
			password: "Password_1",
			err:      user.ErrInvalidEmail,
		},
		{
			desc:     "Invalid password -- too short",
			email:    "valid@mail.com",
			password: "pass",
			err:      user.ErrInvalidPassword,
		},
		{
			desc:     "Invalid password -- no capital letter",
			email:    "valid@mail.com",
			password: "password",
			err:      user.ErrInvalidPassword,
		},
		{
			desc:     "Invalid password -- no numerical",
			email:    "valid@mail.com",
			password: "Password",
			err:      user.ErrInvalidPassword,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			us, err := user.NewUser(tC.email, tC.fullname, tC.password)
			assert.Equalf(t, tC.err, err, "Expected error to be %v, got %v", tC.err, err)               // correct error message
			assert.Equalf(t, tC.email, us.Email, "Expected email to be %v, got %v", tC.email, us.Email) //email is set
			assert.NotEmptyf(t, us.Hash, "Expected hash to be generated")                               //hash is not empty
			assert.NotEmptyf(t, us.CreatedAt, "Expected creation date to be set")                       //created at is set correctly
		})
	}
}

func TestIsValidUser(t *testing.T) {
	validUser := user.User{
		Email:      "valid@mail.com",
		Hash:       "$2a$12$diD46btFldzLxq7IQX5znuMAOOpDDj8iVJM0stBE0uYCE/9c/0qhK",
		ID:         1,
		Fullname:   "User user",
		CreatedAt:  time.Now(),
		VerifiedAt: time.Time{},
		UpdatedAt:  time.Now(),
	}

	invalidEmailUser := validUser
	invalidEmailUser.Email = "valid"

	invalidHashUser := validUser
	invalidHashUser.Hash = ""

	invalidFullnameUser := validUser
	invalidFullnameUser.Fullname = ""

	invalidCreadtedAtUser := validUser
	invalidCreadtedAtUser.CreatedAt = time.Time{}

	invalidMalformedUpdateTime := validUser
	invalidMalformedUpdateTime.UpdatedAt = time.Time{}

	testCases := []struct {
		desc   string
		user   user.User
		result bool
	}{
		{
			desc:   "valid user",
			user:   validUser,
			result: true,
		},
		{
			desc:   "user with invalid hash",
			user:   invalidHashUser,
			result: false,
		},
		{
			desc:   "user with invalid email",
			user:   invalidEmailUser,
			result: false,
		},
		{
			desc:   "user with invalid fullname",
			user:   invalidFullnameUser,
			result: false,
		},
		{
			desc:   "user with invalid createdAt",
			user:   invalidCreadtedAtUser,
			result: false,
		},
		{
			desc:   "malformed UpdatedAt field",
			user:   invalidMalformedUpdateTime,
			result: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := user.IsValidUser(&tC.user)
			assert.Equalf(t, tC.result, res, "Expected the result be %v, got %v, for user: %v", tC.result, res, tC.user)
		})
	}
}
