package main

import (
	"linemessaging/messaging"
	"log"
)

func main() {
	log.Println("broad cast start")

	messaging.Execute()

	log.Println("broad cast end")
}
