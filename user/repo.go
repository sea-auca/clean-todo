package user

import (
	"context"
	"net/mail"

	"github.com/go-rel/changeset"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type Repository interface {
	Create(ctx context.Context, us *User) (*User, error)
	Update(ctx context.Context, us *User, params map[string]interface{}) error
	Delete(ctx context.Context, us *User) error

	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type repo struct {
	r rel.Repository
}

func CreateNewRepo(r rel.Repository) Repository {
	return repo{r}
}

func (r repo) Create(ctx context.Context, us *User) (*User, error) {
	err := r.r.Insert(ctx, us)
	if err != nil {
		return us, err
	}
	err = r.r.Find(ctx, us, where.Eq("email", us.Email))
	return us, err
}

func (r repo) Changeset(us *User, pms map[string]interface{}) *changeset.Changeset {
	ch := changeset.Change(*us, pms)
	changeset.ValidateRequired(ch, []string{"email", "fullname"})
	changeset.ValidateMin(ch, "fullname", 4)
	if _, err := mail.ParseAddress(pms["email"].(string)); err != nil {
		changeset.AddError(ch, "email", "Invalid address format")
	}
	return ch
}

func (r repo) Update(ctx context.Context, us *User, params map[string]interface{}) error {
	ch := r.Changeset(us, params)
	if err := ch.Error(); err != nil {
		return err
	}
	return r.r.Update(ctx, us, ch)
}

func (r repo) Delete(ctx context.Context, us *User) error {
	return r.r.Delete(ctx, us)
}

func (r repo) GetByID(ctx context.Context, id int64) (*User, error) {
	var us User
	err := r.r.Find(ctx, &us, where.Eq("id", id))
	return &us, err
}

func (r repo) GetByEmail(ctx context.Context, email string) (*User, error) {
	var us User
	err := r.r.Find(ctx, &us, where.Eq("email", email))
	return &us, err
}
