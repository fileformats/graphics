package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"errors"
)

type IntegerPropertyAtom struct {
	BasePropertyAtom
	VersionNumber uint8
	Value int32
}

func (n IntegerPropertyAtom) GUID() model.GUID {
	return model.IntegerPropertyAtomElement
}

func (n *IntegerPropertyAtom) Read(c *model.Context) error {
	c.LogGroup("IntegerPropertyAtom")
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
		c.Log("VersionNumber: %d", n.VersionNumber)
	}

	n.Value = c.Data.Int32()
	c.Log("Value: %f", n.Value)

	return c.Data.GetError()
}

func (n *IntegerPropertyAtom) GetPropertyAtom() *BasePropertyAtom {
	return &n.BasePropertyAtom
}

func (n *IntegerPropertyAtom) BaseElement() *JTElement {
	return &n.JTElement
}