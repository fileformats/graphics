package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"errors"
)

// Late Loaded Property Atom Element is a property atom type used to reference an associated piece of atomic
// data in a separate addressable segment of the JT file. The connotation derives from the associated data being
// stored in a separate addressable segment of the JT file, and thus a JT file reader can be structured to support
// the best practice of delaying the loading/reading of the associated data until it is actually needed.
type LateLoadedPropertyAtom struct {
	BasePropertyAtom
	// Version Number is the version identifier for this data collection
	VersionNumber uint8
	// Segment ID is the globally unique identifier for the associated data segment in the JT file
	SegmentId model.GUID
	// Segment Type defines a broad classification of the associated data segment contents
	SegmentType int32
	// Object ID is the identifier for the payload.  Other objects referencing this particular
	// payload will do so using the Object ID
	PayloadObjectId int32
	// Reserved data field that is guaranteed to always be greater than or equal to 1
	Reserved int32
}

func (n LateLoadedPropertyAtom) GUID() model.GUID {
	return model.LateLoadedPropertyAtomElement
}

func (n *LateLoadedPropertyAtom) Read(c *model.Context) error {
	c.LogGroup("LateLoadedPropertyAtom")
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

	c.Data.Unpack(&n.SegmentId)
	c.Log("SegmentId: %s (%s)", n.SegmentId, n.SegmentId.Name())

	n.SegmentType = c.Data.Int32()
	c.Log("SegmentType: %d", n.SegmentType)

	if c.Version.GreaterEqThan(model.V9) {
		n.PayloadObjectId = c.Data.Int32()
		c.Log("PayloadObjectId: %d", n.PayloadObjectId)
		n.Reserved = c.Data.Int32()
		c.Log("Reserved: %d", n.Reserved)
	}

	return c.Data.GetError()
}

func (n *LateLoadedPropertyAtom) GetPropertyAtom() *BasePropertyAtom {
	return &n.BasePropertyAtom
}

func (n *LateLoadedPropertyAtom) BaseElement() *JTElement {
	return &n.JTElement
}