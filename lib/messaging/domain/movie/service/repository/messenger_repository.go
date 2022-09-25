package repository

import (
	"linemessaging/lib/messaging/domain/movie/model/vo"
)

type MessengerRepository interface {
	Broadcast(movie *vo.Movie)
}
