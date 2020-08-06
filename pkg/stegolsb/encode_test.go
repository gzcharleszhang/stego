package stegolsb

import (
	"encoding/binary"
	"github.com/gzcharleszhang/stego/utils"
	"github.com/stretchr/testify/assert"
	"image"
	"testing"
)

func TestPrependMessageSize(t *testing.T) {
	message := []byte("Hello, world!")
	expectedSize := uint32(len(message))
	prependMessageSize(&message)
	sizeBytes := message[:4]
	size := binary.BigEndian.Uint32(sizeBytes)
	assert.Equal(t, expectedSize, size)
}

func TestProcessMessage(t *testing.T) {
	message := "Hello, world!"
	messageData := []byte(message)
	ch := make(chan byte, 100)
	go processMessage(message, ch)

	prependMessageSize(&messageData)

	for _, currByte := range messageData {
		for bit := 0; bit < 8; bit++ {
			b, ok := <-ch
			assert.True(t, ok)
			assert.Equal(t, getBitInByte(currByte, bit), b)
		}
	}
}

func TestEncode(t *testing.T) {
	img, _, err := utils.GetImage(TestImagePath)
	assert.NotNil(t, img)
	assert.Nil(t, err)
	message := "Hello, world!"
	expectedSize := uint32(len([]byte(message)))
	bitsPerByte := 1

	encodedImg, err := Encode(img, message, bitsPerByte)
	assert.Nil(t, err)
	assert.NotNil(t, encodedImg)
	rgba := getRGBAFromImage(encodedImg)

	// check size
	sizeBytes := getMessageSizeFromImage(rgba)
	assert.Equal(t, expectedSize, sizeBytes)

	// check message
	messageBytes := getMessageFromImage(rgba, 4, expectedSize)
	assert.Equal(t, []byte(message), messageBytes)
}

func TestLSBEncode(t *testing.T) {
	testImg, _, err := utils.GetImage(TestImagePath)
	assert.NotNil(t, testImg)
	assert.Nil(t, err)
	msg := "Hello, world!"

	mockEncode := func(img image.Image, message string, bitsPerByte int) (image.Image, error) {
		assert.Equal(t, testImg, img)
		assert.Equal(t, msg, message)
		assert.Equal(t, 1, bitsPerByte)
		return img, nil
	}

	encodeFn = mockEncode
	img, err := LSBEncode(testImg, msg)
	assert.Nil(t, err)
	assert.NotNil(t, img)
	// un-mock
	encodeFn = Encode
}
