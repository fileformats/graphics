package elements

import "github.com/fileformats/graphics/jt/model"

type TriStripSetShapeNode struct {
	VertexShapeNode
}

func (n TriStripSetShapeNode) GUID() model.GUID {
	return model.TriStripSetShapeNodeElement
}

func (n *TriStripSetShapeNode) Read(c *model.Context) error {
	c.LogGroup("TriStripSetShapeNode")
	defer c.LogGroupEnd()

	(&n.VertexShapeNode).Read(c)

	return c.Data.GetError()
}

func (n *TriStripSetShapeNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *TriStripSetShapeNode) BaseElement() *JTElement {
	return &n.JTElement
}