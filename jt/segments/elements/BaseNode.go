package elements

import (
	"errors"

	"github.com/fileformats/graphics/jt/model"
)

type BaseNodeElement interface {
	GetBaseNode() *BaseNode
}

// 10dd1035-2ac8-11d1-9b-6b-00-80-c7-bb-59-97
// Base Node Element represents the simplest form of a node that can exist within the LSG.
// The Base Node Element has no implied LSG semantic behavior nor can it contain any children nodes
type BaseNode struct {
	JTElement
	// Object ID is the identifier for this Object.
	// Other objects referencing this particular object do so using the Object ID
	// NOTE: Deprecated. Only used in version ^8.0
	ObjectId int32
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// Node Flags is a collection of flags.  The flags are combined using the binary OR operator
	NodeFlags uint32
	// Attribute Count indicates the number of Attribute Objects referenced by this Node Object.
	// A node may have zero Attribute Object references
	AttributeCount int32
	// Attribute Object ID is the identifier for a referenced Attribute Object
	AttributeObjectId []int32
	// Reference to the node attributes
	Attributes []BaseAttributeElement
	// Reference to the node property atoms
	Properties map[BasePropertyAtomElement]BasePropertyAtomElement
	// Child nodes
	Children []BaseNodeElement
}

func (n *BaseNode) GUID() model.GUID {
	return model.BaseNodeElement
}

func (n *BaseNode) Ignore() bool {
	return n.NodeFlags&1 == 1
}

func (n *BaseNode) Read(c *model.Context) error {
	c.LogGroup("BaseNode")
	defer c.LogGroupEnd()

	n.ObjectId = c.Data.Int32()
	c.Log("ObjectId: %d", n.ObjectId)

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
		if n.VersionNumber != 1 {
			return errors.New("Invalid version number")
		}
	} else if c.Version.GreaterEqThan(model.V9) {
		n.VersionNumber = uint8(c.Data.Int16())
		if n.VersionNumber != 1 {
			return errors.New("Invalid version number")
		}
	}
	c.Log("VersionNumber: %d", n.VersionNumber)

	n.NodeFlags = c.Data.UInt32()
	c.Log("NodeFlags: %d", n.NodeFlags)

	n.AttributeCount = c.Data.Int32()
	c.Log("AttributeCount: %d", n.AttributeCount)

	for i := 0; i < int(n.AttributeCount); i++ {
		n.AttributeObjectId = append(n.AttributeObjectId, c.Data.Int32())
	}
	c.Log("Attributes: %v", n.AttributeObjectId)

	return c.Data.GetError()
}

func (n *BaseNode) BaseElement() *JTElement {
	return &n.JTElement
}
