package models

import "time"

type Session struct {
	Token     string    `json:"token" db:"token"`
	AccountID int       `json:"account_id" db:"account_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
