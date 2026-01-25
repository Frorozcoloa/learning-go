package health

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/ping", PingHandler)
	req := httptest.NewRequest("GET", "/ping", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "pong")
}
