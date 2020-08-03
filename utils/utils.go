package utils

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
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

func validateAndRegisterFormat(path string) error {
	_, ext := PathWithExtension(path)
	if ext == "jpeg" || ext == "jpg" {
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
		return nil
	}
	if ext == "png" {
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
		return nil
	}
	return errors.New("unsupported image format")
}

func GetImage(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if reader != nil {
		defer reader.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("error opening image: %v\n", err)
	}
	err = validateAndRegisterFormat(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("Error decoding image: %v\n", err)
	}
	return img, nil
}
