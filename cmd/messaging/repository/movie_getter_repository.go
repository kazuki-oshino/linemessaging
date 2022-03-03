package repository

import (
	"linemessaging/cmd/messaging/domain/movie/model"
)

type MovieGetterRepository interface {
	GetLatestMovie(findBy string) *model.Movie
}
