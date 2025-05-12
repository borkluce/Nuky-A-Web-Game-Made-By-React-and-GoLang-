package fiber

import (
	// Package specific
	"github.com/kahleryasla/pkg/go/auth/token"

	// Third
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Example: Extract token from headers
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header required"})
	}

	authHeader = authHeader[7:]

	// Verify token
	_, err := token.VerifyToken(authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Invalid token",
			"details": err.Error(),
			"token":   authHeader,
		})
	}

	return c.Next()
}
