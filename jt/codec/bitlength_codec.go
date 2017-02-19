package codec

import (
	"github.com/fileformats/graphics/jt/model"
	"github.com/cthackers/go/bitstream"
	"bytes"
)

// This is a very simple compression algorithm that runs an adaptive-width bit field encoding for each value.
// As each input value is encountered, the number of bits needed to represent it is calculated and compared
// to the current "field width". The current field width is then adjusted upwards or downwards by a constant
// “step_size” number of bits (i.e. 2 bits for the JT format) to accommodate the input value storage.
// This increment or decrement of the current field width is indicated for each encoded value by a prefix code
// stored with each value
type BitlengtCodec struct {
}

func (n BitlengtCodec) Decode(c *model.Context, cdp *Int32CDP) ([]int32, error) {
	codeText := getCodeText(c)
	bits := newBitBuffer(bitstream.NewReaderLE(bytes.NewReader(codeText)))

	bitFieldWidth := 0
	result := make([]int32, 0)

	for ;bits.Position < int(cdp.CodeTextLen); {
		if (bits.Int32(1) == 0) {
			var decodedSymbol int32 = -1
			if bitFieldWidth == 0 {
				decodedSymbol = 0
			} else {
				decodedSymbol = int32(bits.Int32(bitFieldWidth))
				decodedSymbol <<= uint(32 - bitFieldWidth)
				decodedSymbol >>= uint(32 - bitFieldWidth)
			}
			result = append(result, decodedSymbol)
		} else {
			adjustmentBit := bits.Int32(1)
			for {
				if adjustmentBit == 1 {
					bitFieldWidth += 2
				} else {
					bitFieldWidth -= 2
				}
				if bits.Int32(1) != adjustmentBit {
					break
				}
			}
			var decodedSymbol int32 = -1
			if bitFieldWidth == 0 {
				decodedSymbol = 0
			} else {
				decodedSymbol = int32(bits.Int32(bitFieldWidth))
				decodedSymbol <<= uint(32 - bitFieldWidth)
				decodedSymbol >>= uint(32 - bitFieldWidth)
			}
			result = append(result, decodedSymbol)
		}
	}

	return result, c.Data.GetError()
}
