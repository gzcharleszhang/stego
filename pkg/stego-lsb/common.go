package stego_lsb

import (
	"fmt"
	"image"
)

// returns the bit at pos in the given byte
func getBitInByte(b byte, pos int) byte {
	mask := byte(0x80) // 1000 0000
	mask >>= pos
	bit := mask & b
	if bit == 0 {
		return 0
	}
	return 1
}

// set the bit at pos to the given value in the given byte
func setBitInByte(b *byte, bit byte, pos int) {
	// clear the bit at pos
	longMask := 0xFF7F // 1111 1111 0111 1111
	mask := byte(longMask >> pos)
	*b &= mask
	// set to bit at pos to the given value
	mask = bit << (7 - pos)
	*b |= mask
}

// calculates the maximum size of the encoded message in the given image
// when using the given number of bits per byte to encode
// numEncodeBits: number of bits used to encode in a byte, must be within [1,8]
func MaxEncodeSize(img image.Image, bitsPerByte int) (uint32, error) {
	if bitsPerByte <= 0 || bitsPerByte > 8 {
		return 0, fmt.Errorf("the number of bits used to encode per byte must be within [1,8]")
	}
	bytesPerPixel := 3 // 1 byte of red, green, and blue per pixel
	metadataSize := 4 // int used to store the size of the encoded message
	numPixels := img.Bounds().Dx() * img.Bounds().Dy()
	numBytes := numPixels * bytesPerPixel
	return uint32(numBytes * bitsPerByte - metadataSize), nil
}

// calculates the maximum size of the encoded message in the given image using LSB
func MaxLSBEncodeSize(img image.Image) (uint32, error) {
	return MaxEncodeSize(img, 1)
}
