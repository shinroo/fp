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