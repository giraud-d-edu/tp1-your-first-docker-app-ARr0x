package agecalculator

import (
	"animal-api/config"
	"animal-api/database/dbmodel"
	"animal-api/pkg/model"
	"net/http"

	"github.com/go-chi/render"
)

type AgeCalculatorConfig struct {
	*config.Config
}

func New(configuration *config.Config) *AgeCalculatorConfig {
	return &AgeCalculatorConfig{configuration}
}

func (config *AgeCalculatorConfig) AgeInCatYearsHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.AgeRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	catAge := calculateCatAge(req.HumanAge)
	ageEntry := &dbmodel.AgeEntry{HumanAge: req.HumanAge, CatAge: catAge}
	config.AgeEntryRepository.Create(ageEntry)

	res := &model.AgeResponse{CatAge: catAge}
	render.JSON(w, r, res)
}

func (config *AgeCalculatorConfig) AgeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.AgeEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}
func calculateCatAge(humanAge int) int {
	if humanAge <= 2 {
		return humanAge * 12
	}
	return 24 + (humanAge-2)*4
}
