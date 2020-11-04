package services

import "github.com/danikg/go-todo-rest-api/models"

// IUserService ...
type IUserService interface {
	GetAll() ([]models.User, error)
	GetSingle(id uint) (models.User, error)
	Create(user *models.User) error
	Update(id uint, userData *models.User) (models.User, error)
	Delete(id uint) error
}

// ITodoListService ...
type ITodoListService interface {
	GetAll(userID uint) ([]models.TodoList, error)
	GetSingle(id uint) (models.TodoList, error)
	Create(userID uint, todoList *models.TodoList) error
	Update(id uint, todoListData *models.TodoList) (models.TodoList, error)
	Delete(id uint) error
}

// ITodoItemService ...
type ITodoItemService interface {
	GetAll(listID uint) ([]models.TodoItem, error)
	GetSingle(id uint) (models.TodoItem, error)
	Create(listID uint, todoItem *models.TodoItem) error
	Update(id uint, todoItemData *models.TodoItem) (models.TodoItem, error)
	Delete(id uint) error
}

// ITagService ...
type ITagService interface {
	GetAll(itemID uint) ([]models.Tag, error)
	GetSingle(id uint) (models.Tag, error)
	Create(itemID uint, tag *models.Tag) error
	Update(id uint, tagData *models.Tag) (models.Tag, error)
	Remove(itemID uint, tagID uint) error
	Delete(id uint) error
}
