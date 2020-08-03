package stego_lsb

import (
	"fmt"
	"image"
)

// prepend the size of the message to the message
func prependMessageSize(message *[]byte) {
	messageSize := uint32(len(*message))
	messageSizeBytes := make([]byte, 4)
	// split messageSize into 4 bytes
	mask := uint32(0xFF) // 00...0 1111 1111
	for i := 0; i < 4; i++ {
		shift := 24 - 8 * i
		shifted := messageSize >> shift
		messageSizeBytes[i] = byte(shifted & mask)
	}
	*message = append(messageSizeBytes, *message...)
}

// processes the message and passes the next bit of the message to ch
func processMessage(message string, ch chan byte) {
	defer close(ch)
	messageData := []byte(message)
	prependMessageSize(&messageData)

	for _, currByte := range messageData {
		for bit := 0; bit < 8; bit++ {
			ch <- getBitInByte(currByte, bit)
		}
	}
}

// encodes the message in the given image using the given number of bits per byte
// bitsPerByte: number of bits used to encode in a byte, must be within [1,8]
func Encode(img image.Image, message string, bitsPerByte int) (image.Image, error) {
	if bitsPerByte <= 0 || bitsPerByte > 8 {
		return img, fmt.Errorf("the number of bits used to encode per byte must be within [1,8]")
	}
	messageBytes := uint32(len(message))
	maxEncodeBytes, _ := MaxEncodeSize(img, 1)

	if messageBytes > maxEncodeBytes {
		return img, fmt.Errorf("message size (%d bytes)exceeded the image's maximum encode size (%d bytes)", messageBytes, maxEncodeBytes)
	}

	nextMsgBitCh := make(chan byte, 100)
	go processMessage(message, nextMsgBitCh)

	// make a copy of the image
	rgba := getRGBAFromImage(img)
	width, height := rgba.Bounds().Dx(), rgba.Bounds().Dy()
	for i := 0; i < bitsPerByte; i++ {
		// position of the bit in the byte
		pos := 7 - i
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				// colour at pixel (x,y)
				colour := rgba.RGBAAt(x, y)
				rgbBytes := []*byte{&colour.R, &colour.G, &colour.B}
				for _, currByte := range rgbBytes {

					// get the next bit in the message
					bit, ok := <- nextMsgBitCh
					if !ok {
						// in case previous bytes were not set
						rgba.SetRGBA(x, y, colour)
						return rgba, nil
					}
					setBitInByte(currByte, bit, pos)
				}
				rgba.SetRGBA(x, y, colour)
			}
		}
	}
	return rgba, nil
}

// encodes the message in the given image using LSB
func LSBEncode(img image.Image, message string) (image.Image, error) {
	return Encode(img, message, 1)
}
