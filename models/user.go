package models

import "gorm.io/gorm"

// User model represents a user in db
type User struct {
	gorm.Model
	Username  string
	TodoLists []TodoList `gorm:"constraint:OnDelete:CASCADE;"`
}
