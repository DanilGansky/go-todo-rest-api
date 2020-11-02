package http

import (
	"encoding/json"
	"errors"
	. "net/http"
	"testing"

	"github.com/danikg/go-todo-rest-api/internal/models"
	"github.com/danikg/go-todo-rest-api/internal/utils/test"
	"github.com/stretchr/testify/assert"
)

type todoListTest struct {
	title           string
	method          string
	path            string
	route           string
	body            []byte
	shouldPass      bool
	statusCode      int
	todoListResult  models.TodoList
	todoListsResult []models.TodoList
}

type todoListServiceMock struct{}

func (s *todoListServiceMock) GetAll(userID uint) ([]models.TodoList, error) {
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

func (s *todoListServiceMock) GetSingle(id uint) (models.TodoList, error) {
	if id != 1 {
		return models.TodoList{}, errors.New("not found")
	}

	todoList := models.TodoList{Name: "list1", UserID: 1}
	todoList.ID = 1
	return todoList, nil
}

func (s *todoListServiceMock) Create(userID uint, todoList *models.TodoList) error {
	if userID != 1 {
		return errors.New("err")
	}
	return nil
}

func (s *todoListServiceMock) Update(id uint, todoListData *models.TodoList) (models.TodoList, error) {
	if id != 1 {
		return models.TodoList{}, errors.New("err")
	}

	todoList := models.TodoList{Name: "list1", UserID: 1}
	todoList.ID = 1
	return todoList, nil
}

func (s *todoListServiceMock) Delete(id uint) error {
	if id != 1 {
		return errors.New("err")
	}
	return nil
}

func compareTodoLists(t *testing.T, todoList1 models.TodoList, todoList2 models.TodoList) {
	assert.Equal(t, todoList1.ID, todoList2.ID)
	assert.Equal(t, todoList1.Name, todoList2.Name)
	assert.Equal(t, todoList1.UserID, todoList2.UserID)
}

func testTodoListResult(t *testing.T, tc todoListTest, controller func(ResponseWriter, *Request)) {
	w, r := test.NewRequest(tc.method, tc.path, tc.body)
	test.MakeRequest(tc.route, controller, w, r)
	assert.Equal(t, tc.statusCode, w.Code)

	if tc.shouldPass {
		if len(tc.todoListsResult) != 0 {
			var result []models.TodoList
			json.NewDecoder(w.Body).Decode(&result)
			compareTodoLists(t, tc.todoListsResult[0], result[0])
			compareTodoLists(t, tc.todoListsResult[1], result[1])
		} else {
			var result models.TodoList
			json.NewDecoder(w.Body).Decode(&result)
			compareTodoLists(t, tc.todoListResult, result)
		}
	}
}

func TestTodoListController_GetAll(t *testing.T) {
	tests := []todoListTest{
		{
			title:      "Get all todo lists",
			method:     "GET",
			path:       "/users/1/todo_lists",
			route:      "/users/{user_id}/todo_lists",
			shouldPass: true,
			statusCode: StatusOK,
			todoListsResult: func() []models.TodoList {
				todoLists := []models.TodoList{{Name: "list1", UserID: 1}, {Name: "list2", UserID: 1}}
				todoLists[0].ID = 1
				todoLists[1].ID = 2
				return todoLists
			}(),
		},
		{
			title:      "Get all todo lists, wrong user_id",
			method:     "GET",
			path:       "/users/a/todo_lists",
			route:      "/users/{user_id}/todo_lists",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Get all todo lists, non-existent user_id",
			method:     "GET",
			path:       "/users/2/todo_lists",
			route:      "/users/{user_id}/todo_lists",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	todoListController := NewTodoListController(&todoListServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoListResult(t, tc, todoListController.GetAll)
		})
	}
}

func TestTodoListController_Post(t *testing.T) {
	tests := []todoListTest{
		{
			title:      "Post todo list",
			method:     "POST",
			path:       "/users/1/todo_lists",
			route:      "/users/{user_id}/todo_lists",
			shouldPass: true,
			statusCode: StatusCreated,
			body:       []byte(`{"ID": 1, "Name": "list1", "UserID": 1}`),
			todoListResult: func() models.TodoList {
				todoList := models.TodoList{Name: "list1", UserID: 1}
				todoList.ID = 1
				return todoList
			}(),
		},
		{
			title:      "Post todo list, wrong user_id",
			method:     "POST",
			path:       "/users/a/todo_lists",
			route:      "/users/{user_id}/todo_lists",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte(`{"ID": 1, "Name": "list1", "UserID": 1}`),
		},
		{
			title:      "Post todo list, wrong body",
			method:     "POST",
			path:       "/users/1/todo_lists",
			route:      "/users/{user_id}/todo_lists",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte{},
		},
		{
			title:      "Post todo list, internal error",
			method:     "POST",
			path:       "/users/2/todo_lists",
			route:      "/users/{user_id}/todo_lists",
			shouldPass: false,
			statusCode: StatusInternalServerError,
			body:       []byte(`{"ID": 1, "Name": "list1", "UserID": 1}`),
		},
	}

	todoListController := NewTodoListController(&todoListServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoListResult(t, tc, todoListController.Post)
		})
	}
}

