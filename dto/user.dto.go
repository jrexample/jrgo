package dto

type UserResponseBody struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type UpdateUserNameRequestBody struct {
	Name string `json:"name" validate:"required"`
}
