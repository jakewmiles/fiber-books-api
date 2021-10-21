package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(ctx *fiber.Ctx) error {
	var books []Book
	DB.Find(&books)
	return ctx.JSON(books)
}
func GetBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var book Book
	DB.Find(&book, id)
	return ctx.JSON(book)
}
func PostBook(ctx *fiber.Ctx) error {
	return ctx.SendString("New Book!")
}
func DeleteBook(ctx *fiber.Ctx) error {
	return ctx.SendString("No Book!")
}
