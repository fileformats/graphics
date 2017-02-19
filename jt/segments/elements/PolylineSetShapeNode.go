package elements

import (
	"github.com/fileformats/graphics/jt/model"
)

// A Polyline Set Shape Node Element defines a collection of independent and unconnected polylines.
// Each polyline constitutes one primitive of the set and is defined by one list of vertex coordinates
type PolylineSetShapeNode struct {
	VertexShapeNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// Area Factor specifies a multiplier factor applied to a Polyline Set computed surface area.
	// In JT data viewer applications there may be LOD selection semantics that are based on screen coverage calculations.
	// The socalled ”surface area” of a polyline is computed as if each line segment were a square.  This Area Factor
	// turns each edge into a narrow rectangle.  Valid Area Factor values lie in the range (0,1]
	AreaFactor float32
	// Vertex Bindings is a collection of normal, texture coordinate, and color binding information encoded
	// within a single U64.  All undocumented bits are reserved
	VertexBinding uint64
}

func (n PolylineSetShapeNode) GUID() model.GUID {
	return model.PolylineSetShapeNodeElement
}

func (n *PolylineSetShapeNode) Read(c *model.Context) error {
	c.LogGroup("PolylineSetShapeNode")
	defer c.LogGroupEnd()

	if err := (&n.VertexShapeNode).Read(c); err != nil {
		return err
	}

	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.Equal(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
		}
	}

	n.AreaFactor = c.Data.Float32()
	c.Log("AreaFactor: %f", n.AreaFactor)

	if c.Version.Equal(model.V9) && n.VersionNumber == 1 {
		n.VertexBinding = c.Data.UInt64()
		c.Log("VertexBinding: %d", n.VertexBinding)
	}

	return c.Data.GetError()
}

func (n *PolylineSetShapeNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *PolylineSetShapeNode) BaseElement() *JTElement {
	return &n.JTElement
}
