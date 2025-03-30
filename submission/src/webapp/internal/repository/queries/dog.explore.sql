SELECT
    dog.identifier AS identifier,
    dog.name AS name,
    dog.gender AS gender,
    dog.life_stage AS life_stage,
    dog.image_url AS image_url,
    dog.spca_id AS spca_id,
    COALESCE(dog.breed, 'Unknown breed') AS breed
FROM
    dog
    LEFT JOIN dog_breed ON dog.breed = dog_breed.name
ORDER BY ARRAY [
	COALESCE(dog_breed.affectionate_with_family,1),
	COALESCE(dog_breed.good_with_young_children,1),
	COALESCE(dog_breed.good_with_other_dogs,1), 
	COALESCE(dog_breed.shedding_level,5),
	COALESCE(dog_breed.coat_grooming_frequency,5),
	COALESCE(dog_breed.drooling_level,5),
	COALESCE(dog_breed.coat_length,5),
	COALESCE(dog_breed.openness_to_strangers,1),
	COALESCE(dog_breed.playfulness_level,1),
	COALESCE(dog_breed.watchdog_protective_nature,1),
	COALESCE(dog_breed.adaptability_level,1),
	COALESCE(dog_breed.trainability_level,1),
	COALESCE(dog_breed.energy_level,5),
	COALESCE(dog_breed.barking_level,5),
	COALESCE(dog_breed.mental_stimulation_needs,5)
]::vector <-> $1;