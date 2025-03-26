package models

type DogBreed struct {
	Name string `json:"name" db:"name"`
}
