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
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5
	DB.Create(&book)
	return ctx.JSON(book)
}

func DeleteBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var book Book
	DB.First(&book, id)
	if book.Title == "" {
		return ctx.Status(500).SendString("No book found with ID " + id)
	}
	DB.Delete(&book)
	return ctx.SendString("Book successfully deleted")
}
