package apimodels

import "github.com/shinroo/fp/src/webapp/pkg/models"

type GetProfileResponse struct {
	Profile *models.Profile `json:"profile"`
}

type UpdateProfileRequest struct {
	HasChildren        bool    `json:"has_children"`
	HasActiveLifestyle bool    `json:"has_active_lifestyle"`
	HasLotsOfTime      bool    `json:"has_lots_of_time"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
}

type UpdateProfileResponse struct{}
