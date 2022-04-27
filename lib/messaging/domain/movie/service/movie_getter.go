package service

import (
	"fmt"
	"linemessaging/lib/messaging/domain/movie/model"
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
	// SeroriURL is セロリのURL
	SeroriURL = "https://www.youtube.com/watch?v=M4sWFgBYNbI"
	J         = "j"
	Homosapi  = "homosapi"
	Hige      = "hige"
)

// GodURLList is 神動画リスト
func getGodURLList() []string {
	return []string{
		"https://www.youtube.com/watch?v=vPwaXytZcgI",
		"https://www.youtube.com/watch?v=kOHB85vDuow",
		"https://www.youtube.com/watch?v=rRzxEiBLQCA",
		"https://www.youtube.com/watch?v=XA2YEHn-A8Q",
		"https://www.youtube.com/watch?v=c7rCyll5AeY",
		"https://www.youtube.com/watch?v=3ymwOvzhwHs",
		"https://www.youtube.com/watch?v=fmOEKOjyDxU",
		"https://www.youtube.com/watch?v=sLmLwgxnPUE",
		"https://www.youtube.com/watch?v=CM4CkVFmTds",
		"https://www.youtube.com/watch?v=i0p1bmr0EmE",
		SeroriURL,
	}
}

type likeMovie struct {
	key      string
	timeDiff int
}

func getLikeMovies() []likeMovie {
	return []likeMovie{
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
		return nil, fmt.Errorf("key: %s movie is not published Today.", key)
	}
	return m, nil
}

func (s *MovieService) GetBroadcastMovie() *model.Movie {

	for _, target := range getLikeMovies() {
		movie, err := s.getMoviePublishedToday(target.key, target.timeDiff)
		if err == nil {
			return movie
		}
	}

	rand.Seed(time.Now().UnixNano())
	godURLList := getGodURLList()
	todaysGodURL := godURLList[rand.Intn(len(godURLList))]
	todaysGodMovie, _ := model.NewMovie("神曲", todaysGodURL, nil)
	return todaysGodMovie
}
