package repository

import (
	"Rest-API/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsersRepository(page, limit int) ([]*model.Users, error)
	GetUserByIDRepository(id string) (*model.Users, error)
	UpdateUserByIDRepository(id string, user *model.Users) (*model.Users, error)
	DeleteUserByIDRepository(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUsersRepository(page, limit int) ([]*model.Users, error) {
	var users []*model.Users

	offset := (page - 1) * limit

	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUserByIDRepository(id string) (*model.Users, error) {

	var user model.Users
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %s not found", id)
		}
		return nil, fmt.Errorf("error getting user with ID %s: %s", id, result.Error)
	}
	return &user, nil
}

func (r *userRepository) UpdateUserByIDRepository(id string, user *model.Users) (*model.Users, error) {

	result := r.db.Model(&model.Users{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *userRepository) DeleteUserByIDRepository(id string) error {
	result := r.db.Delete(&model.Users{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
