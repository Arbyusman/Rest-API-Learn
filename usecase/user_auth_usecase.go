package usecase

import (
	"Rest-API/model"
	"Rest-API/repository"
	middlewares "Rest-API/usecase/middleware"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	RegisterUseCase(payload model.Users) (*model.AuthResponse, error)
	RegisterAdminUseCase(payload model.Users) (*model.AuthResponse, error)
	LoginUseCase(payload model.Users) (*model.AuthResponse, error)
}

type authUsecase struct {
	authRepository repository.AuthRepository
}

func NewAuthUsecase(authRepository repository.AuthRepository) *authUsecase {
	return &authUsecase{authRepository: authRepository}
}

func (s *authUsecase) RegisterUseCase(payload model.Users) (*model.AuthResponse, error) {

	lowercasePayload := model.Users{
		Name:     strings.ToLower(payload.Name),
		Email:    strings.ToLower(payload.Email),
		Password: payload.Password,
	}

	if payload.Name == "" || payload.Email == "" {
		return nil, errors.New("name and email are required fields")
	}

	hashedPassword, _ := HashPassword(payload.Password)

	newUserModel := model.Users{
		Name:     lowercasePayload.Name,
		Email:    lowercasePayload.Email,
		Role:     model.USER_TYPE,
		Password: hashedPassword,
	}

	user, err := s.authRepository.RegisterRepository(newUserModel)
	if err != nil {
		return nil, fmt.Errorf("error creating user in database: %w", err)
	}

	token, err := middlewares.CreateToken(*user)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	resp := &model.AuthResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}

	return resp, nil
}
func (s *authUsecase) RegisterAdminUseCase(payload model.Users) (*model.AuthResponse, error) {

	lowercasePayload := model.Users{
		Name:     strings.ToLower(payload.Name),
		Email:    strings.ToLower(payload.Email),
		Password: payload.Password,
	}

	if payload.Name == "" || payload.Email == "" {
		return nil, errors.New("name and email are required fields")
	}

	hashedPassword, _ := HashPassword(payload.Password)

	newUserModel := model.Users{
		Name:     lowercasePayload.Name,
		Email:    lowercasePayload.Email,
		Role:     model.ADMIN_TYPE,
		Password: hashedPassword,
	}

	user, err := s.authRepository.RegisterRepository(newUserModel)
	if err != nil {
		return nil, fmt.Errorf("error creating user in database: %w", err)
	}

	token, err := middlewares.CreateToken(*user)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	resp := &model.AuthResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}

	return resp, nil
}

func (s *authUsecase) LoginUseCase(payload model.Users) (*model.AuthResponse, error) {

	lowercasePayloadEmail := strings.ToLower(payload.Email)

	if payload.Email == "" {
		return nil, errors.New("email field is required")
	}

	if payload.Password == "" {
		return nil, errors.New("password field is required")

	}
	user, err := s.authRepository.GetUserByEmailRepository(lowercasePayloadEmail)
	if err != nil {
		return nil, err
	}

	if !ComparePasswords(user.Password, payload.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := middlewares.CreateToken(*user)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	resp := &model.AuthResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}

	return resp, nil
}

func ComparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
