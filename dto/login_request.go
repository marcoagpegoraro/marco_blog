package dto

type LoginRequest struct {
	Username string `json:"Username" xml:"Username" form:"Username"`
	Password string `json:"Password" xml:"Password" form:"Password"`
}
