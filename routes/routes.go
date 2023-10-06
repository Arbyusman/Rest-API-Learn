package routes

import (
	"Rest-API/controller"
	"Rest-API/repository"
	"Rest-API/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(e *fiber.App, db *gorm.DB) {

	authRepository := repository.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUsecase(authRepository)
	authController := controller.NewAuthController(authUseCase)

	// Users
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)

	// AUTH
	auth := e.Group("/api/v1")
	auth.Post("/login", authController.Login)
	auth.Post("/register", authController.Register)
	auth.Post("/admin/register", authController.RegisterAdmin)

	// User
	// e.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("oke")
	// })

	user := e.Group("/api/v1/")
	user.Get("users", userController.GetAllUsersController)
	user.Get("user/profile", userController.GetProfileUserController)
	user.Put("user", userController.UpdateUserByIDController)
	// Admin
	user.Get("user/:id", userController.GetUserByIdController)
	user.Delete("users/:id", userController.UpdateUserByIDController)

}
