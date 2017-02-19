package elements

import "github.com/fileformats/graphics/jt/model"

type MetaDataNode struct {
	GroupNode
	// Version Number is the version identifier for this data
	VersionNumber uint8
}

func (n MetaDataNode) GUID() model.GUID {
	return model.MetaDataNodeElement
}

func (n *MetaDataNode) Read(c *model.Context) error {
	c.LogGroup("MetaDataNode")
	defer c.LogGroupEnd()

	(&n.GroupNode).Read(c)

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
	} else {
		n.VersionNumber = uint8(c.Data.Int16())
	}
	c.Log("VersionNumber: %d", n.VersionNumber)

	return c.Data.GetError()
}

func (n *MetaDataNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *MetaDataNode) BaseElement() *JTElement {
	return &n.JTElement
}