package api

import (
	"leave-a-message/api/message"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	message.Routes(v1)
}
