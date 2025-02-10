package models

type Account struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
