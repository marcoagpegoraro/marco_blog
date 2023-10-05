package dto

type PostRequest struct {
	PostTitle    string `json:"PostTitle" xml:"PostTitle" form:"PostTitle"`
	PostSubtitle string `json:"PostSubtitle" xml:"PostSubtitle" form:"PostSubtitle"`
	PostTags     string `json:"PostTags" xml:"PostTags" form:"PostTags"`
	PostBody     string `json:"PostBody" xml:"PostBody" form:"PostBody"`
}
