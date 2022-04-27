package service

import (
	"errors"
	"fmt"
	"linemessaging/lib/messaging/domain/movie/model"
	"log"
	"math/rand"
	"time"
)

type MovieService struct {
	movieGetterRepository MovieGetterRepository
}

func NewMovieService(movieGetterRepository MovieGetterRepository) *MovieService {
	return &MovieService{
		movieGetterRepository: movieGetterRepository,
	}
}

const (
	J        = "j"
	Homosapi = "homosapi"
	Hige     = "hige"
)

// GodURLList is 神動画リスト
func getGodURLList() []string {
	return []string{
		"https://www.youtube.com/watch?v=N-bdKXQcGiM",
		"https://www.youtube.com/watch?v=XicdpSmxuT0",
		"https://www.youtube.com/watch?v=B--iJ2pNvLU",
		"https://www.youtube.com/watch?v=E6EItQRTmAI",
		"https://www.youtube.com/watch?v=CbH2F0kXgTY",
	}
}

type LikeMovie struct {
	key      string
	timeDiff int
}

func getLikeMovies() []LikeMovie {
	return []LikeMovie{
		{
			key:      "homosapi",
			timeDiff: 9,
		},
		{
			key:      "j",
			timeDiff: 9,
		},
	}
}

func (s *MovieService) getMoviePublishedToday(key string, timeDiff int) (*model.Movie, error) {
	m := s.movieGetterRepository.GetLatestMovie(key)
	if !m.IsPublishedToday(timeDiff) {
		return nil, fmt.Errorf("key: %s movie is not published Today.Latest publishedDate: %v", key, m.PublishedDate())
	}
	return m, nil
}

func (s *MovieService) GetLikeMovie(movies []LikeMovie) (*model.Movie, error) {
	for _, target := range movies {
		movie, err := s.getMoviePublishedToday(target.key, target.timeDiff)
		if err == nil {
			return movie, nil
		} else {
			log.Println(err.Error())
		}
	}
	return nil, errors.New("Like movie is not exist today.")
}

func (s *MovieService) GetGodMovie(godURLList []string) *model.Movie {
	rand.Seed(time.Now().UnixNano())
	todaysGodURL := godURLList[rand.Intn(len(godURLList))]
	todaysGodMovie, _ := model.NewMovie("神曲", todaysGodURL, nil)
	return todaysGodMovie
}

func (s *MovieService) GetBroadcastMovie() *model.Movie {

	movie, err := s.GetLikeMovie(getLikeMovies())
	if err != nil {
		log.Println(err.Error())
	} else {
		return movie
	}

	return s.GetGodMovie(getGodURLList())
}
