package user

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	//CRUD
	Register(ctx context.Context, email, password, fullname string) (*User, error)
	Update(ctx context.Context, user *User, params map[string]interface{}) (*User, error)
	Delete(ctx context.Context, user *User) error

	//Getteres
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)

	Authenticate(ctx context.Context, email, password string) (string, error)
	Authorize(ctx context.Context, token string) (*User, error)
}

var (
	ErrUserNotFoundEmail  = errors.New("user with such email not found")
	ErrUserNotFoundID     = errors.New("user with such ID not found")
	ErrInvalidCredentials = errors.New("given credentials are not valid")
	ErrInvalidToken       = errors.New("supplied authorization token is invalid")
	ErrSigningFailed      = errors.New("signing jwt token failed")
	ErrTokenInvalidated   = errors.New("this token is not valid anymore - expired or revoked")
)

type userService struct {
	repo        Repository
	private_key *rsa.PrivateKey
}

func NewUserService(repo Repository, private_k *rsa.PrivateKey) Service {
	return userService{repo, private_k}
}

func (u userService) Register(ctx context.Context, email, password, fullname string) (*User, error) {
	user, err := NewUser(email, fullname, password)
	if err != nil {
		return user, err
	}

	return u.repo.Create(ctx, user)
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

type AutheticationClaims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}

func newClaims(id int64) AutheticationClaims {
	return AutheticationClaims{
		id,
		jwt.StandardClaims{
			Issuer:    "clean-todo",
			Audience:  "users",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}
}

func (u userService) Authenticate(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", ErrUserNotFoundEmail
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}
	claims := newClaims(user.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	signed, err := token.SignedString(u.private_key)
	if err != nil {
		return "", ErrSigningFailed
	}
	return signed, nil
}

func (u userService) Authorize(ctx context.Context, token string) (*User, error) {
	tok, err := jwt.ParseWithClaims(token, &AutheticationClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return &u.private_key.PublicKey, nil
	})
	if err != nil {
		log.Println(err)
		return nil, ErrInvalidToken
	}
	claims, ok := tok.Claims.(*AutheticationClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	if err := claims.Valid(); err != nil {
		return nil, ErrTokenInvalidated
	}
	user, err := u.repo.GetByID(ctx, claims.ID)
	if err != nil {
		return nil, ErrUserNotFoundID
	}
	return user, nil
}
