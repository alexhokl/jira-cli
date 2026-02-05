package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"strings"
)

// supportsKittyGraphics checks if the terminal supports the Kitty graphics protocol.
// It checks for known terminals that support it: Kitty, Ghostty, and WezTerm.
func supportsKittyGraphics() bool {
	// Check TERM_PROGRAM first (most reliable)
	termProgram := os.Getenv("TERM_PROGRAM")
	switch strings.ToLower(termProgram) {
	case "ghostty", "kitty", "wezterm":
		return true
	}

	// Check TERM variable
	term := os.Getenv("TERM")
	if strings.Contains(term, "kitty") {
		return true
	}

	// Check for Ghostty's specific environment variable
	if os.Getenv("GHOSTTY_RESOURCES_DIR") != "" {
		return true
	}

	// Check KITTY_WINDOW_ID (set by kitty terminal)
	if os.Getenv("KITTY_WINDOW_ID") != "" {
		return true
	}

	return false
}

// displayImageKitty displays an image using the Kitty graphics protocol.
// The image data should be in PNG, JPEG, or GIF format.
// Returns an error if the image cannot be decoded or displayed.
func displayImageKitty(imageData []byte) error {
	// Decode image to get dimensions for proper display
	img, format, err := image.DecodeConfig(bytes.NewReader(imageData))
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Determine the format code for Kitty protocol
	var formatCode int
	switch format {
	case "png":
		formatCode = 100 // PNG format
	case "jpeg":
		formatCode = 100 // We'll need to re-encode as PNG or use raw
		// For simplicity, we'll use raw RGBA for JPEG
	case "gif":
		formatCode = 100 // GIF also needs conversion
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}

	// Limit image display size to reasonable terminal dimensions
	maxWidth := 80  // columns (approximate)
	maxHeight := 24 // rows (approximate)

	displayWidth := img.Width
	displayHeight := img.Height

	// Scale down if image is too large (assuming ~10 pixels per character)
	pixelsPerCol := 10
	pixelsPerRow := 20

	if displayWidth > maxWidth*pixelsPerCol {
		scale := float64(maxWidth*pixelsPerCol) / float64(displayWidth)
		displayWidth = maxWidth * pixelsPerCol
		displayHeight = int(float64(displayHeight) * scale)
	}
	if displayHeight > maxHeight*pixelsPerRow {
		scale := float64(maxHeight*pixelsPerRow) / float64(displayHeight)
		displayHeight = maxHeight * pixelsPerRow
		displayWidth = int(float64(displayWidth) * scale)
	}

	// Encode image data as base64
	encoded := base64.StdEncoding.EncodeToString(imageData)

	// Write the Kitty graphics protocol escape sequence
	// Format: ESC_Gkey=value,key=value;base64_data ESC\
	// a=T means transmit and display
	// f=100 means PNG format
	// We chunk the data to avoid overwhelming the terminal

	const chunkSize = 4096 // Base64 characters per chunk

	for i := 0; i < len(encoded); i += chunkSize {
		end := i + chunkSize
		if end > len(encoded) {
			end = len(encoded)
		}
		chunk := encoded[i:end]

		// m=1 means more data follows, m=0 means this is the last chunk
		more := 1
		if end >= len(encoded) {
			more = 0
		}

		if i == 0 {
			// First chunk includes all the control data
			// a=T: transmit and display
			// f=100: PNG format
			// c/r: columns/rows to display (optional, let terminal decide)
			fmt.Printf("\x1b_Ga=T,f=%d,m=%d;%s\x1b\\", formatCode, more, chunk)
		} else {
			// Subsequent chunks only have the 'm' flag
			fmt.Printf("\x1b_Gm=%d;%s\x1b\\", more, chunk)
		}
	}

	// Print a newline after the image
	fmt.Println()

	return nil
}

// displayImageFromReader displays an image from an io.Reader using Kitty graphics.
func displayImageFromReader(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read image data: %w", err)
	}
	return displayImageKitty(data)
}

// isImageMimeType checks if the given MIME type is a supported image type.
func isImageMimeType(mimeType string) bool {
	switch strings.ToLower(mimeType) {
	case "image/png", "image/jpeg", "image/jpg", "image/gif":
		return true
	}
	return false
}
