import logging
import json
import re

from base.common import DogScrapingResult, LifeStage, Gender
from base.dog_repository import DogRepository
from base.generic_scraper import GenericScraper
from base.dog_breed_repository import DogBreedRepository
from base.breed_identifier import BreedIdentifier

from bs4 import BeautifulSoup

def none_safe_text(element):
    if element is None:
        return None
    return element.text

def none_to_empty_safe_text(element):
    if element is None:
        return ""
    return element.text

def none_safe_src(element):
    if element is None:
        return None
    return element['src']

def is_valid(element):
    for key, value in element.items():
        if value is None:
            return False
    return True

def age_to_lifestage(age):
    if "week" in age:
        return LifeStage.PUPPY.value
    elif "year" in age:
        return LifeStage.ADULT.value
    else:
        return LifeStage.SENIOR.value
    
def clean_gender(gender):
    if "Female" in gender or "female" in gender:
        return Gender.FEMALE.value
    return Gender.MALE.value

def handle_blank_name(name):
    if name is None or name == "":
        return "Unknown"
    return name

def process_page_source(page_source: str):
    soup = BeautifulSoup(page_source, 'html.parser')

    # Initialize a list to hold all the extracted data
    data = []

    pet_divs = soup.find_all('div', {'class': 'oxy-superbox-secondary'})

    dog_breed_repository = DogBreedRepository()
    breed_identifier = BreedIdentifier(dog_breed_repository.get_all())

    for pet_div in pet_divs:
        pet_internal_id = pet_div['id'].split('-')[-1]

        name_span = pet_div.find('span', {'id': f'span-360-359-{pet_internal_id}'})
        breed_span = pet_div.find('span', {'id': f'span-1288-359-{pet_internal_id}'})
        gender_span = pet_div.find('span', {'id': f'span-372-359-{pet_internal_id}'})
        age_span = pet_div.find('span', {'id': f'span-376-359-{pet_internal_id}'})
        kennel_span = pet_div.find('span', {'id': f'span-380-359-{pet_internal_id}'})
        img = soup.find('img', {'id': f'image-355-359-{pet_internal_id}'})

        identified_breed_name = breed_identifier.identify(none_to_empty_safe_text(breed_span))
        identified_breed = breed_identifier.get_breed_by_name(identified_breed_name)

        pet_data = {
            'name': none_safe_text(name_span),
            'breed': none_safe_text(breed_span),
            'gender': none_safe_text(gender_span),
            'age': none_safe_text(age_span),
            'kennel': none_safe_text(kennel_span),
            'image_url': none_safe_src(img),
            'vector': identified_breed.to_pgvector()
        }
        
        if is_valid(pet_data) and len(pet_data['kennel']) == 4:
            data.append(pet_data)

    print(json.dumps(data, indent=4))

    try:
        dog_repository = DogRepository()
    except Exception as e:
        logging.error(f"Failed to create DogRepository: {e}")
        return

    for pet_data in data:
        life_stage = age_to_lifestage(pet_data['age'])
        gender = clean_gender(pet_data['gender'])
        name = handle_blank_name(pet_data['name'])

        try:
            dog_repository.create(pet_data["kennel"], name, str(gender), str(life_stage), pet_data["image_url"], "3b923b5e-994a-4b41-a1fd-ed060e983f82", pet_data["vector"])
        except Exception as e:
            logging.error(f"Failed to insert data: {e}")

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://lscspca.org.za/animals-for-adoption/")

    page_source = ""
    try:
        logging.info("Scraping")
        page_source = scraper.scrape("oxy-superbox-secondary")
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        scraper.close()

    try:
        logging.info("Processing page source")
        process_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process page source: {e}")
    
    