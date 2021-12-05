package main

import (

	// "github.com/ankithans/secureX/secure"

	"log"

	"github.com/ankithans/secureX/api/routes"
	"github.com/ankithans/secureX/api/utils"
	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := utils.GoDotEnvVariable("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&models.AuditLogs{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": 200, "message": "Welcome to SecureX"})
	})
	app.Get("/api/v1/login", func(c *fiber.Ctx) error {
		return routes.Login(c, db)
	})

	log.Fatal(app.Listen(":6000"))
}
