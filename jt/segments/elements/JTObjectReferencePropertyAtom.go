package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"errors"
)

// JT Object Reference Property Atom Element represents a property atom whose value is an object ID
// for another object within the JT file
type JTObjectReferencePropertyAtom struct {
	BasePropertyAtom
	VersionNumber uint8
	// Object ID specifies the identifier within the JT file for the referenced object.
	ObjectId int32
}

func (n JTObjectReferencePropertyAtom) GUID() model.GUID {
	return model.JTObjectReferencePropertyAtomElement
}

func (n *JTObjectReferencePropertyAtom) Read(c *model.Context) error {
	c.LogGroup("JTObjectReferencePropertyAtom")
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

	n.ObjectId = c.Data.Int32()
	c.Log("ObjectId: %f", n.ObjectId)

	return c.Data.GetError()
}

func (n *JTObjectReferencePropertyAtom) Value() interface {} {
	return n.ObjectId
}


func (n *JTObjectReferencePropertyAtom) GetPropertyAtom() *BasePropertyAtom {
	return &n.BasePropertyAtom
}

func (n *JTObjectReferencePropertyAtom) BaseElement() *JTElement {
	return &n.JTElement
}