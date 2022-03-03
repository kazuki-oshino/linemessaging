package main

import (
	"linemessaging/cmd/messaging/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("broad cast start")
	godotenv.Load("../../.env")
	app.PublishMessage()
	log.Println("broad cast end")
}
