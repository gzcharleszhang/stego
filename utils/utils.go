package utils

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
)

// PathWithExtension splits the given file path into the path name and file extension
func PathWithExtension(p string) (path string, extension string) {
	parts := strings.Split(p, ".")
	path = strings.Join(parts[:len(parts)-1], ".")
	extension = parts[len(parts)-1]
	return
}

func validateAndRegisterFormat(path string) (string, error) {
	_, ext := PathWithExtension(path)
	if ext == "png" {
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
		return "png", nil
	}
	return "", errors.New("unsupported image format")
}

// GetImage returns an image object from the given path
func GetImage(path string) (image.Image, string, error) {
	reader, err := os.Open(path)
	if reader != nil {
		defer reader.Close()
	}
	if err != nil {
		return nil, "", fmt.Errorf("error opening image: %v", err)
	}
	format, err := validateAndRegisterFormat(path)
	if err != nil {
		return nil, "", err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, "", fmt.Errorf("error decoding image: %v", err)
	}
	return img, format, nil
}
