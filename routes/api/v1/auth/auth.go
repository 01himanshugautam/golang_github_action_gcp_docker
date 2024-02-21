package auth

import (
	"github.com/gofiber/fiber/v2"
	handler "go_test/http/api/v1/service"
	"go_test/http/api/v1/service/auth_service"
)

func Routes(v1 fiber.Router) {
	authService := auth_service.NewAuthService()
	auth := v1.Group("/auth")
	auth.Get("/send-otp", authService.SendOTP)
	auth.Get("/login", handler.List)
}
