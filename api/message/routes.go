package message

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Get("/message", GetMessages)
	route.Post("/message", NewMessage)

	//route.Get("/books/:id", GetBook)
	//route.Put("/books/:id", UpdateBook)
	//route.Delete("/books/:id", DeleteBook)
}
