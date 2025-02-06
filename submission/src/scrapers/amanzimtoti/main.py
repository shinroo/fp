import logging
import json

from selenium.webdriver import Chrome
from selenium.webdriver import ChromeOptions
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from bs4 import BeautifulSoup

# Add this function to set headers
def set_headers(driver):
    headers = {
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "accept-encoding": "gzip, deflate, br, zstd",
        "accept-language": "en-ZA,en-GB;q=0.9,en-US;q=0.8,en;q=0.7",
        "cookie": "PHPSESSID=867cd1fcf18688b18d8e39778fd301e9",
        "dnt": "1",
        "priority": "u=0, i",
        "referer": "https://www.spcatoti.co.za/",
        "sec-ch-ua": '"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"',
        "sec-ch-ua-mobile": "?1",
        "sec-ch-ua-platform": '"Android"',
        "sec-fetch-dest": "document",
        "sec-fetch-mode": "navigate",
        "sec-fetch-site": "same-origin",
        "sec-fetch-user": "?1",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Mobile Safari/537.36"
    }

    # Use Chrome DevTools Protocol to set headers
    driver.execute_cdp_cmd("Network.enable", {})
    driver.execute_cdp_cmd("Network.setExtraHTTPHeaders", {"headers": headers})

def scrape(driver) -> str:
    try:
        driver.get("https://www.spcatoti.co.za/adopt-a-pet.php?id=ad")
    except Exception as e:
        logging.error(f"Failed to get website: {e}")
        raise e

    html = ""

    try:
        # Wait for a specific element to be present (important for dynamic content)
        element_present = EC.presence_of_element_located((By.CLASS_NAME, 'mypets'))
        WebDriverWait(driver, 10).until(element_present) # Wait up to 10 seconds
        html = driver.page_source
        logging.info("Successfully got page HTML")
    except Exception as e:
        logging.error(f"Failed to get page HTML: {e}")
        raise e

    return html

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

    options = ChromeOptions()
    options.add_argument("--no-sandbox")
    options.add_argument("--window-size=1280,720")
    options.add_argument("--disable-dev-shm-usage")
    options.add_argument("--headless")

    driver = None

    try:
        logging.info("Starting chrome webdriver")
        driver = Chrome(options=options)
    except Exception as e:
        logging.error(f"Failed to start driver: {e}")
        quit()

    try:
        logging.info("Setting request headers")
        set_headers(driver)
    except Exception as e:
        logging.error(f"Failed to set headers: {e}")
        quit()

    page_source = ""

    try:
        logging.info("Scraping")
        page_source = scrape(driver)
    except Exception as e:
        logging.error(f"Failed to scrape: {e}")
    finally:
        logging.info("Quitting the driver")
        driver.quit()
        logging.info("Driver quit")

    try:
        logging.info("Processing page source")
        process_page_source(page_source)
    except Exception as e:
        logging.error(f"Failed to process page source: {e}")
    
    