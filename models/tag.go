package models

import "gorm.io/gorm"

// Tag model represents a tag in db
type Tag struct {
	gorm.Model
	Text string
}

// ITagService ...
type ITagService interface {
	GetAll(itemID uint) ([]Tag, error)
	GetSingle(id uint) (Tag, error)
	Create(itemID uint, tag *Tag) error
	Update(id uint, tagData *Tag) (Tag, error)
	Remove(itemID uint, tagID uint) error
	Delete(id uint) error
}

// ITagRepository ...
type ITagRepository interface {
	GetAll(todoItem *TodoItem) ([]Tag, error)
	GetSingle(id uint) (Tag, error)
	Create(todoItem *TodoItem, tag *Tag) error
	Update(id uint, tagData *Tag) (Tag, error)
	Remove(todoItem *TodoItem, tagID uint) error
	Delete(id uint) error
}
