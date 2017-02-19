package elements

import "github.com/fileformats/graphics/jt/model"

// A Polygon Set Shape Node Element defines a collection of independent and unconnected polygons.
// Each polygon constitutes one primitive of the set and is defined by one list of vertex coordinates
type PolygonSetShapeNode struct {
	VertexShapeNode
}

func (n PolygonSetShapeNode) GUID() model.GUID {
	return model.PolygonSetShapeNodeElement
}

func (n *PolygonSetShapeNode) Read(c *model.Context) error {
	c.LogGroup("PolygonSetShapeNode")
	defer c.LogGroupEnd()

	if err := (&n.VertexShapeNode).Read(c); err != nil {
		return err
	}

	return c.Data.GetError()
}

func (n *PolygonSetShapeNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *PolygonSetShapeNode) BaseElement() *JTElement {
	return &n.JTElement
}