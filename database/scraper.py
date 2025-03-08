import re
import time

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import WebDriverWait

# Initialize the WebDriver (Chrome in this example)
driver = webdriver.Chrome()
wait = WebDriverWait(driver, 10)

# Open the page containing the fighter thumbnails
driver.get("https://tekken.com/fighters")
time.sleep(3)  # Wait for page load

# We'll use a more flexible regex to match url("...") or url('...')
pattern = r'url\(["\']?(.*?)["\']?\)'

# Lists to store the URLs
background_urls = []
page_asset_urls = []

# Get the total count of thumbnail containers
thumbnail_count = len(wait.until(EC.presence_of_all_elements_located((By.CSS_SELECTOR, ".thumbnail__container"))))

# Loop over indexes so we re-find elements on every iteration
for i in range(thumbnail_count):
    try:
        # Re-locate the container for the current index
        containers = wait.until(EC.presence_of_all_elements_located((By.CSS_SELECTOR, ".thumbnail__container")))
        container = containers[i]

        # Locate the child element with the background image
        image_div = container.find_element(By.CSS_SELECTOR, ".thumbnail__image")
        style_attr = image_div.get_attribute("style")
        match = re.search(pattern, style_attr)
        if match:
            bg_url = match.group(1)
            background_urls.append(bg_url)
            print("Extracted background URL:", bg_url)
        else:
            print("No URL found in style attribute for:", style_attr)

        # Click the container to navigate to the subsequent page
        container.click()
        time.sleep(3)  # Wait for navigation

        try:
            # Instead of grabbing the <img> src, locate the <picture> element
            picture_elem = wait.until(EC.presence_of_element_located((By.TAG_NAME, "picture")))
            # Find all <source> tags within the picture element
            source_elems = picture_elem.find_elements(By.TAG_NAME, "source")
            
            # Select the <source> with the desired media attribute
            target_src = None
            for source in source_elems:
                media_attr = source.get_attribute("media")
                srcset_attr = source.get_attribute("srcset")
                # Change the condition below based on which media query you want
                if media_attr and "(min-width: 1079px)" in media_attr:
                    target_src = srcset_attr
                    break
            
            if target_src:
                page_asset_urls.append(target_src)
                print("Extracted asset URL from <source> tag:", target_src)
            else:
                print("No matching <source> found, falling back to <img>")
                # Optionally fallback to the <img> tag
                image_elem = picture_elem.find_element(By.TAG_NAME, "img")
                asset_url = image_elem.get_attribute("src")
                page_asset_urls.append(asset_url)
                print("Extracted asset URL from <img> tag:", asset_url)
        except Exception as inner_e:
            print("Error extracting asset from subsequent page:", inner_e)
    except Exception as outer_e:
        print("Error processing container:", outer_e)
    finally:
        # Navigate back to the thumbnails page and wait for it to load
        driver.back()
        time.sleep(3)

print("\nAll background URLs:")
for url in background_urls:
    print(url)

print("\nAll asset URLs from subsequent pages:")
for url in page_asset_urls:
    print(url)

driver.quit()
