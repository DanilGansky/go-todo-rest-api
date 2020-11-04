package webservices

import (
	"testing"

	"github.com/danikg/go-todo-rest-api/models"
	mocks "github.com/danikg/go-todo-rest-api/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUserService_GetAll(t *testing.T) {
	userService := NewUserService(&mocks.UserRepositoryMock{})
	users, err := userService.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, users)

	userService = NewUserService(&mocks.UserRepositoryMock{GenerateErr: true})
	users, err = userService.GetAll()
	assert.Error(t, err)
	assert.Empty(t, users)
}

func TestUserService_GetSingle(t *testing.T) {
	userService := NewUserService(&mocks.UserRepositoryMock{})
	user, err := userService.GetSingle(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	user, err = userService.GetSingle(2)
	assert.Error(t, err)
	assert.Empty(t, user)
}

func TestUserService_Create(t *testing.T) {
	userService := NewUserService(&mocks.UserRepositoryMock{})
	user := models.User{Username: "user1"}
	user.ID = 1

	err := userService.Create(&user)
	assert.NoError(t, err)

	userService = NewUserService(&mocks.UserRepositoryMock{GenerateErr: true})
	err = userService.Create(&user)
	assert.Error(t, err)
}

func TestUserService_Update(t *testing.T) {
	userService := NewUserService(&mocks.UserRepositoryMock{})
	user := models.User{Username: "user1"}
	user.ID = 1

	resultUser, err := userService.Update(1, &user)
	assert.NoError(t, err)
	assert.NotEmpty(t, resultUser)

	resultUser, err = userService.Update(2, &user)
	assert.Error(t, err)
	assert.Empty(t, &resultUser)
}

func TestUserService_Delete(t *testing.T) {
	userService := NewUserService(&mocks.UserRepositoryMock{})
	assert.NoError(t, userService.Delete(1))
	assert.Error(t, userService.Delete(2))
}
