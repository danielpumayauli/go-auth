package main

import (
	// "fmt"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@/go_auth"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	fmt.Println(db)

	app := fiber.New()

	app.Get("/", home)

	app.Listen(":8000")
}

func home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
