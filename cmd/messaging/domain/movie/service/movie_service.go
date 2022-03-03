package service

import (
	"linemessaging/cmd/messaging/domain/movie/model"
	"linemessaging/cmd/messaging/repository"
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
	// todo(kazuki): 投稿頻度が半端ないため一旦やめておく
	// higeMovie := s.movieGetterRepository.GetLatestMovie("hige")
	// if higeMovie.IsPublishedToday(9) {
	// 	return higeMovie
	// }
	rand.Seed(time.Now().UnixNano())
	godURLList := getGodURLList()
	todaysGodURL := godURLList[rand.Intn(len(godURLList))]
	todaysGodMovie, _ := model.NewMovie("神曲", todaysGodURL, nil)
	return todaysGodMovie
}
