package models

import "time"

type SimilarityAlert struct {
	ID                  int       `json:"id" db:"id"`
	AccountID           int       `json:"account_id" db:"account_id"`
	SimilarityThreshold float32   `json:"similarity_threshold" db:"similarity_threshold"`
	Actioned            bool      `json:"actioned" db:"actioned"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
}
