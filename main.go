package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	Connect()
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	log.Fatal(app.Listen(":3000"))
}
