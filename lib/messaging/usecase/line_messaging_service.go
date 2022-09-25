package usecase

import (
	"linemessaging/lib/messaging/domain/movie/service"
	"linemessaging/lib/messaging/infrastructure/repository_impl"
)

func PublishMessage() {

	// make movie
	movieGetterRepository := repository_impl.NewFeedMovieGetterRepository()
	movieService := service.NewMovieService(movieGetterRepository)
	movie := movieService.GetBroadcastMovie()

	// broadcast movie
	messengerRepository := repository_impl.NewLineMessengerRepository()
	messengerRepository.Broadcast(movie)
}
