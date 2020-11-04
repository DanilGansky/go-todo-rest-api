package http

import (
	"encoding/json"
	. "net/http"
	"testing"

	"github.com/danikg/go-todo-rest-api/models"
	"github.com/danikg/go-todo-rest-api/todoitem/service/mock"
	"github.com/danikg/go-todo-rest-api/utils/test"
	"github.com/stretchr/testify/assert"
)

type todoItemTest struct {
	title           string
	method          string
	path            string
	route           string
	body            []byte
	shouldPass      bool
	statusCode      int
	todoItemResult  models.TodoItem
	todoItemsResult []models.TodoItem
}

func compareTodoItems(t *testing.T, todoItem1 models.TodoItem, todoItem2 models.TodoItem) {
	assert.Equal(t, todoItem1.ID, todoItem2.ID)
	assert.Equal(t, todoItem1.Title, todoItem2.Title)
	assert.Equal(t, todoItem1.Description, todoItem2.Description)
}

func testTodoItemResult(t *testing.T, tc todoItemTest, controller func(ResponseWriter, *Request)) {
	w, r := test.NewRequest(tc.method, tc.path, tc.body)
	test.MakeRequest(tc.route, controller, w, r)
	assert.Equal(t, tc.statusCode, w.Code)

	if tc.shouldPass {
		if len(tc.todoItemsResult) != 0 {
			var result []models.TodoItem
			json.NewDecoder(w.Body).Decode(&result)
			compareTodoItems(t, tc.todoItemsResult[0], result[0])
			compareTodoItems(t, tc.todoItemsResult[1], result[1])
		} else {
			var result models.TodoItem
			json.NewDecoder(w.Body).Decode(&result)
			compareTodoItems(t, tc.todoItemResult, result)
		}
	}
}

func TestTodoItemController_GetAll(t *testing.T) {
	tests := []todoItemTest{
		{
			title:      "Get all todo items",
			method:     "GET",
			path:       "/todo_lists/1/todo_items",
			route:      "/todo_lists/{list_id}/todo_items",
			shouldPass: true,
			statusCode: StatusOK,
			todoItemsResult: func() []models.TodoItem {
				todoItems := []models.TodoItem{
					{Title: "item1", Description: ""},
					{Title: "item2", Description: "desc2"},
				}

				todoItems[0].ID = 1
				todoItems[1].ID = 2
				return todoItems
			}(),
		},
		{
			title:      "Get all todo items, wrong list_id",
			method:     "GET",
			path:       "/todo_lists/a/todo_items",
			route:      "/todo_lists/{list_id}/todo_items",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Get all todo items, non-existent list_id",
			method:     "GET",
			path:       "/todo_lists/2/todo_items",
			route:      "/todo_lists/{list_id}/todo_items",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	todoItemController := NewTodoItemController(&mock.TodoItemServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoItemResult(t, tc, todoItemController.GetAll)
		})
	}
}

func TestTodoItemController_Post(t *testing.T) {
	tests := []todoItemTest{
		{
			title:      "Post todo item",
			method:     "POST",
			path:       "/todo_lists/1/todo_items",
			route:      "/todo_lists/{list_id}/todo_items",
			shouldPass: true,
			statusCode: StatusCreated,
			body:       []byte(`{"ID": 1, "Title": "item1", "Description": ""}`),
			todoItemResult: func() models.TodoItem {
				todoItem := models.TodoItem{Title: "item1", Description: ""}
				todoItem.ID = 1
				return todoItem
			}(),
		},
		{
			title:      "Post todo item, wrong list_id",
			method:     "POST",
			path:       "/todo_lists/a/todo_items",
			route:      "/todo_lists/{list_id}/todo_items",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte(`{"ID": 1, "Title": "item1", "Description": ""}`),
		},
		{
			title:      "Post todo item, wrong body",
			method:     "POST",
			path:       "/todo_lists/1/todo_items",
			route:      "/todo_lists/{list_id}/todo_items",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte{},
		},
		{
			title:      "Post todo item, internal error",
			method:     "POST",
			path:       "/todo_lists/2/todo_items",
			route:      "/todo_lists/{list_id}/todo_items",
			shouldPass: false,
			statusCode: StatusInternalServerError,
			body:       []byte(`{"ID": 1, "Title": "item1", "Description": ""}`),
		},
	}

	todoItemController := NewTodoItemController(&mock.TodoItemServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoItemResult(t, tc, todoItemController.Post)
		})
	}
}

func TestTodoItemController_GetSingle(t *testing.T) {
	tests := []todoItemTest{
		{
			title:      "Get todo item",
			method:     "GET",
			path:       "/todo_items/1",
			route:      "/todo_items/{id}",
			shouldPass: true,
			statusCode: StatusOK,
			todoItemResult: func() models.TodoItem {
				todoItem := models.TodoItem{Title: "item1", Description: ""}
				todoItem.ID = 1
				return todoItem
			}(),
		},
		{
			title:      "Get todo item, wrong id",
			method:     "GET",
			path:       "/todo_items/a",
			route:      "/todo_items/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Get todo item, non-existent id",
			method:     "GET",
			path:       "/todo_items/2",
			route:      "/todo_items/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	todoItemController := NewTodoItemController(&mock.TodoItemServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoItemResult(t, tc, todoItemController.GetSingle)
		})
	}
}

func TestTodoItemController_Put(t *testing.T) {
	tests := []todoItemTest{
		{
			title:      "Put todo item",
			method:     "PUT",
			path:       "/todo_items/1",
			route:      "/todo_items/{id}",
			shouldPass: true,
			statusCode: StatusOK,
			body:       []byte(`{"ID": 1, "Title": "item1", "Description": ""}`),
			todoItemResult: func() models.TodoItem {
				todoItem := models.TodoItem{Title: "item1", Description: ""}
				todoItem.ID = 1
				return todoItem
			}(),
		},
		{
			title:      "Put todo item, wrong id",
			method:     "PUT",
			path:       "/todo_items/a",
			route:      "/todo_items/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte(`{"ID": 1, "Title": "item1", "Description": ""}`),
		},
		{
			title:      "Put todo item, wrong body",
			method:     "PUT",
			path:       "/todoItems/1",
			route:      "/todoItems/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte{},
		},
		{
			title:      "Put todo item, non-existent id",
			method:     "PUT",
			path:       "/todo_items/2",
			route:      "/todo_items/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
			body:       []byte(`{"ID": 1, "Title": "item1", "Description": ""}`),
		},
	}

	todoItemController := NewTodoItemController(&mock.TodoItemServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoItemResult(t, tc, todoItemController.Put)
		})
	}
}

func TestTodoItemController_Delete(t *testing.T) {
	tests := []todoItemTest{
		{
			title:      "Delete todo item",
			method:     "DELETE",
			path:       "/todo_items/1",
			route:      "/todo_items/{id}",
			shouldPass: true,
			statusCode: StatusNoContent,
		},
		{
			title:      "Delete todo item, wrong id",
			method:     "DELETE",
			path:       "/todo_items/a",
			route:      "/todo_items/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Delete todo item, non-existent id",
			method:     "DELETE",
			path:       "/todo_items/2",
			route:      "/todo_items/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	todoItemController := NewTodoItemController(&mock.TodoItemServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoItemResult(t, tc, todoItemController.Delete)
		})
	}
}
