package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	log.Println("broad cast start")
	bot, err := linebot.New(os.Getenv("SECRET"), os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	message1 := linebot.NewTextMessage("mazi test.")
	message2 := linebot.NewTextMessage("https://www.youtube.com/watch?v=M4sWFgBYNbI")
	if _, err := bot.BroadcastMessage(message1, message2).Do(); err != nil {
		log.Fatalln(err)
	}
	log.Println("broad cast end")
}
