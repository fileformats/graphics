package codec

import (
	"encoding/binary"

	"github.com/cthackers/go/bitstream"
)

type BitBuffer struct {
	Position   int
	BitBuffer  int
	NBits      int

	Buffer    *bitstream.BitReader
	ByteOrder binary.ByteOrder
}

func newBitBuffer(data *bitstream.BitReader) *BitBuffer {
	return &BitBuffer{
		Position:   0,
		Buffer:     data,
		ByteOrder:  data.GetByteOrder(),
	}
}

func (b *BitBuffer) Byte(bitCount int) byte {
	return byte(b.Int32(bitCount))
}

func (b *BitBuffer) Int32(bitCount int) int32 {
	if bitCount <= 0 {
		return 0
	}
	bPos := 0
	var result int32 = 0
	var length = bPos + bitCount

	for length > 0 {
		if b.NBits == 0 {
			b.BitBuffer = int(b.Buffer.Int8())
			b.NBits = 8
			b.BitBuffer &= 0xFF
		}
		if bPos == 0 {
			result <<= 1
			result |= int32(b.BitBuffer) >> 7
		} else {
			bPos--
		}
		b.BitBuffer <<= 1
		b.BitBuffer &= 0xFF
		b.NBits--
		length--
		b.Position++
	}
	return result
}