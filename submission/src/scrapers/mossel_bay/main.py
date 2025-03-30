import logging
import json
import re

from base.common import DogScrapingResult, LifeStage, Gender
from base.dog_repository import DogRepository
from base.generic_scraper import GenericScraper
from base.dog_breed_repository import DogBreedRepository
from base.breed_identifier import BreedIdentifier

from bs4 import BeautifulSoup

def parse_age(description):
    age = description.lower()
    if "week" in age or "weeks" in age:
        return LifeStage.PUPPY
    if "month" in age or "months" in age:
        return LifeStage.PUPPY
    if "year" in age or "years" in age:
        return LifeStage.ADULT
    if "adult" in age:
        return LifeStage.ADULT
    if "puppy" in age:
        return LifeStage.PUPPY
    return LifeStage.SENIOR

def parse_gender(description):
    gender = description.lower()
    if "girl" in gender or "female" in gender:
        return Gender.FEMALE
    if "boy" in gender or "male" in gender:
        return Gender.MALE
    return Gender.MALE

def process_page_source(page_source: str):
    soup = BeautifulSoup(page_source, 'html.parser')

    # Initialize a list to hold all the extracted data
    data = []

    try:
        dog_breed_repository = DogBreedRepository()
        breed_identifier = BreedIdentifier(dog_breed_repository.get_all())
    except Exception as e:
        logging.error(f"Failed to create breed identifir instance: {e}")
        raise

    figures = soup.find_all("figure", {"class": "gallery-item"})

    for figure in figures:
        fig_captions = figure.find_all("figcaption")
        if fig_captions:
            if len(fig_captions) == 1:
                print(figure)

                # id
                data_id = fig_captions[0]["id"].replace("gallery", "mossel-bay")

                # image
                image_url = figure.find("a")["href"]

                # description
                description = fig_captions[0].text

                # breed
                try:
                    identified_breed_name = breed_identifier.identify(description.replace("old", "").replace("Old", ""))
                    identified_breed = breed_identifier.get_breed_by_name(identified_breed_name)
                except Exception as e:
                    logging.error(f"Failed to identify breed: {e}")
                    return

                data.append({
                    "id": data_id,
                    "image_url": image_url,
                    "description": description,
                    "breed": identified_breed.name,
                    "life_stage": parse_age(description),
                    "gender": parse_gender(description)
                })

    print(json.dumps(data, indent=4))

    try:
        dog_repository = DogRepository()
    except Exception as e:
        logging.error(f"Failed to create DogRepository: {e}")
        return

    for pet_data in data:
        try:
            dog_repository.create(pet_data["id"], f"Dog #{pet_data['id'].split('-')[-1]}", str(pet_data["gender"].value), str(pet_data["life_stage"].value), pet_data["image_url"], "a9e9e834-ef83-4009-b0c5-548f04440f2a", "", pet_data["breed"])
        except Exception as e:
            logging.error(f"Failed to insert data: {e}")

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://www.grspca.org/mossel-bay-dogs/")

    page_source = ""
    try:
        logging.info("Scraping")
        page_source = scraper.scrape("gallery-item")
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        scraper.close()

    try:
        logging.info("Processing page source")
        process_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process page source: {e}")
    
    