package service

import (
	"errors"

	"github.com/danikg/go-todo-rest-api/models"
)

// UserServiceMock ...
type UserServiceMock struct {
	GenerateErr bool
}

// GetAll ...
func (s *UserServiceMock) GetAll() ([]models.User, error) {
	if s.GenerateErr {
		return []models.User{}, errors.New("err")
	}

	users := []models.User{
		models.User{Username: "user1"},
		models.User{Username: "user2"},
	}

	users[0].ID = 1
	users[1].ID = 2
	return users, nil
}

// GetSingle ...
func (s *UserServiceMock) GetSingle(id uint) (models.User, error) {
	if id != 1 {
		return models.User{}, errors.New("err")
	}

	user := models.User{Username: "user1"}
	user.ID = 1
	return user, nil
}

// Create ...
func (s *UserServiceMock) Create(user *models.User) error {
	if s.GenerateErr {
		return errors.New("err")
	}
	return nil
}

// Update ...
func (s *UserServiceMock) Update(id uint, userData *models.User) (models.User, error) {
	if id != 1 {
		return models.User{}, errors.New("err")
	}

	user := models.User{Username: "user1"}
	user.ID = 1
	return user, nil
}

// Delete ...
func (s *UserServiceMock) Delete(id uint) error {
	if id != 1 {
		return errors.New("err")
	}
	return nil
}
