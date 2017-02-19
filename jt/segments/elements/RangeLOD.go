package elements

import "github.com/fileformats/graphics/jt/model"

// Range LOD Nodes hold a list of alternate representations and the ranges over which those representations
// are appropriate. Range Limits indicate the distance between a specified center point and the eye point,
// within which the corresponding alternate representation is appropriate
type RangeLODNode struct {
	LODNode
	// Range Limits indicate the WCS distance between a specified centre point and the eye point,
	// within which the corresponding alternate representation is appropriate
	RangeLimits   model.VectorF32
	// Centre specifies the X,Y,Z coordinates for the MCS centre point upon which alternative
	// representation selection eye distance computations are based
	Center        model.Vector3D
}

func (n RangeLODNode) GUID() model.GUID {
	return model.RangeLodNodeElement
}

func (n *RangeLODNode) Read(c *model.Context) error {
	c.LogGroup("RangeLODNode")
	defer c.LogGroupEnd()

	if err := (&n.LODNode).Read(c); err != nil {
		return err
	}
	(&n.RangeLimits).Read(c)
	c.Log("RangeLimits: %v", n.RangeLimits)

	c.Data.Unpack(&n.Center)
	c.Log("Center: %s", n.Center)

	return c.Data.GetError()
}

func (n *RangeLODNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *RangeLODNode) BaseElement() *JTElement {
	return &n.JTElement
}