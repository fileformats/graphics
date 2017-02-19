package quantize

import "github.com/fileformats/graphics/jt/model"

// A Point Quantizer Data collection is made up of three Uniform Quantizer Data collections; there is a separate
// Uniform Quantizer Data collection for the X, Y, and Z values of point coordinates
type PointQuantizer struct {
	XQuantizerData UniformQuantizer
	YQuantizerData UniformQuantizer
	ZQuantizerData UniformQuantizer
}

func (n *PointQuantizer) Read(c *model.Context) error {
	c.Log("PointQuantizer")
	if err := (&n.XQuantizerData).Read(c); err != nil {
		return err
	}
	if err := (&n.YQuantizerData).Read(c); err != nil {
		return err
	}
	if err := (&n.ZQuantizerData).Read(c); err != nil {
		return err
	}
	return c.Data.GetError()
}