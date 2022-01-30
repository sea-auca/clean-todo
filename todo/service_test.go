package todo_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sea-auca/clean-todo/todo"
	"github.com/sea-auca/clean-todo/todo/mocks"
	"github.com/sea-auca/clean-todo/user"
	"github.com/stretchr/testify/assert"
)

func setupService(t *testing.T) (*mocks.MockRepository, todo.Service) {
	ctrl := gomock.NewController(t)
	m := mocks.NewMockRepository(ctrl)
	todo := todo.NewService(m)
	return m, todo
}

var bg = context.Background()

func TestCreate(t *testing.T) {
	defaultInput := todo.Todo{
		Name:        "test",
		Description: "test",
		DueTo:       time.Date(2050, 12, 1, 1, 1, 1, 1, time.Now().Location()),
		Author: &user.User{
			ID:         1,
			Fullname:   "Rasulov-Emirlan",
			Email:      "go@gmail.com",
			Hash:       "$2a$04$Wjv8rtuvD9zBzjnKYA88Auv5eoDWc1REy6G7ZRXxrRXK0P5QCja/C",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			VerifiedAt: time.Now(),
		},
	}

	emptyNameInput := defaultInput
	emptyNameInput.Name = ""

	type wantedResult struct {
		err   error
		isnil bool
	}

	testCases := []struct {
		desc         string
		input        *todo.Todo
		wantedResult wantedResult
	}{
		{
			desc:  "default",
			input: &defaultInput,
			wantedResult: wantedResult{
				err:   nil,
				isnil: false,
			},
		},
		{
			desc:  "empty name",
			input: &emptyNameInput,
			wantedResult: wantedResult{
				err:   todo.ErrEmptyName,
				isnil: true,
			},
		},
	}

	m, todoService := setupService(t)
	m.EXPECT().Create(bg, gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, tt *todo.Todo) (*todo.Todo, error) {
		return tt, nil
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			todos, err := todoService.Create(
				bg,
				tC.input.Name,
				tC.input.Description,
				tC.input.DueTo,
				*tC.input.Author,
			)

			assert.Equalf(t, tC.wantedResult.err, err, "Expected a different error")
			if tC.wantedResult.isnil {
				assert.Nilf(t, todos, "Expected todos to be nil")
				return
			}
			if !tC.wantedResult.isnil {
				assert.NotNilf(t, todos, "Epected todos not to be nil")
			}
		})
	}
}

func TestListByUserID(t *testing.T) {
	type input struct {
		userID int64
		limit  int
		offset int
	}
	type wantedResult struct {
		err   error
		isnil bool
	}

	testCases := []struct {
		desc         string
		input        input
		wantedResult wantedResult
	}{
		{
			desc: "default",
			input: input{
				userID: 1,
				limit:  10,
				offset: 0,
			},
			wantedResult: wantedResult{
				err:   nil,
				isnil: false,
			},
		},
		{
			desc: "negative userID",
			input: input{
				userID: -1,
				limit:  10,
				offset: 0,
			},
			wantedResult: wantedResult{
				err:   todo.ErrNegativeUserID,
				isnil: true,
			},
		},
		{
			desc: "-1, -1 on limit and offset",
			input: input{
				userID: 1,
				limit:  -1,
				offset: -1,
			},
			wantedResult: wantedResult{
				err:   todo.ErrNegativeLimit,
				isnil: true,
			},
		},
		{
			desc: "negative limit",
			input: input{
				userID: 1,
				limit:  -2,
				offset: 0,
			},
			wantedResult: wantedResult{
				err:   todo.ErrNegativeLimit,
				isnil: true,
			},
		},
		{
			desc: "negative offset",
			input: input{
				userID: 1,
				limit:  0,
				offset: -2,
			},
			wantedResult: wantedResult{
				err:   todo.ErrNegativeOffset,
				isnil: true,
			},
		},
	}

	m, todoService := setupService(t)
	m.EXPECT().ListByUserID(bg, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, userID int64, limit, offset int) ([]*todo.Todo, error) {
		var todos []*todo.Todo
		todos = append(todos, &todo.Todo{
			ID:          1,
			Name:        "test",
			Description: "test",
		})
		return todos, nil
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			todos, err := todoService.ListByUserID(
				bg,
				tC.input.userID,
				tC.input.limit,
				tC.input.offset,
			)

			assert.Equalf(t, tC.wantedResult.err, err, "Expected a different error")
			if tC.wantedResult.isnil {
				assert.Nilf(t, todos, "Expected todos to be nil")
				return
			}
			if !tC.wantedResult.isnil {
				assert.NotNilf(t, todos, "Epected todos not to be nil")
			}
		})
	}
}

