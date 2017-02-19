package elements

import (
	"github.com/fileformats/graphics/jt/model"
)

type PartNode struct {
	MetaDataNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
	Reserved      int32
}

func (n PartNode) GUID() model.GUID {
	return model.PartNodeElement
}

func (n *PartNode) Read(c *model.Context) error {
	c.LogGroup("PartNode")
	defer c.LogGroupEnd()

	(&n.MetaDataNode).Read(c)

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
	} else {
		n.VersionNumber = uint8(c.Data.Int16())
	}

	c.Log("VersionNumber: %d", n.VersionNumber)

	n.Reserved = c.Data.Int32()
	c.Log("Reserved: %d", n.Reserved)

	return c.Data.GetError()
}

func (n *PartNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *PartNode) BaseElement() *JTElement {
	return &n.JTElement
}
