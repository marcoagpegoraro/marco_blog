package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Id    uint   `gorm:"autoIncrement:false"`
	Name  string `gorm:"primaryKey"`
	Posts []Post `gorm:"many2many:posts_tags;"`
}
