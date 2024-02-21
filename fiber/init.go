package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"go_test/routes/api"
	"go_test/routes/web"
)

func InitFiber() *fiber.App {
	engine := html.New("./public/views", ".html")

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
		Views:         engine,
	})

	app.Static("/static", "./public/static")
	app.Use(cors.New())
	api.InitRoutes(app)
	web.InitRoutes(app)
	return app
}
