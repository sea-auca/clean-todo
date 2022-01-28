package todo

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type Repository interface {
	Create(ctx context.Context, todo *Todo) (*Todo, error)
	ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error)
	SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error)
	Update(ctx context.Context, todo *Todo) (*Todo, error)
	Delete(ctx context.Context, todo *Todo) error
}
type repository struct {
	r rel.Repository
}

func CreateNewRepo(r rel.Repository) Repository {
	return &repository{r}
}

func (r *repository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	err := r.r.Insert(ctx, todo)
	return todo, err
}

func (r *repository) ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error) {
	var todos []*Todo
	err := r.r.Find(
		ctx,
		&todos,
		where.Eq("user_id", userID),
		rel.Select().Limit(limit).Offset(offset),
	)
	return todos, err
}

func (r *repository) SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error) {
	var todos []*Todo
	err := r.r.FindAll(
		ctx,
		&todos,
		where.Eq("name", text).Or(where.Eq("description", text)),
		rel.Select().Limit(limit).Offset(offset),
	)
	return todos, err
}

func (r *repository) Update(ctx context.Context, todo *Todo) (*Todo, error) {
	err := r.r.Update(ctx, &todo)
	return todo, err
}

func (r *repository) Delete(ctx context.Context, todo *Todo) error {
	// TODO: maybe it would be good to specify user_id in delete
	return r.r.Delete(ctx, &todo)
}
