package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id        uint `gorm:"primaryKey"`
	Title     string
	Subtitle  string
	Body      string
	IsDraft   bool
	Language  string
	Tags      []Tag `gorm:"many2many:posts_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
