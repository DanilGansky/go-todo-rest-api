package http

import (
	"encoding/json"
	"net/http"

	"github.com/danikg/go-todo-rest-api/models"
	"github.com/danikg/go-todo-rest-api/services"
	"github.com/danikg/go-todo-rest-api/utils/response"
	"github.com/danikg/go-todo-rest-api/utils/route"
)

// TodoItemController ...
type TodoItemController struct {
	TodoItemService services.ITodoItemService
}

// NewTodoItemController ...
func NewTodoItemController(todoItemService services.ITodoItemService) *TodoItemController {
	return &TodoItemController{TodoItemService: todoItemService}
}

// GetAll returns all todo items by todo list id
func (c *TodoItemController) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		listID    uint
		todoItems []models.TodoItem
		err       error
	)

	if listID, err = route.GetRouteVar(r, "list_id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if todoItems, err = c.TodoItemService.GetAll(listID); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, todoItems, 0)
}

// Post creates a new todo item
func (c *TodoItemController) Post(w http.ResponseWriter, r *http.Request) {
	var (
		todoItem models.TodoItem
		listID   uint
		err      error
	)

	if listID, err = route.GetRouteVar(r, "list_id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&todoItem); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.TodoItemService.Create(listID, &todoItem); err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponse(w, todoItem, http.StatusCreated)
}

// GetSingle returns a single todo item by id
func (c *TodoItemController) GetSingle(w http.ResponseWriter, r *http.Request) {
	var (
		id       uint
		todoItem models.TodoItem
		err      error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if todoItem, err = c.TodoItemService.GetSingle(id); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, todoItem, 0)
}

// Put updates the todo item by id
func (c *TodoItemController) Put(w http.ResponseWriter, r *http.Request) {
	var (
		id           uint
		todoItemData models.TodoItem
		todoItem     models.TodoItem
		err          error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&todoItemData); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if todoItem, err = c.TodoItemService.Update(id, &todoItemData); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, todoItem, 0)
}

// Delete removes the todo item by id
func (c *TodoItemController) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		id  uint
		err error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.TodoItemService.Delete(id); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
