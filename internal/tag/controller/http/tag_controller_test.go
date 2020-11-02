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

type tagTest struct {
	title      string
	method     string
	path       string
	route      string
	body       []byte
	shouldPass bool
	statusCode int
	tagResult  models.Tag
	tagsResult []models.Tag
}

type tagServiceMock struct{}

func (s *tagServiceMock) GetAll(itemID uint) ([]models.Tag, error) {
	if itemID != 1 {
		return []models.Tag{}, errors.New("err")
	}

	tags := []models.Tag{{Text: "tag1"}, {Text: "tag2"}}
	tags[0].ID = 1
	tags[1].ID = 2
	return tags, nil
}

func (s *tagServiceMock) GetSingle(id uint) (models.Tag, error) {
	if id != 1 {
		return models.Tag{}, errors.New("not found")
	}

	tag := models.Tag{Text: "tag1"}
	tag.ID = 1
	return tag, nil
}

func (s *tagServiceMock) Create(itemID uint, tag *models.Tag) error {
	if itemID != 1 {
		return errors.New("err")
	}
	return nil
}

func (s *tagServiceMock) Update(id uint, tagData *models.Tag) (models.Tag, error) {
	if id != 1 {
		return models.Tag{}, errors.New("err")
	}

	tag := models.Tag{Text: "tag1"}
	tag.ID = 1
	return tag, nil
}

func (s *tagServiceMock) Remove(itemID uint, tagID uint) error {
	if itemID != 1 {
		return errors.New("err")
	}
	return nil
}

func (s *tagServiceMock) Delete(id uint) error {
	if id != 1 {
		return errors.New("err")
	}
	return nil
}

func compareTags(t *testing.T, expect models.Tag, actual models.Tag) {
	assert.Equal(t, expect.ID, actual.ID)
	assert.Equal(t, expect.Text, actual.Text)
}

func testTagResult(t *testing.T, tc tagTest, controller func(ResponseWriter, *Request)) {
	w, r := test.NewRequest(tc.method, tc.path, tc.body)
	test.MakeRequest(tc.route, controller, w, r)
	assert.Equal(t, tc.statusCode, w.Code)

	if tc.shouldPass {
		if len(tc.tagsResult) != 0 {
			var result []models.Tag
			json.NewDecoder(w.Body).Decode(&result)
			compareTags(t, tc.tagsResult[0], result[0])
			compareTags(t, tc.tagsResult[1], result[1])
		} else {
			var result models.Tag
			json.NewDecoder(w.Body).Decode(&result)
			compareTags(t, tc.tagResult, result)
		}
	}
}

func TestTagController_GetAll(t *testing.T) {
	tests := []tagTest{
		{
			title:      "Get all tags",
			method:     "GET",
			path:       "/todo_items/1/tags",
			route:      "/todo_items/{item_id}/tags",
			shouldPass: true,
			statusCode: StatusOK,
			tagsResult: func() []models.Tag {
				tags := []models.Tag{{Text: "tag1"}, {Text: "tag2"}}
				tags[0].ID = 1
				tags[1].ID = 2
				return tags
			}(),
		},
		{
			title:      "Get all tags, wrong item_id",
			method:     "GET",
			path:       "/todo_items/a/tags",
			route:      "/todo_items/{item_id}/tags",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Get all tags, non-existent item_id",
			method:     "GET",
			path:       "/todo_items/2/tags",
			route:      "/todo_items/{item_id}/tags",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	tagController := NewTagController(&tagServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTagResult(t, tc, tagController.GetAll)
		})
	}
}

func TestTagController_Post(t *testing.T) {
	tests := []tagTest{
		{
			title:      "Post tag",
			method:     "POST",
			path:       "/todo_items/1/tags",
			route:      "/todo_items/{item_id}/tags",
			shouldPass: true,
			statusCode: StatusCreated,
			body:       []byte(`{"ID": 1, "Text": "tag1"}`),
			tagResult: func() models.Tag {
				tag := models.Tag{Text: "tag1"}
				tag.ID = 1
				return tag
			}(),
		},
		{
			title:      "Post tag, wrong item_id",
			method:     "POST",
			path:       "/todo_items/a/tags",
			route:      "/todo_items/{item_id}/tags",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte(`{"ID": 1, "Text": "tag1"}`),
		},
		{
			title:      "Post tag, wrong body",
			method:     "POST",
			path:       "/todo_items/1/tags",
			route:      "/todo_items/{item_id}/tags",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte{},
		},
		{
			title:      "Post tag, internal error",
			method:     "POST",
			path:       "/todo_items/2/tags",
			route:      "/todo_items/{item_id}/tags",
			shouldPass: false,
			statusCode: StatusInternalServerError,
			body:       []byte(`{"ID": 1, "Text": "tag1"}`),
		},
	}

	tagController := NewTagController(&tagServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTagResult(t, tc, tagController.Post)
		})
	}
}

