package model

import (
	"errors"
	"net/http"
)

type AgeRequest struct {
	HumanAge int `json:"human_age"`
}

func (a *AgeRequest) Bind(r *http.Request) error {
	if a.HumanAge < 0 {
		return errors.New("human_age must be a positive integer")
	}
	return nil
}

type AgeResponse struct {
	CatAge int `json:"cat_age"`
}
