from enum import Enum

class StrEnum(str, Enum):
    pass

class Gender(StrEnum):
    MALE = "Male"
    FEMALE = "Female"

class LifeStage(StrEnum):
    PUPPY = "Puppy"
    ADULT = "Adult"
    SENIOR = "Senior"

class DogScrapingResult:
    def __init__(self, identifier: str, name: str, gender: Gender, life_stage: LifeStage, image_url: str, spca_id):
        self.identifier = identifier
        self.name = name
        self.gender = gender
        self.life_stage = life_stage
        self.image_url = image_url
        self.spca_id = spca_id

class DogBreed:
    def __init__(
        self, 
        name: str,
        aliases: str,
        affectionate_with_family: int,
        good_with_young_children: int,
        good_with_other_dogs: int,
        shedding_level: int,
        coat_grooming_frequency: int,
        drooling_level: int,
        coat_length: int,
        openness_to_strangers: int,
        playfulness_level: int,
        watchdog_protective_nature: int,
        adaptability_level: int,
        trainability_level: int,
        energy_level: int,
        barking_level: int,
        mental_stimulation_needs: int
    ):
        self.name = name
        self.aliases = aliases
        self.affectionate_with_family = affectionate_with_family
        self.good_with_young_children = good_with_young_children
        self.good_with_other_dogs = good_with_other_dogs
        self.shedding_level = shedding_level
        self.coat_grooming_frequency = coat_grooming_frequency
        self.drooling_level = drooling_level
        self.coat_length = coat_length
        self.openness_to_strangers = openness_to_strangers
        self.playfulness_level = playfulness_level
        self.watchdog_protective_nature = watchdog_protective_nature
        self.adaptability_level = adaptability_level
        self.trainability_level = trainability_level
        self.energy_level = energy_level
        self.barking_level = barking_level
        self.mental_stimulation_needs = mental_stimulation_needs

    def to_pgvector(self) -> str:
        return f"[{self.affectionate_with_family}, {self.good_with_young_children}, {self.good_with_other_dogs}, {self.shedding_level}, {self.coat_grooming_frequency}, {self.drooling_level}, {self.coat_length}, {self.openness_to_strangers}, {self.playfulness_level}, {self.watchdog_protective_nature}, {self.adaptability_level}, {self.trainability_level}, {self.energy_level}, {self.barking_level}, {self.mental_stimulation_needs}]"
    
class SpecificAlert:
    def __init__(self, id, breed, gender: Gender, life_stage: LifeStage, email):
        self.id = id
        self.breed = breed
        self.gender = gender
        self.life_stage = life_stage
        self.email = email

class SimilarityAlert:
    def __init__(self, id, threshold, email, account_id):
        self.id = id
        self.threshold = threshold
        self.email = email
        self.account_id = account_id

class Dog:
    def __init__(self, name, breed, spca_name, image_url):
        self.name = name
        self.breed = breed
        self.spca_name = spca_name
        self.image_url = image_url