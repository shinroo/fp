package models

import "fmt"

type Profile struct {
	AccountID                int     `json:"account_id" db:"account_id"`
	AffectionateWithFamily   int     `json:"affectionate_with_family" db:"affectionate_with_family"`
	GoodWithYoungChildren    int     `json:"good_with_young_children" db:"good_with_young_children"`
	GoodWithOtherDogs        int     `json:"good_with_other_dogs" db:"good_with_other_dogs"`
	SheddingLevel            int     `json:"shedding_level" db:"shedding_level"`
	CoatGroomingFrequency    int     `json:"coat_grooming_frequency" db:"coat_grooming_frequency"`
	DroolingLevel            int     `json:"drooling_level" db:"drooling_level"`
	CoatLength               int     `json:"coat_length" db:"coat_length"`
	OpennessToStrangers      int     `json:"openness_to_strangers" db:"openness_to_strangers"`
	PlayfulnessLevel         int     `json:"playfulness_level" db:"playfulness_level"`
	WatchdogProtectiveNature int     `json:"watchdog_protective_nature" db:"watchdog_protective_nature"`
	AdaptabilityLevel        int     `json:"adaptability_level" db:"adaptability_level"`
	TrainabilityLevel        int     `json:"trainability_level" db:"trainability_level"`
	EnergyLevel              int     `json:"energy_level" db:"energy_level"`
	BarkingLevel             int     `json:"barking_level" db:"barking_level"`
	MentalStimulationNeeds   int     `json:"mental_stimulation_needs" db:"mental_stimulation_needs"`
	Latitude                 float64 `json:"latitude" db:"latitude"`
	Longitude                float64 `json:"longitude" db:"longitude"`
	NearestSPCAID            string  `json:"nearest_spca" db:"nearest_spca"`
}

func (p *Profile) GetPreferredDogVector() string {
	return fmt.Sprintf(
		"[%d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d]",
		p.AffectionateWithFamily,
		p.GoodWithYoungChildren,
		p.GoodWithOtherDogs,
		p.SheddingLevel,
		p.CoatGroomingFrequency,
		p.DroolingLevel,
		p.CoatLength,
		p.OpennessToStrangers,
		p.PlayfulnessLevel,
		p.WatchdogProtectiveNature,
		p.AdaptabilityLevel,
		p.TrainabilityLevel,
		p.EnergyLevel,
		p.BarkingLevel,
		p.MentalStimulationNeeds,
	)
}
