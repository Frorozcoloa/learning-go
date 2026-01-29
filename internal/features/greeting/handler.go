package greeting

import (
	"github.com/gofiber/fiber/v2"
)

func HelloHandler(c *fiber.Ctx) error {
	name := c.Params("userName")
	return c.JSON(
		fiber.Map{
			"message": "Hola, " + name,
		},
	)
}
