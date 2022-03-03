package repository

import (
	"linemessaging/cmd/messaging/domain/movie/model"
)

type MessengerRepository interface {
	Broadcast(movie *model.Movie)
}
