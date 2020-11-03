package pg

import (
	"github.com/danikg/go-todo-rest-api/models"
	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository ...
func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{Conn: conn}
}

// GetAll returns all users from the db
func (u *UserRepository) GetAll() ([]models.User, error) {
	users := []models.User{}
	err := u.Conn.Preload("TodoLists").Find(&users).Error
	return users, err
}

// GetSingle returns a user by id
func (u *UserRepository) GetSingle(id uint) (models.User, error) {
	user := models.User{}
	err := u.Conn.Preload("TodoLists").First(&user, id).Error
	return user, err
}

// Create creates a new user
func (u *UserRepository) Create(user *models.User) error {
	return u.Conn.Create(user).Error
}

// Update updates the user
func (u *UserRepository) Update(id uint, userData *models.User) (models.User, error) {
	user, err := u.GetSingle(id)
	if err != nil {
		return user, err
	}

	err = u.Conn.Model(&user).Update("Username", userData.Username).Error
	return user, err
}

// Delete removes the user
func (u *UserRepository) Delete(id uint) error {
	user, err := u.GetSingle(id)
	if err != nil {
		return err
	}
	return u.Conn.Unscoped().Delete(&user).Error
}
