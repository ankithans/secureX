package repository

import "github.com/ankithans/secureX/api/models"

var Users []models.User = []models.User{
	{
		Username: "ankithans",
		Password: "abc@def",
		Name:     "Ankit Hans",
		Phone:    8877898893,
		Address:  "3971 James Avenue, RANIER, Minnesota",
	},
	{
		Username: "johndoe01",
		Password: "abc@de",
		Name:     "John Doe",
		Phone:    8877898893,
		Address:  "3971 James Avenue, RANIER, Minnesota",
	},
	{
		Username: "ap23",
		Password: "abc@de",
		Name:     "Aryamaan",
		Phone:    8877898893,
		Address:  "3971 James Avenue, RANIER, Minnesota",
	},
	{
		Username: "pateladit01",
		Password: "abc@def",
		Name:     "Adit Patel",
		Phone:    8877898893,
		Address:  "3971 James Avenue, RANIER, Minnesota",
	},
}

func FindUsername(username string) bool {
	for _, user := range Users {
		if user.Username == username {
			return true
		}
	}
	return false
}

func MatchPassword(username string, password string) bool {
	for _, user := range Users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
