package elements

import "github.com/fileformats/graphics/jt/model"

// An LOD Node holds a list of alternate representations. The list is represented as the
// children of a base group node, however, there are no implicit semantics associated with the ordering.
// Traversers of LSG may apply semantics to the ordering as part of alternative representation selection.
type LODNode struct {
	GroupNode
	// Version Number is the version identifier for this node
	Version uint8
	// Reserved Field is a vector data field reserved for future JT format expansion. Deprecated in ^v10
	ReservedRangeField model.VectorF32
	// Reserved Field is a data field reserved for future JT format expansion. Deprecated in ^v10
	ReservedField int32
}

func (n LODNode) GUID() model.GUID {
	return model.LodNodeElement
}

func (n *LODNode) Read(c *model.Context) error {
	c.LogGroup("LODNode")
	defer c.LogGroupEnd()

	if err := (&n.GroupNode).Read(c); err != nil {
		return err
	}

	if c.Version.Equal(model.V10) {
		n.Version = c.Data.UInt8()
	}
	if c.Version.Equal(model.V9) {
		n.Version = uint8(c.Data.UInt16())
	}

	if !c.Version.Equal(model.V10) {
		(&n.ReservedRangeField).Read(c)
		n.ReservedField = c.Data.Int32()
	}

	return c.Data.GetError()
}

func (n *LODNode) BaseElement() *JTElement {
	return &n.JTElement
}
