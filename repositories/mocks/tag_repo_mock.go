package mocks

import (
	"errors"

	"github.com/danikg/go-todo-rest-api/models"
)

// TagRepositoryMock ...
type TagRepositoryMock struct{}

// GetAll ...
func (s *TagRepositoryMock) GetAll(todoItem *models.TodoItem) ([]models.Tag, error) {
	if todoItem.ID != 1 {
		return []models.Tag{}, errors.New("err")
	}

	tags := []models.Tag{{Text: "tag1"}, {Text: "tag2"}}
	tags[0].ID = 1
	tags[1].ID = 2
	return tags, nil
}

// GetSingle ...
func (s *TagRepositoryMock) GetSingle(id uint) (models.Tag, error) {
	if id != 1 {
		return models.Tag{}, errors.New("not found")
	}

	tag := models.Tag{Text: "tag1"}
	tag.ID = 1
	return tag, nil
}

// Create ...
func (s *TagRepositoryMock) Create(todoItem *models.TodoItem, tag *models.Tag) error {
	if todoItem.ID != 1 {
		return errors.New("err")
	}
	return nil
}

// Update ...
func (s *TagRepositoryMock) Update(id uint, tagData *models.Tag) (models.Tag, error) {
	if id != 1 {
		return models.Tag{}, errors.New("err")
	}

	tag := models.Tag{Text: "tag1"}
	tag.ID = 1
	return tag, nil
}

// Remove ...
func (s *TagRepositoryMock) Remove(todoItem *models.TodoItem, tagID uint) error {
	if tagID != 1 {
		return errors.New("err")
	}
	return nil
}

// Delete ...
func (s *TagRepositoryMock) Delete(id uint) error {
	if id != 1 {
		return errors.New("err")
	}
	return nil
}
