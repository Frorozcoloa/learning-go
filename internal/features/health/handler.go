package health

import "github.com/gofiber/fiber/v2"

func PingHandler(c *fiber.Ctx) error {
	response := HealthResponse{
		Status:  "ok",
		Message: "pong",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
