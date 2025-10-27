package resources

import "golang-fiber-base-project/app/models"

type UserResource struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	IsAdmin   bool    `json:"is_admin"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

func NewUserResource(user *models.User) UserResource {
	var deletedAt *string
	if user.DeletedAt.Valid {
		formatted := user.DeletedAt.Time.Format("2006-01-02 15:04:05")
		deletedAt = &formatted
	}

	return UserResource{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

func NewUserResources(users []models.User) []UserResource {
	userResources := make([]UserResource, len(users))
	for i := range users {
		userResources[i] = NewUserResource(&users[i])
	}
	return userResources

	// userResources := []UserResource{}

	// for _, user := range users {
	// 	userResource := NewUserResource(user)
	// 	userResources = append(userResources, userResource)
	// }

}
