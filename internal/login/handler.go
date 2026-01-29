package login

import (
	"learning-go/internal/platform/auth"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	token, _ := auth.GenerateToken("Fredy")
	return c.SendString(token)
}
