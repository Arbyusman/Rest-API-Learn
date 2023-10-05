package controller

import (
	"Rest-API/model"
	"Rest-API/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetAllUsersController(c *fiber.Ctx) error
	// GetUserByIdController(c *fiber.Ctx) error
}

type userController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		UserUsecase: userUsecase,
	}
}

func (ctrl *userController) GetAllUsersController(c *fiber.Ctx) error {

	// userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	// if userId == "" {
	// 	return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
	// 		StatusCode: http.StatusUnauthorized,
	// 		Message:    "Token unauthorized",
	// 	})
	// }

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	users, err := ctrl.UserUsecase.GetAllUsersUseCase(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to get users: %v", err),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    "Successfully get users",
		},
		Data: users,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
		},
	})
}

// func (ctrl *userController) GetUserByIdController(c *fiber.Ctx) error {
// 	// userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
// 	// if userId == "" {
// 	// 	return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
// 	// 		StatusCode: http.StatusUnauthorized,
// 	// 		Message:    "token unauthorized",
// 	// 	})

// 	// }

// 	user, err := ctrl.UserUsecase.GetUserByIDUseCase('2')
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
// 			StatusCode: http.StatusNotFound,
// 			Message:    err.Error(),
// 		})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
// 		MetaData: model.MetaData{
// 			StatusCode: http.StatusOK,
// 			Message:    fmt.Sprintf("Successfully get user with ID: %s", user.ID),
// 		},
// 		Data: user,
// 	})

// }
