package services

import (
	"context"
	"errors"
	"golang-fiber-base-project/app/helpers"
	"golang-fiber-base-project/app/http/requests"
	"golang-fiber-base-project/app/models"
	"golang-fiber-base-project/app/repositories"
	"golang-fiber-base-project/config"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, email, password string) (*helpers.JwtTokenDetail, error)
	Register(ctx context.Context, request *requests.RegisterRequest) error
	// RefreshToken(ctx context.Context) error
}

type AuthService struct {
	userRepository repositories.UserRepositoryInterface
	cfg            *config.AppConfig
}

func NewAuthService(userRepository repositories.UserRepositoryInterface, cfg *config.AppConfig) AuthServiceInterface {
	return &AuthService{userRepository, cfg}
}

func (service *AuthService) Login(ctx context.Context, email, password string) (*helpers.JwtTokenDetail, error) {
	user, err := service.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return helpers.GenerateToken(user, service.cfg.JWTSecret)
}

func (service *AuthService) Register(ctx context.Context, request *requests.RegisterRequest) error {
	_, err := service.userRepository.FindByEmail(ctx, request.Email)
	if err == nil {
		return err
	}

	password, err := helpers.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	return service.userRepository.Register(ctx, user)
}
