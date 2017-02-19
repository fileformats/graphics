package codec

import "github.com/fileformats/graphics/jt/model"

type HuffmanCodec struct {}

func (p HuffmanCodec) Decode(c *model.Context, cdp *Int32CDP) ([]int32, error) {
	return []int32{}, c.Data.GetError()
}