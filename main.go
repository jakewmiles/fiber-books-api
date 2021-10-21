package main

import "github.com/gofiber/fiber/v2"

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func router(app *fiber.App) {
	app.Get("/", helloWorld)
}

func main() {
	app := fiber.New()

	router(app)

	app.Listen(":3000")
}