func TestTagController_GetSingle(t *testing.T) {
	tests := []tagTest{
		{
			title:      "Get tag",
			method:     "GET",
			path:       "/tags/1",
			route:      "/tags/{id}",
			shouldPass: true,
			statusCode: StatusOK,
			tagResult: func() models.Tag {
				tag := models.Tag{Text: "tag1"}
				tag.ID = 1
				return tag
			}(),
		},
		{
			title:      "Get tag, wrong id",
			method:     "GET",
			path:       "/tags/a",
			route:      "/tags/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Get tag, non-existent id",
			method:     "GET",
			path:       "/tags/2",
			route:      "/tags/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	tagController := NewTagController(&tagServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTagResult(t, tc, tagController.GetSingle)
		})
	}
}

func TestTagController_Put(t *testing.T) {
	tests := []tagTest{
		{
			title:      "Put tag",
			method:     "PUT",
			path:       "/tags/1",
			route:      "/tags/{id}",
			shouldPass: true,
			statusCode: StatusOK,
			body:       []byte(`{"ID": 1, "Text": "tag1"}`),
			tagResult: func() models.Tag {
				tag := models.Tag{Text: "tag1"}
				tag.ID = 1
				return tag
			}(),
		},
		{
			title:      "Put tag, wrong item_id",
			method:     "PUT",
			path:       "/tags/a",
			route:      "/tags/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte(`{"ID": 1, "Text": "tag1"}`),
		},
		{
			title:      "Put tag, wrong body",
			method:     "PUT",
			path:       "/tags/1",
			route:      "/tags/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
			body:       []byte{},
		},
		{
			title:      "Put tag, non-existent id",
			method:     "PUT",
			path:       "/tags/2",
			route:      "/tags/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
			body:       []byte(`{"ID": 1, "Text": "tag1"}`),
		},
	}

	tagController := NewTagController(&tagServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTagResult(t, tc, tagController.Put)
		})
	}
}

func TestTagController_Remove(t *testing.T) {
	tests := []tagTest{
		{
			title:      "Remove tag",
			method:     "DELETE",
			path:       "/todo_items/1/tags/1",
			route:      "/todo_items/{item_id}/tags/{tag_id}",
			shouldPass: true,
			statusCode: StatusNoContent,
		},
		{
			title:      "Remove tag, wrong item_id",
			method:     "DELETE",
			path:       "/todo_items/a/tags/1",
			route:      "/todo_items/{item_id}/tags/{tag_id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Remove tag, wrong tag_id",
			method:     "DELETE",
			path:       "/todo_items/1/tags/a",
			route:      "/todo_items/{item_id}/tags/{tag_id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Remove tag, non-existent item_id or tag_id",
			method:     "DELETE",
			path:       "/todo_items/2/tags/2",
			route:      "/todo_items/{item_id}/tags/{tag_id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	tagController := NewTagController(&tagServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTagResult(t, tc, tagController.Remove)
		})
	}
}

func TestTagController_Delete(t *testing.T) {
	tests := []tagTest{
		{
			title:      "Delete tag",
			method:     "DELETE",
			path:       "/tags/1",
			route:      "/tags/{id}",
			shouldPass: true,
			statusCode: StatusNoContent,
		},
		{
			title:      "Delete tag, wrong id",
			method:     "DELETE",
			path:       "/tags/a",
			route:      "/tags/{id}",
			shouldPass: false,
			statusCode: StatusBadRequest,
		},
		{
			title:      "Delete tag, non-existent id",
			method:     "DELETE",
			path:       "/tags/2",
			route:      "/tags/{id}",
			shouldPass: false,
			statusCode: StatusNotFound,
		},
	}

	tagController := NewTagController(&tagServiceMock{})
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			testTagResult(t, tc, tagController.Delete)
		})
	}
}