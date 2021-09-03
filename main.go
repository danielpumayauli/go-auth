package main

import (
	"github.com/danielpumayauli/go-auth/database"
	"github.com/danielpumayauli/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)
	app.Get("/", home)

	app.Listen(":8000")
}

func home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
