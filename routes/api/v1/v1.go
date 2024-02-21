package v1

import (
	"github.com/gofiber/fiber/v2"
	"go_test/routes/api/v1/auth"
)

func InitRoutes(api fiber.Router) {
	v1 := api.Group("/v1")
	auth.Routes(v1)
}
