package middlewares

import (
	"Rest-API/config"
	"Rest-API/model"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Middlewares interface {
	CreateToken(user model.Users) (string, error)
	ExtractTokenUserId(userType string, c *fiber.Ctx) string
	AuthMiddleware() fiber.Handler
}

func CreateToken(user model.Users) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = user.ID
	claims["user_type"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.SecretJWT))
}

func ExtractTokenUserId(userType string, c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.SecretJWT), nil
	})
	if err != nil || !token.Valid {
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)
	if userType == model.ALL_TYPE {
		return userId
	} else if claims["user_type"].(string) == userType {
		return userId
	} else {
		return ""
	}
}

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid Authorization header")
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.SecretJWT), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
		}

		claims := token.Claims.(jwt.MapClaims)
		userId, ok := claims["userId"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid user ID")
		}

		c.Locals("userId", userId)

		return c.Next()
	}
}
