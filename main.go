package main

import (
	"github.com/danielpumayauli/go-auth/database"
	"github.com/danielpumayauli/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	app.Get("/", home)

	app.Listen(":8000")
}

func home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
