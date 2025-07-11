#!/bin/bash
set -e

# Start Xvfb (Virtual Frame Buffer)
Xvfb :99 -screen 0 1920x1080x24 &

# Wait for X server to start
sleep 2

# Start screenshot capture in background
/bin/screenshot.sh &

# Start your Fyne application
# Replace this with your actual Fyne app command
# /app/fyne-app

# For demo purposes, start a simple X application
# Remove this and uncomment your app above
xterm -e "echo 'Fyne app would run here. Replace with your app.' && sleep infinity" &

# Keep container running
wait