package apimodels

import "github.com/shinroo/fp/src/webapp/pkg/models"

type SearchRequest struct {
	Keywords  string            `json:"keywords"`
	Gender    *models.Gender    `json:"gender"`
	LifeStage *models.LifeStage `json:"life_stage"`
}

type SearchResponse struct {
	Results []*models.Dog `json:"results"`
}
