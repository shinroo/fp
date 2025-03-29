from .common import DogBreed

import nltk
from nltk import word_tokenize

nltk.download('punkt')
nltk.download('punkt_tab')

class BreedIdentifier:
    similarity_threshold = 0.7

    # If we can't identify a breed
    # we'll use this as a fallback
    #
    # this has "safe" default values
    fallback_breed = DogBreed(
        "Unknown",
        "",
        1,
        1,
        1,
        5,
        5,
        5,
        5,
        1,
        1,
        1,
        1,
        1,
        5,
        5,
        5
    )

    def __init__(self, dog_breeds: list[DogBreed]):
        self.dog_breeds = dog_breeds
        self.dog_breeds_map = {breed.name: breed for breed in dog_breeds}
        self.tokenization_map = {}

        for breed in self.dog_breeds:
            breed_words = word_tokenize(breed.name) + word_tokenize(breed.aliases)
            breed_words = [word.lower() for word in breed_words]
            breed_words = [word for word in breed_words if word.isalnum()]
            self.tokenization_map[breed.name] = breed_words

    @staticmethod
    def string_similarity(s1, s2):
        distance = nltk.edit_distance(s1, s2)
        max_len = max(len(s1), len(s2))
        similarity = 1 - (distance / max_len)
        return similarity
    
    def get_breed_by_name(self, name):
        found = self.dog_breeds_map.get(name)
        if found is not None:
            return found
        return self.fallback_breed

    def identify(self, text):
        text_words = word_tokenize(text)
        text_words = [word.lower() for word in text_words]
        text_words = [word for word in text_words if word.isalnum()]

        matches = []

        for name, breed_words in self.tokenization_map.items():
            for text_word in text_words:
                for breed_word in breed_words:
                    similarity = self.string_similarity(text_word, breed_word)
                    if similarity >= self.similarity_threshold:
                        matches.append({
                            "name": name,
                            "similarity": similarity
                        })
        
        max_similarity = 0
        max_name = ""

        for match in matches:
            if match["similarity"] >= max_similarity:
                max_similarity = match["similarity"]
                max_name = match["name"]

        return max_name