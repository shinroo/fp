package models

type Dog struct {
	Identifier string    `db:"identifier" json:"identifier"`
	Name       string    `db:"name" json:"name"`
	Gender     Gender    `db:"gender" json:"gender"`
	LifeStage  LifeStage `db:"life_stage" json:"life_stage"`
	ImageUrl   string    `db:"image_url" json:"image_url"`
	SpcaId     string    `db:"spca_id" json:"spca_id"`
}
