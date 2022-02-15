package messaging

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mmcdole/gofeed"
)

const (
	// HomosapiFeedURL is ホモサピ動画FEED URL
	HomosapiFeedURL = "https://www.youtube.com/feeds/videos.xml?channel_id=UCd0hscDvJvzRbo8Rk7JPQMA"

	// SeroriURL is セロリのURL
	SeroriURL = "https://www.youtube.com/watch?v=M4sWFgBYNbI"

	// HigeSoriFeedURL is ひげそりのFEED URL
	HigeSoriFeedURL = "https://www.youtube.com/feeds/videos.xml?channel_id=UCVI4ZUakZBLvdgb0ltKPS8Q"
)

// TargetLatestMovie is movie info struct.
type TargetLatestMovie struct {
	title         string
	url           string
	publishedDate time.Time
}

// NewTarget is make TargetLatestMovie struct.
func NewTarget(item gofeed.Item) *TargetLatestMovie {
	return &TargetLatestMovie{
		title:         item.Title,
		url:           item.Link,
		publishedDate: *item.PublishedParsed,
	}
}

// GetTargetLatestMovieByFeed is get TargetLatestMovie by feed url.
func GetTargetLatestMovieByFeed(feedURL string) *TargetLatestMovie {
	feed, err := gofeed.NewParser().ParseURL(feedURL)
	if err != nil {
		log.Fatal("failed to parse feed URL.")
	}

	if len(feed.Items) == 0 {
		log.Fatal("target channel does't have Movie.")
	}

	return NewTarget(*feed.Items[0])
}

// BroadcastInfo is BroadcastInfo struct.
type BroadcastInfo struct {
	bot     *linebot.Client
	message string
	url     string
}

// NewBroadCastInfo is make NewBroadCastInfo.
func NewBroadCastInfo(message, url string) *BroadcastInfo {
	bot, err := linebot.New(os.Getenv("SECRET"), os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	return &BroadcastInfo{
		bot:     bot,
		message: message,
		url:     url,
	}
}

// BroadCast is broadcast by line messaging api.
func (b *BroadcastInfo) BroadCast() {
	message1 := linebot.NewTextMessage(b.message)
	message2 := linebot.NewTextMessage(b.url)
	if _, err := b.bot.BroadcastMessage(message1, message2).Do(); err != nil {
		log.Fatalln(err)
	}
}

// checkTargetAndBroadCast is check target latest movie, if published date is today broadcast this movie.
func checkTargetAndBroadCast(targetURL string, foundMessage string, timeDiff int) bool {
	feed := GetTargetLatestMovieByFeed(targetURL)
	var broadCastInfo *BroadcastInfo
	if feed.publishedDate.Add(time.Hour*time.Duration(timeDiff)).Day() == time.Now().UTC().Add(time.Hour*9).Day() {
		broadCastInfo = NewBroadCastInfo(foundMessage, feed.url)
		broadCastInfo.BroadCast()
		log.Println(foundMessage)
		return true
	}
	return false
}

// Execute is to execute Line Messaging API to push message.
func Execute() {
	godotenv.Load(".env")

	isHomosapiBroadCast := checkTargetAndBroadCast(HomosapiFeedURL, "ホモサピの動画あったよ！嬉しいね！", 9)
	if isHomosapiBroadCast {
		return
	}

	isHigeBroadCast := checkTargetAndBroadCast(HigeSoriFeedURL, "ひげの最新動画来たでー", 0)
	if isHigeBroadCast {
		return
	}

	broadCastInfo := NewBroadCastInfo("今日はホモサピもひげもないよ・・・代わりにパセリ聞いてね", SeroriURL)
	broadCastInfo.BroadCast()
	log.Println("今日はホモサピもひげもないよ・・・代わりにパセリ聞いてね")
}
