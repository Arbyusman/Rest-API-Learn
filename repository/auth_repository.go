package repository

import (
	"Rest-API/model"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterRepository(user model.Users) (*model.Users, error)
	LoginRepository(user *model.Users) (*model.Users, error)
	GetUserByEmailRepository(email string) (*model.Users, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) RegisterRepository(user model.Users) (*model.Users, error) {
	var count int64
	r.db.Model(&model.Users{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		return nil, errors.New("email already exists")
	}
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) LoginRepository(user *model.Users) (*model.Users, error) {
	result := r.db.Where("email = ? AND password = ?", user.Email, user.Password).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("incorrect email or password")
		}
		return nil, errors.New("failed to get user")
	}

	return user, nil
}

func (r *authRepository) GetUserByEmailRepository(email string) (*model.Users, error) {
	var user model.Users
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
