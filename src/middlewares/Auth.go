package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.Token)

	if !ok {
		return fmt.Errorf("Invalid credentials")
	}

	claims := user.Claims.(jwt.MapClaims)

	companyId := claims["companyId"].(string)
	if !ok {
		return fmt.Errorf("Invalid credentials[0]")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		userId = ""
	}

	userEmail, ok := claims["userEmail"].(string)
	if !ok {
		userEmail = ""
	}

	userName, ok := claims["userName"].(string)
	if !ok {
		userName = ""
	}

	c.Locals("companyId", companyId)
	c.Locals("userId", userId)
	c.Locals("userEmail", userEmail)
	c.Locals("userName", userName)

	return c.Next()
}
