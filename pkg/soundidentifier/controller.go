package soundidentifier

import (
	"animal-api/database"
	"animal-api/database/dbmodel"
	"animal-api/pkg/model"
	"net/http"

	"github.com/go-chi/render"
)

func getAnimalSound(animal string) string {
	switch animal {
	case "chat":
		return "Miaou"
	case "chien":
		return "Wouf"
	case "vache":
		return "Meuh"
	case "cheval":
		return "Hiiii"
	case "mouton":
		return "Bêêê"
	case "coq":
		return "Cocorico"
	default:
		return "Inconnu"
	}
}

func AnimalSoundHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.SoundRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	animal := req.AnimalNames[0]
	sound := getAnimalSound(animal)

	res := &model.SoundResponse{AnimalSound: sound}

	animalSound := dbmodel.AnimalSound{AnimalName: req.AnimalNames[0], Sound: res.AnimalSound}
	database.DB.Create(&animalSound)

	render.JSON(w, r, res)
}

func MultipleAnimalSoundsHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.SoundRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	sounds := make(map[string]string)
	for _, animal := range req.AnimalNames {
		sounds[animal] = getAnimalSound(animal)
	}

	res := &model.MultipleSoundsResponse{AnimalSounds: sounds}

	for i, _ := range req.AnimalNames {
		animalSound := dbmodel.AnimalSound{AnimalName: req.AnimalNames[i], Sound: sounds[req.AnimalNames[i]]}
		database.DB.Create(&animalSound)
	}

	render.JSON(w, r, res)
}

func AgeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	var animalSounds []dbmodel.AnimalSound
	database.DB.Find(&animalSounds)

	render.JSON(w, r, animalSounds)
}
