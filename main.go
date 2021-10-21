package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jakewmiles/fiber-books-api/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func router(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBook)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book/", book.PostBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=fiber_books_api port=5432 sslmode=disable TimeZone=Europe/London"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&book.Book{})
	DB = db
	fmt.Println("Connected to database!")
}

func main() {
	app := fiber.New()

	ConnectDatabase()
	router(app)

	app.Listen(":3000")
}
