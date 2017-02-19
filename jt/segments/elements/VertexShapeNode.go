package elements

import (
	"errors"

	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments/quantize"
)

// Vertex Shape Node Element represents shapes defined by collections of vertices.
type VertexShapeNode struct {
	BaseShapeNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// Vertex Bindings is a collection of normal, texture coordinate, and colour binding information encoded within
	// a single U64. All bits fields that are not defined as in use should be set to 0
	VertexBinding  uint64
	VertexBinding2 uint64
	// Normal Binding specifies how (at what granularity) normal vector data is supplied (“bound”) for the
	// shape in the associated/referenced Shape LOD Element
	// = 0 − None.  Shape has no normal data.
	// = 1 − Per Vertex.  Shape has a normal vector for every vertex.
	// = 2 − Per Facet.  Shape has a normal vector for every face/polygon.
	// = 3 − Per Primitive. Shape has a normal vector for each shape primitive
	NormalBinding int32
	// Texture Coord Binding specifies how (at what granularity) texture coordinate data is supplied (“bound”)
	// for the shape in the associated/referenced Shape LOD Element
	TextureCoordBinding int32
	// Color Binding specifies how (at what granularity) color data is supplied (“bound”) for the shape in the
	// associated/referenced Shape LOD Element
	ColorBinding int32
	// Quantization Parameters specifies for each shape data type grouping (i.e. Vertex, Normal, Texture
	// Coordinates, Color) the number of quantization bits used for given qualitative compression level
	QuantizationParameters quantize.QuantizationParam
}

func (n VertexShapeNode) GUID() model.GUID {
	return model.VertexShapeNodeElement
}

func (n *VertexShapeNode) Read(c *model.Context) error {
	c.LogGroup("VertexShapeNode")
	defer c.LogGroupEnd()

	if err := (&n.BaseShapeNode).Read(c); err != nil {
		return err
	}

	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.GreaterEqThan(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
		}
		if n.VersionNumber != 1 {
			return errors.New("Invalid version number")
		}
		c.Log("VersionNumber: %d", n.VersionNumber)

		n.VertexBinding = c.Data.UInt64()
		c.Log("VertexBinding: %d", n.VertexBinding)

		if c.Version.Equal(model.V9) {
			if err := (&n.QuantizationParameters).Read(c); err != nil {
				return err
			}
			if n.VersionNumber != 1 {
				n.VertexBinding2 = c.Data.UInt64()
				c.Log("VertexBinding2: %d", n.VertexBinding2)
			}
		}
	} else {
		n.NormalBinding = c.Data.Int32()
		c.Log("NormalBinding: %d", n.NormalBinding)

		n.TextureCoordBinding = c.Data.Int32()
		c.Log("TextureCoordBinding: %d", n.TextureCoordBinding)

		n.ColorBinding = c.Data.Int32()
		c.Log("ColorBinding: %d", n.ColorBinding)

		if err := (&n.QuantizationParameters).Read(c); err != nil {
			return err
		}
	}
	return c.Data.GetError()
}

func (n *VertexShapeNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}
func (n *VertexShapeNode) BaseElement() *JTElement {
	return &n.JTElement
}
