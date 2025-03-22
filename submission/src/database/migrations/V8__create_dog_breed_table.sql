CREATE TABLE IF NOT EXISTS dog_breed (
    name TEXT PRIMARY KEY,
    affectionate_with_family INT,
    good_with_young_children INT,
    good_with_other_dogs INT,
    shedding_level INT,
    coat_grooming_frequency INT,
    drooling_level INT,
    coat_length INT,
    openness_to_strangers INT,
    playfulness_level INT,
    watchdog_protective_nature INT,
    adaptability_level INT,
    trainability_level INT,
    energy_level INT,
    barking_level INT,
    mental_stimulation_needs INT
);

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Retrievers (Labrador)',
        5,
        5,
        5,
        4,
        2,
        2,
        1,
        5,
        5,
        3,
        5,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'French Bulldogs',
        5,
        5,
        4,
        3,
        1,
        3,
        1,
        5,
        5,
        3,
        5,
        4,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'German Shepherd Dogs',
        5,
        5,
        3,
        4,
        2,
        2,
        3,
        3,
        4,
        5,
        5,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Retrievers (Golden)',
        5,
        5,
        5,
        4,
        2,
        2,
        3,
        5,
        4,
        3,
        5,
        5,
        3,
        1,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bulldogs',
        4,
        3,
        3,
        3,
        3,
        3,
        1,
        4,
        4,
        3,
        3,
        4,
        3,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Poodles',
        5,
        5,
        3,
        1,
        4,
        1,
        5,
        5,
        5,
        5,
        4,
        5,
        4,
        4,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Beagles',
        3,
        5,
        5,
        3,
        2,
        1,
        1,
        3,
        4,
        2,
        4,
        3,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Rottweilers',
        5,
        3,
        3,
        3,
        1,
        3,
        1,
        3,
        4,
        5,
        4,
        5,
        3,
        1,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pointers (German Shorthaired)',
        5,
        5,
        4,
        3,
        2,
        2,
        1,
        4,
        4,
        4,
        4,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Dachshunds',
        5,
        3,
        4,
        2,
        2,
        2,
        1,
        4,
        4,
        4,
        4,
        4,
        3,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pembroke Welsh Corgis',
        5,
        3,
        4,
        4,
        2,
        1,
        1,
        4,
        4,
        5,
        4,
        4,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Australian Shepherds',
        3,
        5,
        3,
        3,
        2,
        1,
        3,
        3,
        4,
        3,
        3,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Yorkshire Terriers',
        5,
        5,
        3,
        1,
        5,
        1,
        5,
        5,
        4,
        5,
        5,
        4,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Boxers',
        4,
        5,
        3,
        2,
        2,
        3,
        1,
        4,
        4,
        4,
        3,
        4,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Great Danes',
        5,
        3,
        3,
        3,
        1,
        4,
        1,
        3,
        4,
        5,
        4,
        3,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Siberian Huskies',
        5,
        5,
        5,
        4,
        2,
        1,
        3,
        5,
        5,
        1,
        4,
        3,
        5,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Cavalier King Charles Spaniels',
        5,
        5,
        5,
        2,
        2,
        2,
        3,
        4,
        3,
        3,
        3,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Doberman Pinschers',
        5,
        5,
        3,
        4,
        1,
        2,
        1,
        4,
        4,
        5,
        4,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Miniature Schnauzers',
        5,
        5,
        3,
        3,
        4,
        2,
        3,
        3,
        4,
        4,
        5,
        5,
        3,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Shih Tzu',
        5,
        5,
        5,
        1,
        4,
        1,
        5,
        3,
        3,
        3,
        5,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Boston Terriers',
        5,
        5,
        4,
        2,
        2,
        1,
        1,
        5,
        5,
        3,
        3,
        4,
        4,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bernese Mountain Dogs',
        5,
        5,
        5,
        5,
        3,
        3,
        3,
        4,
        4,
        3,
        4,
        4,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pomeranians',
        5,
        3,
        3,
        2,
        3,
        1,
        5,
        3,
        3,
        4,
        4,
        3,
        3,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Havanese',
        5,
        5,
        5,
        2,
        3,
        1,
        5,
        5,
        5,
        3,
        5,
        4,
        3,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Cane Corso',
        4,
        3,
        3,
        2,
        1,
        3,
        1,
        3,
        3,
        5,
        3,
        4,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (English Springer)',
        5,
        3,
        4,
        3,
        2,
        3,
        3,
        4,
        4,
        3,
        4,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Shetland Sheepdogs',
        5,
        5,
        5,
        3,
        3,
        1,
        5,
        2,
        5,
        5,
        5,
        5,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Brittanys',
        3,
        4,
        4,
        3,
        3,
        1,
        1,
        3,
        4,
        3,
        3,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pugs',
        5,
        5,
        4,
        4,
        2,
        1,
        1,
        5,
        5,
        3,
        5,
        4,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (Cocker)',
        4,
        5,
        5,
        3,
        4,
        2,
        5,
        4,
        3,
        3,
        4,
        4,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Miniature American Shepherds',
        5,
        5,
        5,
        3,
        3,
        2,
        3,
        3,
        4,
        3,
        4,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Border Collies',
        5,
        3,
        3,
        3,
        3,
        1,
        3,
        4,
        5,
        3,
        5,
        5,
        5,
        4,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Mastiffs',
        5,
        5,
        3,
        3,
        1,
        4,
        1,
        3,
        3,
        5,
        4,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Chihuahuas',
        4,
        1,
        3,
        2,
        1,
        1,
        1,
        2,
        4,
        4,
        4,
        3,
        4,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Vizslas',
        5,
        5,
        4,
        3,
        2,
        2,
        1,
        4,
        5,
        3,
        5,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Basset Hounds',
        3,
        5,
        5,
        2,
        3,
        4,
        1,
        3,
        3,
        3,
        3,
        3,
        2,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Belgian Malinois',
        3,
        3,
        3,
        3,
        2,
        1,
        1,
        3,
        3,
        4,
        3,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Maltese',
        5,
        3,
        3,
        1,
        4,
        1,
        5,
        3,
        3,
        4,
        4,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Weimaraners',
        5,
        5,
        3,
        3,
        2,
        2,
        1,
        3,
        5,
        5,
        4,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Collies',
        4,
        5,
        3,
        3,
        3,
        2,
        1,
        3,
        4,
        3,
        4,
        4,
        3,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Newfoundlands',
        5,
        5,
        5,
        3,
        2,
        5,
        3,
        5,
        3,
        5,
        4,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Rhodesian Ridgebacks',
        5,
        5,
        3,
        3,
        2,
        2,
        1,
        3,
        3,
        5,
        4,
        4,
        3,
        2,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Shiba Inu',
        5,
        3,
        3,
        3,
        2,
        1,
        1,
        3,
        3,
        5,
        3,
        2,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'West Highland White Terriers',
        5,
        5,
        3,
        3,
        3,
        1,
        3,
        4,
        5,
        5,
        4,
        3,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bichons Frises',
        5,
        5,
        5,
        1,
        5,
        1,
        5,
        5,
        4,
        2,
        4,
        4,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bloodhounds',
        4,
        3,
        3,
        3,
        2,
        5,
        1,
        3,
        3,
        2,
        3,
        4,
        3,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (English Cocker)',
        5,
        5,
        5,
        3,
        3,
        2,
        3,
        4,
        3,
        3,
        4,
        4,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Akitas',
        3,
        3,
        1,
        3,
        3,
        1,
        3,
        2,
        3,
        5,
        3,
        3,
        4,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Portuguese Water Dogs',
        5,
        5,
        4,
        2,
        4,
        2,
        5,
        4,
        5,
        4,
        5,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Retrievers (Chesapeake Bay)',
        4,
        3,
        3,
        3,
        3,
        2,
        3,
        3,
        3,
        4,
        4,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Dalmatians',
        5,
        3,
        3,
        4,
        2,
        2,
        1,
        4,
        4,
        4,
        4,
        4,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'St. Bernards',
        5,
        5,
        3,
        3,
        2,
        5,
        1,
        3,
        3,
        5,
        4,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Papillons',
        5,
        5,
        3,
        3,
        2,
        1,
        3,
        5,
        5,
        4,
        5,
        5,
        4,
        5,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Australian Cattle Dogs',
        3,
        3,
        3,
        3,
        1,
        1,
        1,
        3,
        3,
        4,
        3,
        4,
        5,
        1,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bullmastiffs',
        4,
        3,
        3,
        3,
        1,
        3,
        1,
        3,
        3,
        5,
        3,
        4,
        4,
        1,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Samoyeds',
        5,
        5,
        3,
        3,
        3,
        1,
        5,
        5,
        5,
        4,
        4,
        4,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Scottish Terriers',
        5,
        3,
        2,
        2,
        3,
        2,
        3,
        3,
        4,
        5,
        4,
        3,
        3,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Soft Coated Wheaten Terriers',
        5,
        5,
        3,
        1,
        4,
        2,
        3,
        3,
        3,
        3,
        3,
        3,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Whippets',
        5,
        5,
        5,
        2,
        1,
        1,
        1,
        3,
        4,
        3,
        3,
        3,
        4,
        1,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pointers (German Wirehaired)',
        5,
        3,
        3,
        2,
        2,
        2,
        3,
        4,
        4,
        3,
        4,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Chinese Shar-Pei',
        4,
        3,
        3,
        3,
        1,
        3,
        1,
        3,
        3,
        4,
        4,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Airedale Terriers',
        3,
        3,
        3,
        1,
        3,
        1,
        1,
        3,
        3,
        5,
        3,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Wirehaired Pointing Griffons',
        5,
        5,
        3,
        3,
        2,
        3,
        3,
        5,
        4,
        3,
        4,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bull Terriers',
        4,
        3,
        1,
        3,
        2,
        1,
        1,
        4,
        4,
        3,
        3,
        3,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Alaskan Malamutes',
        3,
        3,
        3,
        3,
        3,
        1,
        3,
        3,
        3,
        4,
        3,
        5,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Cardigan Welsh Corgis',
        4,
        4,
        3,
        3,
        2,
        1,
        3,
        4,
        4,
        3,
        3,
        4,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Giant Schnauzers',
        5,
        3,
        3,
        3,
        4,
        2,
        3,
        3,
        4,
        5,
        4,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Old English Sheepdogs',
        5,
        5,
        3,
        3,
        4,
        3,
        5,
        3,
        4,
        4,
        4,
        4,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Italian Greyhounds',
        5,
        3,
        5,
        3,
        1,
        1,
        1,
        5,
        4,
        3,
        4,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Great Pyrenees',
        5,
        3,
        3,
        3,
        2,
        3,
        3,
        3,
        3,
        5,
        3,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Dogues de Bordeaux',
        5,
        3,
        3,
        4,
        1,
        5,
        1,
        3,
        3,
        5,
        4,
        4,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Russell Terriers',
        5,
        3,
        5,
        3,
        2,
        1,
        1,
        5,
        5,
        4,
        4,
        3,
        5,
        4,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Cairn Terriers',
        4,
        3,
        3,
        2,
        2,
        1,
        3,
        3,
        4,
        4,
        3,
        3,
        3,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Irish Wolfhounds',
        5,
        3,
        4,
        3,
        2,
        2,
        3,
        3,
        3,
        5,
        3,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Setters (Irish)',
        5,
        5,
        5,
        3,
        3,
        2,
        3,
        5,
        5,
        3,
        4,
        4,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Greater Swiss Mountain Dogs',
        5,
        5,
        3,
        3,
        2,
        3,
        1,
        5,
        4,
        4,
        3,
        4,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Miniature Pinschers',
        5,
        3,
        4,
        3,
        1,
        1,
        1,
        3,
        4,
        5,
        4,
        3,
        5,
        5,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Lhasa Apsos',
        5,
        3,
        3,
        1,
        3,
        1,
        5,
        3,
        3,
        5,
        5,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Chinese Crested',
        4,
        3,
        3,
        1,
        2,
        1,
        1,
        4,
        3,
        3,
        4,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Coton de Tulear',
        5,
        5,
        5,
        2,
        4,
        1,
        5,
        5,
        4,
        3,
        4,
        4,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Staffordshire Bull Terriers',
        5,
        5,
        3,
        2,
        2,
        3,
        1,
        4,
        4,
        5,
        5,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'American Staffordshire Terriers',
        5,
        3,
        3,
        2,
        1,
        1,
        1,
        4,
        3,
        5,
        3,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Rat Terriers',
        5,
        5,
        3,
        3,
        2,
        1,
        1,
        5,
        5,
        4,
        4,
        5,
        4,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Chow Chows',
        4,
        3,
        2,
        3,
        3,
        3,
        3,
        2,
        3,
        5,
        3,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Anatolian Shepherd Dogs',
        1,
        3,
        3,
        3,
        2,
        1,
        1,
        1,
        3,
        5,
        3,
        2,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Basenjis',
        3,
        3,
        3,
        2,
        1,
        1,
        1,
        3,
        3,
        3,
        3,
        2,
        4,
        1,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (Boykin)',
        3,
        5,
        5,
        3,
        3,
        1,
        3,
        3,
        4,
        3,
        4,
        4,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Lagotti Romagnoli',
        5,
        3,
        5,
        1,
        2,
        2,
        3,
        4,
        4,
        3,
        4,
        4,
        4,
        2,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Brussels Griffons',
        4,
        3,
        3,
        3,
        3,
        1,
        1,
        4,
        4,
        3,
        4,
        4,
        3,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Retrievers (Nova Scotia Duck Tolling)',
        5,
        5,
        4,
        3,
        2,
        2,
        3,
        3,
        5,
        3,
        4,
        5,
        5,
        2,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Norwegian Elkhounds',
        5,
        3,
        3,
        3,
        2,
        2,
        3,
        4,
        4,
        5,
        4,
        4,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Standard Schnauzers',
        5,
        5,
        3,
        1,
        3,
        3,
        3,
        3,
        4,
        5,
        4,
        5,
        3,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Dogo Argentinos',
        5,
        3,
        3,
        4,
        1,
        3,
        1,
        4,
        4,
        5,
        4,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bouviers des Flandres',
        3,
        3,
        3,
        3,
        4,
        2,
        3,
        3,
        3,
        4,
        3,
        4,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pekingese',
        5,
        3,
        3,
        3,
        3,
        1,
        5,
        3,
        4,
        4,
        4,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Keeshonden',
        5,
        5,
        5,
        3,
        3,
        2,
        5,
        5,
        5,
        5,
        5,
        5,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Border Terriers',
        4,
        5,
        3,
        2,
        2,
        1,
        1,
        4,
        4,
        3,
        4,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Leonbergers',
        5,
        5,
        3,
        4,
        4,
        3,
        5,
        5,
        3,
        4,
        4,
        5,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Tibetan Terriers',
        5,
        3,
        3,
        3,
        3,
        2,
        5,
        3,
        3,
        4,
        4,
        3,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Neapolitan Mastiffs',
        3,
        3,
        3,
        3,
        2,
        5,
        1,
        3,
        2,
        5,
        3,
        3,
        2,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Setters (English)',
        5,
        4,
        4,
        3,
        3,
        3,
        3,
        4,
        4,
        3,
        4,
        4,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Retrievers (Flat-Coated)',
        5,
        5,
        5,
        3,
        2,
        2,
        3,
        5,
        5,
        3,
        5,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Borzois',
        3,
        3,
        3,
        3,
        2,
        1,
        3,
        3,
        3,
        3,
        3,
        2,
        4,
        2,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Fox Terriers (Wire)',
        5,
        5,
        3,
        2,
        3,
        2,
        3,
        5,
        4,
        4,
        4,
        4,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Miniature Bull Terriers',
        5,
        3,
        3,
        2,
        1,
        2,
        1,
        3,
        4,
        4,
        4,
        3,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Belgian Tervuren',
        3,
        3,
        3,
        3,
        2,
        1,
        3,
        3,
        3,
        4,
        3,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Setters (Gordon)',
        5,
        3,
        3,
        3,
        2,
        4,
        3,
        3,
        3,
        4,
        4,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Silky Terriers',
        5,
        3,
        3,
        2,
        3,
        1,
        5,
        4,
        4,
        5,
        4,
        4,
        3,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Norwich Terriers',
        5,
        5,
        3,
        3,
        2,
        1,
        1,
        5,
        4,
        4,
        4,
        3,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spinoni Italiani',
        5,
        3,
        4,
        3,
        2,
        3,
        3,
        3,
        3,
        3,
        3,
        4,
        3,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Japanese Chin',
        5,
        3,
        5,
        3,
        2,
        1,
        3,
        3,
        3,
        3,
        4,
        3,
        3,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Welsh Terriers',
        5,
        5,
        3,
        2,
        3,
        2,
        3,
        5,
        4,
        4,
        4,
        4,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Toy Fox Terriers',
        5,
        3,
        3,
        3,
        2,
        1,
        1,
        5,
        4,
        5,
        4,
        5,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Schipperkes',
        5,
        3,
        3,
        3,
        2,
        1,
        1,
        3,
        4,
        5,
        4,
        4,
        3,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Parson Russell Terriers',
        5,
        3,
        5,
        2,
        2,
        1,
        1,
        3,
        4,
        3,
        4,
        4,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pointers',
        5,
        3,
        5,
        3,
        2,
        2,
        1,
        4,
        4,
        3,
        4,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Belgian Sheepdogs',
        3,
        3,
        3,
        3,
        2,
        1,
        3,
        3,
        3,
        4,
        3,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Tibetan Spaniels',
        5,
        5,
        3,
        3,
        2,
        2,
        3,
        4,
        4,
        4,
        4,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'American Eskimo Dogs',
        5,
        5,
        3,
        3,
        3,
        1,
        3,
        5,
        3,
        3,
        4,
        4,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Irish Terriers',
        5,
        5,
        1,
        2,
        1,
        1,
        3,
        3,
        3,
        5,
        4,
        3,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Beaucerons',
        3,
        3,
        3,
        4,
        3,
        1,
        1,
        2,
        3,
        4,
        3,
        3,
        5,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Afghan Hounds',
        3,
        3,
        3,
        1,
        4,
        1,
        5,
        3,
        3,
        3,
        3,
        1,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Boerboels',
        5,
        4,
        2,
        3,
        2,
        3,
        1,
        3,
        3,
        5,
        3,
        4,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Fox Terriers (Smooth)',
        5,
        3,
        3,
        3,
        2,
        1,
        1,
        3,
        4,
        5,
        4,
        3,
        4,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bearded Collies',
        4,
        5,
        5,
        3,
        4,
        1,
        5,
        4,
        4,
        3,
        4,
        3,
        4,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Black Russian Terriers',
        3,
        3,
        3,
        3,
        4,
        3,
        3,
        2,
        3,
        5,
        4,
        4,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Black and Tan Coonhounds',
        4,
        5,
        5,
        3,
        2,
        3,
        1,
        3,
        3,
        2,
        4,
        3,
        3,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (Welsh Springer)',
        5,
        5,
        4,
        3,
        2,
        2,
        3,
        1,
        3,
        3,
        4,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'American Hairless Terriers',
        5,
        5,
        3,
        1,
        1,
        1,
        1,
        3,
        3,
        3,
        5,
        5,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Norfolk Terriers',
        5,
        5,
        3,
        3,
        2,
        1,
        1,
        5,
        4,
        4,
        4,
        3,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Xoloitzcuintli',
        5,
        3,
        3,
        1,
        1,
        1,
        1,
        3,
        4,
        3,
        4,
        4,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Manchester Terriers',
        5,
        4,
        3,
        2,
        2,
        1,
        1,
        3,
        4,
        4,
        4,
        4,
        4,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Kerry Blue Terriers',
        5,
        4,
        2,
        1,
        3,
        2,
        3,
        3,
        3,
        5,
        3,
        3,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Australian Terriers',
        3,
        5,
        3,
        1,
        2,
        1,
        1,
        3,
        4,
        4,
        3,
        4,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (Clumber)',
        4,
        3,
        3,
        3,
        2,
        4,
        3,
        4,
        3,
        3,
        4,
        4,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Lakeland Terriers',
        5,
        3,
        3,
        2,
        2,
        1,
        1,
        3,
        4,
        3,
        4,
        3,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bluetick Coonhounds',
        3,
        3,
        5,
        3,
        2,
        2,
        1,
        3,
        3,
        2,
        3,
        4,
        4,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'English Toy Spaniels',
        5,
        5,
        5,
        3,
        3,
        2,
        3,
        3,
        4,
        3,
        4,
        3,
        3,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'German Pinschers',
        5,
        3,
        3,
        3,
        1,
        1,
        1,
        3,
        4,
        4,
        3,
        5,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Tibetan Mastiffs',
        4,
        3,
        3,
        4,
        3,
        3,
        3,
        1,
        3,
        5,
        3,
        3,
        3,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bedlington Terriers',
        3,
        3,
        3,
        1,
        3,
        1,
        3,
        3,
        3,
        4,
        3,
        3,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Greyhounds',
        4,
        3,
        4,
        2,
        1,
        1,
        1,
        3,
        3,
        3,
        4,
        3,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pulik',
        5,
        3,
        3,
        1,
        5,
        2,
        5,
        3,
        3,
        5,
        4,
        5,
        3,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Salukis',
        5,
        3,
        3,
        2,
        1,
        1,
        1,
        3,
        3,
        1,
        3,
        3,
        4,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Barbets',
        4,
        5,
        5,
        1,
        3,
        1,
        3,
        3,
        3,
        3,
        3,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Redbone Coonhounds',
        5,
        5,
        5,
        3,
        2,
        3,
        1,
        3,
        3,
        3,
        4,
        3,
        3,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Swedish Vallhunds',
        5,
        5,
        3,
        4,
        2,
        2,
        1,
        3,
        4,
        3,
        4,
        4,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Sealyham Terriers',
        5,
        3,
        3,
        3,
        3,
        1,
        3,
        4,
        3,
        5,
        4,
        4,
        3,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spanish Water Dogs',
        5,
        4,
        3,
        1,
        4,
        2,
        3,
        3,
        4,
        3,
        3,
        4,
        4,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Briards',
        3,
        3,
        3,
        1,
        4,
        2,
        5,
        3,
        3,
        4,
        3,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Berger Picards',
        3,
        3,
        3,
        3,
        1,
        1,
        3,
        3,
        3,
        4,
        3,
        4,
        4,
        2,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Entlebucher Mountain Dogs',
        5,
        3,
        5,
        3,
        1,
        2,
        1,
        3,
        3,
        5,
        4,
        3,
        5,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Treeing Walker Coonhounds',
        5,
        5,
        5,
        3,
        1,
        3,
        1,
        3,
        4,
        3,
        4,
        5,
        5,
        4,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Icelandic Sheepdogs',
        5,
        5,
        4,
        3,
        2,
        2,
        3,
        3,
        3,
        4,
        4,
        5,
        4,
        4,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Wirehaired Vizslas',
        5,
        5,
        3,
        3,
        1,
        2,
        1,
        5,
        5,
        3,
        4,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pumik',
        5,
        3,
        3,
        1,
        2,
        2,
        3,
        3,
        4,
        4,
        4,
        5,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Portuguese Podengo Pequenos',
        5,
        5,
        5,
        3,
        2,
        1,
        3,
        3,
        4,
        3,
        4,
        3,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (American Water)',
        3,
        3,
        3,
        1,
        3,
        1,
        3,
        3,
        3,
        3,
        3,
        5,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Retrievers (Curly-Coated)',
        5,
        5,
        3,
        2,
        1,
        1,
        1,
        3,
        4,
        4,
        4,
        4,
        4,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (Field)',
        5,
        5,
        4,
        3,
        2,
        2,
        3,
        4,
        3,
        3,
        4,
        5,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Lowchen',
        5,
        5,
        3,
        2,
        2,
        1,
        5,
        3,
        4,
        4,
        4,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Nederlandse Kooikerhondjes',
        5,
        3,
        3,
        3,
        2,
        2,
        3,
        4,
        3,
        3,
        4,
        5,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Affenpinschers',
        3,
        3,
        3,
        3,
        3,
        1,
        1,
        5,
        3,
        3,
        4,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Petits Bassets Griffons Vendeens',
        4,
        5,
        5,
        2,
        2,
        2,
        5,
        4,
        4,
        4,
        4,
        3,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Finnish Lapphunds',
        5,
        4,
        3,
        4,
        2,
        2,
        3,
        4,
        3,
        3,
        4,
        4,
        3,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Scottish Deerhounds',
        5,
        3,
        5,
        3,
        1,
        1,
        3,
        3,
        3,
        3,
        3,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Plott Hounds',
        5,
        3,
        5,
        2,
        1,
        2,
        1,
        4,
        3,
        4,
        4,
        4,
        4,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Norwegian Buhunds',
        5,
        3,
        3,
        3,
        2,
        1,
        3,
        3,
        3,
        4,
        4,
        3,
        4,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Glen of Imaal Terriers',
        5,
        3,
        3,
        2,
        3,
        2,
        3,
        3,
        3,
        3,
        4,
        3,
        3,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Setters (Irish Red and White)',
        5,
        5,
        5,
        2,
        2,
        2,
        3,
        5,
        3,
        3,
        5,
        5,
        5,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Ibizan Hounds',
        5,
        3,
        5,
        2,
        1,
        1,
        1,
        3,
        4,
        3,
        4,
        3,
        5,
        3,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (Sussex)',
        5,
        3,
        3,
        3,
        3,
        3,
        3,
        3,
        4,
        3,
        4,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Bergamasco Sheepdogs',
        3,
        3,
        3,
        1,
        1,
        2,
        5,
        3,
        3,
        4,
        3,
        3,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Spaniels (Irish Water)',
        5,
        3,
        3,
        1,
        3,
        2,
        3,
        3,
        5,
        3,
        4,
        5,
        3,
        1,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Polish Lowland Sheepdogs',
        5,
        3,
        3,
        2,
        4,
        2,
        5,
        3,
        4,
        5,
        5,
        4,
        3,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Otterhounds',
        5,
        3,
        3,
        2,
        2,
        3,
        3,
        4,
        3,
        3,
        4,
        4,
        3,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Kuvaszok',
        5,
        3,
        3,
        3,
        2,
        2,
        3,
        3,
        3,
        5,
        4,
        5,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Komondorok',
        5,
        3,
        2,
        1,
        4,
        2,
        5,
        3,
        3,
        5,
        3,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Cirnechi dell’Etna',
        4,
        3,
        4,
        1,
        1,
        1,
        1,
        3,
        4,
        3,
        4,
        4,
        3,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pharaoh Hounds',
        5,
        3,
        5,
        3,
        2,
        1,
        1,
        3,
        3,
        3,
        3,
        4,
        4,
        3,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Dandie Dinmont Terriers',
        4,
        3,
        3,
        2,
        3,
        1,
        3,
        4,
        3,
        4,
        3,
        4,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Pyrenean Shepherds',
        5,
        3,
        3,
        3,
        2,
        2,
        3,
        3,
        5,
        5,
        4,
        5,
        5,
        4,
        5
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Skye Terriers',
        5,
        3,
        3,
        3,
        3,
        1,
        5,
        3,
        3,
        3,
        3,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Canaan Dogs',
        3,
        3,
        3,
        4,
        2,
        1,
        1,
        3,
        3,
        4,
        3,
        4,
        3,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'American English Coonhounds',
        3,
        3,
        5,
        2,
        1,
        1,
        1,
        3,
        3,
        3,
        3,
        3,
        4,
        4,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Chinooks',
        4,
        5,
        5,
        3,
        3,
        1,
        3,
        3,
        3,
        4,
        4,
        4,
        3,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Finnish Spitz',
        5,
        5,
        4,
        3,
        2,
        2,
        3,
        3,
        3,
        3,
        4,
        3,
        5,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Grand Basset Griffon Vendeens',
        5,
        3,
        4,
        3,
        2,
        2,
        3,
        5,
        4,
        3,
        4,
        3,
        5,
        4,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Sloughis',
        4,
        3,
        3,
        3,
        1,
        1,
        1,
        2,
        3,
        2,
        3,
        3,
        4,
        2,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Harriers',
        5,
        5,
        5,
        3,
        1,
        2,
        1,
        4,
        4,
        3,
        4,
        4,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Cesky Terriers',
        4,
        5,
        3,
        2,
        2,
        1,
        3,
        4,
        3,
        3,
        4,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'American Foxhounds',
        3,
        5,
        5,
        3,
        1,
        1,
        1,
        3,
        3,
        3,
        3,
        3,
        4,
        5,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Azawakhs',
        3,
        3,
        3,
        2,
        2,
        1,
        1,
        1,
        3,
        3,
        3,
        2,
        3,
        1,
        3
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'English Foxhounds',
        5,
        5,
        5,
        3,
        1,
        2,
        1,
        4,
        4,
        3,
        4,
        4,
        4,
        5,
        4
    ) ON CONFLICT (name) DO NOTHING;

INSERT INTO
    dog_breed (
        name,
        affectionate_with_family,
        good_with_young_children,
        good_with_other_dogs,
        shedding_level,
        coat_grooming_frequency,
        drooling_level,
        coat_length,
        openness_to_strangers,
        playfulness_level,
        watchdog_protective_nature,
        adaptability_level,
        trainability_level,
        energy_level,
        barking_level,
        mental_stimulation_needs
    )
VALUES
    (
        'Norwegian Lundehunds',
        3,
        3,
        3,
        3,
        2,
        1,
        1,
        3,
        3,
        3,
        3,
        3,
        3,
        3,
        3
    ) ON CONFLICT (name) DO NOTHING;