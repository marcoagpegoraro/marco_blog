package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id       uint `gorm:"primaryKey"`
	Title    string
	Subtitle string
	Body     string

	CreatedAt time.Time
	UpdatedAt time.Time
}