func TestSearchByText(t *testing.T) {
	type input struct {
		text   string
		userID int64
		limit  int
		offset int
	}
	type wantedResult struct {
		err   error
		isnil bool
	}

	testCases := []struct {
		desc         string
		input        input
		wantedResult wantedResult
	}{
		{
			desc: "default",
			input: input{
				text:   "text",
				userID: 1,
				limit:  10,
				offset: 0,
			},
			wantedResult: wantedResult{
				err:   nil,
				isnil: false,
			},
		},
		{
			desc: "empty text",
			input: input{
				text:   "",
				userID: 1,
				limit:  10,
				offset: 0,
			},
			wantedResult: wantedResult{
				err:   todo.ErrEmptyText,
				isnil: true,
			},
		},
		{
			desc: "negative userID",
			input: input{
				text:   "text",
				userID: -1,
				limit:  10,
				offset: 0,
			},
			wantedResult: wantedResult{
				err:   todo.ErrNegativeUserID,
				isnil: true,
			},
		},
		{
			desc: "-1, -1 on limit and offset",
			input: input{
				text:   "text",
				userID: 1,
				limit:  -1,
				offset: -1,
			},
			wantedResult: wantedResult{
				err:   todo.ErrNegativeLimit,
				isnil: true,
			},
		},
		{
			desc: "negative limit",
			input: input{
				text:   "text",
				userID: 1,
				limit:  -2,
				offset: 0,
			},
			wantedResult: wantedResult{
				err:   todo.ErrNegativeLimit,
				isnil: true,
			},
		},
		{
			desc: "negative offset",
			input: input{
				text:   "text",
				userID: 1,
				limit:  0,
				offset: -2,
			},
			wantedResult: wantedResult{
				err:   todo.ErrNegativeOffset,
				isnil: true,
			},
		},
	}

	m, todoService := setupService(t)
	m.EXPECT().SearchByText(bg, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, text string, userID int64, limit, offset int) ([]*todo.Todo, error) {
		var todos []*todo.Todo
		todos = append(todos, &todo.Todo{
			ID:          1,
			Name:        "test",
			Description: "test",
		})
		return todos, nil
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			todos, err := todoService.SearchByText(
				bg,
				tC.input.text,
				tC.input.userID,
				tC.input.limit,
				tC.input.offset,
			)

			assert.Equalf(t, tC.wantedResult.err, err, "Expected a different error")
			if tC.wantedResult.isnil {
				assert.Nilf(t, todos, "Expected todos to be nil")
				return
			}
			if !tC.wantedResult.isnil {
				assert.NotNilf(t, todos, "Epected todos not to be nil")
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	defaultInput := todo.Todo{
		ID:          1,
		Name:        "test",
		Description: "test",
	}

	nullIDInput := defaultInput
	nullIDInput.ID = 0

	emptyNameInput := defaultInput
	emptyNameInput.Name = ""

	emptyDescriptionInput := defaultInput
	emptyDescriptionInput.Description = ""

	type wantedResult struct {
		err   error
		isnil bool
	}

	testCases := []struct {
		desc         string
		input        *todo.Todo
		wantedResult wantedResult
	}{
		{
			desc:  "default",
			input: &defaultInput,
			wantedResult: wantedResult{
				err:   nil,
				isnil: false,
			},
		},
		{
			desc:  "null todo.ID",
			input: &nullIDInput,
			wantedResult: wantedResult{
				err:   todo.ErrNegativeID,
				isnil: true,
			},
		},
		{
			desc:  "empty name",
			input: &emptyNameInput,
			wantedResult: wantedResult{
				err:   todo.ErrEmptyName,
				isnil: true,
			},
		},
		{
			desc:  "empty description",
			input: &emptyDescriptionInput,
			wantedResult: wantedResult{
				err:   todo.ErrEmptyDescription,
				isnil: true,
			},
		},
	}

	m, todoService := setupService(t)
	m.EXPECT().Update(bg, gomock.AssignableToTypeOf(&defaultInput)).AnyTimes().DoAndReturn(func(ctx context.Context, t *todo.Todo) (*todo.Todo, error) {
		return t, nil
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			todos, err := todoService.Update(
				bg,
				tC.input,
			)

			assert.Equalf(t, tC.wantedResult.err, err, "Expected a different error")
			if tC.wantedResult.isnil {
				assert.Nilf(t, todos, "Expected todos to be nil")
				return
			}
			if !tC.wantedResult.isnil {
				assert.NotNilf(t, todos, "Epected todos not to be nil")
			}
		})
	}
}

func TestDelete(t *testing.T) {
	defaultInput := todo.Todo{
		ID:          1,
		Name:        "test",
		Description: "test",
	}

	nullIDInput := defaultInput
	nullIDInput.ID = 0

	emptyNameInput := defaultInput
	emptyNameInput.Name = ""

	emptyDescriptionInput := defaultInput
	emptyDescriptionInput.Description = ""

	testCases := []struct {
		desc  string
		input *todo.Todo
		err   error
	}{
		{
			desc:  "default",
			input: &defaultInput,
			err:   nil,
		},
		{
			desc:  "null todo.ID",
			input: &nullIDInput,
			err:   todo.ErrNegativeID,
		},
	}

	m, todoService := setupService(t)
	m.EXPECT().Delete(bg, gomock.AssignableToTypeOf(&defaultInput)).AnyTimes().DoAndReturn(func(ctx context.Context, t *todo.Todo) error {
		return nil
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := todoService.Delete(
				bg,
				tC.input,
			)

			assert.Equalf(t, tC.err, err, "Expected a different error")
		})
	}
}
