package repository

import (
	"fmt"
	"hacktiv8-lc3-p2/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Register(request entity.UserRegisterPayload) (*entity.UserResponse, error) {
	newUser := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := ur.DB.Where("email = ?", newUser.Email).First(&newUser).Error; err == nil {
		return nil, fmt.Errorf("409 | Email already exists")
	}

	if err := ur.DB.Create(&newUser).Error; err != nil {
		return nil, fmt.Errorf("500 | Internal server error")
	}

	user := entity.UserResponse{
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return &user, nil
}
