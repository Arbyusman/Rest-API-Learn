package controller

import (
	"Rest-API/model"
	"Rest-API/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	LoginController(c fiber.Ctx) error
	RegisterController(c fiber.Ctx) error
}

type authController struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) *authController {
	return &authController{
		authUsecase: authUsecase,
	}
}

func (u *authController) LoginController(c fiber.Ctx) error {
	var payload model.Users

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	user, err := u.authUsecase.LoginUseCase(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Login successfully",
		},
		Data: user,
	})

}

func (u *authController) RegisterController(c fiber.Ctx) error {
	var payload model.Users

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	user, err := u.authUsecase.RegisterUseCase(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Register successfully",
		},
		Data: user,
	})

}
