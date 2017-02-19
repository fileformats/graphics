package elements

import "github.com/fileformats/graphics/jt/model"

// The Switch Node is very much like a Group Node in that it contains an ordered list of references to other nodes,
// called the children nodes.  The difference is that a Switch Node also contains additional data indicating
// which child (one or none) a LSG traverser should process/traverse
type SwitchNode struct {
	GroupNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// Selected Child is the index for the selected child node.
	// Valid Selected Child values reside within the following range: “-1 < Selected Child < Child Count”.
	// Where “-1” indicates that no child is to be selected and “Child Count” is the data field value from GroupNode
	SelectedChild int32
}

func (n SwitchNode) GUID() model.GUID {
	return model.SwitchNodeElement
}

func (n *SwitchNode) Read(c *model.Context) error {
	c.LogGroup("SwitchNode")
	defer c.LogGroupEnd()

	if err := (&n.GroupNode).Read(c); err != nil {
		return err
	}

	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.Equal(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.UInt16())
		}
	}

	n.SelectedChild = c.Data.Int32()
	c.Log("SelectedChild: %d", n.SelectedChild)

	return c.Data.GetError()
}

func (n *SwitchNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *SwitchNode) BaseElement() *JTElement {
	return &n.JTElement
}