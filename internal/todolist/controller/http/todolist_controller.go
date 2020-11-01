package http

import (
	"encoding/json"
	"net/http"

	"github.com/danikg/go-todo-rest-api/internal/models"
	"github.com/danikg/go-todo-rest-api/internal/utils/response"
	"github.com/danikg/go-todo-rest-api/internal/utils/route"
)

// TodoListController ...
type TodoListController struct {
	todoListService models.ITodoListService
}

// NewTodoListController ...
func NewTodoListController(todoListService models.ITodoListService) *TodoListController {
	return &TodoListController{todoListService: todoListService}
}

// GetAll returns all todo lists by user id
func (c *TodoListController) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		todoLists []models.TodoList
		userID    uint
		err       error
	)

	if userID, err = route.GetRouteVar(r, "user_id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if todoLists, err = c.todoListService.GetAll(userID); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, todoLists, 0)
}

// Post creates a new todo list
func (c *TodoListController) Post(w http.ResponseWriter, r *http.Request) {
	var (
		todoList models.TodoList
		userID   uint
		err      error
	)

	if userID, err = route.GetRouteVar(r, "user_id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&todoList); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.todoListService.Create(userID, &todoList); err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponse(w, todoList, http.StatusCreated)
}

// GetSingle returns a single todo list by id
func (c *TodoListController) GetSingle(w http.ResponseWriter, r *http.Request) {
	var (
		id       uint
		todoList models.TodoList
		err      error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if todoList, err = c.todoListService.GetSingle(id); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, todoList, 0)
}

// Put updates the todo list by id
func (c *TodoListController) Put(w http.ResponseWriter, r *http.Request) {
	var (
		id           uint
		todoListData models.TodoList
		todoList     models.TodoList
		err          error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&todoListData); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	todoList, err = c.todoListService.Update(id, &todoListData)
	if err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, todoList, 0)
}

// Delete removes the todo list by id
func (c *TodoListController) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		id  uint
		err error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.todoListService.Delete(id); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
