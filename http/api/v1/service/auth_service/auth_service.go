package auth_service

import "github.com/gofiber/fiber/v2"

// AuthService Service Auth represents the authentication module
type AuthService struct{}

// NewAuthService NewAuth creates a new instance of Auth
func NewAuthService() *AuthService {
	return &AuthService{}
}

// SendOTP sends an OTP to the user
func (a *AuthService) SendOTP(c *fiber.Ctx) error {
	return c.SendString("OTP Sent")
}
