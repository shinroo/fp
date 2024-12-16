import os
import re
import json
from bs4 import BeautifulSoup

# Define the path to the HTML file
html_file_path = os.path.join(os.path.dirname(__file__), 'dbn_spca.html')

# Read the HTML file
with open(html_file_path, 'r', encoding='utf-8') as file:
    html_content = file.read()

# Parse the HTML content using BeautifulSoup
soup = BeautifulSoup(html_content, 'html.parser')

# Extract the anchor tags that contain dog information from the HTML content
dog_info_tags = soup.find_all('a', {'data-gal':'prettyPhoto[adopt]'})

# Define function to determine gender based on description
def sex_from_description(description: str):
    if "boy" in description or "Boy" in description:
        return "male"
    elif "girl" in description or "Girl" in description:
        return "female"
    return None

# Define function to determine age based on description
def age_from_description(description: str):
    # Match ages like X years
    age_pattern_years = r'\b(\d+)\s+years?\b'
    match_years = re.search(age_pattern_years, description)
    if match_years:
        return float(match_years.group(1))
    
    # Match ages like X month
    age_pattern_months = r'\b(\d+)\s+month?\b'
    match_month = re.search(age_pattern_months, description)
    if match_month:
        return float(int(match_month.group(1))/12)
    
    return None


# Define a function to extract the kennel number from the description
def kennel_from_description(description: str):
    kennel_pattern = r'\bKennel\s+([A-Z]\d+)\b'
    match = re.search(kennel_pattern, description)
    if match:
        return match.group(1)
    return None

# Define a function to extract the reference number from the description
def reference_from_description(description: str):
    reference_pattern = r'\bRef:\s*(\d+)\b'
    match = re.search(reference_pattern, description)
    if match:
        return match.group(1)
    return None

# Define a function to extract pet type from the HTML
def pet_type_from_html(tag):
    is_dog = tag.find_parent(id='gallery1')
    is_cat = tag.find_parent(id='gallery2')

    if is_dog:
        return "Dog"
    elif is_cat:
        return "Cat"
    else:
        return None

# Create JSON structured data from each tag
json_data = []
for tag in dog_info_tags:
    json_data.append({
        # This information is extracted from the anchor tag and its attributes
        'name': tag.find('h3').text,
        'image': f"https://spcadbn.org.za/{tag.find('img')['src']}",
        'description': tag['title'],
        'sex': sex_from_description(tag['title']),
        'age': age_from_description(tag['title']),
        'kennel': kennel_from_description(tag['title']),
        'reference': reference_from_description(tag['title']),
        'type': pet_type_from_html(tag),

        # The following information is known because we are only working with one SPCA for this example
        'spca_branch': 'Durban and Coast SPCA',
        'province': 'KwaZulu-Natal',
        'latitude': -29.80363108540157,
        'longitude': 30.99372385667752
    })

print(json.dumps(json_data, indent=2))