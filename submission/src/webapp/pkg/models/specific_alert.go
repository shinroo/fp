package models

import "time"

type SpecificAlert struct {
	ID        int       `json:"id" db:"id"`
	AccountID int       `json:"account_id" db:"account_id"`
	Breed     string    `json:"breed" db:"breed"`
	LifeStage LifeStage `json:"life_stage" db:"life_stage"`
	Gender    Gender    `json:"gender" db:"gender"`
	Actioned  bool      `json:"actioned" db:"actioned"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
