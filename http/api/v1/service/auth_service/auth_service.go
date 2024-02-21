package auth_service

import "github.com/gofiber/fiber/v2"

// Service Auth represents the authentication module
type Service struct{}

// NewAuthService NewAuth creates a new instance of Auth
func NewAuthService() *Service {
	return &Service{}
}

// SendOTP sends an OTP to the user
func (a *Service) sendOTP(c *fiber.Ctx) error {
	return c.SendString("OTP Sent")
}
