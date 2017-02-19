package codec

import "github.com/fileformats/graphics/jt/model"

type MoveToFrontCodec struct {}

func (n MoveToFrontCodec) Decode(c *model.Context, cdp *Int32CDP) ([]int32, error) {
	return []int32{}, c.Data.GetError()
}