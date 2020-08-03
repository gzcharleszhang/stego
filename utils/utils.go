package utils

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
)

func PathWithExtension(p string) (path string, extension string) {
	parts := strings.Split(p, ".")
	path = strings.Join(parts[:len(parts)-1], ".")
	extension = parts[len(parts)-1]
	return
}

func validateFormat(path string) error {
	_, ext := PathWithExtension(path)
	if ext != "png" {
		return errors.New("unsupported image format, image must be in png")
	}
	return nil
}

func GetImage(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if reader != nil {
		defer reader.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("error opening png file: %v\n", err)
	}
	err = validateFormat(path)
	if err != nil {
		return nil, err
	}
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("Error decoding image: %v\n", err)
	}
	return img, nil
}
