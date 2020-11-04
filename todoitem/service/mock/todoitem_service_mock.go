package mock

import (
	"errors"

	"github.com/danikg/go-todo-rest-api/models"
)

// TodoItemServiceMock ...
type TodoItemServiceMock struct{}

// GetAll ...
func (s *TodoItemServiceMock) GetAll(listID uint) ([]models.TodoItem, error) {
	if listID != 1 {
		return []models.TodoItem{}, errors.New("err")
	}

	todoItems := []models.TodoItem{
		{Title: "item1", Description: ""},
		{Title: "item2", Description: "desc2"},
	}

	todoItems[0].ID = 1
	todoItems[1].ID = 2
	return todoItems, nil
}

// GetSingle ...
func (s *TodoItemServiceMock) GetSingle(id uint) (models.TodoItem, error) {
	if id != 1 {
		return models.TodoItem{}, errors.New("not found")
	}

	todoItem := models.TodoItem{Title: "item1", Description: ""}
	todoItem.ID = 1
	return todoItem, nil
}

// Create ...
func (s *TodoItemServiceMock) Create(listID uint, todoItem *models.TodoItem) error {
	if listID != 1 {
		return errors.New("err")
	}
	return nil
}

// Update ...
func (s *TodoItemServiceMock) Update(id uint, todoItemData *models.TodoItem) (models.TodoItem, error) {
	if id != 1 {
		return models.TodoItem{}, errors.New("err")
	}

	todoItem := models.TodoItem{Title: "item1", Description: ""}
	todoItem.ID = 1
	return todoItem, nil
}

// Delete ...
func (s *TodoItemServiceMock) Delete(id uint) error {
	if id != 1 {
		return errors.New("err")
	}
	return nil
}