func TestTodoListController_GetSingle(t *testing.T) {
	tests := []todoListTest{
		{
			title:      "Get todo list",
			method:     "GET",
			path:       "/todo_lists/1",
			route:      "/todo_lists/{id}",
			shouldPass: true,
			statusCode: StatusOK,
			todoListResult: func() models.TodoList {
				todoList := models.TodoList{Name: "list1", UserID: 1}
				todoList.ID = 1
				return todoList
			}(),
		},
		{
			title:      "Get todo list, wrong id",
			method:     "GET",
			path:       "/todo_lists/a",
			route:      "/todo_lists/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Get todo list, non-existent id",
			method:     "GET",
			path:       "/todo_lists/2",
			route:      "/todo_lists/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	todoListController := NewTodoListController(&todoListServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoListResult(t, tc, todoListController.GetSingle)
		})
	}
}

func TestTodoListController_Put(t *testing.T) {
	tests := []todoListTest{
		{
			title:      "Put todo list",
			method:     "PUT",
			path:       "/todo_lists/1",
			route:      "/todo_lists/{id}",
			shouldPass: true,
			statusCode: StatusOK,
			body:       []byte(`{"ID": 1, "Name": "list1", "UserID": 1}`),
			todoListResult: func() models.TodoList {
				todoList := models.TodoList{Name: "list1", UserID: 1}
				todoList.ID = 1
				return todoList
			}(),
		},
		{
			title:      "Put todo list, wrong id",
			method:     "PUT",
			path:       "/todo_lists/a",
			route:      "/todo_lists/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte(`{"ID": 1, "Name": "list1", "UserID": 1}`),
		},
		{
			title:      "Put todo list, wrong body",
			method:     "PUT",
			path:       "/todo_lists/1",
			route:      "/todo_lists/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte{},
		},
		{
			title:      "Put todo list, non-existent id",
			method:     "PUT",
			path:       "/todo_lists/2",
			route:      "/todo_lists/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
			body:       []byte(`{"ID": 1, "Name": "list1", "UserID": 1}`),
		},
	}

	todoListController := NewTodoListController(&todoListServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoListResult(t, tc, todoListController.Put)
		})
	}
}

func TestTodoListController_Delete(t *testing.T) {
	tests := []todoListTest{
		{
			title:      "Delete todo list",
			method:     "DELETE",
			path:       "/todo_lists/1",
			route:      "/todo_lists/{id}",
			shouldPass: true,
			statusCode: StatusNoContent,
		},
		{
			title:      "Delete todo list, wrong id",
			method:     "DELETE",
			path:       "/todo_lists/a",
			route:      "/todo_lists/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Delete todo list, non-existent id",
			method:     "DELETE",
			path:       "/todo_lists/2",
			route:      "/todo_lists/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	todoListController := NewTodoListController(&todoListServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTodoListResult(t, tc, todoListController.Delete)
		})
	}
}
