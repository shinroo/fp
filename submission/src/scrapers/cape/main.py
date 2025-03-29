import logging
import json
import re

from base.common import DogScrapingResult, LifeStage, Gender
from base.generic_scraper import GenericScraper
from base.dog_repository import DogRepository
from base.dog_breed_repository import DogBreedRepository
from base.breed_identifier import BreedIdentifier

from bs4 import BeautifulSoup

def extract_name_and_id(url):
    # Match the last part of the URL before optional trailing slash and optional '-z'
    match = re.search(r'/([a-zA-Z-]+)-([a-zA-Z]\d+)(?:-z)?/?$', url)
    if match:
        name = match.group(1)
        url_id = match.group(2)
        return name, url_id
    return None, None

def parse_age(s):
    if "mths" in s or "mnths" in s:
        return LifeStage.PUPPY
    return LifeStage.ADULT

def parse_gender(s):
    if "Female" in s or "female" in s:
        return Gender.FEMALE
    return Gender.MALE
    
def process_dogs_page_source(page_source: str) -> list:
    links = []

    soup = BeautifulSoup(page_source, 'html.parser')

    for article_tag in soup.find_all('article', {'class': 'category-dogs-for-adoption'}):
        for anchor_tag in article_tag.find_all('a', {'class': 'elementor-post__thumbnail__link'}):
            links.append(anchor_tag['href'])

    return links

def process_dog_subpage_source(page_source: str, sub_page_url: str, breed_identifier: BreedIdentifier, dog_repository: DogRepository):
    soup = BeautifulSoup(page_source, 'html.parser')

    name, url_id = extract_name_and_id(sub_page_url)

    attribute_id_map = {
        "age": "32b83ea3",
        "breed": "4eeed9c3",
        "gender": "22f2b1f6",
        "image_url": "71e74449"
    }

    # image
    try:
        image_div = soup.find("div", {"data-id": attribute_id_map["image_url"]})
        img = image_div.find("img")
        image_url = img["nitro-lazy-src"]
    except Exception as e:
        logging.error(f"Failed to get image_url: {e}")
        return

    # age
    try:
        age_section = soup.find("section", {"data-id": attribute_id_map["age"]})
        age_p_tags = age_section.find_all("p")
        age = ""
        if len(age_p_tags) == 2:
            age = age_p_tags[1].text
    except Exception as e:
        logging.error(f"Failed to get age: {e}")
        return
    
    # breed
    try:
        breed_section = soup.find("section", {"data-id": attribute_id_map["breed"]})
        breed_p_tags = breed_section.find_all("p")
        breed = ""
        if len(breed_p_tags) == 2:
            breed = breed_p_tags[1].text
    except Exception as e:
        logging.error(f"Failed to get breed: {e}")
        return

    # gender
    try:
        gender_section = soup.find("section", {"data-id": attribute_id_map["gender"]})
        gender_p_tags = gender_section.find_all("p")
        print(gender_p_tags)
        gender = ""
        if len(gender_p_tags) == 2:
            gender = gender_p_tags[1].text
        else:
            gender_div_tags = gender_section.find_all("div", {"class": "elementor-widget-container"})
            if len(gender_div_tags) == 2:
                gender = gender_div_tags[1].text
    except Exception as e:
        logging.error(f"Failed to get gender: {e}")
        return
    
    try:
        identified_breed_name = breed_identifier.identify(breed)
        identified_breed = breed_identifier.get_breed_by_name(identified_breed_name)
    except Exception as e:
        logging.error(f"Failed to identify breed: {e}")
        return

    print(json.dumps({
        "name": name.capitalize(),
        "id": url_id,
        "image_url": image_url,
        "age": parse_age(age),
        "breed": breed,
        "gender": parse_gender(gender),
        "matched_breed": identified_breed.name
    }, indent=4))

    try:
        dog_repository.create(url_id, name.capitalize(), str(parse_gender(gender).value), str(parse_age(age).value), image_url, "a2c421f2-d4a0-4340-94fd-8ff804b07a06", "", identified_breed.name)
    except Exception as e:
        logging.error(f"Failed to insert data: {e}")

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://capespca.co.za/adopt/dogs/")

    page_source = ""
    try:
        logging.info("Scraping dogs page")
        page_source = scraper.scrape("category-dogs-for-adoption")
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
            sub_page_source = scraper.scrape("elementor-element-32b83ea3")
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

    
    