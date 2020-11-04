package webservices

import (
	"github.com/danikg/go-todo-rest-api/models"
	repos "github.com/danikg/go-todo-rest-api/repositories"
)

// UserService ...
type UserService struct {
	UserRepo repos.IUserRepository
}

// NewUserService ...
func NewUserService(userRepo repos.IUserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// GetAll returns all users from the db
func (u *UserService) GetAll() ([]models.User, error) {
	return u.UserRepo.GetAll()
}

// GetSingle returns a user by id
func (u *UserService) GetSingle(id uint) (models.User, error) {
	return u.UserRepo.GetSingle(id)
}

// Create creates a new user
func (u *UserService) Create(user *models.User) error {
	return u.UserRepo.Create(user)
}

// Update updates the user
func (u *UserService) Update(id uint, userData *models.User) (models.User, error) {
	return u.UserRepo.Update(id, userData)
}

// Delete removes the user
func (u *UserService) Delete(id uint) error {
	return u.UserRepo.Delete(id)
}
