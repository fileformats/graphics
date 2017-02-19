package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"errors"
)

// A NULL Shape Node Element defines a shape which has no direct geometric primitive representation (i.e. it is empty/NULL).
// NULL Shape Node Elements are often used as “proxy/placeholder” nodes within the serialized LSG when the actual
// Shape LOD data is run time generated (i.e. not persisted).
type NullShapeNode struct {
	BaseShapeNode
	// Version Number is the version identifier for this node
	VersionNumber uint8
}

func (n NullShapeNode) GUID() model.GUID {
	return model.NullShapeNodeElement
}

func (n *NullShapeNode) Read(c *model.Context) error {
	c.LogGroup("NullShapeNode")
	defer c.LogGroupEnd()

	if err := (&n.BaseShapeNode).Read(c); err != nil {
		return err
	}
	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.GreaterEqThan(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
		}
		if n.VersionNumber != 1 {
			return errors.New("Invalid version number")
		}
	}
	c.Log("VersionNumber: %d", n.VersionNumber)

	return c.Data.GetError()
}

func (n *NullShapeNode) GetBaseNode() *BaseNode {
	return &n.BaseNode
}

func (n *NullShapeNode) BaseElement() *JTElement {
	return &n.JTElement
}
