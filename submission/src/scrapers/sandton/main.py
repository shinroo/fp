import logging
import json
import re

from base.common import DogScrapingResult, LifeStage, Gender
from base.generic_scraper import GenericScraper
from base.dog_repository import DogRepository
from base.dog_breed_repository import DogBreedRepository
from base.breed_identifier import BreedIdentifier

from bs4 import BeautifulSoup

def extract_number(text):
    match = re.search(r'\d+', text)
    return int(match.group()) if match else None

def process_age(s):
    if "month" in s or "months" in s:
        num = extract_number(s)
        if num:
            if num < 12:
                return LifeStage.PUPPY
        return LifeStage.ADULT
    if "year" in s or "years" in s:
        return LifeStage.ADULT
    return LifeStage.SENIOR

def process_sex(s):
    if s == "":
        return Gender.FEMALE
    if "Female" in s or "female" in s:
        return Gender.FEMALE
    return Gender.MALE

    
def process_dogs_page_source(page_source: str) -> list:
    links = []

    soup = BeautifulSoup(page_source, 'html.parser')

    for link_div in soup.find_all("div", {"class": "raven-column-link"}):
        links.append(link_div["data-raven-link"])

    return links

def process_dog_subpage_source(page_source: str, sub_page_url: str, breed_identifier: BreedIdentifier, dog_repository: DogRepository):
    soup = BeautifulSoup(page_source, 'html.parser')

    dynamic_field_divs = soup.find_all("div", {"class": "jet-listing-dynamic-field__content"})

    # name
    try:
        name_span = soup.find("span", {"class": "raven-heading-title"})
        name = ""
        if name_span:
            name = name_span.text
        print(name)
    except Exception as e:
        logging.error(f"Failed to get name: {e}")
        return

    # image
    try:
        image_div = soup.find("div", {"class": "raven-image"})
        img = image_div.find("img")
        image_url = ""
        if img:
            image_url = img["data-src"]
        print(image_url)
    except Exception as e:
        logging.error(f"Failed to get image: {e}")
        return
    
    # img id
    try:
        number = 0
        for class_name in img['class']:
            match = re.search(r'wp-image-(\d+)', class_name)
            if match:
                number = match.group(1)
                print(number)
                break
    except Exception as e:
        logging.error(f"Failed to get image id: {e}")
        return

    # age
    try:
        age = ""
        age_div = dynamic_field_divs[2]
        print(age_div)
        if age_div:
            age = process_age(age_div.text)
        print(age)
    except Exception as e:
        logging.error(f"Failed to get age: {e}")
        return

    # sex
    try:
        sex = ""
        sex_div = dynamic_field_divs[3]
        print(sex_div)
        if sex_div:
            sex = process_sex(sex_div.text)
        print(sex)
    except Exception as e:
        logging.error(f"Failed to get sex: {e}")
        return

    # breed
    try:
        breed = ""
        breed_div = dynamic_field_divs[4]
        print(breed_div)
        if breed_div:
            breed = breed_div.text
            breed = breed.replace("-", " ")
        print(breed)
    except Exception as e:
        logging.error(f"Failed to get breed: {e}")
        return
    
    try:
        identified_breed_name = breed_identifier.identify(breed)
        identified_breed = breed_identifier.get_breed_by_name(identified_breed_name)
    except Exception as e:
        logging.error(f"Failed to identify breed: {e}")
        return

    print(json.dumps({
        "name": name.capitalize(),
        "id": f"sandton-{number}" if number != 0 else f"sandton-{name}",
        "image_url": image_url,
        "age": age.value,
        "breed": breed,
        "gender": sex.value,
        "matched_breed": identified_breed.name
    }, indent=4))

    if breed in ["DSH", "DMH", "DLH"]:
        logging.info(f"Skipping cat: {breed}")
        return

    try:
        dog_repository.create(f"sandton-{number}" if number != 0 else f"sandton-{name}", name.capitalize(), str(sex.value), str(age.value), image_url, "16db9057-21d2-4a32-bb35-8ebd304695ea", "", identified_breed.name)
    except Exception as e:
        logging.error(f"Failed to insert data: {e}")

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://sandtonspca.org.za/adoptions/")

    page_source = ""
    try:
        logging.info("Scraping dogs page")
        page_source = scraper.scrape("raven-column-link ")
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        scraper.close()

    sub_pages = []
    try:
        logging.info("Processing dogs page source")
        sub_pages = process_dogs_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process dogs page source: {e}")

    try:
        dog_breed_repository = DogBreedRepository()
        breed_identifier = BreedIdentifier(dog_breed_repository.get_all())
    except Exception as e:
        logging.error(f"Failed to create breed identifir instance: {e}")
        raise

    try:
        dog_repository = DogRepository()
    except Exception as e:
        logging.error(f"Failed to create DogRepository: {e}")
        raise

    dogs = []
    for sub_page in sub_pages:
        logging.info(f"Processing sub page: {sub_page}")

        scraper = GenericScraper(sub_page)

        sub_page_source = ""
        try:
            logging.info(f"Scraping sub page: {sub_page}")
            sub_page_source = scraper.scrape("raven-image")
        except Exception as e:
            logging.error(f"Failed to scrape sub page: {e}")
        finally:
            scraper.close()

        if not sub_page_source:
            continue

        try:
            logging.info(f"Processing sub page source: {sub_page}")
            process_dog_subpage_source(sub_page_source, sub_page, breed_identifier, dog_repository)
        except Exception as e:
            logging.error(f"Failed to process sub page source: {e}")
    