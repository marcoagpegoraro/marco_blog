package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID    uint
	Title string
	Body  string
}
