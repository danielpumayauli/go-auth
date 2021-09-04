package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"net/smtp"

	"github.com/danielpumayauli/go-auth/database"
	"github.com/danielpumayauli/go-auth/models"
	"github.com/gofiber/fiber/v2"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := RandStringRunes(12)

	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	if data["email"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email not found",
		})
	}

	database.DB.Create(&passwordReset)

	from := "admin@example.com"

	to := []string{
		data["email"],
	}
	// URL Client
	url := "http://localhost:3000/reset/" + token

	message := []byte("Click <a href=\"" + url + "\">to reset your password!</a>")

	err := smtp.SendMail("localhost:1025", nil, from, to, message)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error sending mail")
		log.Fatal(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"email":   data["email"],
	})

}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
