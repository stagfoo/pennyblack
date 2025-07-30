#!bin/bash
watch -n 5 "scrot -o -q - | convert - -resize 400x300! -dither FloydSteinberg -monochrome GRAY:- | python3 /path/to/your/buffer-test-1.py"
