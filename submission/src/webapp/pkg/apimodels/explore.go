package apimodels

import "github.com/shinroo/fp/src/webapp/pkg/models"

type ExploreResponse struct {
	Results    []*models.Dog           `json:"results"`
	SPCALookup map[string]*models.SPCA `json:"spca_lookup"`
}
