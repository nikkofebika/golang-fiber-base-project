package requests

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email,unique=users.email"`
	Password string `json:"password" validate:"required,min=6"`
	IsAdmin  bool   `json:"is_admin" validate:"required,boolean"`
}

type UserUpdateRequest struct {
	Name     *string `json:"name" validate:""`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password" validate:"min=6"`
	IsAdmin  *bool   `json:"is_admin" validate:"boolean"`
}
