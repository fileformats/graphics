package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"errors"
)

type FloatingPointPropertyAtom struct {
	BasePropertyAtom
	VersionNumber uint8
	FloatValue float32
}

func (n FloatingPointPropertyAtom) GUID() model.GUID {
	return model.FloatingPointPropertyAtomElement
}

func (n *FloatingPointPropertyAtom) Read(c *model.Context) error {
	c.LogGroup("FloatingPointPropertyAtom")
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

	n.FloatValue = c.Data.Float32()
	c.Log("Value: %f", n.Value)

	return c.Data.GetError()
}

func (n *FloatingPointPropertyAtom) Value() interface{} {
	return n.Value
}

func (n *FloatingPointPropertyAtom) GetPropertyAtom() *BasePropertyAtom {
	return &n.BasePropertyAtom
}

func (n *FloatingPointPropertyAtom) BaseElement() *JTElement {
	return &n.JTElement
}