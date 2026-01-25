package middleware

import (
	"time"

	"learning-go/internal/platform/logger" // Importa tu paquete de zap

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ZapLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		logger.Log.Info("Petición HTTP",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", duration),
			zap.String("ip", c.IP()),
		)

		return err
	}
}
