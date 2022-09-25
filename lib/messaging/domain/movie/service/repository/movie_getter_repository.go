package repository

import (
	"linemessaging/lib/messaging/domain/movie/model/vo"
)

type MovieGetterRepository interface {
	GetLatestMovie(findBy string) *vo.Movie
}
