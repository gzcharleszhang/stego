package utils

import (
	"github.com/stretchr/testify/assert"
	"image"
	"os"
	"testing"
)

func TestPathWithExtension(t *testing.T) {
	path, ext := PathWithExtension("/test/example.png")
	assert.Equal(t, "/test/example", path)
	assert.Equal(t, "png", ext)

	path, ext = PathWithExtension("./test/example.jpg")
	assert.Equal(t, "./test/example", path)
	assert.Equal(t, "jpg", ext)

	path, ext = PathWithExtension("/.example/test.example.jpeg")
	assert.Equal(t, "/.example/test.example", path)
	assert.Equal(t, "jpeg", ext)
}

func TestGetImage(t *testing.T) {
	path := "../test/testdata/butterfly.png"
	expectedFormat := "png"
	reader, err := os.Open(path)
	if reader != nil {
		defer reader.Close()
	}
	assert.Nil(t, err)
	assert.Nil(t, err)
	expectedImg, _, err := image.Decode(reader)
	assert.Nil(t, err)

	img, format, err := GetImage(path)
	assert.Nil(t, err)
	assert.Equal(t, expectedFormat, format)
	assert.Equal(t, expectedImg, img)

	_, _, err = GetImage("fake.gif")
	assert.NotNil(t, err)
}
