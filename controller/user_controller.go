package controller

import (
	"Rest-API/model"
	"Rest-API/usecase"
	"Rest-API/usecase/cloudinary"
	middlewares "Rest-API/usecase/middleware"
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
	userId := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	users, count, err := ctrl.UserUsecase.GetAllUsersUseCase(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to get users: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: fiber.StatusOK,
			Message:    "Successfully get users",
		},
		Data: users,
		Pagination: &model.Pagination{
			Page:  page,
			Limit: limit,
			Count: count,
		},
	})
}

func (ctrl *userController) GetUserByIdController(c *fiber.Ctx) error {

	userId := c.Params("id")

	response, err := ctrl.UserUsecase.GetUserByIDUseCase(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    fmt.Sprintf("Successfully get user with ID: %s", response.ID),
		},
		Data: response,
	})

}
func (ctrl *userController) GetProfileUserController(c *fiber.Ctx) error {

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	response, err := ctrl.UserUsecase.GetUserByIDUseCase(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: http.StatusOK,
			Message:    fmt.Sprintf("Successfully get user with ID: %s", response.ID),
		},
		Data: response,
	})

}

func (ctrl *userController) UpdateUserByIDController(c *fiber.Ctx) error {

	userId := middlewares.ExtractTokenUserId(model.USER_TYPE, c)
	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	var payload model.Users

	// Parse the form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	imageFiles := form.File["image"]

	if len(imageFiles) > 0 {
		file, _ := imageFiles[0].Open()
		defer file.Close()

		imageURL, err := cloudinary.ImageUploadHelper(file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			})
		}
		payload.Image = imageURL
	}
	payload.Name = getFormValue(form.Value, "name")
	payload.Address = getFormValue(form.Value, "address")
	payload.Phone = getFormValue(form.Value, "phone")

	response, err := ctrl.UserUsecase.UpdateUserByIDUseCase(userId, &payload)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: fiber.StatusOK,
			Message:    "User Updated successfully",
		},
		Data: response,
	})
}

func (ctrl *userController) DeleteUserByIDController(c *fiber.Ctx) error {

	user := middlewares.ExtractTokenUserId(model.ADMIN_TYPE, c)
	if user == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	userId := c.Params("id")
	err := ctrl.UserUsecase.DeleteUserByIDUseCase(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.HttpResponse{
		MetaData: model.MetaData{
			StatusCode: fiber.StatusOK,
			Message:    "User Deleted successfully",
		},
	})

}

func getFormValue(form map[string][]string, field string) string {
	if values, exists := form[field]; exists && len(values) > 0 {
		return values[0]
	}
	return ""
}
