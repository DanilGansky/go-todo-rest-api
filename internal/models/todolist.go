package models

import "gorm.io/gorm"

// TodoList represents a todo list in db
type TodoList struct {
	gorm.Model
	Name   string
	UserID uint
}

// ITodoListService ...
type ITodoListService interface {
	GetAll(userID uint) ([]TodoList, error)
	GetSingle(id uint) (TodoList, error)
	Create(userID uint, todoList *TodoList) error
	Update(id uint, todoListData *TodoList) (TodoList, error)
	Delete(id uint) error
}

// ITodoListRepository ...
type ITodoListRepository interface {
	GetAll(userID uint) ([]TodoList, error)
	GetSingle(id uint) (TodoList, error)
	Create(userID uint, todoList *TodoList) error
	Update(id uint, todoListData *TodoList) (TodoList, error)
	Delete(id uint) error
}
