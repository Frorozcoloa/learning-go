package main

import (
	"learning-go/internal/features/health"
	"learning-go/internal/platform/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	app := fiber.New()

	logger.Log.Info("Iniciando servidor en el puerto :3000")
	app.Get("/ping", health.PingHandler)
	app.Listen(":3000")
}
