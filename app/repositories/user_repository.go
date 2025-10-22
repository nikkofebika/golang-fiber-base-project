package repositories

import (
	"golang-fiber-base-project/app/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db}
}

func (repository *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := repository.DB.Find(&users).Error
	return users, err
}

func (repository *userRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	err := repository.DB.Take(&user, id).Error
	return user, err
}

func (repository *userRepository) Create(user *models.User) error {
	return repository.DB.Create(user).Error
}

func (repository *userRepository) Update(user *models.User) error {
	return repository.DB.Save(user).Error
}

func (repository *userRepository) Delete(id uint) error {
	return repository.DB.Delete(&models.User{}, id).Error
}
