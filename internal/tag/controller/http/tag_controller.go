package http

import (
	"encoding/json"
	"net/http"

	"github.com/danikg/go-todo-rest-api/internal/models"
	"github.com/danikg/go-todo-rest-api/internal/utils/response"
	"github.com/danikg/go-todo-rest-api/internal/utils/route"
)

// TagController ...
type TagController struct {
	TagService models.ITagService
}

// NewTagController ...
func NewTagController(tagService models.ITagService) *TagController {
	return &TagController{TagService: tagService}
}

// GetAll returns all tags by todo item id
func (c *TagController) GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		itemID uint
		tags   []models.Tag
		err    error
	)

	if itemID, err = route.GetRouteVar(r, "item_id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if tags, err = c.TagService.GetAll(itemID); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, tags, 0)
}

// Post creates a new tag
func (c *TagController) Post(w http.ResponseWriter, r *http.Request) {
	var (
		tag    models.Tag
		itemID uint
		err    error
	)

	if itemID, err = route.GetRouteVar(r, "item_id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&tag); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.TagService.Create(itemID, &tag); err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponse(w, tag, http.StatusCreated)
}

// GetSingle returns a single tag by id
func (c *TagController) GetSingle(w http.ResponseWriter, r *http.Request) {
	var (
		id  uint
		tag models.Tag
		err error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if tag, err = c.TagService.GetSingle(id); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, tag, 0)
}

// Put updates the tag by id
func (c *TagController) Put(w http.ResponseWriter, r *http.Request) {
	var (
		id      uint
		tagData models.Tag
		tag     models.Tag
		err     error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&tagData); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if tag, err = c.TagService.Update(id, &tagData); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	response.SendResponse(w, tag, 0)
}

// Remove removes the tag from the todo item
func (c *TagController) Remove(w http.ResponseWriter, r *http.Request) {
	var (
		itemID uint
		tagID  uint
		err    error
	)

	if itemID, err = route.GetRouteVar(r, "item_id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if tagID, err = route.GetRouteVar(r, "tag_id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.TagService.Remove(itemID, tagID); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete removes the tag from the db
func (c *TagController) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		id  uint
		err error
	)

	if id, err = route.GetRouteVar(r, "id"); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = c.TagService.Delete(id); err != nil {
		response.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
