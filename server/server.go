package server

import (
	"fmt"
	"os"

	"leave-a-message/database"
	"leave-a-message/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func setupMiddlewares(app *fiber.App) {
	//app.Use(helmet.New())
	//app.Use(recover.New())
	app.Use(cors.New())
	//app.Use(compress.New(compress.Config{
	//	Level: compress.LevelBestSpeed, // 1
	//}))
	//app.Use(etag.New())
	//if os.Getenv("ENABLE_LIMITER") != "" {
	//	app.Use(limiter.New())
	//}
	if os.Getenv("ENABLE_LOGGER") != "" {
		app.Use(logger.New())
	}
}

func Create() *fiber.App {
	database.SetupDatabase()

	// Setup fiber html template engine (official)
	engine := html.New("./templates", ".html")
	// optional for development: engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*pkg.Error); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*fiber.Error); ok {
				return ctx.Status(e.Code).JSON(pkg.Error{Status: e.Code, Code: "internal-server", Message: e.Message})
			} else {
				return ctx.Status(500).JSON(pkg.Error{Status: 500, Code: "internal-server", Message: err.Error()})
			}
		},
	})

	setupMiddlewares(app)

	// Serve static assets from ./static
	app.Static("/static", "./static")

	// Render minimal Cloudflare-style frontend using ctx.Render
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"Title": "Leave a Message"})
	})

	return app
}

func Listen(app *fiber.App) error {

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	return app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort))
}
