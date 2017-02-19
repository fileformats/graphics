package elements

import (
	"errors"
	"time"

	"github.com/fileformats/graphics/jt/model"
)

// Date Property Atom Element represents a property atom whose value is a “date”
type DatePropertyAtom struct {
	BasePropertyAtom
	VersionNumber uint8
	Year          int16
	Month         int16
	Day           int16
	Hour          int16
	Minute        int16
	Second        int16
}

func (n DatePropertyAtom) GUID() model.GUID {
	return model.DatePropertyAtomElement
}

func (n *DatePropertyAtom) Read(c *model.Context) error {
	c.LogGroup("DatePropertyAtom")
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
		if n.VersionNumber != 1 {
			return errors.New("Invalid version number")
		}
		c.Log("VersionNumber: %d", n.VersionNumber)
	}

	n.Year = c.Data.Int16()
	n.Month = c.Data.Int16()
	n.Day = c.Data.Int16()
	n.Hour = c.Data.Int16()
	n.Minute = c.Data.Int16()
	n.Second = c.Data.Int16()

	c.Log("Date: %d/%d/%d %d:%d:%d", n.Year, n.Month, n.Day, n.Hour, n.Minute, n.Second)

	return c.Data.GetError()
}

func (n *DatePropertyAtom) Time() time.Time {
	return time.Date(
		int(n.Year), time.Month(n.Month), int(n.Day),
		int(n.Hour), int(n.Minute), int(n.Second),
		0,
		time.UTC,
	)
}

func (n *DatePropertyAtom) GetPropertyAtom() *BasePropertyAtom {
	return &n.BasePropertyAtom
}

func (n *DatePropertyAtom) BaseElement() *JTElement {
	return &n.JTElement
}
