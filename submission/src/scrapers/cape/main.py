import logging
import json

from base.common import DogScrapingResult
from base.generic_scraper import GenericScraper

from bs4 import BeautifulSoup

def process_page_source(page_source: str):
    return

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)

    logging.info("Scraper starting")

    scraper = GenericScraper("https://capespca.co.za/adopt/dogs/")

    page_source = ""
    try:
        logging.info("Scraping")
        page_source = scraper.scrape("category-dogs-for-adoption")
        print(page_source)
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        scraper.close()

    try:
        logging.info("Processing page source")
        process_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process page source: {e}")
    
    