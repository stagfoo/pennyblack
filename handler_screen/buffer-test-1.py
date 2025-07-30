#!/usr/bin/python
# -*- coding:utf-8 -*-
import sys
import logging
from waveshare_epd import epd4in2_V2
from PIL import Image, ImageOps # Import ImageOps
import traceback

logging.basicConfig(level=logging.INFO)

# Display dimensions
WIDTH = 400
HEIGHT = 300

# The path to the screenshot file
IMAGE_PATH = "/tmp/screenshot.png"

try:
    # --- Open the image file ---
    logging.info("Opening image file...")
    image = Image.open(IMAGE_PATH)

    # Convert to grayscale before inverting for best results
    image = image.convert('L')
    
    # Invert the image colors (black -> white, white -> black)
    image = ImageOps.invert(image) 

    # --- Initialize and display ---
    logging.info("Initializing display...")
    epd = epd4in2_V2.EPD()
    epd.init_fast(epd.Seconds_1_5S)
    WhiteImage = Image.new('1', (WIDTH, HEIGHT), 255)
    WhiteImage.paste(image, (0,0))
    logging.info("Displaying processed image...")
    epd.display_Fast(epd.getbuffer(WhiteImage))

    logging.info("Putting display to sleep.")
    epd.sleep()

except IOError as e:
    logging.error(f"Error opening or processing image: {e}")
    traceback.print_exc()

except Exception as e:
    logging.error(e)
    traceback.print_exc()

except KeyboardInterrupt:
    logging.info("Exiting on Ctrl+C")
    epd4in2_V2.epdconfig.module_exit(cleanup=True)
    exit()