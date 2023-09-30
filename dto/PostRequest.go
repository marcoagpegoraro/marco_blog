package dto

type PostRequest struct {
	PostTitle    string `json:"PostTitle" xml:"PostTitle" form:"PostTitle"`
	PostSubtitle string `json:"PostSubtitle" xml:"PostSubtitle" form:"PostSubtitle"`
	PostBody     string `json:"PostBody" xml:"PostBody" form:"PostBody"`
}
