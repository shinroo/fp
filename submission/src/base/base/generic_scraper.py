from selenium.webdriver import Chrome
from selenium.webdriver import ChromeOptions
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

class GenericScraper:
    driver = None

    def __init__(self, url):
        self.url = url

        options = ChromeOptions()
        options.add_argument("--no-sandbox")
        options.add_argument("--window-size=1280,720")
        options.add_argument("--disable-dev-shm-usage")
        options.add_argument("--headless")

        self.driver = Chrome(options=options)
        self.set_headers()

    def set_headers(self):
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
        self.driver.execute_cdp_cmd("Network.enable", {})
        self.driver.execute_cdp_cmd("Network.setExtraHTTPHeaders", {"headers": headers})

    def scrape(self, wait_class_name: str) -> str:
        self.driver.get(self.url)

        element_present = EC.presence_of_element_located((By.CLASS_NAME, wait_class_name))
        WebDriverWait(self.driver, 10).until(element_present)

        return self.driver.page_source
    
    def close(self):
        self.driver.quit()

    