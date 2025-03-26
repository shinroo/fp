package apimodels

import "github.com/shinroo/fp/src/webapp/pkg/models"

type CreateSpecificAlertRequest struct {
	Breed     string           `json:"breed"`
	Gender    models.Gender    `json:"gender"`
	LifeStage models.LifeStage `json:"life_stage"`
}

type CreateSpecificAlertResponse struct{}
