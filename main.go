package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakewmiles/fiber-books-api/book"
)

func router(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book/", book.PostBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	book.ConnectDatabase()
	router(app)

	app.Listen(":3000")
}
