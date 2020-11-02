package repository

import (
	"github.com/danikg/go-todo-rest-api/internal/models"
	"gorm.io/gorm"
)

// TodoItemRepository ...
type TodoItemRepository struct {
	Conn *gorm.DB
}

// NewTodoItemRepository ...
func NewTodoItemRepository(conn *gorm.DB) *TodoItemRepository {
	return &TodoItemRepository{Conn: conn}
}

// GetAll returns all todo items by todo list id
func (t *TodoItemRepository) GetAll(listID uint) ([]models.TodoItem, error) {
	todoItems := []models.TodoItem{}
	return todoItems, t.Conn.Joins("TodoList").Preload("Tags").Find(&todoItems, "todo_list_id = ?", listID).Error
}

// GetSingle returns a todo item by id
func (t *TodoItemRepository) GetSingle(id uint) (models.TodoItem, error) {
	todoItem := models.TodoItem{}
	return todoItem, t.Conn.Joins("TodoList").Preload("Tags").First(&todoItem, id).Error
}

// Create creates a new todo item
func (t *TodoItemRepository) Create(listID uint, todoItem *models.TodoItem) error {
	todoItem.TodoListID = listID
	return t.Conn.Create(todoItem).Error
}

// Update updates the todo item
func (t *TodoItemRepository) Update(id uint, todoItemData *models.TodoItem) (models.TodoItem, error) {
	todoItem, err := t.GetSingle(id)
	if err != nil {
		return todoItem, err
	}
	return todoItem, t.Conn.Model(&todoItem).Updates(
		models.TodoItem{
			Title:       todoItemData.Title,
			Description: todoItemData.Description,
		},
	).Error
}

// Delete removes the todo item
func (t *TodoItemRepository) Delete(id uint) error {
	todoItem, err := t.GetSingle(id)
	if err != nil {
		return err
	}
	return t.Conn.Unscoped().Delete(&todoItem).Error
}
