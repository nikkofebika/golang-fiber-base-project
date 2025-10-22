package services

import (
	"golang-fiber-base-project/app/http/requests"
	"golang-fiber-base-project/app/http/resources"
	"golang-fiber-base-project/app/models"
	"golang-fiber-base-project/app/repositories"
)

type UserServiceInterface interface {
	FindAll() ([]resources.UserResource, error)
	FindByID(id uint) (resources.UserResource, error)
	Create(request requests.UserCreateRequest) error
	Update(user *models.User) error
	Delete(id uint) error
}

type userService struct {
	repository repositories.UserRepositoryInterface
}

func NewUserService(repository repositories.UserRepositoryInterface) UserServiceInterface {
	return &userService{repository}
}

func (service *userService) FindAll() ([]resources.UserResource, error) {
	users, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return resources.ToUserResources(users), nil
}

func (service *userService) FindByID(id uint) (resources.UserResource, error) {
	user, err := service.repository.FindByID(id)

	return resources.ToUserResource(&user), err
}

func (service *userService) Create(request requests.UserCreateRequest) error {
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		IsAdmin:  request.IsAdmin,
	}

	return service.repository.Create(&user)
}

func (service *userService) Update(user *models.User) error {
	return service.repository.Update(user)
}

func (service *userService) Delete(id uint) error {
	return service.repository.Delete(id)
}
