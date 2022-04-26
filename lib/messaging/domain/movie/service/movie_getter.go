package service

import (
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

func (s *MovieService) GetBroadcastMovie() *model.Movie {

	homosapiMovie := s.movieGetterRepository.GetLatestMovie("homosapi")
	if homosapiMovie.IsPublishedToday(9) {
		return homosapiMovie
	}
	jeradonMovie := s.movieGetterRepository.GetLatestMovie("j")
	if jeradonMovie.IsPublishedToday(9) {
		return jeradonMovie
	}
	rand.Seed(time.Now().UnixNano())
	godURLList := getGodURLList()
	todaysGodURL := godURLList[rand.Intn(len(godURLList))]
	todaysGodMovie, _ := model.NewMovie("神曲", todaysGodURL, nil)
	return todaysGodMovie
}
