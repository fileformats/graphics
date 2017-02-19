package codec

import "github.com/fileformats/graphics/jt/model"

type NullCodec struct {}

func (n NullCodec) Decode(c *model.Context, cdp *Int32CDP) ([]int32, error) {
	length := c.Data.Int32()
	decoded := make([]int32, length / 4)
	c.Data.Unpack(&decoded)
	return decoded, c.Data.GetError()
}
