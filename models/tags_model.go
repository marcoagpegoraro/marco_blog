package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Id        uint   `gorm:"autoIncrement:false"`
	Name      string `gorm:"primaryKey"`
	Posts     []Post `gorm:"many2many:posts_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t Tag) String() string {
	return fmt.Sprintf("Tag: Id: %d - Name: %s - Posts: %s - CreatedAt: %s - UpdatedAt: %s", t.Id, t.Name, t.Posts, t.CreatedAt, t.UpdatedAt)
}
