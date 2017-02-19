package elements

import (
	"errors"
	"github.com/fileformats/graphics/jt/model"
)

type BasePropertyAtomElement interface {
	GetPropertyAtom() *BasePropertyAtom
}

// Base Property Atom Element represents the simplest form of a property that can exist within
// the LSG and has no type specific value data associated with it.
type BasePropertyAtom struct {
	JTElement
	// Object ID is the identifier for this Object.
	// Other objects referencing this particular object do so using the Object ID
	// NOTE: Deprecated. Only used in version ^8.0
	ObjectId int32
	// Version Number is the version identifier for this data collection
	VersionNumber uint8
	// State Flags is a collection of flags.
	// The flags are combined using the binary OR operator and store various state information for property atoms.
	// Bits 0 – 7 are freely available for an application to store whatever property atom information desired
	StateFlags uint32
}

func (n BasePropertyAtom) GUID() model.GUID {
	return model.BasePropertyAtomElement
}

func (n *BasePropertyAtom) Read(c *model.Context) error {
	c.LogGroup("BasePropertyAtom")
	defer c.LogGroupEnd()

	n.ObjectId = c.Data.Int32()
	c.Log("ObjectId: %d", n.ObjectId)

	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.GreaterEqThan(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
			// @TODO: Check this
			// c.Data.Skip(2)
		}
		if n.VersionNumber != 1 {
			return errors.New("Invalid version number")
		}
		c.Log("VersionNumber: %d", n.VersionNumber)
	}

	n.StateFlags = c.Data.UInt32()
	c.Log("StateFlags: %d", n.StateFlags)

	return c.Data.GetError()
}

func (n *BasePropertyAtom) BaseElement() *JTElement {
	return &n.JTElement
}
