package v1

import (
	"github.com/gofiber/fiber/v2"
	handler "go_test/http/api/v1/service"
)

func AuthRoutes(api fiber.Router) {
	api.Get("/send-otp", handler.List)
	api.Get("/login", handler.List)
}
