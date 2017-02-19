package elements

import (
	"errors"

	"github.com/fileformats/graphics/jt/model"
)

// Base Shape Node Element represents the simplest form of a shape node that can exist within the LSG.
type BaseShapeNode struct {
	BaseNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// The Transformed BBox is an axis-aligned NCS bounding box and represents the transformed
	// geometry extents for all geometry contained in the Shape Node.
	TransformedBox model.BoundingBox
	// The Untransformed BBox is an axis-aligned LCS bounding box and represents the untransformed geometry
	// extents for all geometry contained in the Shape Node
	UntransformedBox model.BoundingBox
	// Area is the total surface area for this node and all of its descendents.
	// This value is stored in NCS coordinate space (i.e. values scaled by NCS scaling).
	Area float32
	// Vertex Count Range is the aggregate minimum and maximum vertex count for this Shape Node.
	// There is a minimum and maximum value to accommodate shape types that can themselves generate varying
	// representations. The minimum value represents the least vertex count that can be achieved by the
	// Shape Node. The maximum value represents the greatest vertex count that can be achieved by the Shape Node
	VertexCountRange model.Int32Range
	// Node Count Range is the aggregate minimum and maximum count of all node descendants of the Shape Node.
	// The minimum value represents the least node count that can be achieved by the Shape Node’s descendants.
	// The maximum value represents the greatest node count that can be achieved by Shape Node’s descendants.
	// For Shape Nodes the minimum and maximum count values should always be equal to “1”
	NodeCountRange model.Int32Range
	// Polygon Count Range is the aggregate minimum and maximum polygon count for this Shape Node.
	// There is a minimum and maximum value to accommodate shape types that can themselves generate varying
	// representations. The minimum value represents the least polygon count that can be achieved by the Shape Node.
	// The maximum value represents the greatest polygon count that can be achieved by the Shape Node.
	PolygonCountRange model.Int32Range
	// Size specifies the in memory length in bytes of the associated/referenced Shape LOD Element.
	// This Size value has no relevancy to the on-disk (JT File) size of the associated/referenced Shape LOD Element.
	// A value of zero indicates that the in memory size is unknown
	Size uint32
	// Compression Level specifies the qualitative compression level applied to the associated/referenced Shape LOD Element
	// = 0.0 − “Lossless” compression used.
	// = 0.1 − “Minimally Lossy” compression used
	// = 0.5 − “Moderate Lossy” compression used
	// = 1.0 − “Aggressive Lossy” compression used
	CompressionLevel float32
}

func (n BaseShapeNode) GUID() model.GUID {
	return model.BaseShapeNodeElement
}

func (n *BaseShapeNode) Read(c *model.Context) error {
	c.LogGroup("BaseShapeNode")
	defer c.LogGroupEnd()

	if err := (&n.BaseNode).Read(c); err != nil {
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
	}
	c.Log("VersionNumber: %d", n.VersionNumber)

	if c.Version.LessThan(model.V10) {
		c.Data.Unpack(&n.TransformedBox)
		c.Log("TransformedBox: %s", n.TransformedBox)
	}

	c.Data.Unpack(&n.UntransformedBox)
	c.Log("UntransformedBox: %s", n.UntransformedBox)

	n.Area = c.Data.Float32()
	c.Log("Area: %f", n.Area)

	c.Data.Unpack(&n.VertexCountRange)
	c.Log("VertexCountRange: %v", n.VertexCountRange)

	c.Data.Unpack(&n.NodeCountRange)
	c.Log("NodeCountRange: %v", n.NodeCountRange)

	c.Data.Unpack(&n.PolygonCountRange)
	c.Log("PolygonCountRange: %v", n.PolygonCountRange)

	if c.Version.GreaterEqThan(model.V10) {
		n.Size = c.Data.UInt32()
	} else {
		n.Size = uint32(c.Data.Int32())
	}
	c.Log("Size: %d", n.Size)

	n.CompressionLevel = c.Data.Float32()
	c.Log("CompressionLevel: %f", n.CompressionLevel)

	return c.Data.GetError()
}

func (n *BaseShapeNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *BaseShapeNode) BaseElement() *JTElement {
	return &n.JTElement
}
