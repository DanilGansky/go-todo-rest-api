package http

import (
	"encoding/json"
	"net/http"

	"github.com/danikg/go-todo-rest-api/internal/models"
	"github.com/danikg/go-todo-rest-api/internal/utils/response"
	"github.com/danikg/go-todo-rest-api/internal/utils/route"
)

// UserController ...
type UserController struct {
	UserService models.IUserService
}

// NewUserController ...
func NewUserController(userService models.IUserService) *UserController {
	return &UserController{UserService: userService}
}

// GetAll returns all users from the db
func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserService.GetAll()
	if err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, users, 0)
}

// Post creates a new user
func (c *UserController) Post(w http.ResponseWriter, r *http.Request) {
	var (
		user models.User
		err  error
	)

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.UserService.Create(&user); err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponse(w, user, http.StatusCreated)
}

// GetSingle returns a single user by id
func (c *UserController) GetSingle(w http.ResponseWriter, r *http.Request) {
	var (
		id   uint
		user models.User
		err  error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if user, err = c.UserService.GetSingle(id); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, user, 0)
}

// Put updates the user by id
func (c *UserController) Put(w http.ResponseWriter, r *http.Request) {
	var (
		id       uint
		userData models.User
		user     models.User
		err      error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&userData); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if user, err = c.UserService.Update(id, &userData); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, user, 0)
}

// Delete removes the user by id
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		id  uint
		err error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.UserService.Delete(id); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
