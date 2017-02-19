package elements

import (
	"github.com/fileformats/graphics/jt/model"
)

// Light Set Attribute Element holds an unordered list of Lights.
// JT format LSG traversal semantics state that light set attributes accumulate down the LSG through
// addition of lights to an attribute list.
type LightSetAttribute struct {
	BaseAttribute
	// Version Number is the version identifier for this element
	VersionNumber uint8
	// Light Count specifies the number of lights in the Light Set
	LightCount int32
	// Light Object ID is the identifier for a referenced Light Object
	LightObjectIds []int32
}

func (n LightSetAttribute) GUID() model.GUID {
	return model.LightSetAttributeElement
}

func (n *LightSetAttribute) Read(c *model.Context) error {
	c.LogGroup("LightSetAttribute")
	defer c.LogGroupEnd()

	if err := (&n.BaseAttribute).Read(c); err != nil {
		return err
	}

	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.GreaterEqThan(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
		}
		c.Log("VersionNumber: %d", n.VersionNumber)
	}

	n.LightCount = c.Data.Int32()
	c.Log("LightCount: %d", n.LightCount)

	for i := 0; i < int(n.LightCount); i++ {
		n.LightObjectIds = append(n.LightObjectIds, c.Data.Int32())
	}
	c.Log("LightObjectIds: %s", n.LightObjectIds)

	return c.Data.GetError()
}

func (n *LightSetAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *LightSetAttribute) BaseElement() *JTElement {
	return &n.JTElement
}
