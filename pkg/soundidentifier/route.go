package soundidentifier

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	router := chi.NewRouter()
	router.Post("/animal-sound", AnimalSoundHandler)
	router.Post("/multiple-animal-sounds", MultipleAnimalSoundsHandler)
	router.Get("/history", AgeHistoryHandler)
	return router
}
