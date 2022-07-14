package repository

import (
	"fmt"
	"linemessaging/lib/messaging/domain/movie/model"
	"linemessaging/lib/messaging/domain/movie/service"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

type lineMessengerRepository struct{}

func NewLineMessengerRepository() service.MessengerRepository {

	return &lineMessengerRepository{}
}

func (repo *lineMessengerRepository) Broadcast(movie *model.Movie) {
	bot, err := linebot.New(os.Getenv("SECRET"), os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	message1 := linebot.NewTextMessage(fmt.Sprintf("今日の動画は「%s」やで!?!?", movie.Title()))
	message2 := linebot.NewTextMessage(movie.URL())
	if _, err := bot.BroadcastMessage(message1, message2).Do(); err != nil {
		log.Fatalln(err)
	}
}
