package repositories

import (
	"context"
	"fmt"
	"golang-fiber-base-project/app/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAll(ctx context.Context) ([]models.User, error)
	FindByID(ctx context.Context, id uint) (models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
	Register(ctx context.Context, user *models.User) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db}
}

func (repository *userRepository) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := repository.DB.Find(&users).Error
	return users, err
}

func (repository *userRepository) FindByID(ctx context.Context, id uint) (models.User, error) {
	var user models.User
	err := repository.DB.Take(&user, id).Error
	return user, err
}

func (repository *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := repository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *userRepository) Create(ctx context.Context, user *models.User) error {
	return repository.DB.Create(user).Error
}

func (repository *userRepository) Update(ctx context.Context, user *models.User) error {
	fmt.Println("Update repo", user)
	return repository.DB.Save(user).Error
}

func (repository *userRepository) Delete(ctx context.Context, id uint) error {
	return repository.DB.Delete(&models.User{}, id).Error
}

func (repository *userRepository) Register(ctx context.Context, user *models.User) error {
	return repository.DB.Create(user).Error
}
