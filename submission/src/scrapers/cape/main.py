import logging
import json

from base.common import DogScrapingResult
from base.generic_scraper import GenericScraper

from bs4 import BeautifulSoup

def process_dogs_page_source(page_source: str) -> list:
    links = []

    soup = BeautifulSoup(page_source, 'html.parser')

    for article_tag in soup.find_all('article', {'class': 'category-dogs-for-adoption'}):
        for anchor_tag in article_tag.find_all('a', {'class': 'elementor-post__thumbnail__link'}):
            links.append(anchor_tag['href'])

    return links

def process_dog_subpage_source(page_source: str):
    soup = BeautifulSoup(page_source, 'html.parser')

    print(soup.prettify())

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

    for sub_page in sub_pages:
        logging.info(f"Processing sub page: {sub_page}")

        scraper = GenericScraper(sub_page)

        sub_page_source = ""
        try:
            logging.info(f"Scraping sub page: {sub_page}")
        except Exception as e:
            logging.error(f"Failed to scrape sub page: {e}")
        finally:
            scraper.close()

        if not sub_page_source:
            continue

        try:
            logging.info(f"Processing sub page source: {sub_page}")
            process_dog_subpage_source(sub_page_source)
        except Exception as e:
            logging.error(f"Failed to process sub page source: {e}")

    
    