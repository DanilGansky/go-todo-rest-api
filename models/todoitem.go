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

// ITodoItemService ...
type ITodoItemService interface {
	GetAll(listID uint) ([]TodoItem, error)
	GetSingle(id uint) (TodoItem, error)
	Create(listID uint, todoItem *TodoItem) error
	Update(id uint, todoItemData *TodoItem) (TodoItem, error)
	Delete(id uint) error
}

// ITodoItemRepository ...
type ITodoItemRepository interface {
	GetAll(listID uint) ([]TodoItem, error)
	GetSingle(id uint) (TodoItem, error)
	Create(listID uint, todoItem *TodoItem) error
	Update(id uint, todoItemData *TodoItem) (TodoItem, error)
	Delete(id uint) error
}
