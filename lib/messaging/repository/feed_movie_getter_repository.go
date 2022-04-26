package repository

import (
	"linemessaging/lib/messaging/domain/movie/model"
	"linemessaging/lib/messaging/domain/movie/service"
	"log"

	"github.com/mmcdole/gofeed"
)

const (
	// HomosapiFeedURL is ホモサピ動画FEED URL
	HomosapiFeedURL = "https://www.youtube.com/feeds/videos.xml?channel_id=UCd0hscDvJvzRbo8Rk7JPQMA"

	// HigeSoriFeedURL is ひげそりのFEED URL
	HigeSoriFeedURL = "https://www.youtube.com/feeds/videos.xml?channel_id=UCVI4ZUakZBLvdgb0ltKPS8Q"
)

type feedMovieGetterRepository struct{}

func NewFeedMovieGetterRepository() service.MovieGetterRepository {
	return &feedMovieGetterRepository{}
}

func (repo *feedMovieGetterRepository) GetLatestMovie(findBy string) *model.Movie {

	url := getURLByFeedKey(findBy)
	if url == "" {
		log.Fatal("target is not existed.")
	}

	feed, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		log.Fatal("failed to parse feed URL.")
	}

	if len(feed.Items) == 0 {
		log.Fatal("target channel does't have Movie.")
	}
	movie, err := model.NewMovie(feed.Items[0].Title, feed.Items[0].Link, feed.Items[0].PublishedParsed)
	if err != nil {
		log.Fatal("url not found.")
	}
	return movie
}

func getURLByFeedKey(key string) string {
	feedMap := map[string]string{"homosapi": HomosapiFeedURL, "hige": HigeSoriFeedURL}
	return feedMap[key]
}
