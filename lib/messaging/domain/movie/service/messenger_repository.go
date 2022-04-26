package service

import (
	"linemessaging/lib/messaging/domain/movie/model"
)

type MessengerRepository interface {
	Broadcast(movie *model.Movie)
}
