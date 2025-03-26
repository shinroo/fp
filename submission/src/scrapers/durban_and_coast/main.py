import logging
import json
import re

from base.common import DogScrapingResult, LifeStage, Gender
from base.dog_repository import DogRepository
from base.generic_scraper import GenericScraper
from base.dog_breed_repository import DogBreedRepository
from base.breed_identifier import BreedIdentifier

from bs4 import BeautifulSoup

# Define function to determine gender based on description
def sex_from_description(description: str):
    if "boy" in description or "Boy" in description:
        return "male"
    elif "girl" in description or "Girl" in description:
        return "female"
    return None

# Define function to determine age based on description
def age_from_description(description: str):
    # Match ages like X years
    age_pattern_years = r'\b(\d+)\s+years?\b'
    match_years = re.search(age_pattern_years, description)
    if match_years:
        return float(match_years.group(1))
    
    # Match ages like X month
    age_pattern_months = r'\b(\d+)\s+month?\b'
    match_month = re.search(age_pattern_months, description)
    if match_month:
        return float(int(match_month.group(1))/12)
    
    return None

# Define a function to extract the kennel number from the description
def kennel_from_description(description: str):
    kennel_pattern = r'\bKennel\s+([A-Z]\d+)\b'
    match = re.search(kennel_pattern, description)
    if match:
        return match.group(1)
    return None

# Define a function to extract the reference number from the description
def reference_from_description(description: str):
    reference_pattern = r'\bRef:\s*(\d+)\b'
    match = re.search(reference_pattern, description)
    if match:
        return match.group(1)
    return None

# Define a function to extract the part of a description that contains "is a"
def extract_is_a(description: str):
    try:
        if "He is" in description:
            extracted = description.split("He is")[1].split(".")[0]
            extracted = extracted.replace("old", "")
            extracted = extracted.replace("year", "")
            extracted = extracted.replace("month", "")
            extracted = extracted.replace("He is", "")
            return extracted
        elif "She is" in description:
            extracted = description.split("She is")[1].split(".")[0]
            extracted = extracted.replace("old", "")
            extracted = extracted.replace("year", "")
            extracted = extracted.replace("month", "")
            extracted = extracted.replace("She is", "")
            return extracted
        else:
            return ""
    except:
        return ""

def process_page_source(page_source: str):
    soup = BeautifulSoup(page_source, 'html.parser')

    # Initialize a list to hold all the extracted data
    data = []

    dogs_section = soup.find('section', {'id': 'gallery1'})

    dog_links = dogs_section.find_all('a', {'data-gal': 'prettyPhoto[adopt]'})

    dog_breed_repository = DogBreedRepository()
    breed_identifier = BreedIdentifier(dog_breed_repository.get_all())

    for dog_link in dog_links:
        identified_breed_name = breed_identifier.identify(extract_is_a(dog_link['title']))
        identified_breed = breed_identifier.get_breed_by_name(identified_breed_name)

        data.append({
            'name': dog_link.find('h3').text,
            'image': f"https://spcadbn.org.za/{dog_link.find('img')['src']}",
            'description': extract_is_a(dog_link['title']),
            'sex': sex_from_description(dog_link['title']),
            'age': age_from_description(dog_link['title']),
            'kennel': kennel_from_description(dog_link['title']),
            'reference': reference_from_description(dog_link['title']),
            'vector': identified_breed.to_pgvector(),
            'matched_breed': identified_breed.name
        })

    try:
        dog_repository = DogRepository()
    except Exception as e:
        logging.error(f"Failed to create DogRepository: {e}")
        return

    for pet_detail in data:
        life_stage = LifeStage.ADULT.value
        if pet_detail["age"] < 1:
            life_stage = LifeStage.PUPPY.value
        if pet_detail["age"] > 10:
            life_stage = LifeStage.SENIOR.value
        
        gender = Gender.FEMALE.value
        if pet_detail["sex"] == "male":
            gender = Gender.MALE.value

        try:
            dog_repository.create(pet_detail["reference"], pet_detail["name"], str(gender), str(life_stage), pet_detail["image"], "c0054c53-c8aa-4124-a7e4-ba379028de4d", pet_detail["vector"], pet_detail["matched_breed"])
        except Exception as e:
            logging.error(f"Failed to insert data: {e}")

    data = json.dumps(data, indent=4)
    print(data)

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://www.spcadbn.org.za/adopt.php")

    page_source = ""
    try:
        logging.info("Scraping")
        page_source = scraper.scrape("gallery")
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        scraper.close()

    try:
        logging.info("Processing page source")
        process_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process page source: {e}")
    
    