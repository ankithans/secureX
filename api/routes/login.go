package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ankithans/secureX/api/repository"
	"github.com/ankithans/secureX/api/utils"
	"github.com/ankithans/secureX/secure"
	"github.com/gofiber/fiber/v2"
)

var loginCountByPort = make(map[string]int)
var lastFraudLoginTime = time.Now()

func Login(c *fiber.Ctx) error {
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println(c.Port())

	timeNow := time.Now()
	timeDiff := timeNow.Sub(lastFraudLoginTime)
	if timeDiff.Minutes() >= 1 {
		// close docker container
		go secure.StopApiContainer()
	}

	if loginCountByPort[c.Port()] >= 3 {
		lastFraudLoginTime = time.Now()
		fmt.Println("Intruder detected; redirecting to decoy")

		secure.RunApiContainer()

		response, err := http.Get("http://127.0.0.1:8080/api/v1/login?username=" + username + "&password=" + password)
		if err != nil {
			fmt.Println(err, 1)
			log.Fatal(err)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err, 2)
			log.Fatal(err)
		}

		var jsonResult map[string]interface{}
		json.Unmarshal(responseData, &jsonResult)
		return c.JSON(jsonResult)
	}

	loginCountByPort[c.Port()]++

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
