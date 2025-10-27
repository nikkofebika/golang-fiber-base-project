package services

import (
	"context"
	"golang-fiber-base-project/app/exceptions"
	"golang-fiber-base-project/app/helpers"
	"golang-fiber-base-project/app/http/requests"
	"golang-fiber-base-project/app/http/resources"
	"golang-fiber-base-project/app/models"
	"golang-fiber-base-project/app/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	FindAll(ctx context.Context) ([]resources.UserResource, error)
	FindByID(ctx context.Context, id uint) (resources.UserResource, error)
	Create(ctx context.Context, request *requests.UserCreateRequest) error
	Update(ctx context.Context, id uint, user *requests.UserUpdateRequest) error
	Delete(ctx context.Context, id uint) error
}

type userService struct {
	repository repositories.UserRepositoryInterface
}

func NewUserService(repository repositories.UserRepositoryInterface) UserServiceInterface {
	return &userService{repository}
}

func (service *userService) FindAll(ctx context.Context) ([]resources.UserResource, error) {
	users, err := service.repository.FindAll(ctx)
	if err != nil {
		return nil, exceptions.NewDatabaseException(err)
	}

	return resources.NewUserResources(users), nil
}

func (service *userService) FindByID(ctx context.Context, id uint) (resources.UserResource, error) {
	user, err := service.repository.FindByID(ctx, id)

	return resources.NewUserResource(&user), err
}

func (service *userService) Create(ctx context.Context, request *requests.UserCreateRequest) error {
	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		IsAdmin:  request.IsAdmin,
	}

	return service.repository.Create(ctx, &user)
}

func (service *userService) Update(ctx context.Context, id uint, request *requests.UserUpdateRequest) error {
	user, err := service.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if request.Email != nil {
		existUser, _ := service.repository.FindByEmail(ctx, *request.Email)
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
		if err != nil {
			return err
		}

		user.Password = string(bytes)
	}

	if request.IsAdmin != nil {
		user.IsAdmin = *request.IsAdmin
	}

	return service.repository.Update(ctx, &user)
}

func (service *userService) Delete(ctx context.Context, id uint) error {
	_, err := service.repository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return service.repository.Delete(ctx, id)
}
