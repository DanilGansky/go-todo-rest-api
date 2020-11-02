package repository

import (
	"github.com/danikg/go-todo-rest-api/models"
	"gorm.io/gorm"
)

// TodoListRepository ...
type TodoListRepository struct {
	Conn *gorm.DB
}

// NewTodoListRepository ...
func NewTodoListRepository(conn *gorm.DB) *TodoListRepository {
	return &TodoListRepository{Conn: conn}
}

// GetAll returns all todo lists by user id
func (t *TodoListRepository) GetAll(userID uint) ([]models.TodoList, error) {
	todoLists := []models.TodoList{}
	return todoLists, t.Conn.Find(&todoLists, "user_id = ?", userID).Error
}

// GetSingle returns a todo list by id
func (t *TodoListRepository) GetSingle(id uint) (models.TodoList, error) {
	todoList := models.TodoList{}
	return todoList, t.Conn.First(&todoList, id).Error
}

// Create creates a new todo list
func (t *TodoListRepository) Create(userID uint, todoList *models.TodoList) error {
	todoList.UserID = userID
	return t.Conn.Create(todoList).Error
}

// Update updates the todo list
func (t *TodoListRepository) Update(id uint, todoListData *models.TodoList) (models.TodoList, error) {
	todoList, err := t.GetSingle(id)
	if err != nil {
		return todoList, err
	}
	return todoList, t.Conn.Model(&todoList).Update("Name", todoListData.Name).Error
}

// Delete removes the todo list
func (t *TodoListRepository) Delete(id uint) error {
	todoList, err := t.GetSingle(id)
	if err != nil {
		return err
	}
	return t.Conn.Unscoped().Delete(&todoList).Error
}
