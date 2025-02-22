package models

import "database/sql"

type SPCA struct {
	ID        string         `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Latitude  float64        `json:"latitude" db:"latitude"`
	Longitude float64        `json:"longitude" db:"longitude"`
	Website   sql.NullString `json:"website" db:"website"`
	Address   string         `json:"address" db:"address"`
}
