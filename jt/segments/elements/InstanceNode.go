package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"errors"
)

// An Instance Node contains a single reference to another node. Their purpose is to allow sharing of
// nodes and assignment of instance-specific attributes for the instanced node.
// Instance Nodes may not contain references to themselves or their ancestors.
type InstanceNode struct {
	BaseNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// Child Node Object ID is the identifier for the instanced Node Object
	ChildNodeObjectId int32
}

func (n InstanceNode) GUID() model.GUID {
	return model.InstanceNodeElement
}

func (n *InstanceNode) Read(c *model.Context) error {
	c.LogGroup("InstanceNode")
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
		if n.VersionNumber != 1  {
			return errors.New("Invalid version number")
		}
		c.Log("VersionNumber: %d", n.VersionNumber)
	}

	n.ChildNodeObjectId = c.Data.Int32()
	c.Log("ChildNodeObjectId: %d", n.ChildNodeObjectId)

	return c.Data.GetError()
}

func (n *InstanceNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *InstanceNode) BaseElement() *JTElement {
	return &n.JTElement
}