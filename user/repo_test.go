package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/rel/where"
	"github.com/go-rel/reltest"
	"github.com/sea-auca/clean-todo/user"
	"github.com/stretchr/testify/assert"
)

func setupRepo() (r *reltest.Repository, repo user.Repository) {
	r = reltest.New()
	repo = user.CreateNewRepo(r)
	return
}

func createValidUser() *user.User {
	u, _ := user.NewUser("valid@mail.com", "User user", "Paswwrod_valid1")
	return u
}

func TestRepoCreate(t *testing.T) {
	m, r := setupRepo()

	m.ExpectInsert().ForType("user.User").Success().Once()
	usr := createValidUser()
	usr_copy := *usr
	usr_copy.ID = 1
	m.ExpectFind(where.Eq("email", usr.Email)).Result(usr_copy).Once()

	usr, err := r.Create(context.Background(), usr)
	assert.NoErrorf(t, err, "Expected error to be nil")
	assert.Equal(t, int64(1), usr.ID)

}

func TestRepoCreateError(t *testing.T) {
	m, r := setupRepo()

	m.ExpectInsert().ForType("user.User").Error(errors.New("test error"))
	usr := createValidUser()
	m.ExpectFind(where.Eq("email", usr.Email)).NotFound().Times(0)

	_, err := r.Create(context.Background(), usr)
	assert.Errorf(t, err, "Expected error to be not nil")
}

func TestRepoUpdateChangesetError(t *testing.T) {
	_, r := setupRepo()

	usr := createValidUser()
	params := make(map[string]interface{})
	params["email"] = "new_invalid_email@mail."

	err := r.Update(context.Background(), usr, params)
	assert.Errorf(t, err, "Expected error to be not nil")
}

func TestRepoUpdateSuccess(t *testing.T) {
	m, r := setupRepo()

	usr := createValidUser()
	params := make(map[string]interface{})
	params["email"] = "new_valid_email@mail.com"

	m.ExpectUpdate().ForType("user.User").Success().Once()

	err := r.Update(context.Background(), usr, params)
	assert.NoErrorf(t, err, "Expected error to be not nil")
}

func TestRepoDeleteUser(t *testing.T) {
	m, r := setupRepo()
	usr := createValidUser()
	m.ExpectDelete().ForType("user.User").Success().Once()

	err := r.Delete(context.Background(), usr)
	assert.NoError(t, err)
}

func TestRepoGetUseByID(t *testing.T) {
	m, r := setupRepo()
	usr := createValidUser()
	usr.ID = 1

	m.ExpectFind(where.Eq("id", usr.ID))
	r.GetByID(context.Background(), usr.ID)
}

func TestRepoGetUseByEmail(t *testing.T) {
	m, r := setupRepo()
	usr := createValidUser()
	usr.ID = 1

	m.ExpectFind(where.Eq("email", usr.Email))
	r.GetByEmail(context.Background(), usr.Email)
}
