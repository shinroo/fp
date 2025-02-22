package models

type Profile struct {
	AccountID          int     `json:"account_id" db:"account_id"`
	HasChildren        bool    `json:"has_children" db:"has_children"`
	HasActiveLifestyle bool    `json:"has_active_lifestyle" db:"has_active_lifestyle"`
	HasLotsOfTime      bool    `json:"has_lots_of_time" db:"has_lots_of_time"`
	Latitude           float64 `json:"latitude" db:"latitude"`
	Longitude          float64 `json:"longitude" db:"longitude"`
	NearestSPCAID      string  `json:"nearest_spca" db:"nearest_spca"`
}
