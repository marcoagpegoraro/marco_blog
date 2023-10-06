package dto

import "github.com/marcoagpegoraro/marco_blog/enum"

type PostRequest struct {
	PostId       uint          `json:"PostId" xml:"PostId" form:"PostId"`
	PostTitle    string        `json:"PostTitle" xml:"PostTitle" form:"PostTitle"`
	PostSubtitle string        `json:"PostSubtitle" xml:"PostSubtitle" form:"PostSubtitle"`
	PostTags     string        `json:"PostTags" xml:"PostTags" form:"PostTags"`
	PostLanguage enum.Language `json:"PostLanguage" xml:"PostLanguage" form:"PostLanguage"`
	PostBody     string        `json:"PostBody" xml:"PostBody" form:"PostBody"`
	PostIsDraft  bool          `json:"PostIsDraft" xml:"PostIsDraft" form:"PostIsDraft"`
}
