#!/bin/bash
set -e

mkdir -p /screenshots
counter=1

while true; do
    timestamp=\$(date +"%Y%m%d_%H%M%S")
    filename="/screenshots/screenshot_\${timestamp}_\${counter}.png"
    scrot "\$filename"
    echo "Screenshot saved: \$filename"
    counter=\$((counter + 1))
    sleep 1
done