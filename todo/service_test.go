package todo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	defaultTodo := Todo{
		Name:        "hello",
		Description: "testing",
	}
	defaultExpected := Todo{
		Name:        "hello",
		Description: "testing",
	}

	noName := defaultTodo
	noName.Name = ""
	noNameExpected := noName

	noDescription := defaultTodo
	noDescription.Description = ""
	noDescriptionExpected := noDescription

	testCases := []struct {
		desc string
		ctx  context.Context
		*Todo
		err    error
		expect *Todo
	}{
		{
			desc:   "usual",
			ctx:    context.Background(),
			Todo:   &defaultTodo,
			err:    nil,
			expect: &defaultExpected,
		},
		{
			desc:   "no name",
			ctx:    context.Background(),
			Todo:   &noName,
			err:    nil,
			expect: &noNameExpected,
		},
		{
			desc:   "no description",
			ctx:    context.Background(),
			Todo:   &noDescription,
			err:    nil,
			expect: &noDescriptionExpected,
		},
	}

	mockRepo := NewMockRepository()
	service := NewService(mockRepo)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			todo, err := service.Create(tC.ctx, tC.Todo)
			assert.Equalf(t, tC.err, err, "Could not create todo on %s. Due to error: %v. Result is: %v", tC.desc, tC.err, todo)
			assert.Equalf(t, tC.expect, todo, "Input into Create was mutated. Result: %v. Wanted: %v", todo, tC.expect)
		})
	}
}

func TestListByUserID(t *testing.T) {
	var defaultExpected []*Todo = nil

	type input struct {
		userID int64
		limit  int
		offset int
	}
	testCases := []struct {
		desc   string
		ctx    context.Context
		input  input
		err    error
		expect []*Todo
	}{
		{
			desc: "usual",
			ctx:  context.Background(),
			input: input{
				userID: 1,
				limit:  10,
				offset: 0,
			},
			err:    nil,
			expect: defaultExpected,
		},
		{
			desc: "negative userID",
			ctx:  context.Background(),
			input: input{
				userID: -1,
				limit:  10,
				offset: 0,
			},
			err:    ErrNegativeUserID,
			expect: defaultExpected,
		},
		{
			desc: "nil userID",
			ctx:  context.Background(),
			input: input{
				userID: 0,
				limit:  10,
				offset: 0,
			},
			err:    ErrNullUserID,
			expect: defaultExpected,
		},
	}

	mockRepo := NewMockRepository()
	service := NewService(mockRepo)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			todos, err := service.ListByUserID(tC.ctx, tC.input.userID, tC.input.limit, tC.input.offset)
			assert.Equalf(t, tC.err, err, "Could not list todos on %s. Due to error: %v. Result is: %v", tC.desc, tC.err, todos)
			assert.Equalf(t, tC.expect, todos, "Input into ListByUserID was mutated. Result: %v. Wanted: %v", todos, tC.expect)
		})
	}
}

func TestSearchByText(t *testing.T) {
	var defaultExpected []*Todo = nil

	type input struct {
		text   string
		userID int64
		limit  int
		offset int
	}
	testCases := []struct {
		desc   string
		ctx    context.Context
		input  input
		err    error
		expect []*Todo
	}{
		{
			desc: "usual",
			ctx:  context.Background(),
			input: input{
				text:   "hello",
				userID: 1,
				limit:  10,
				offset: 0,
			},
			err:    nil,
			expect: defaultExpected,
		},
		{
			desc: "negative userID",
			ctx:  context.Background(),
			input: input{
				text:   "hello",
				userID: -1,
				limit:  10,
				offset: 0,
			},
			err:    ErrNegativeUserID,
			expect: defaultExpected,
		},
		{
			desc: "nil userID",
			ctx:  context.Background(),
			input: input{
				text:   "hello",
				userID: 0,
				limit:  10,
				offset: 0,
			},
			err:    ErrNullUserID,
			expect: defaultExpected,
		},
	}

	mockRepo := NewMockRepository()
	service := NewService(mockRepo)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			todos, err := service.SearchByText(tC.ctx, tC.input.text, tC.input.userID, tC.input.limit, tC.input.offset)
			assert.Equalf(t, tC.err, err, "Could not list todos on %s. Due to error: %v. Result is: %v", tC.desc, tC.err, todos)
			assert.Equalf(t, tC.expect, todos, "Input into ListByUserID was mutated. Result: %v. Wanted: %v", todos, tC.expect)
		})
	}
}

func TestUpdate(t *testing.T) {
	defaultTodo := Todo{
		Name:        "hello",
		Description: "testing",
	}
	defaultExpected := Todo{
		Name:        "hello",
		Description: "testing",
	}

	noName := defaultTodo
	noName.Name = ""
	noNameExpected := noName

	noDescription := defaultTodo
	noDescription.Description = ""
	noDescriptionExpected := noDescription

	testCases := []struct {
		desc string
		ctx  context.Context
		*Todo
		err    error
		expect *Todo
	}{
		{
			desc:   "usual",
			ctx:    context.Background(),
			Todo:   &defaultTodo,
			err:    nil,
			expect: &defaultExpected,
		},
		{
			desc:   "no name",
			ctx:    context.Background(),
			Todo:   &noName,
			err:    nil,
			expect: &noNameExpected,
		},
		{
			desc:   "no description",
			ctx:    context.Background(),
			Todo:   &noDescription,
			err:    nil,
			expect: &noDescriptionExpected,
		},
	}

	mockRepo := NewMockRepository()
	service := NewService(mockRepo)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			todo, err := service.Update(tC.ctx, tC.Todo)
			assert.Equalf(t, tC.err, err, "Could not Update todo on %s. Due to error: %v. Result is: %v", tC.desc, tC.err, todo)
			assert.Equalf(t, tC.expect, todo, "Input into Update was mutated. Result: %v. Wanted: %v", todo, tC.expect)
		})
	}
}

func TestDelete(t *testing.T) {
	defaultTodo := Todo{
		Name:        "hello",
		Description: "testing",
	}
	defaultExpected := Todo{
		Name:        "hello",
		Description: "testing",
	}

	noName := defaultTodo
	noName.Name = ""
	noNameExpected := noName

	noDescription := defaultTodo
	noDescription.Description = ""
	noDescriptionExpected := noDescription

	testCases := []struct {
		desc string
		ctx  context.Context
		*Todo
		err    error
		expect *Todo
	}{
		{
			desc:   "usual",
			ctx:    context.Background(),
			Todo:   &defaultTodo,
			err:    nil,
			expect: &defaultExpected,
		},
		{
			desc:   "no name",
			ctx:    context.Background(),
			Todo:   &noName,
			err:    nil,
			expect: &noNameExpected,
		},
		{
			desc:   "no description",
			ctx:    context.Background(),
			Todo:   &noDescription,
			err:    nil,
			expect: &noDescriptionExpected,
		},
	}

	mockRepo := NewMockRepository()
	service := NewService(mockRepo)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := service.Delete(tC.ctx, tC.Todo)
			assert.Equalf(t, tC.err, err, "Could not Delete todo on %s. Due to error: %v", tC.desc, tC.err)
		})
	}
}
