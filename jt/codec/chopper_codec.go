package codec

import "github.com/fileformats/graphics/jt/model"

type ChopperCodec struct {}

func (n ChopperCodec) Decode(c *model.Context, cdp *Int32CDP) ([]int32, error) {
	bitsToChop := c.Data.UInt8()
	if bitsToChop == 0 {
		return n.Decode(c, cdp)
	}

	valueBias := c.Data.Int32()
	valueSpan := c.Data.UInt8()

	choppedMSB, err := n.Decode(c, cdp)
	if err != nil {
		return choppedMSB, err
	}
	choppedLSB, err := n.Decode(c, cdp)
	if err != nil {
		return choppedLSB, err
	}
	decoded := make([]int32, len(choppedMSB))
	for i := 0; i < len(choppedMSB); i++ {
		x := (choppedLSB[i] | (choppedMSB[i] << uint(valueSpan - bitsToChop))) + int32(valueBias)
		decoded = append(decoded, int32(x))
	}
	return decoded, c.Data.GetError()
}
