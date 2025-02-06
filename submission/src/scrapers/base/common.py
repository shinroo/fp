from enum import Enum

class StrEnum(str, Enum):
    pass

class SPCA(StrEnum):
    AMANZIMTOTI = "Amanzimtoti"
    CAPE = "Cape"

class Gender(StrEnum):
    MALE = "Male"
    FEMALE = "Female"

class DogScrapingResult:
    def __init__(self, identifier: str, name: str, gender: Gender, age: int, image: str, spca: SPCA):
        self.identifier = identifier
        self.name = name
        self.gender = gender
        self.age = age
        self.image = image
        self.spca = spca