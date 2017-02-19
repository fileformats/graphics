package quantize

import "github.com/fileformats/graphics/jt/model"

// Quantization Parameters specifies for each shape data type grouping (i.e. Vertex, Normal, Texture Coordinates,
// Color) the number of quantization bits used for given qualitative compression level.
// Although these Quantization Parameters values are saved in the associated/referenced Shape LOD Element,
// they are also saved here so that a JT File loader/reader does not have to load the Shape LOD Element
// in order to determine the Shape quantization level
type QuantizationParam struct {
	// Bits Per Vertex specifies the number of quantization bits per vertex coordinate component.
	// Value shall be within range [0:24] inclusive
	BitsPerVertex uint8
	// Normal Bits Factor is a parameter used to calculate the number of quantization bits for normal vectors.
	// Value shall be within range [0:13] inclusive. The actual number of quantization bits per normal is
	// computed using this factor and the following formula:   BitsPerNormal = 6 + 2 * Normal Bits Factor
	NormalBitsFactor uint8
	BitsPerNormal uint8
	// Bits Per Texture Coord specifies the number of quantization bits per texture coordinate component.
	// Value shall be within range [0:24] inclusive
	BitsPerTextureCoord uint8
	// Bits Per Colour specifies the number of quantization bits per colour component.
	// Value shall be within range [0:24] inclusive
	BitsPerColor uint8
}

func (n *QuantizationParam) Read(c *model.Context) error {
	c.LogGroup("QuantizationParam")
	defer c.LogGroupEnd()

	n.BitsPerVertex = c.Data.UInt8()
	c.Log("BitsPerVertex: %d", n.BitsPerVertex)

	n.NormalBitsFactor = c.Data.UInt8()
	c.Log("NormalBitsFactor: %d", n.NormalBitsFactor)

	n.BitsPerNormal = 6 + 2 * n.NormalBitsFactor
	c.Log("BitsPerNormal: %d", n.BitsPerNormal)

	n.BitsPerTextureCoord = c.Data.UInt8()
	c.Log("BitsPerTextureCoord: %d", n.BitsPerTextureCoord)

	n.BitsPerColor = c.Data.UInt8()
	c.Log("BitsPerColor: %d", n.BitsPerColor)

	return c.Data.GetError()
}