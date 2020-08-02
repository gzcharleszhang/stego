package stego_lsb

import (
	"github.com/magiconair/properties/assert"
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
