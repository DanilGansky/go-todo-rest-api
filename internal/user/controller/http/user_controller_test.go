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

type userTest struct {
	title       string
	method      string
	path        string
	route       string
	body        []byte
	shouldPass  bool
	statusCode  int
	userResult  models.User
	usersResult []models.User
}

type userServiceMock struct {
	generateErr bool
}

func (s *userServiceMock) GetAll() ([]models.User, error) {
	if s.generateErr {
		return []models.User{}, errors.New("err")
	}

	users := []models.User{
		models.User{Username: "user1"},
		models.User{Username: "user2"},
	}

	users[0].ID = 1
	users[1].ID = 2
	return users, nil
}

func (s *userServiceMock) GetSingle(id uint) (models.User, error) {
	if id != 1 {
		return models.User{}, errors.New("err")
	}

	user := models.User{Username: "user1"}
	user.ID = 1
	return user, nil
}

func (s *userServiceMock) Create(user *models.User) error {
	if s.generateErr {
		return errors.New("err")
	}
	return nil
}

func (s *userServiceMock) Update(id uint, userData *models.User) (models.User, error) {
	if id != 1 {
		return models.User{}, errors.New("err")
	}

	user := models.User{Username: "user1"}
	user.ID = 1
	return user, nil
}

func (s *userServiceMock) Delete(id uint) error {
	if id != 1 {
		return errors.New("err")
	}
	return nil
}

func compareUsers(t *testing.T, user1 models.User, user2 models.User) {
	assert.Equal(t, user1.ID, user2.ID)
	assert.Equal(t, user1.Username, user2.Username)
}

func testUserResult(t *testing.T, tc userTest, controller func(ResponseWriter, *Request)) {
	w, r := test.NewRequest(tc.method, tc.path, tc.body)
	test.MakeRequest(tc.route, controller, w, r)
	assert.Equal(t, tc.statusCode, w.Code)

	if tc.shouldPass {
		if len(tc.usersResult) != 0 {
			var result []models.User
			json.NewDecoder(w.Body).Decode(&result)
			compareUsers(t, tc.usersResult[0], result[0])
			compareUsers(t, tc.usersResult[1], result[1])
		} else {
			var result models.User
			json.NewDecoder(w.Body).Decode(&result)
			compareUsers(t, tc.userResult, result)
		}
	}
}

func TestUserController_GetAll(t *testing.T) {
	tests := []userTest{
		{
			title:      "Get all users",
			method:     "GET",
			path:       "/users",
			route:      "/users",
			shouldPass: true,
			statusCode: StatusOK,
			usersResult: func() []models.User {
				users := []models.User{{Username: "user1"}, {Username: "user2"}}
				users[0].ID = 1
				users[1].ID = 2
				return users
			}(),
		},
		{
			title:      "Get all users, internal error",
			method:     "GET",
			path:       "/users",
			route:      "/users",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	userController := NewUserController(&userServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			if !tc.shouldPass {
				userController = NewUserController(&userServiceMock{true})
			}

			testUserResult(t, tc, userController.GetAll)
		})
	}
}

func TestUserController_Post(t *testing.T) {
	tests := []userTest{
		{
			title:      "Post user",
			method:     "POST",
			path:       "/users",
			route:      "/users",
			shouldPass: true,
			statusCode: StatusCreated,
			body:       []byte(`{"ID": 1, "Username": "user1"}`),
			userResult: func() models.User {
				user := models.User{Username: "user1"}
				user.ID = 1
				return user
			}(),
		},
		{
			title:      "Post user, wrong body",
			method:     "POST",
			path:       "/users",
			route:      "/users",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte{},
		},
		{
			title:      "Post user, internal error",
			method:     "POST",
			path:       "/users",
			route:      "/users",
			shouldPass: false,
			statusCode: StatusInternalServerError,
			body:       []byte(`{"ID": 1, "Username": "user1"}`),
		},
	}

	userController := NewUserController(&userServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			if !tc.shouldPass {
				userController = NewUserController(&userServiceMock{true})
			}

			testUserResult(t, tc, userController.Post)
		})
	}
}

func TestUserController_GetSingle(t *testing.T) {
	tests := []userTest{
		{
			title:      "Get user",
			method:     "GET",
			path:       "/users/1",
			route:      "/users/{id}",
			shouldPass: true,
			statusCode: StatusOK,
			userResult: func() models.User {
				todoList := models.User{Username: "user1"}
				todoList.ID = 1
				return todoList
			}(),
		},
		{
			title:      "Get user, wrong id",
			method:     "GET",
			path:       "/users/a",
			route:      "/users/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Get user, non-existent id",
			method:     "GET",
			path:       "/users/2",
			route:      "/users/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	userController := NewUserController(&userServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testUserResult(t, tc, userController.GetSingle)
		})
	}
}

func TestUserController_Put(t *testing.T) {
	tests := []userTest{
		{
			title:      "Put user",
			method:     "PUT",
			path:       "/users/1",
			route:      "/users/{id}",
			shouldPass: true,
			statusCode: StatusOK,
			body:       []byte(`{"ID": 1, "Username": "user1"}`),
			userResult: func() models.User {
				todoList := models.User{Username: "user1"}
				todoList.ID = 1
				return todoList
			}(),
		},
		{
			title:      "Put user, wrong id",
			method:     "PUT",
			path:       "/users/a",
			route:      "/users/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte(`{"ID": 1, "Username": "user1"}`),
		},
		{
			title:      "Put user, wrong body",
			method:     "PUT",
			path:       "/users/1",
			route:      "/users/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte{},
		},
		{
			title:      "Put user, non-existent id",
			method:     "PUT",
			path:       "/users/2",
			route:      "/users/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
			body:       []byte(`{"ID": 1, "Username": "user1"}`),
		},
	}

	userController := NewUserController(&userServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testUserResult(t, tc, userController.Put)
		})
	}
}

func TestUserController_Delete(t *testing.T) {
	tests := []userTest{
		{
			title:      "Delete user",
			method:     "DELETE",
			path:       "/users/1",
			route:      "/users/{id}",
			shouldPass: true,
			statusCode: StatusNoContent,
		},
		{
			title:      "Delete user, wrong id",
			method:     "DELETE",
			path:       "/users/a",
			route:      "/users/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Delete user, non-existent id",
			method:     "DELETE",
			path:       "/users/2",
			route:      "/users/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	userController := NewUserController(&userServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testUserResult(t, tc, userController.Delete)
		})
	}
}
