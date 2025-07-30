#!/usr/bin/python
# -*- coding:utf-8 -*-
import sys
import os
import time
import logging
import traceback
from PIL import Image
from waveshare_epd import epd4in2_V2

# --- Setup paths ---
picdir = os.path.join(os.path.dirname(os.path.dirname(os.path.realpath(__file__))), 'pic')
libdir = os.path.join(os.path.dirname(os.path.dirname(os.path.realpath(__file__))), 'lib')
if os.path.exists(libdir):
    sys.path.append(libdir)

logging.basicConfig(level=logging.DEBUG)

# --- Global EPD object ---
# We initialize it once here.
try:
    epd = epd4in2_V2.EPD()
    epd.Init_4Gray()
    epd.Clear()

except Exception as e:
    logging.error(f"Failed to initialize EPD: {e}")
    exit()

def clear_white():
    epd.init_fast(epd.Seconds_1_5S)
    WhiteImage = Image.new('1', (epd.height, epd.width), 255)
    epd.display_Fast(epd.getbuffer(WhiteImage))
    epd.sleep()

def update_display(bmp):
    """
    Handles the complete cycle of waking, updating, and sleeping the display.
    This is called every 1 second with a different image.
    """
    try:
        # 1. Wake up the display for updating
        epd.init_fast(epd.Seconds_1_5S)
        WhiteImage = Image.new('1', (epd.width,epd.height), 255)  # 255: clear the frame
        # bmp = Image.open(os.path.join(picdir, image_filename))
        WhiteImage.paste(bmp, (0,0))
        epd.display_4Gray(epd.getbuffer_4Gray(WhiteImage))

    except FileNotFoundError:
        logging.error(f"Image file not found")
    except Exception as e:
        logging.error(f"An error occurred during display update: {e}")


def main():
    """
    Main function to cycle through images every 1 second.
    """
    
    # List of images to cycle through

    image1 = Image.open(os.path.join(picdir, 'page_1.png'))
    image2 = Image.open(os.path.join(picdir, 'page_2.png'))
    image3= Image.open(os.path.join(picdir, 'page_3.png'))
    image4 = Image.open(os.path.join(picdir, 'page_4.png'))
    image_list = [
        image1,
        image2,
        image3,
        image4,
    ]
    
    try:
        logging.info("--- Starting EPD Image Viewer (Auto-cycling) ---")
        
        image_index = 0
        clear_white()
        while True:
            # Get current image from the list
            current_image = image_list[image_index]
            
            # Update the display with current image
            update_display(current_image)
            
            # Move to next image (cycle back to 0 when we reach the end)
            image_index = (image_index + 1) % len(image_list)
            
            # Wait 1 second before next update
            logging.info("Waiting 1 second before next image...")

    except IOError as e:
        logging.error(e)
        traceback.print_exc()

    except KeyboardInterrupt:
        logging.info("Exiting on Ctrl+C")
        # Ensure GPIO pins and resources are released cleanly
        epd4in2_V2.epdconfig.module_exit(cleanup=True)
        exit()

if __name__ == '__main__':
    main()