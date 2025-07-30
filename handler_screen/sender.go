package main

import (
	"os"
)

func main() {
	// A 4.2inch e-Paper display is 400x300 pixels.
	// At 1 bit per pixel, the buffer size is (400 * 300) / 8 = 15,000 bytes.
	const bufferSize = 15000
	buffer := make([]byte, bufferSize)

	// Create a simple test pattern: top half black, bottom half white.
	for i := 0; i < bufferSize; i++ {
		if i < bufferSize/2 {
			buffer[i] = 0xFF // Black
		} else {
			buffer[i] = 0xFF // White
		}
	}

	// Write the raw byte buffer directly to standard output.
	os.Stdout.Write(buffer)
}
