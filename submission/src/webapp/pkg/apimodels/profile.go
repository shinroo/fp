package apimodels

import "github.com/shinroo/fp/src/webapp/pkg/models"

type GetProfileResponse struct {
	Profile *models.Profile `json:"profile"`
}

type UpdateProfileRequest struct {
	AffectionateWithFamily   int     `json:"affectionate_with_family"`
	GoodWithYoungChildren    int     `json:"good_with_young_children"`
	GoodWithOtherDogs        int     `json:"good_with_other_dogs"`
	SheddingLevel            int     `json:"shedding_level"`
	CoatGroomingFrequency    int     `json:"coat_grooming_frequency"`
	DroolingLevel            int     `json:"drooling_level"`
	CoatLength               int     `json:"coat_length"`
	OpennessToStrangers      int     `json:"openness_to_strangers"`
	PlayfulnessLevel         int     `json:"playfulness_level"`
	WatchdogProtectiveNature int     `json:"watchdog_protective_nature"`
	AdaptabilityLevel        int     `json:"adaptability_level"`
	TrainabilityLevel        int     `json:"trainability_level"`
	EnergyLevel              int     `json:"energy_level"`
	BarkingLevel             int     `json:"barking_level"`
	MentalStimulationNeeds   int     `json:"mental_stimulation_needs"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
}

type UpdateProfileResponse struct{}
