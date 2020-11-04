package models

import "gorm.io/gorm"

// TodoItem represents a todo item in db
type TodoItem struct {
	gorm.Model
	Title       string
	Description string
	TodoListID  uint
	TodoList    TodoList `gorm:"constraint:OnDelete:CASCADE;"`
	Tags        []Tag    `gorm:"many2many:todo_item_tags;constraint:OnDelete:CASCADE;"`
}
