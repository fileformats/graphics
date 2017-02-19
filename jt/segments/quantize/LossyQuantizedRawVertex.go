package quantize

import (
	"github.com/fileformats/graphics/jt/codec"
	"github.com/fileformats/graphics/jt/model"
)

// The Lossy Quantized Raw Vertex Data collection contains all the per-vertex information
type LossyQuantizedRawVertex struct {
	// The Quantized Vertex Coord Array data collection contains the quantization data/representation for a set of
	// vertex coordinates
	QuantVertexCoord QuantizedVertexCoordArray
	// The Quantized Vertex Normal Array data collection contains the quantization data/representation for a set of
	// vertex normals. Quantized Vertex Normal Array data collection is only present if previously read Normal Binding
	// value is not equal to zero
	QuantVertexNorm QuantizedVertexNormalArray
	// Vertex Data Indices is a vector of indices (one per vertex) into the uncompressed/dequantized unique vertex data
	// arrays (Vertex Coords, Vertex Normals, Vertex Texture Coords, Vertex Colors) identifying each Vertex’s data
	// (i.e. for each Vertex there is an index identifying the location within the unique arrays of the particular
	// Vertex’s data). The Compressed Vertex Index List uses the Int32 version of the CODEC to compress and encode data
	VertexDataIndices []int32
}

func (n *LossyQuantizedRawVertex) Read(c *model.Context, normalBinding, textureCoordBinding, colorBinding uint8) error {
	c.Log("LossyQuantizedRawVertex")
	if err := (&n.QuantVertexCoord).Read(c); err != nil {
		return err
	}

	if normalBinding != 0 {
		if err := (&n.QuantVertexNorm).Read(c); err != nil {
			return err
		}
	}

	if textureCoordBinding != 0 {
		// @todo: read texture coordinates
	}

	if colorBinding != 0 {
		var quantVertex QuantizedVertexNormalArray
		if err := (&quantVertex).Read(c); err != nil {
			return err
		}
	}

	val, err := (&codec.Int32CDP{}).ReadVecI32(c)
	if err != nil {
		return err
	}
	codec.UnpackResidual(val, codec.StripIndex)
	n.VertexDataIndices = val
	c.Log("VertexDataIndices: %v", n.VertexDataIndices)

	return c.Data.GetError()
}