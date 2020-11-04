package webservices

import (
	"testing"

	"github.com/danikg/go-todo-rest-api/models"

	"github.com/stretchr/testify/assert"

	mocks "github.com/danikg/go-todo-rest-api/repositories/mocks"
)

func TestTodoItemService_GetAll(t *testing.T) {
	todoItemService := NewTodoItemService(&mocks.TodoItemRepositoryMock{}, &mocks.TodoListRepositoryMock{})
	todoItems, err := todoItemService.GetAll(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, todoItems)

	todoItems, err = todoItemService.GetAll(2)
	assert.Error(t, err)
	assert.Empty(t, todoItems)
}

func TestTodoItemService_GetSingle(t *testing.T) {
	todoItemService := NewTodoItemService(&mocks.TodoItemRepositoryMock{}, &mocks.TodoListRepositoryMock{})
	todoItem, err := todoItemService.GetSingle(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, todoItem)

	todoItem, err = todoItemService.GetSingle(2)
	assert.Error(t, err)
	assert.Empty(t, todoItem)
}

func TestTodoItemService_Create(t *testing.T) {
	todoItemService := NewTodoItemService(&mocks.TodoItemRepositoryMock{}, &mocks.TodoListRepositoryMock{})
	todoItem := models.TodoItem{Title: "item"}
	todoItem.ID = 1

	err := todoItemService.Create(1, &todoItem)
	assert.NoError(t, err)

	err = todoItemService.Create(2, &todoItem)
	assert.Error(t, err)
}

func TestTodoItemService_Update(t *testing.T) {
	todoItemService := NewTodoItemService(&mocks.TodoItemRepositoryMock{}, &mocks.TodoListRepositoryMock{})
	todoItem := models.TodoItem{Title: "item"}
	todoItem.ID = 1

	resultTodoItem, err := todoItemService.Update(1, &todoItem)
	assert.NoError(t, err)
	assert.NotEmpty(t, resultTodoItem)

	resultTodoItem, err = todoItemService.Update(2, &todoItem)
	assert.Error(t, err)
	assert.Empty(t, &resultTodoItem)
}

func TestTodoItemService_Delete(t *testing.T) {
	todoItemService := NewTodoItemService(&mocks.TodoItemRepositoryMock{}, &mocks.TodoListRepositoryMock{})
	assert.NoError(t, todoItemService.Delete(1))
	assert.Error(t, todoItemService.Delete(2))
}
