package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danikg/go-todo-rest-api/models"
	tagMocks "github.com/danikg/go-todo-rest-api/tag/repository/mock"
	tiMocks "github.com/danikg/go-todo-rest-api/todoitem/repository/mock"
)

func TestTagService_GetAll(t *testing.T) {
	tagService := NewTagService(&tagMocks.TagRepositoryMock{}, &tiMocks.TodoItemRepositoryMock{})
	tags, err := tagService.GetAll(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, tags)

	tags, err = tagService.GetAll(2)
	assert.Error(t, err)
	assert.Empty(t, tags)
}

func TestTagService_GetSingle(t *testing.T) {
	tagService := NewTagService(&tagMocks.TagRepositoryMock{}, &tiMocks.TodoItemRepositoryMock{})
	tag, err := tagService.GetSingle(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, tag)

	tag, err = tagService.GetSingle(2)
	assert.Error(t, err)
	assert.Empty(t, tag)
}

func TestTagService_Create(t *testing.T) {
	tagService := NewTagService(&tagMocks.TagRepositoryMock{}, &tiMocks.TodoItemRepositoryMock{})
	tag := models.Tag{Text: "tag"}
	tag.ID = 1

	err := tagService.Create(1, &tag)
	assert.NoError(t, err)

	err = tagService.Create(2, &tag)
	assert.Error(t, err)
}

func TestTagService_Update(t *testing.T) {
	tagService := NewTagService(&tagMocks.TagRepositoryMock{}, &tiMocks.TodoItemRepositoryMock{})
	tag := models.Tag{Text: "tag"}
	tag.ID = 1

	resultTag, err := tagService.Update(1, &tag)
	assert.NoError(t, err)
	assert.NotEmpty(t, resultTag)

	resultTag, err = tagService.Update(2, &tag)
	assert.Error(t, err)
	assert.Empty(t, &resultTag)
}

func TestTagService_Remove(t *testing.T) {
	tagService := NewTagService(&tagMocks.TagRepositoryMock{}, &tiMocks.TodoItemRepositoryMock{})
	assert.NoError(t, tagService.Remove(1, 1))
	assert.Error(t, tagService.Remove(2, 1))
	assert.Error(t, tagService.Remove(1, 2))
}

func TestTagService_Delete(t *testing.T) {
	tagService := NewTagService(&tagMocks.TagRepositoryMock{}, &tiMocks.TodoItemRepositoryMock{})
	assert.NoError(t, tagService.Delete(1))
	assert.Error(t, tagService.Delete(2))
}
