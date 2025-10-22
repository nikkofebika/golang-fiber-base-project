package services

import (
	"errors"
	"fmt"
	"golang-fiber-base-project/app/http/requests"
	"golang-fiber-base-project/app/http/resources"
	"golang-fiber-base-project/app/models"
	"golang-fiber-base-project/app/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	FindAll() ([]resources.UserResource, error)
	FindByID(id uint) (resources.UserResource, error)
	Create(request *requests.UserCreateRequest) error
	Update(id uint, user *requests.UserUpdateRequest) error
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

func (service *userService) Create(request *requests.UserCreateRequest) error {
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		IsAdmin:  request.IsAdmin,
	}

	return service.repository.Create(&user)
}

func (service *userService) Update(id uint, request *requests.UserUpdateRequest) error {
	user, err := service.repository.FindByID(id)
	if err != nil {
		return err
	}

	fmt.Println("Update service", user)

	if request.Email != nil {
		existUser, _ := service.repository.FindByEmail(*request.Email)
		if existUser != nil && existUser.ID != user.ID {
			return err
		}

		user.Email = *request.Email
	}

	if request.Name != nil {
		user.Name = *request.Name
	}

	if request.Password != nil {
		bytes, err := bcrypt.GenerateFromPassword([]byte(*request.Password), bcrypt.DefaultCost)
		fmt.Println("Update service bytes", string(bytes))
		if err != nil {
			return err
		}

		user.Password = string(bytes)
	}

	if request.IsAdmin != nil {
		user.IsAdmin = *request.IsAdmin
	}
	fmt.Println("Update service SELESAI", user)

	return service.repository.Update(&user)
}

func (service *userService) Delete(id uint) error {
	user, err := service.repository.FindByID(id)
	if err != nil {
		return err
	}

	if &user == nil {
		return errors.New("user not found")
	}

	return service.repository.Delete(id)
}
