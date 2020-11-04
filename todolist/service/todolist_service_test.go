package service

import (
	"testing"

	"github.com/danikg/go-todo-rest-api/models"

	"github.com/stretchr/testify/assert"

	tlMocks "github.com/danikg/go-todo-rest-api/todolist/repository/mock"
	userMocks "github.com/danikg/go-todo-rest-api/user/repository/mock"
)

func TestTodoListService_GetAll(t *testing.T) {
	todoListService := NewTodoListService(&userMocks.UserRepositoryMock{}, &tlMocks.TodoListRepositoryMock{})
	todoLists, err := todoListService.GetAll(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, todoLists)

	todoLists, err = todoListService.GetAll(2)
	assert.Error(t, err)
	assert.Empty(t, todoLists)
}

func TestTodoListService_GetSingle(t *testing.T) {
	todoListService := NewTodoListService(&userMocks.UserRepositoryMock{}, &tlMocks.TodoListRepositoryMock{})
	todoList, err := todoListService.GetSingle(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, todoList)

	todoList, err = todoListService.GetSingle(2)
	assert.Error(t, err)
	assert.Empty(t, todoList)
}

func TestTodoListService_Create(t *testing.T) {
	todoListService := NewTodoListService(&userMocks.UserRepositoryMock{}, &tlMocks.TodoListRepositoryMock{})
	todoList := models.TodoList{Name: "list"}
	todoList.ID = 1

	err := todoListService.Create(1, &todoList)
	assert.NoError(t, err)

	err = todoListService.Create(2, &todoList)
	assert.Error(t, err)
}

func TestTodoListService_Update(t *testing.T) {
	todoListService := NewTodoListService(&userMocks.UserRepositoryMock{}, &tlMocks.TodoListRepositoryMock{})
	todoList := models.TodoList{Name: "list"}
	todoList.ID = 1

	resultTodoList, err := todoListService.Update(1, &todoList)
	assert.NoError(t, err)
	assert.NotEmpty(t, resultTodoList)

	resultTodoList, err = todoListService.Update(2, &todoList)
	assert.Error(t, err)
	assert.Empty(t, &resultTodoList)
}

func TestTodoListService_Delete(t *testing.T) {
	todoListService := NewTodoListService(&userMocks.UserRepositoryMock{}, &tlMocks.TodoListRepositoryMock{})
	assert.NoError(t, todoListService.Delete(1))
	assert.Error(t, todoListService.Delete(2))
}
