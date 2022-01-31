package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mmcdole/gofeed"
)

func main() {
	log.Println("broad cast start")

	godotenv.Load(".env")
	bot, err := linebot.New(os.Getenv("SECRET"), os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	feed, err := gofeed.NewParser().ParseURL("https://www.youtube.com/feeds/videos.xml?channel_id=UCd0hscDvJvzRbo8Rk7JPQMA")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	latestItem := feed.Items[0]
	publishedDate := latestItem.PublishedParsed
	if publishedDate.Add(time.Hour*9).Day() == time.Now().UTC().Add(time.Hour*9).Day() {
		message1 := linebot.NewTextMessage("ホモサピの最新動画が来ているよ！")
		message2 := linebot.NewTextMessage(latestItem.Link)
		if _, err := bot.BroadcastMessage(message1, message2).Do(); err != nil {
			log.Fatalln(err)
		}
	} else {
		message1 := linebot.NewTextMessage("今日はホモサピの動画あがっていないよ・・・代わりにパセリ聞いてね")
		message2 := linebot.NewTextMessage("https://www.youtube.com/watch?v=M4sWFgBYNbI")
		if _, err := bot.BroadcastMessage(message1, message2).Do(); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("broad cast end")
}
