package usecase

import (
	"Rest-API/model"
	"Rest-API/repository"
	"errors"
	"fmt"
	"strings"
	"time"
)

type UserUsecase interface {
	GetAllUsersUseCase(page, limit int) ([]*model.UserResponse, int, error)
	GetUserByIDUseCase(userId string) (*model.UserResponse, error)
	UpdateUserByIDUseCase(userId string, payload *model.Users) (*model.UserResponse, error)
	DeleteUserByIDUseCase(userId string) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uc *userUsecase) GetAllUsersUseCase(page, limit int) ([]*model.UserResponse, int, error) {

	users, totalCount, err := uc.userRepository.GetAllUsersRepository(page, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all users: %v", err)
	}

	resp := make([]*model.UserResponse, 0, len(users))
	for _, user := range users {
		resp = append(resp, &model.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Role:      user.Role,
			Address:   user.Address,
			Image:     user.Image,
			UpdatedAt: user.UpdatedAt,
			CreatedAt: user.CreatedAt,
		})
	}

	return resp, totalCount, nil
}

func (uc *userUsecase) GetUserByIDUseCase(userId string) (*model.UserResponse, error) {

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	resp := &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		Address:   user.Address,
		Image:     user.Image,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}

	return resp, nil
}

func (uc *userUsecase) UpdateUserByIDUseCase(userId string, payload *model.Users) (*model.UserResponse, error) {

	lowercasePayload := &model.Users{
		Name:    strings.ToLower(payload.Name),
		Email:   strings.ToLower(payload.Email),
		Address: strings.ToLower(payload.Address),
		Phone:   payload.Phone,
	}

	user, err := uc.userRepository.GetUserByIDRepository(userId)
	if err != nil {
		return nil, fmt.Errorf("User Not Found: %v", err)
	}

	user.Name = lowercasePayload.Name
	user.Address = lowercasePayload.Address
	user.Phone = lowercasePayload.Phone
	user.Image = payload.Image
	user.UpdatedAt = time.Now()

	updatedUser, err := uc.userRepository.UpdateUserByIDRepository(userId, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	resp := &model.UserResponse{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Phone:     updatedUser.Phone,
		Role:      updatedUser.Role,
		Address:   updatedUser.Address,
		Image:     updatedUser.Image,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return resp, nil
}

func (uc *userUsecase) DeleteUserByIDUseCase(userId string) error {
	err := uc.userRepository.DeleteUserByIDRepository(userId)
	if err != nil {
		return errors.New("user not found")
	}
	return err
}
