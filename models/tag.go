package models

import "gorm.io/gorm"

// Tag model represents a tag in db
type Tag struct {
	gorm.Model
	Text string
}
