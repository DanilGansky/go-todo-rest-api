package pg

import (
	"github.com/danikg/go-todo-rest-api/models"
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
	err := t.Conn.Joins("TodoList").Preload("Tags").Find(&todoItems, "todo_list_id = ?", listID).Error
	return todoItems, err
}

// GetSingle returns a todo item by id
func (t *TodoItemRepository) GetSingle(id uint) (models.TodoItem, error) {
	todoItem := models.TodoItem{}
	err := t.Conn.Joins("TodoList").Preload("Tags").First(&todoItem, id).Error
	return todoItem, err
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
	err = t.Conn.Model(&todoItem).
		Updates(models.TodoItem{
			Title:       todoItemData.Title,
			Description: todoItemData.Description}).Error
	return todoItem, err
}

// Delete removes the todo item
func (t *TodoItemRepository) Delete(id uint) error {
	todoItem, err := t.GetSingle(id)
	if err != nil {
		return err
	}
	return t.Conn.Unscoped().Delete(&todoItem).Error
}
