package routes

import (
	"fmt"

	"github.com/ankithans/secureX/api/repository"
	"github.com/ankithans/secureX/api/utils"
	"github.com/gofiber/fiber/v2"
)

var colorReset = "\033[0m"

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

func Login(c *fiber.Ctx) error {
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println(string(colorBlue), c.Port(), string(colorReset))

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
