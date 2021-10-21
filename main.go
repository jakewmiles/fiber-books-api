package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakewmiles/fiber-books-api/book"
)

func helloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}

func router(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBook)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book/", book.PostBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	router(app)

	app.Listen(":3000")
}
