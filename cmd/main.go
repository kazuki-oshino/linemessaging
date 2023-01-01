package main

import (
	"linemessaging/lib/messaging/usecase"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("broad cast start!")
	godotenv.Load(".env")
	usecase.PublishMessage()
	log.Println("broad cast end")
}
