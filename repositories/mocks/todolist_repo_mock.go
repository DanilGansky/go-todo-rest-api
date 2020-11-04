package mocks

import (
	"errors"

	"github.com/danikg/go-todo-rest-api/models"
)

// TodoListRepositoryMock ...
type TodoListRepositoryMock struct{}

// GetAll ...
func (s *TodoListRepositoryMock) GetAll(userID uint) ([]models.TodoList, error) {
	if userID != 1 {
		return []models.TodoList{}, errors.New("err")
	}

	todoLists := []models.TodoList{
		models.TodoList{Name: "list1", UserID: 1},
		models.TodoList{Name: "list2", UserID: 1},
	}

	todoLists[0].ID = 1
	todoLists[1].ID = 2
	return todoLists, nil
}

// GetSingle ...
func (s *TodoListRepositoryMock) GetSingle(id uint) (models.TodoList, error) {
	if id != 1 {
		return models.TodoList{}, errors.New("not found")
	}

	todoList := models.TodoList{Name: "list1", UserID: 1}
	todoList.ID = 1
	return todoList, nil
}

// Create ...
func (s *TodoListRepositoryMock) Create(userID uint, todoList *models.TodoList) error {
	if userID != 1 {
		return errors.New("err")
	}
	return nil
}

// Update ...
func (s *TodoListRepositoryMock) Update(id uint, todoListData *models.TodoList) (models.TodoList, error) {
	if id != 1 {
		return models.TodoList{}, errors.New("err")
	}

	todoList := models.TodoList{Name: "list1", UserID: 1}
	todoList.ID = 1
	return todoList, nil
}

// Delete ...
func (s *TodoListRepositoryMock) Delete(id uint) error {
	if id != 1 {
		return errors.New("err")
	}
	return nil
}
