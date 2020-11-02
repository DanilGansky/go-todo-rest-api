package models

import "gorm.io/gorm"

// User model represents a user in db
type User struct {
	gorm.Model
	Username  string
	TodoLists []TodoList `gorm:"constraint:OnDelete:CASCADE;"`
}

// IUserService ...
type IUserService interface {
	GetAll() ([]User, error)
	GetSingle(id uint) (User, error)
	Create(user *User) error
	Update(id uint, userData *User) (User, error)
	Delete(id uint) error
}

// IUserRepository ...
type IUserRepository interface {
	GetAll() ([]User, error)
	GetSingle(id uint) (User, error)
	Create(user *User) error
	Update(id uint, userData *User) (User, error)
	Delete(id uint) error
}
