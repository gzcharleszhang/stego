package stegolsb

import (
	"github.com/stretchr/testify/assert"
	"image"
	"image/png"
	"os"
	"testing"
)

func TestGetBitInByte(t *testing.T) {
	b := byte(0x80) // 1000 0000
	assert.Equal(t, getBitInByte(b, 0), byte(1))

	b = byte(0x10) // 0001 0000
	assert.Equal(t, getBitInByte(b, 0), byte(0))
	assert.Equal(t, getBitInByte(b, 3), byte(1))

	b = byte(0x01) // 0000 0001
	assert.Equal(t, getBitInByte(b, 7), byte(1))
}

func TestSetBitInByte(t *testing.T) {
	b := byte(0) // 0000 0000
	setBitInByte(&b, 1, 0)
	assert.Equal(t, b, byte(0x80))
	setBitInByte(&b, 0, 0)
	assert.Equal(t, b, byte(0))
	setBitInByte(&b, 1, 7)
	assert.Equal(t, b, byte(1))
	setBitInByte(&b, 1, 4)
	assert.Equal(t, b, byte(9))
}

func TestMaxEncodeSize(t *testing.T) {
	reader, err := os.Open("../../test/testdata/butterfly.png")
	if reader != nil {
		defer reader.Close()
	}
	if err != nil {
		t.Errorf("Error opening png file: %v\n", err)
	}
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	img, _, err := image.Decode(reader)
	if err != nil {
		t.Errorf("Error decoding image: %v\n", err)
	}
	// 3 bytes per pixel, 2 bits per byte, 4 bytes for metadata
	expectedSize := uint32(img.Bounds().Dx()*img.Bounds().Dy()*3*2 - 4)
	maxSize, err := MaxEncodeSize(img, 2)
	assert.Equal(t, maxSize, expectedSize)
	assert.Nil(t, err)

	// test error handling
	_, err = MaxEncodeSize(img, 9)
	assert.NotNil(t, err)
}

func TestMaxLSBEncodeSize(t *testing.T) {
	reader, err := os.Open("../../test/testdata/butterfly.png")
	if reader != nil {
		defer reader.Close()
	}
	if err != nil {
		t.Errorf("Error opening png file: %v\n", err)
	}
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	img, _, err := image.Decode(reader)
	if err != nil {
		t.Errorf("Error decoding image: %v\n", err)
	}
	// 3 bytes per pixel, 1 bit per byte, 4 bytes for metadata
	expectedSize := uint32(img.Bounds().Dx()*img.Bounds().Dy()*3 - 4)
	maxSize, _ := MaxEncodeSize(img, 1)
	assert.Equal(t, maxSize, expectedSize)
}
