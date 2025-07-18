#!/usr/bin/python
# -*- coding:utf-8 -*-
import sys
import os
import time
import logging
import traceback
from PIL import Image
from waveshare_epd import epd4in2_V2

# Import the 'keyboard' library to listen for system-wide key presses.
# You may need to install it: pip install keyboard
# On Linux, this script must be run as root for the keyboard library to work.
import keyboard

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
except Exception as e:
    logging.error(f"Failed to initialize EPD: {e}")
    exit()

def update_display(image_filename):
    """
    Handles the complete cycle of waking, updating, and sleeping the display.
    This is called whenever a registered key is pressed.
    """
    image_path = os.path.join(picdir, image_filename)
    logging.info(f"Attempting to display image: {image_path}")

    try:
        # 1. Wake up the display for updating
        logging.info("Initializing display (waking up)...")
        epd.init()

        # 2. Load the specified BMP image
        image = Image.open(image_path)

        # 3. Display the image on the screen
        logging.info("Sending image to display...")
        epd.Init_4Gray()
        epd.display_4Gray(epd.getbuffer_4Gray(image))

        # 4. Put the display back to sleep to conserve power
        logging.info("Putting display to sleep.")
        epd.sleep()

    except FileNotFoundError:
        logging.error(f"Image file not found at: {image_path}")
    except Exception as e:
        logging.error(f"An error occurred during display update: {e}")


def main():
    """
    Main function to set up the initial display and keyboard listeners.
    """
    try:
        logging.info("--- Starting EPD Image Viewer ---")
        logging.info("Displaying initial image...")
        # Display a default image ('penny.bmp') on startup
        update_display('penny.bmp')

        # --- Setup Keyboard Listeners ---
        # The old 'while True' loop with epd.get_key() has been replaced.
        # We now use the 'keyboard' library to trigger updates in the background.
        # A lambda function is used to pass the correct image filename to our handler.
        logging.info("Setting up keyboard listeners for 'q', 'w', 'e', 'r', 't'...")
        keyboard.on_press_key("q", lambda _: update_display('q_image.bmp'))
        keyboard.on_press_key("w", lambda _: update_display('w_image.bmp'))
        keyboard.on_press_key("e", lambda _: update_display('e_image.bmp'))
        keyboard.on_press_key("r", lambda _: update_display('r_image.bmp'))
        keyboard.on_press_key("t", lambda _: update_display('mushishi-page.jpg'))

        logging.info("Ready for keyboard input. Press Ctrl+C to exit.")
        # Keep the script running to listen for key presses
        while True:
            time.sleep(1)

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