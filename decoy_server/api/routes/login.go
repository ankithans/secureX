package routes

import (
	"fmt"

	"github.com/ankithans/secureX/api/models"
	"github.com/ankithans/secureX/api/repository"
	"github.com/ankithans/secureX/api/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var colorReset = "\033[0m"

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

func Login(c *fiber.Ctx, db *gorm.DB) error {

	username := c.Query("username")
	password := c.Query("password")
	clientIp := c.Context().RemoteAddr()

	auditLog := models.AuditLogs{
		RemoteAddress: clientIp.String(),
		Ip:            c.IP(),
		Port:          c.Port(),
		Network:       clientIp.Network(),
		Status:        "danger",
		Description:   "Logging in with username: " + username + "and password: " + password,
		Location:      "decoy",
	}
	db.Create(&auditLog)

	// var audit models.AuditLogs
	// db.First(&audit)

	// fmt.Println(audit)

	fmt.Println(clientIp.String())

	// check username in repository
	if !repository.FindUsername(username) {
		return c.JSON(fiber.Map{"status": 404, "message": "username not found"})
	}

	// match password in repository
	if !repository.MatchPassword(username, password) {
		return c.JSON(fiber.Map{"status": 401, "message": "wrong password"})
	}

	return c.JSON(fiber.Map{"status": 200, "message": "successfully logged in", "username": username, "access_token": utils.RandStringBytes(18)})
}
