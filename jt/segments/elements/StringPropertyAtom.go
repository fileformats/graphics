package elements

import (
	"errors"

	"github.com/fileformats/graphics/jt/model"
)

// String Property Atom Element represents a character string property atom.
type StringPropertyAtom struct {
	BasePropertyAtom
	// Version Number is the version identifier for this data collection
	VersionNumber uint8
	// Value contains the character string value for this property atom
	Value model.MbString
}

func (n StringPropertyAtom) GUID() model.GUID {
	return model.StringPropertyAtomElement
}

func (n *StringPropertyAtom) Read(c *model.Context) error {
	c.LogGroup("StringPropertyAtom")
	defer c.LogGroupEnd()

	if err := (&n.BasePropertyAtom).Read(c); err != nil {
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
	}
	c.Log("VersionNumber: %d", n.VersionNumber)

	if err := (&n.Value).Read(c); err != nil {
		return err
	}
	c.Log("Value: %s", n.Value)

	return c.Data.GetError()
}

func (n *StringPropertyAtom) GetPropertyAtom() *BasePropertyAtom {
	return &n.BasePropertyAtom
}

func (n *StringPropertyAtom) BaseElement() *JTElement {
	return &n.JTElement
}