package elements

import (
	"github.com/fileformats/graphics/jt/model"
)

// A Primitive Set Shape Node Element represents a list/set of primitive shapes (e.g. box, cylinder, sphere, etc.)
// whose LODs can be procedurally generated. Procedurally generate means that the raw geometric shape definition data
// (e.g. vertices, polygons, normals, etc) for LODs is not directly stored; instead some basic shape information is
// stored (e.g. sphere centre and radius) from which LODs can be generated
type PrimitiveSetShapeNode struct {
	BaseShapeNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// Vertex Bindings is a collection of normal, texture coordinate, and colour binding information encoded within a single U64.
	// All bits fields that are not defined as in use should be set to
	VertexBindings uint64
	// Texture Coord Gen Type specifies how a texture is applied to each face of the primitive. Single tile means one
	// copy of the texture will be stretched to fit the face, isotropic means that the texture will be duplicated on
	// the longer dimension of the face in order to maintain the texture‘s aspect ratio
	// = 0 Single Tile…Indicates that a single copy of a texture image will be applied to significant primitive features
	//     (i.e. cube face, cylinder wall, end cap) no matter how eccentrically shaped.
	// = 1 Isotropic…Implies that multiple copies of a texture image may be mapped onto eccentric surfaces such that a
	//     mapped texel stays approximately square
	TexCoordGenType int32
	// Texture Coord Binding specifies how (at what granularity) texture coordinate data is supplied (“bound”) for the
	// shape in the associated/referenced Shape LOD Element.  Valid values are as follows:
	// = 0 None.  Shape has no texture coordinate data.
	// = 1 Per Vertex.  Shape has texture coordinates for every vertex
	TextureCoordBinding int32
	// Color Binding specifies how (at what granularity) color data is supplied (“bound”) for the shape
	// in the associated/referenced Shape LOD Element
	ColorBinding int32
	// Bits Per Vertex specifies the number of quantization bits per vertex coordinate component.
	// Value must be within range [0:24] inclusive
	BitsPerVertex uint8
	// Bits Per Color specifies the number of quantization bits per color component.
	// Value must be within range [0:24] inclusive
	BitsPerColor uint8

}

func (n PrimitiveSetShapeNode) GUID() model.GUID {
	return model.PrimitiveSetShapeNodeElement
}

func (n *PrimitiveSetShapeNode) Read(c *model.Context) error {
	c.LogGroup("PrimitiveSetShapeNode")
	defer c.LogGroupEnd()

	if err := (&n.BaseShapeNode).Read(c); err != nil {
		return err
	}

	switch {
	case c.Version.Equal(model.V8):
		n.TextureCoordBinding = c.Data.Int32()
		c.Log("TextureCoordBinding: %d", n.TextureCoordBinding)
		n.ColorBinding = c.Data.Int32()
		c.Log("ColorBinding: %d", n.ColorBinding)
		n.BitsPerVertex = c.Data.UInt8()
		c.Log("BitsPerVertex: %d", n.BitsPerVertex)
		n.BitsPerColor = c.Data.UInt8()
		c.Log("BitsPerColor: %d", n.BitsPerColor)
		if c.Data.Int16() == 1 {
			n.TexCoordGenType  = c.Data.Int32()
			c.Log("TexCoordGenType: %d", n.TexCoordGenType)
		}

	case c.Version.Equal(model.V9):
		n.VersionNumber = uint8(c.Data.Int16())
		n.TextureCoordBinding = c.Data.Int32()
		c.Log("TextureCoordBinding: %d", n.TextureCoordBinding)
		n.ColorBinding = c.Data.Int32()
		c.Log("ColorBinding: %d", n.ColorBinding)
		n.TexCoordGenType  = c.Data.Int32()
		c.Log("TexCoordGenType: %d", n.TexCoordGenType)
		c.Data.Int16() // version 2
		n.BitsPerVertex = c.Data.UInt8()
		c.Log("BitsPerVertex: %d", n.BitsPerVertex)
		n.BitsPerColor = c.Data.UInt8()
		c.Log("BitsPerColor: %d", n.BitsPerColor)

	case c.Version.Equal(model.V10):
		n.VersionNumber = c.Data.UInt8()
		n.VertexBindings = c.Data.UInt64()
		c.Log("VertexBindings: %d", n.VertexBindings)
		n.TexCoordGenType = c.Data.Int32()
		c.Log("TexCoordGenType: %d", n.TexCoordGenType)
		c.Data.UInt8() // version 2
		n.BitsPerVertex = c.Data.UInt8()
		c.Log("BitsPerVertex: %d", n.BitsPerVertex)
		n.BitsPerColor = c.Data.UInt8()
		c.Log("BitsPerColor: %d", n.BitsPerColor)
	}

	return c.Data.GetError()
}

func (n *PrimitiveSetShapeNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *PrimitiveSetShapeNode) BaseElement() *JTElement {
	return &n.JTElement
}
