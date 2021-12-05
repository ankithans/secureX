package main

import (
	"log"

	"github.com/ankithans/secureX/api/routes"
	"github.com/ankithans/secureX/api/utils"
	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func main() {

	// db.AutoMigrate(&models.AuditLogs{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": 200, "message": "Welcome to SecureX"})
	})

	app.Get("/api/v1/login", func(c *fiber.Ctx) error {
		dsn := utils.GoDotEnvVariable("POSTGRES_URI")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		return routes.Login(c, db)
	})

	log.Fatal(app.Listen(":8080"))
}
