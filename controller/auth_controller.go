package controller

import (
	"Rest-API/model"
	"Rest-API/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authController struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) *authController {
	return &authController{
		authUsecase: authUsecase,
	}
}

func (u *authController) Login(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Login successfully",
		},
		Data: user,
	})

}

func (u *authController) Register(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Register successfully",
		},
		Data: user,
	})

}

func (u *authController) RegisterAdmin(c *fiber.Ctx) error {
	var payload model.Users

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	user, err := u.authUsecase.RegisterAdminUseCase(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "User Register successfully",
		},
		Data: user,
	})

}
