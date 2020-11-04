package webservices

import (
	"github.com/danikg/go-todo-rest-api/models"
	repos "github.com/danikg/go-todo-rest-api/repositories"
)

// TagService ...
type TagService struct {
	TagRepo      repos.ITagRepository
	TodoItemRepo repos.ITodoItemRepository
}

// NewTagService ...
func NewTagService(tagRepo repos.ITagRepository, todoItemRepo repos.ITodoItemRepository) *TagService {
	return &TagService{
		TagRepo:      tagRepo,
		TodoItemRepo: todoItemRepo,
	}
}

// GetAll returns all tags by todo item id
func (t *TagService) GetAll(itemID uint) ([]models.Tag, error) {
	todoItem, err := t.TodoItemRepo.GetSingle(itemID)
	if err != nil {
		return []models.Tag{}, err
	}
	return t.TagRepo.GetAll(&todoItem)
}

// GetSingle returns a tag by id
func (t *TagService) GetSingle(id uint) (models.Tag, error) {
	return t.TagRepo.GetSingle(id)
}

// Create creates a new tag
func (t *TagService) Create(itemID uint, tag *models.Tag) error {
	todoItem, err := t.TodoItemRepo.GetSingle(itemID)
	if err != nil {
		return err
	}
	return t.TagRepo.Create(&todoItem, tag)
}

// Update updates the tag
func (t *TagService) Update(id uint, tagData *models.Tag) (models.Tag, error) {
	return t.TagRepo.Update(id, tagData)
}

// Remove removes the tag from the todo item
func (t *TagService) Remove(itemID uint, tagID uint) error {
	todoItem, err := t.TodoItemRepo.GetSingle(itemID)
	if err != nil {
		return err
	}
	return t.TagRepo.Remove(&todoItem, tagID)
}

// Delete removes the tag from the db
func (t *TagService) Delete(id uint) error {
	return t.TagRepo.Delete(id)
}
