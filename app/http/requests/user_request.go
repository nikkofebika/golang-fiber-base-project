package requests

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	IsAdmin  bool   `json:"is_admin" validate:"required,boolean"`
}
