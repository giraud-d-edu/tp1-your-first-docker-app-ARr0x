package model

import (
	"errors"
	"net/http"
)

type SoundRequest struct {
	AnimalNames []string `json:"animals"`
}

func (a *SoundRequest) Bind(r *http.Request) error {
	if len(a.AnimalNames) == 0 {
		return errors.New("le champ 'animals' est manquant ou vide")
	}
	return nil
}

type SoundResponse struct {
	AnimalSound string `json:"sound"`
}

type MultipleSoundsResponse struct {
	AnimalSounds map[string]string `json:"sounds"`
}
