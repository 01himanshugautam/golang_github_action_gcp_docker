package fiber

import (
	"chumbak_be/routes"
	"chumbak_be/service"
	"chumbak_be/utils/abstraction"
	"chumbak_be/utils/handler"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html"
	"gorm.io/gorm"
	"os"
)

func InitFiber(psqlDatabase *gorm.DB, requestDebugFile *os.File) *fiber.App {
	appContext := abstraction.Context{
		Db: psqlDatabase,
	}

	engine := html.New("./views", ".html")

	// start fiber server
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./public")
	app.Use(cors.New())

	loggerConfig := fiberLogger.Config{
		// https://github.com/gofiber/fiber/issues/2372
		// hack: regex for sanitization \\u\S+\s
		Format: "url: ${url} | latency: ${latency} | time: ${time} | status: ${status}\n",
		/*Done: func(_ *fiber.Ctx, logString []byte) {
			msg := fmt.Sprintf("%s ", logString)
			logger.LogDebugMessage(msg)
		},*/
		Output: requestDebugFile,
	}
	app.Use(fiberLogger.New(loggerConfig))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(fiberRecover.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title": "Hello from team chumbak",
		})
	})

	// serves the templates
	app.Get("/slider/:template_id", func(ctx *fiber.Ctx) error {
		templateId := ctx.Params("template_id")
		template, err := service.GetTemplateFromTemplateId(abstraction.CreateNewAppContext(ctx, psqlDatabase), templateId)
		fmt.Printf("template is %+v\n", template)
		if err != nil {
			return ctx.SendString(err.Error())
		}
		if template == nil {
			return ctx.SendString("Record not found")
		}

		switch template.TemplateType {
		case "image":
			return ctx.Render("index", fiber.Map{
				"imgUrl": template.Url,
			}, "layouts/slider")
		case "video":
			return ctx.Render("index", fiber.Map{
				"videoUrl": template.Url,
			}, "layouts/video_template")
		default:
			return ctx.SendString("unsupported template type")
		}
	})

	routes.AuthRoutes(app, handler.AuthHandler{
		AppContext: appContext,
	})
	routes.AdsRoute(app, handler.AdsHandler{
		AppContext: appContext,
	})
	routes.MetricRoutes(app, handler.MetricHandler{
		AppContext: appContext,
	})
	return app
}
