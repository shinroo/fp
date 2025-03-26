import logging
import json

from base.common import DogScrapingResult, LifeStage, Gender
from base.dog_repository import DogRepository
from base.generic_scraper import GenericScraper
from base.dog_breed_repository import DogBreedRepository
from base.breed_identifier import BreedIdentifier

from bs4 import BeautifulSoup

def age_to_lifestage(age):
    if "week" in age or "month" in age:
        return LifeStage.PUPPY.value
    elif "year" in age or "yr" in age or "years" in age or "yrs" in age: 
        return LifeStage.ADULT.value
    else:
        return LifeStage.SENIOR.value
    
def clean_gender(gender):
    if "Female" in gender or "female" in gender:
        return Gender.FEMALE.value
    return Gender.MALE.value

def strip_query_from_image(image_url):
    return image_url.split("/v1")[0]

def process_page_source(page_source: str):
    soup = BeautifulSoup(page_source, 'html.parser')

    # Initialize a list to hold all the extracted data
    data = []

    pet_divs = soup.find_all('div', {'class': 'wixui-repeater__item'})

    attribute_index_to_label = {
        0: "kennel",
        2: "breed",
        3: "age",
        4: "sex"
    }

    dog_breed_repository = DogBreedRepository()
    breed_identifier = BreedIdentifier(dog_breed_repository.get_all())

    for pet_div in pet_divs:
        data_divs = pet_div.find_all('div', {'class': 'wixui-rich-text'})
        if len(data_divs) < 3:
            continue
        name_div = data_divs[0]
        data_div = data_divs[2]

        # Get name
        innermost_span = name_div.find_all('span')[-1]
        name = innermost_span.text

        pet_data = {
            "name": name
        }

        # Get image
        img_tag = pet_div.find('img')
        if img_tag:
            pet_data["image"] = strip_query_from_image(img_tag['src'])
        
        # Get attributes
        p_tags = data_div.find_all('p')
        for index, p_tag in enumerate(p_tags):
            if index not in attribute_index_to_label:
                continue

            value_span = p_tag.find_all('span')[-1]
            value = value_span.text
            pet_data[attribute_index_to_label[index]] = value

        identified_breed_name = breed_identifier.identify(pet_data["breed"])
        identified_breed = breed_identifier.get_breed_by_name(identified_breed_name)

        pet_data["vector"] = identified_breed.to_pgvector()
        pet_data["matched_breed"] = identified_breed.name

        data.append(pet_data)

    print(json.dumps(data, indent=4))

    try:
        dog_repository = DogRepository()
    except Exception as e:
        logging.error(f"Failed to create DogRepository: {e}")
        return

    for pet_data in data:
        life_stage = age_to_lifestage(pet_data['age'])
        gender = clean_gender(pet_data['sex'])

        try:
            dog_repository.create(f"randburg-{pet_data['kennel']}", pet_data["name"], str(gender), str(life_stage), pet_data["image"], "a7c6c894-aa83-4fdd-8dc2-512b425128b2", pet_data["vector"], pet_data["matched_breed"])
        except Exception as e:
            logging.error(f"Failed to insert data: {e}")

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://www.spca-rbg.org.za/our-dogs")

    page_source = ""
    try:
        logging.info("Scraping")
        page_source = scraper.scrape("wixui-repeater__item")
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        scraper.close()

    try:
        logging.info("Processing page source")
        process_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process page source: {e}")
    
    