package todo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/sea-auca/clean-todo/user"
)

type Repository interface {
	Create(ctx context.Context, to *Todo) (*Todo, error)
	ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error)
	SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error)
	Update(ctx context.Context, to *Todo) error
	Delete(ctx context.Context, to *Todo) error
}

type repository struct {
	r pgx.Conn
}

func CreateNewRepo(r pgx.Conn) Repository {
	return &repository{r}
}

func (r *repository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	rows, err := r.r.Query(ctx, `
	INSERT INTO todos (name, description, author_id, due_to, is_done, created_at, updatet_at)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING id, name, description, due_to, is_done, created_at, updatet_at;
	`, todo.Name, todo.Description, todo.Author.ID, todo.DueTo, todo.IsDone, todo.CreatedAt, todo.UpdatedAt)

	if err != nil {
		return nil, err
	}

	var newTodo Todo

	if rows.Next() {
		if err := rows.Scan(
			&newTodo.ID,
			&newTodo.Name,
			&newTodo.Description,
			&newTodo.DueTo,
			&newTodo.IsDone,
			&newTodo.CreatedAt,
			&newTodo.UpdatedAt,
		); err != nil {
			return nil, err
		}
	}
	newTodo.Author = todo.Author
	return &newTodo, nil
}

func (r *repository) ListByUserID(ctx context.Context, userID int64, limit, offset int) ([]*Todo, error) {
	rows, err := r.r.Query(ctx, `
	SELECT id, name, description, due_to, is_done, created_at, updated_at
	FROM todos
	WHERE user_id = $1
	LIMIT $2 OFFSET $3;
	`, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		todos                       []*Todo
		id                          int64
		name, description           string
		isDone                      bool
		dueTo, createdAt, updatedAt time.Time
	)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&description,
			&dueTo,
			&isDone,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}
		todos = append(todos, &Todo{
			ID:          id,
			Name:        name,
			Description: description,
			DueTo:       dueTo,
			IsDone:      isDone,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}
	return todos, nil
}

func (r *repository) SearchByText(ctx context.Context, text string, userID int64, limit, offset int) ([]*Todo, error) {
	rows, err := r.r.Query(ctx, `
	SELECT id, name, description, due_to, is_done, created_at, updated_at
	FROM todos
	WHERE user_id = $1
	AND to_tsvector(name || ' ' || description) @@ to_tsquery($2)
	LIMIT $3 OFFSET $4;
	`, userID, text, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		todos                       []*Todo
		id                          int64
		name, description           string
		isDone                      bool
		dueTo, createdAt, updatedAt time.Time
	)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&description,
			&dueTo,
			&isDone,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}
		todos = append(todos, &Todo{
			ID:          id,
			Name:        name,
			Description: description,
			DueTo:       dueTo,
			IsDone:      isDone,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Author:      &user.User{ID: userID},
		})
	}
	return todos, nil
}

func (r *repository) Update(ctx context.Context, todo *Todo) error {
	_, err := r.r.Query(ctx, `
	UPDATE todos
	SET name=$1, description=$2, due_to=$3, is_done=$4, created_at=$5, updated_at=$6
	WHERE id = $7;
	`, todo.Name, todo.Description, todo.DueTo, todo.IsDone, todo.CreatedAt, todo.UpdatedAt)
	return err
}

func (r *repository) Delete(ctx context.Context, todo *Todo) error {
	_, err := r.r.Query(ctx, `
	DELETE FROM todos
	WHERE id = $2;
	`, todo.ID)
	return err
}
