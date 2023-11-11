package repository

import (
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

//func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
//	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
//		return err
//	}
//	return nil
//}

func (ur *userRepository) CreateUser(user model.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
