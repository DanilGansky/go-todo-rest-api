package repository

import (
	"github.com/danikg/go-todo-rest-api/models"
	"gorm.io/gorm"
)

// TagRepository ...
type TagRepository struct {
	Conn *gorm.DB
}

// NewTagRepository ...
func NewTagRepository(conn *gorm.DB) *TagRepository {
	return &TagRepository{Conn: conn}
}

// GetAll returns all tags by todo item id
func (t *TagRepository) GetAll(todoItem *models.TodoItem) ([]models.Tag, error) {
	tags := []models.Tag{}
	return tags, t.Conn.Model(&todoItem).Association("Tags").Find(&tags)
}

// GetSingle returns a tag by id
func (t *TagRepository) GetSingle(id uint) (models.Tag, error) {
	tag := models.Tag{}
	return tag, t.Conn.First(&tag, id).Error
}

// Create creates a new tag
func (t *TagRepository) Create(todoItem *models.TodoItem, tag *models.Tag) error {
	return t.Conn.Model(todoItem).Association("Tags").Append(tag)
}

// Update updates the tag
func (t *TagRepository) Update(id uint, tagData *models.Tag) (models.Tag, error) {
	tag, err := t.GetSingle(id)
	if err != nil {
		return tag, err
	}
	return tag, t.Conn.Model(&tag).Update("text", tagData.Text).Error
}

// Remove removes the tag from the todo item
func (t *TagRepository) Remove(todoItem *models.TodoItem, tagID uint) error {
	tag, err := t.GetSingle(tagID)
	if err != nil {
		return err
	}
	return t.Conn.Model(&todoItem).Association("Tags").Delete(&tag)
}

// Delete removes the tag from the db
func (t *TagRepository) Delete(id uint) error {
	tag, err := t.GetSingle(id)
	if err != nil {
		return err
	}
	return t.Conn.Unscoped().Delete(&tag).Error
}
