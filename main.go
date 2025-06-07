package main

import (
	"log"

	"Saturobot/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go func() {
		if err := config.StartTelegramBot(); err != nil {
			log.Fatalf("Failed to start Telegram bot: %v", err)
		}
	}()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Telegram Bot is running...")
	})

	app.Post("/webhook", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Testing a POST request"})
	})
	log.Fatal(app.Listen(":3000"))
}
