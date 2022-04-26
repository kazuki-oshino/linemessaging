package service

import (
	"linemessaging/lib/messaging/domain/movie/model"
)

type MovieGetterRepository interface {
	GetLatestMovie(findBy string) *model.Movie
}
