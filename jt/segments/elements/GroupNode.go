package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"errors"
)

// Group Nodes contain an ordered list of references to other nodes, called the group‘s children.
// Group nodes may contain zero or more children; the children may be of any node type.
// Group nodes may not contain references to themselves or their ancestors.
type GroupNode struct {
	BaseNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// Child Count indicates the number of child nodes for this Group Node Object
	ChildNodes int32
	// Child Node Object ID is the identifier for the referenced Node Object
	ChildNodeIds []int32
}

func (n *GroupNode) GUID() model.GUID {
	return model.GroupNodeElement
}

func (n *GroupNode) Read(c *model.Context) error {
	c.LogGroup("GroupNode")
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

	n.ChildNodes = c.Data.Int32()
	c.Log("Child Nodes Count: %d", n.ChildNodes)

	for i := 0; i < int(n.ChildNodes); i++ {
		n.ChildNodeIds = append(n.ChildNodeIds, c.Data.Int32())
	}
	c.Log("Child Nodes: %v", n.ChildNodeIds)

	return c.Data.GetError()
}

func (n *GroupNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *GroupNode) BaseElement() *JTElement {
	return &n.JTElement
}
