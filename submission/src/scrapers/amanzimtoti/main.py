import logging
import json

from base.common import DogScrapingResult
from base.generic_scraper import GenericScraper

from bs4 import BeautifulSoup

def process_page_source(page_source: str):
    soup = BeautifulSoup(page_source, 'html.parser')

    # Initialize a list to hold all the extracted data
    data = []

    # Find all the main tables
    pet_table = soup.find('div', {'id': 'sec1'})

    # Find all images:
    images = []
    ref_numbers = []
    for img_anchor_tag in pet_table.find_all('a', {'rel': 'lightbox-1'}):
        images.append(f"https://www.spcatoti.co.za/{img_anchor_tag['href']}")
        ref_numbers.append(img_anchor_tag['title'])

    # Find corresponding pet details:
    pet_details = []
    for image, ref_number in zip(images, ref_numbers):
        # Find the <td> element with the known text
        td_element = soup.find('td', text=ref_number)

        # Check if the <td> element was found
        if td_element:
            # Navigate to the closest parent <table> element
            parent_table = td_element.find_parent('table')
            
            # Check if the parent <table> element was found
            if parent_table:
                data_tds = parent_table.find_all('td')

                details = {
                    "number": ref_number,
                    "name": data_tds[5].text,
                    "breed": data_tds[7].text,
                    "approximate_age": data_tds[9].text,
                    "gender": data_tds[11].text,
                    "image": image
                }
                pet_details.append(details)
            else:
                raise Exception("No parent <table> element found.")
        else:
            raise Exception(f"No <td> element with text '{ref_number}' found.")

    data = json.dumps(pet_details, indent=4)
    print(data)

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://www.spcatoti.co.za/adopt-a-pet.php?id=ad")

    page_source = ""
    try:
        logging.info("Scraping")
        page_source = scraper.scrape("mypets")
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        scraper.close()

    try:
        logging.info("Processing page source")
        process_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process page source: {e}")
    
    