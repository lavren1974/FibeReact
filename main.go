package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./frontend/dist")

	app.Get("/app", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"app": "FibeReact",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
