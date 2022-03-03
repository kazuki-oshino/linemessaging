package model

import (
	"log"
	"time"
)

// Movie is movie info struct.
type Movie struct {
	title         string
	url           string
	publishedDate *time.Time
}

// NewTarget is make TargetLatestMovie struct.
func NewMovie(title, url string, publishedDate *time.Time) (movie *Movie) {
	if url == "" {
		log.Fatalln("urlは必須です")
	}
	return &Movie{
		title:         title,
		url:           url,
		publishedDate: publishedDate,
	}
}

func (m *Movie) Title() string {
	return m.title
}

func (m *Movie) URL() string {
	return m.url
}

func (m *Movie) IsPublishedToday(timeDiff int) bool {
	if m.publishedDate.UTC().Add(time.Hour*time.Duration(timeDiff)).Day() == time.Now().UTC().Add(time.Hour*9).Day() {
		return true
	} else {
		return false
	}
}
