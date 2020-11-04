package mocks

import (
	"errors"

	"github.com/danikg/go-todo-rest-api/models"
)

// TagServiceMock ...
type TagServiceMock struct{}

// GetAll ...
func (s *TagServiceMock) GetAll(itemID uint) ([]models.Tag, error) {
	if itemID != 1 {
		return []models.Tag{}, errors.New("err")
	}

	tags := []models.Tag{{Text: "tag1"}, {Text: "tag2"}}
	tags[0].ID = 1
	tags[1].ID = 2
	return tags, nil
}

// GetSingle ...
func (s *TagServiceMock) GetSingle(id uint) (models.Tag, error) {
	if id != 1 {
		return models.Tag{}, errors.New("not found")
	}

	tag := models.Tag{Text: "tag1"}
	tag.ID = 1
	return tag, nil
}

// Create ...
func (s *TagServiceMock) Create(itemID uint, tag *models.Tag) error {
	if itemID != 1 {
		return errors.New("err")
	}
	return nil
}

// Update ...
func (s *TagServiceMock) Update(id uint, tagData *models.Tag) (models.Tag, error) {
	if id != 1 {
		return models.Tag{}, errors.New("err")
	}

	tag := models.Tag{Text: "tag1"}
	tag.ID = 1
	return tag, nil
}

// Remove ...
func (s *TagServiceMock) Remove(itemID uint, tagID uint) error {
	if itemID != 1 {
		return errors.New("err")
	}
	return nil
}

// Delete ...
func (s *TagServiceMock) Delete(id uint) error {
	if id != 1 {
		return errors.New("err")
	}
	return nil
}
