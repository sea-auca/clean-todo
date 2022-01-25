package user

import (
	"errors"
	"net/mail"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//General user structure.
//Verified at is empty initially, as this field is set upon verification
type User struct {
	Email      string
	Fullname   string
	Hash       string
	ID         int64
	VerifiedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

var ErrInvalidEmail = errors.New("invalid email format")
var ErrInvalidPassword = errors.New("invalid password format")

func NewUser(email, fullname, password string) (*User, error) {
	user := User{Email: email, Fullname: fullname}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return &user, err
	}

	user.Hash = string(hash)
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	if err := checkEmail(email); err != nil {
		return &user, err
	}

	if err := checkPassword(password); err != nil {
		return &user, err
	}

	return &user, nil
}

func IsValidUser(u *User) bool {
	if err := checkEmail(u.Email); err != nil {
		return false
	}

	if len(u.Hash) != 60 {
		return false
	}

	if u.Fullname == "" {
		return false
	}

	if u.CreatedAt.IsZero() {
		return false
	}

	if u.CreatedAt.After(u.UpdatedAt) {
		return false
	}

	return true
}

func checkPassword(password string) error {
	if len(password) < 8 {
		return ErrInvalidPassword
	}

	if strings.ToLower(password) == password {
		return ErrInvalidPassword
	}

	if !strings.ContainsAny(password, "1234567890") {
		return ErrInvalidPassword
	}
	return nil
}

func checkEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil
}
