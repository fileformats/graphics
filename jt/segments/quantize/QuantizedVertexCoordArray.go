package quantize

import (
	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/codec"
)

// The Quantized Vertex Coord Array data collection contains the quantization data/representation for a set
// of vertex coordinates.
type QuantizedVertexCoordArray struct {
	PointQuantizer PointQuantizer
	// Vertex Count specifies the count (number of unique) vertices in the Vertex Codes arrays.
	VertexCount int32
	// X Vertex Coord Codes is a vector of quantizer “codes” for all the X-components of a set of vertex coordinates
	XVertexCoordCodes []float32
	// Y Vertex Coord Codes is a vector of quantizer “codes” for all the Y-components of a set of vertex coordinates
	YVertexCoordCodes []float32
	// Z Vertex Coord Codes is a vector of quantizer “codes” for all the Z-components of a set of vertex coordinates
	ZVertexCoordCodes []float32
}

func (n *QuantizedVertexCoordArray) Read(c *model.Context) error {
	c.Log("QuantizedVertexCoordArray")
	if err := (&n.PointQuantizer).Read(c); err != nil {
		return err
	}
	n.VertexCount = c.Data.Int32()
	c.Log("VertexCount: %d", n.VertexCount)

	xvals, err := (&codec.Int32CDP{}).ReadVecI32(c)
	if err != nil {
		return err
	}
	codec.UnpackResidual(xvals, codec.Lag1)
	n.XVertexCoordCodes = n.PointQuantizer.XQuantizerData.Dequantize(xvals)
	c.Log("XVertexCoordCodes: %v", n.XVertexCoordCodes)

	yvals, err := (&codec.Int32CDP{}).ReadVecI32(c)
	if err != nil {
		return err
	}
	codec.UnpackResidual(yvals, codec.Lag1)
	n.YVertexCoordCodes = n.PointQuantizer.YQuantizerData.Dequantize(yvals)
	c.Log("YVertexCoordCodes: %v", n.YVertexCoordCodes)

	zvals, err := (&codec.Int32CDP{}).ReadVecI32(c)
	if err != nil {
		return err
	}
	codec.UnpackResidual(zvals, codec.Lag1)
	n.ZVertexCoordCodes = n.PointQuantizer.ZQuantizerData.Dequantize(zvals)
	c.Log("ZVertexCoordCodes: %v", n.ZVertexCoordCodes)

	return c.Data.GetError()
}