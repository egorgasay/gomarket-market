package repository

import (
	"go-rest-api/internal/constants"
	"go-rest-api/internal/domains"
	"go-rest-api/internal/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domains.IRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user model.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByUsername(user *model.User, username string) error {
	if err := ur.db.Where("username=?", username).Limit(3).Find(&user).Error; err != nil {
		return constants.ErrRecordNotFound
	}

	return nil
}
