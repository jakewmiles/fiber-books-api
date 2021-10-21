package book

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(ctx *fiber.Ctx) error {
	return ctx.SendString("All Book!")
}
func GetBook(ctx *fiber.Ctx) error {
	return ctx.SendString("One Book!")
}
func PostBook(ctx *fiber.Ctx) error {
	return ctx.SendString("New Book!")
}
func DeleteBook(ctx *fiber.Ctx) error {
	return ctx.SendString("No Book!")
}
