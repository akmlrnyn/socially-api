package repository

import (
	"gin_social/configs"
	"gin_social/entities"

	"gorm.io/gorm"
)

type AuthRepository interface{
	EmailExist(email string) bool
	Register(req *entities.User) error
}

type authRepository struct{
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) EmailExist(email string) bool {
	var user entities.User
	err := r.db.First(&user, "email = ?", email).Error

	return err == nil 
}

func (r *authRepository) Register(user *entities.User) error {
	err := configs.DB.Create(&user).Error

	return err
}