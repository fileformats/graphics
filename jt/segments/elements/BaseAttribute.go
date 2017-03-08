package elements

import (
	"errors"

	"github.com/fileformats/graphics/jt/model"
)

type BaseAttributeElement interface {
	GetBaseAttribute() *BaseAttribute
}

// Attribute Elements (e.g. colour, texture, material, lights, etc.) are placed in LSG as objects associated with nodes.
// Attribute Elements are not nodes themselves, but can be associated with any node.
type BaseAttribute struct {
	JTElement `json:"-"`
	// Object ID is the identifier for this Object.
	// Other objects referencing this particular object do so using the Object ID
	// NOTE: Deprecated. Only used in version ^8.0
	ObjectId int32 `json:"-"`
	// Version Number is the version identifier for this node
	VersionNumber uint8 `json:"-"`
	// State Flags is a collection of flags. The flags are combined using the binary OR operator and
	// store various state information for Attribute Elements; such as indicating that the attributes
	// accumulation is final.  All bits fields that are not defined as in use should be set to 0.
	//   0x01 - Unused for jt files >= v10
	//          = 0 – Accumulation is to occur normally
	//          = 1 – Accumulation is “final”
	//   0x02 - Accumulation Force flag. Provides a way to assign nodes in LSG, attributes that
	//          shall not be overridden by ancestors
	//          = 0 – Accumulation of this attribute obeys ancestor‘s Final flag setting
	//          = 1 – Accumulation of this attribute is forced
	//   0x03 - Accumulation Ignore Flag. Provides a way to indicate that the attribute is to be ignored.
	//          = 0 – Attribute is to be accumulated normally (subject to values of Force/Final flags)
	//          = 1 – Attribute is to be ignored.
	//   0x08 - Attribute Persistable Flag. Provides a way to indicate that the attribute is to be persistable to a JT file.
	//          = 0 – Attribute is to be non-persistable.
	//          = 1 – Attribute is to be persistable.
	StateFlags uint8 `json:"-"`
	// Field Inhibit Flags is a collection of flags, each flag corresponding to a collection of state data
	// within a particular Attribute type.
	FieldInhibitFlags uint32 `json:"-"`
	// Field Final Flags is a collection of flags, each flag being parallel to the corresponding flag in the
	// Field Inhibit Flags. If the field‘s bit in Field Final Flags is set, then that field within the
	// Attribute will become ―final and will not allow any subsequent accumulation into the specified field
	FieldFinalFlags uint32 `json:"-"`
}

func (n *BaseAttribute) GUID() model.GUID {
	return model.BaseAttributeData
}

func (n *BaseAttribute) Read(c *model.Context) error {
	c.LogGroup("BaseAttribute")
	defer c.LogGroupEnd()

	n.ObjectId = c.Data.Int32()
	c.Log("ObjectId: %d", n.ObjectId)

	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.Equal(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
			// @TODO: must check here the 2 extra flags
		}
		if n.VersionNumber != 1 {
			return errors.New("Invalid version number")
		}
	}

	n.StateFlags = c.Data.UInt8()
	c.Log("StateFlags: %d", n.StateFlags)

	n.FieldInhibitFlags = c.Data.UInt32()
	c.Log("FieldInhibitFlags: %d", n.FieldInhibitFlags)

	if c.Version.GreaterEqThan(model.V10) {
		n.FieldFinalFlags = c.Data.UInt32()
		c.Log("FieldFinalFlags: %d", n.FieldFinalFlags)
	}

	return c.Data.GetError()
}

func (n *BaseAttribute) ReadLight(c *model.Context) error {
	c.LogGroup("BaseAttribute")
	defer c.LogGroupEnd()

	n.ObjectId = c.Data.Int32()
	c.Log("ObjectId: %d", n.ObjectId)

	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.Equal(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
			// @TODO: must check here the 2 extra flags
		}
		if n.VersionNumber != 1 {
			return errors.New("Invalid version number")
		}
	}

	n.StateFlags = c.Data.UInt8()
	c.Log("StateFlags: %d", n.StateFlags)

	if c.Version.GreaterEqThan(model.V10) {
		n.FieldFinalFlags = c.Data.UInt32()
		c.Log("FieldFinalFlags: %d", n.FieldFinalFlags)
	}

	return c.Data.GetError()
}

func (n *BaseAttribute) BaseElement() *JTElement {
	return &n.JTElement
}
