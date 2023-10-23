package models

import (
	"fmt"

	"gorm.io/gorm"
)

type PostsTags struct {
	gorm.Model
	TagName string
	PostId  uint
}

func (p PostsTags) String() string {
	return fmt.Sprintf("PostsTags: TagName: %s - PostId: %d", p.TagName, p.PostId)
}
