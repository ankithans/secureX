package main

import (

	// "github.com/ankithans/secureX/secure"
	"log"

	"github.com/ankithans/secureX/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// secure.RunApiContainer()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": 200, "message": "Welcome to SecureX"})
	})
	app.Get("/api/v1/login", routes.Login)

	log.Fatal(app.Listen(":6000"))
}
