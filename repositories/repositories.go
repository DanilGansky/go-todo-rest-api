package repositories

import "github.com/danikg/go-todo-rest-api/models"

// IUserRepository ...
type IUserRepository interface {
	GetAll() ([]models.User, error)
	GetSingle(id uint) (models.User, error)
	Create(user *models.User) error
	Update(id uint, userData *models.User) (models.User, error)
	Delete(id uint) error
}

// ITodoListRepository ...
type ITodoListRepository interface {
	GetAll(userID uint) ([]models.TodoList, error)
	GetSingle(id uint) (models.TodoList, error)
	Create(userID uint, todoList *models.TodoList) error
	Update(id uint, todoListData *models.TodoList) (models.TodoList, error)
	Delete(id uint) error
}

// ITodoItemRepository ...
type ITodoItemRepository interface {
	GetAll(listID uint) ([]models.TodoItem, error)
	GetSingle(id uint) (models.TodoItem, error)
	Create(listID uint, todoItem *models.TodoItem) error
	Update(id uint, todoItemData *models.TodoItem) (models.TodoItem, error)
	Delete(id uint) error
}

// ITagRepository ...
type ITagRepository interface {
	GetAll(todoItem *models.TodoItem) ([]models.Tag, error)
	GetSingle(id uint) (models.Tag, error)
	Create(todoItem *models.TodoItem, tag *models.Tag) error
	Update(id uint, tagData *models.Tag) (models.Tag, error)
	Remove(todoItem *models.TodoItem, tagID uint) error
	Delete(id uint) error
}
