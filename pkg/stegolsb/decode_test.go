package stegolsb

import (
	"github.com/gzcharleszhang/stego/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitsToByte(t *testing.T) {
	// 1001 0110
	bits := []byte{1, 0, 0, 1, 0, 1, 1, 0}
	expectedByte := byte(0x96)
	assert.Equal(t, expectedByte, bitsToByte(bits))
}

func TestDecode(t *testing.T) {
	img, _, _ := utils.GetImage(TestImagePath)
	expectedMessage := "Hello, world!"
	outImg, _ := LSBEncode(img, expectedMessage)
	message, err := Decode(outImg)
	assert.Nil(t, err)
	assert.Equal(t, expectedMessage, message)
}
