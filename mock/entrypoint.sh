#!/bin/bash

# --- Configuration ---
# Set the display to use
DISPLAY_NUM=99
export DISPLAY=:${DISPLAY_NUM}

# Your Fyne application command
APP_COMMAND="./main" # Or "./your-compiled-app"

# --- Script Logic ---
ls -la
echo "Starting Xvfb on display ${DISPLAY}"
Xvfb ${DISPLAY} -screen 0 400x600x24 &
XVFBPID=$!

echo "Starting Fyne app..."
${APP_COMMAND} &
APPPID=$!



echo "Starting screenshot loop..."
i=0
while kill -0 $APPPID 2>/dev/null; do
  # Capture the screen, convert, and save as a PNG
  xwd -root -display ${DISPLAY} | magick convert xwd:- /screenshots/screenshot_${i}.png
  echo "Screenshot saved: /screenshots/screenshot_${i}.png"
  i=$((i+1))
  sleep 1
done

echo "App has closed. Cleaning up..."
kill $XVFBPID