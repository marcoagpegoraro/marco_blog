package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string `gorm:"primaryKey"`
	Posts []Post `gorm:"many2many:posts_tags;"`
}
