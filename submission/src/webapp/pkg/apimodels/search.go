package apimodels

import "github.com/shinroo/fp/src/webapp/pkg/models"

type SearchRequest struct {
	Keywords  string            `json:"keywords"`
	Gender    *models.Gender    `json:"gender"`
	LifeStage *models.LifeStage `json:"life_stage"`
	Breed     *string           `json:"breed"`
}

type SearchResponse struct {
	Results    []*models.Dog           `json:"results"`
	SPCALookup map[string]*models.SPCA `json:"spca_lookup"`
}
