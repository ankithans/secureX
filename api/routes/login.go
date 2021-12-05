package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/ankithans/secureX/api/models"
	"github.com/ankithans/secureX/api/repository"
	"github.com/ankithans/secureX/api/utils"
	"github.com/ankithans/secureX/secure"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var loginCountByPort = make(map[string]int)
var lastFraudLoginTime = time.Now()

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

	fmt.Println("User from port:", string(colorBlue), c.Context().RemoteAddr().String(), string(colorReset))

	timeNow := time.Now()
	timeDiff := timeNow.Sub(lastFraudLoginTime)
	if timeDiff.Minutes() >= 1 {
		// close docker container
		go secure.StopApiContainer()
	}

	if loginCountByPort[c.Port()] >= 3 {

		clientIp := c.Context().RemoteAddr()

		message := "Hi Team,\nA fraudulent activity has been captured on Login API. Please find the details below. For more information please refer to Audit Logs in Database.\n\n" + "RemoteAddress: " + clientIp.String() + "\nIp: " + c.IP() + "\nDescription: " + "Logging in with username: " + username + "and password: " + password + "\nNetwork: " + clientIp.Network() + "\nStatus: " + "danger" + "\n\nThanks & regards\nDecoy Team"
		go sendMail(message)
		auditLog := models.AuditLogs{
			RemoteAddress: clientIp.String(),
			Ip:            c.IP(),
			Port:          c.Port(),
			Network:       clientIp.Network(),
			Status:        "danger",
			Description:   "Logging in with username: " + username + "and password: " + password,
			Location:      "server",
		}
		go db.Create(&auditLog)

		// var audit models.AuditLogs
		// db.First(&audit)

		// fmt.Println(audit)

		lastFraudLoginTime = time.Now()
		fmt.Println(string(colorYellow), "Intruder detected; redirecting to decoy")

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

func sendMail(message string) {
	from := goDotEnvVariable("EMAIL")
	password := goDotEnvVariable("PASSWD")
	toList := []string{"ankithans1947@gmail.com"}
	host := "smtp.gmail.com"
	port := "587"

	fro := fmt.Sprintf("From: <%s>\r\n", "dserver03@gmail.com")
	to := fmt.Sprintf("To: <%s>\r\n", "ankithans1947@gmail.com")
	subject := "Subject: no-reply-decoy-server-team\r\n"

	msg := fro + to + subject + "\r\n" + message
	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// handling the errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
