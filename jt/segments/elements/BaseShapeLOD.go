package elements

import "github.com/fileformats/graphics/jt/model"

type BaseShapeLOD struct {
	JTElement
	VersionNumber uint8
}

func (n BaseShapeLOD) GUID() model.GUID {
	return model.BaseShapeLODElement
}

func (n *BaseShapeLOD) Read(c *model.Context) error {
	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
	} else {
		n.VersionNumber = uint8(c.Data.Int16())
	}
	c.Log("VersionNumber: %d", n.VersionNumber)
	return c.Data.GetError()
}

func (n *BaseShapeLOD) BaseElement() *JTElement {
	return &n.JTElement
}