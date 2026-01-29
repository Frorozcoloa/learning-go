package main

import (
	"learning-go/internal/features/greeting"
	"learning-go/internal/features/health"
	"learning-go/internal/login"
	"learning-go/internal/platform/logger"
	"learning-go/internal/platform/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	app := fiber.New()
	app.Use(middleware.ZapLogger())

	logger.Log.Info("Iniciando servidor en el puerto :3000")
	app.Get("/login", login.LoginHandler)
	app.Get("/hello/:userName", middleware.ProtectJWT(), greeting.HelloHandler)
	app.Get("/ping", health.PingHandler)
	app.Listen(":3000")
}
