package models

import (
	"time"

	"github.com/marcoagpegoraro/marco_blog/enum"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id        uint `gorm:"primaryKey"`
	Title     string
	Subtitle  string
	Body      string
	IsDraft   bool
	Language  enum.Language
	Tags      []Tag `gorm:"many2many:posts_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}