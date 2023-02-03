package handler

type LoginRes struct {
	ID          uint   `json:"id" form:"id"`
	Email       string `json:"email" form:"email"`
	Fullname    string `json:"fullname" form:"fullname"`
	Username    string `json:"username" form:"username"`
	Gender      string `json:"gender" form:"gender"`
	Avatar      string `json:"avatar" form:"avatar"`
	Sampul      string `json:"sampul" form:"sampul"`
	DateOfBirth string `json:"dateofbirth" form:"dateofbirth"`
	Bio         string `json:"bio" form:"bio"`
}
