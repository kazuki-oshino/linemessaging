package service

import (
	"fmt"
	"linemessaging/lib/messaging/domain/movie/model/vo"
	"linemessaging/lib/messaging/domain/movie/service/repository"
	"math/rand"
	"time"
)

type MovieService struct {
	movieGetterRepository repository.MovieGetterRepository
}

func NewMovieService(movieGetterRepository repository.MovieGetterRepository) *MovieService {
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
		"https://www.youtube.com/watch?v=F-QTb-0wRGk",
		"https://www.youtube.com/watch?v=RLlA18piUVo",
		"https://www.youtube.com/watch?v=HG0M_eZduxY",
		"https://www.youtube.com/watch?v=C47K1TX9PAA",
		"https://www.youtube.com/watch?v=kNH3eExqWFw",
	}
}

type LikeMovie struct {
	key      string
	timeDiff int
}

func getLikeMovies() []LikeMovie {
	return []LikeMovie{
		{
			key:      Homosapi,
			timeDiff: 9,
		},
	}
}

func (s *MovieService) getMoviePublishedToday(key string, timeDiff int) (*vo.Movie, error) {
	m := s.movieGetterRepository.GetLatestMovie(key)
	if !m.IsPublishedToday(timeDiff) {
		return nil, fmt.Errorf("key: %s movie is not published Today.Latest publishedDate: %v", key, m.PublishedDate())
	}
	return m, nil
}

func (s *MovieService) GetLikeMovie(movies []LikeMovie) *vo.Movie {
	for _, target := range movies {
		movie, err := s.getMoviePublishedToday(target.key, target.timeDiff)
		if err == nil {
			return movie
		}
	}

	return nil
}

func (s *MovieService) GetGodMovie(godURLList []string) *vo.Movie {
	rand.Seed(time.Now().UnixNano())
	todaysGodURL := godURLList[rand.Intn(len(godURLList))]
	todaysGodMovie, _ := vo.NewMovie("神曲", todaysGodURL, nil)
	return todaysGodMovie
}

func (s *MovieService) GetBroadcastMovie() *vo.Movie {

	movie := s.GetLikeMovie(getLikeMovies())
	if movie != nil {
		return movie
	}

	return s.GetGodMovie(getGodURLList())
}
