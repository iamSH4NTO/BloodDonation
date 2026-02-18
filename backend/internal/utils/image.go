package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // Support PNG decoding
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

const (
	MaxUploadSize = 5 * 1024 * 1024 // 5MB
	ImageQuality  = 75              // 75% quality for compression
	MaxImageWidth = 800             // Resize if wider than 800px
)

func ValidateImage(file io.Reader) (string, error) {
	// Read first 512 bytes to detect content type
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	contentType := http.DetectContentType(buffer[:n])

	if contentType != "image/jpeg" && contentType != "image/png" {
		return "", fmt.Errorf("invalid file type: %s. Only JPEG and PNG are allowed", contentType)
	}

	return contentType, nil
}

func OptimizeImage(input io.Reader, outputPath string) error {
	img, _, err := image.Decode(input)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// Resize if too wide
	bounds := img.Bounds()
	if bounds.Dx() > MaxImageWidth {
		img = resize.Resize(MaxImageWidth, 0, img, resize.Lanczos3)
	}

	// Create directory if not exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Encode to JPEG with quality reduction
	err = jpeg.Encode(out, img, &jpeg.Options{Quality: ImageQuality})
	if err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}

	return nil
}

func SaveImage(file io.Reader, filename string, subDir string) (string, error) {
	uploadDir := filepath.Join("uploads", subDir)
	outputPath := filepath.Join(uploadDir, filename)

	err := OptimizeImage(file, outputPath)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
