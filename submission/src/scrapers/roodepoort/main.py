import logging
import json
import re

from base.common import DogScrapingResult, LifeStage, Gender
from base.dog_repository import DogRepository
from base.generic_scraper import GenericScraper
from base.dog_breed_repository import DogBreedRepository
from base.breed_identifier import BreedIdentifier

from bs4 import BeautifulSoup

def parse_gender(s):
    if "Male" in s:
        return Gender.MALE
    return Gender.FEMALE

def parse_life_stage(s):
    if "Puppy" in s or "Month" in s or "Months" in s:
        return LifeStage.PUPPY
    if "Adult" in s:
        return LifeStage.ADULT
    return LifeStage.SENIOR

def strip_gender(s, gender):
    return s.replace(gender.value, "")

def extract_numeric_id(url):
    match = re.search(r'/(\d+)_', url)
    if match:
        return match.group(1)
    return None

def process_page_source(page_source: str):
    soup = BeautifulSoup(page_source, 'html.parser')

    # Initialize a list to hold all the extracted data
    data = []

    dog_breed_repository = DogBreedRepository()
    breed_identifier = BreedIdentifier(dog_breed_repository.get_all())

    contents = soup.find_all("div", class_="et_pb_blurb_content")

    for content_div in contents:
        img = content_div.find("img")
        img_url = ""
        if img:
            img_url = img.attrs["src"]

        description_div = content_div.find("div", class_="et_pb_blurb_description")

        if description_div:
            p_tags = description_div.find_all("p")
            if len(p_tags) != 2:
                continue

            gender = parse_gender(p_tags[0].text)
            life_stage = parse_life_stage(p_tags[1].text)
            breed_str = strip_gender(p_tags[0].text, gender)

            identified_breed_name = breed_identifier.identify(breed_str)
            identified_breed = breed_identifier.get_breed_by_name(identified_breed_name)

            numeric_id = extract_numeric_id(img_url)
            if numeric_id is None:
                continue

            data.append({
                "id": f"roodepoort-{numeric_id}",
                "gender": gender.value,
                "life_stage": life_stage.value,
                "breed": breed_str,
                "image_url": img_url,
                "matched_breed": identified_breed.name
            })

    try:
        dog_repository = DogRepository()
    except Exception as e:
        logging.error(f"Failed to create DogRepository: {e}")
        return
    
    for pet_data in data:
        try:
            dog_repository.create(pet_data["id"], "Unknown", gender, life_stage, pet_data["image_url"], "833506bb-2f21-48e9-afe6-f12b4be04141", "", pet_data["matched_breed"])
        except Exception as e:
            logging.error(f"Failed to insert data: {e}")

    print(json.dumps(data, indent=4))

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://spcaroodepoort.org.za/adoption-process-dogs/")

    page_source = ""
    try:
        logging.info("Scraping")
        page_source = scraper.scrape("et_pb_blurb_content")
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        scraper.close()

    try:
        logging.info("Processing page source")
        process_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process page source: {e}")
    
    