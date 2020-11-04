package models

import "gorm.io/gorm"

// TodoList represents a todo list in db
type TodoList struct {
	gorm.Model
	Name   string
	UserID uint
}
