package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tom-draper/api-analytics/server/database"
	"github.com/tom-draper/api-analytics/server/email"
	"github.com/tom-draper/api-analytics/server/tools/usage/usage"
	"os"
)

func getEmailAddress() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	address := os.Getenv("EMAIL_ADDRESS")
	return address
}

func emailBody(users []database.UserRow, usage []usage.UserCount) string {
	return fmt.Sprintf("%d new users\n%d requests", len(users), len(usage))
}

func main() {
	users, err := usage.DailyUsers()
	if err != nil {
		panic(err)
	}
	usage, err := usage.DailyUsage()
	if err != nil {
		panic(err)
	}
	body := emailBody(users, usage)
	address := getEmailAddress()
	println(address)
	email.SendEmail("API Analytics", body, address)
}
