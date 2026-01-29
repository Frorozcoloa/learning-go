package greeting_test

import (
	"net/http/httptest"
	"testing"
	"learning-go/internal/features/greeting"
	"learning-go/internal/platform/auth"
	"learning-go/internal/platform/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)


func TestHelloHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/hello/:userName", middleware.ProtectJWT(), greeting.HelloHandler)

	t.Run("Debe retornar 200 y el saludo correcto con un JWT válido", func(t *testing.T) {
		token, _ := auth.GenerateToken("Fredy")
		req := httptest.NewRequest("GET", "/hello/Fredy", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, _ := app.Test(req)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Debe retornar 401 si no se envía el token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/hello/Fredy", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, 401, resp.StatusCode)
	})
}