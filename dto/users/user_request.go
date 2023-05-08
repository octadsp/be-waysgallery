package usersdto

type UserUpdateRequest struct {
	FullName string `json:"fullName" form:"fullName"`
	Greeting string `json:"greeting" form:"greeting"`
	Avatar   string `json:"image" form:"image"`
}
