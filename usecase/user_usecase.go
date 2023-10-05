package usecase

import (
	"Rest-API/model"
	"Rest-API/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetAllUsersUseCase(page, limit int) ([]*model.UserResponse, error)
	GetUserByIDUseCase(userId string) (*model.UserResponse, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uc *userUsecase) GetAllUsersUseCase(page, limit int) ([]*model.UserResponse, error) {

	users, err := uc.userRepository.GetAllUsersRepository(page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}

	resp := make([]*model.UserResponse, 0, len(users))
	for _, user := range users {
		resp = append(resp, &model.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return resp, nil
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
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}

	return resp, nil
}

func CompareHashPin(hashPin string, pin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPin), []byte(pin))
	return err == nil
}

func HashPin(pin string) string {
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPin)
}
