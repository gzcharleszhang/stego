package stegolsb

import (
	"encoding/binary"
	"image"
)

func bitsToByte(bits []byte) byte {
	b := byte(0)
	for i, bit := range bits {
		shift := 7 - i
		bit <<= shift
		b |= bit
	}
	return b
}

// returns the message in the image using the LSB scheme
// offset: the number of bytes to skip
// length: the number of bytes to collect
func getMessageFromImage(img *image.RGBA, offset uint32, length uint32) []byte {
	counter, offsetBits, lengthBits := uint32(0), offset*8, length*8
	var message []byte
	var bits []byte
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	for pos := 7; pos >= 0; pos-- {
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				// colour at pixel (x,y)
				colour := img.RGBAAt(x, y)
				rgbBytes := []*byte{&colour.R, &colour.G, &colour.B}
				for _, currByte := range rgbBytes {
					if counter < offsetBits {
						// skip this bit if we haven't offset enough
						counter++
						continue
					}
					if counter >= lengthBits+offsetBits {
						// terminate if we've collected enough bytes
						return message
					}
					bits = append(bits, getBitInByte(*currByte, pos))
					if len(bits) == 8 {
						// if we collected 8 bits, then we append to the message
						message = append(message, bitsToByte(bits))
						bits = nil
					}
					counter++
				}
			}
		}
	}
	return message
}

// when the image was encoded, the size of the image was prepended to the image
// this function will return the size of the message
func getMessageSizeFromImage(img *image.RGBA) uint32 {
	// message size should be stored as an uint32, which is 4 bytes
	messageSize := getMessageFromImage(img, 0, 4)
	return binary.BigEndian.Uint32(messageSize)
}

// Decode tries to extract the hidden message from an image
func Decode(img image.Image) (string, error) {
	rgba := getRGBAFromImage(img)
	messageBytes := getMessageSizeFromImage(rgba)
	// offset the 32-bit message size
	message := getMessageFromImage(rgba, 4, messageBytes)
	return string(message), nil
}
