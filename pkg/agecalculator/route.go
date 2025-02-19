package agecalculator

import (
	"animal-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	ageCalculatorConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/age-in-cat-years", ageCalculatorConfig.AgeInCatYearsHandler)
	router.Get("/history", ageCalculatorConfig.AgeHistoryHandler)

	return router
}
