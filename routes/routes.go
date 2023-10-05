package routes

import (
	"Rest-API/controller"
	"Rest-API/repository"
	"Rest-API/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(e *fiber.App, db *gorm.DB) {

	// authRepository := repository.NewAuthRepository(db)
	// authUseCase := usecase.NewAuthUsecase(authRepository)
	// authController := controller.NewAuthController(authUseCase)

	// Users
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)

	// AUTH
	// auth := e.Group("/api/v1")
	// auth.POST("/login", authController.Login)
	// auth.POST("/register", authController.Register)

	// User
	e.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("oke")
	})

	user := e.Group("/api/v1/")
	user.Get("users", userController.GetAllUsersController)

}
