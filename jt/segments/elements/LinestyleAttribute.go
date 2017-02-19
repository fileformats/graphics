package elements

import "github.com/fileformats/graphics/jt/model"

// Linestyle Attribute Element contains information defining the graphical properties to be used for rendering polylines.
// JT format LSG traversal semantics state that Linestyle attributes accumulate down the LSG by replacement.
type LinestyleAttribute struct {
	BaseAttribute
	// Version Number is the version identifier for this node
	VersionNumber uint8
	// Data Flags is a collection of flags and line type data. The flags and line type data are combined using the
	// binary OR operator and store various polyline rendering attributes. All bits fields that are not defined as
	// in use should be set to
	// 0x0F Line Type (stored in bits 0 – 3). Line type specifies the polyline rendering stipple-pattern.
	//      0 - Solid
	//      1 - Dash
	//      2 - Dot
	//      3 - Dash-Dot
	//      4 - Dash-Dot-Dot
	//      5 = Long-Dash
	//      6 = Center-Dash
	//      7 = Center-Dash=Dash
	// 0x10 Antialiasing Flag (stored in bit 4). Indicates if antialiasing should be applied as part of rendering polylines.
	// = 0 – Antialiasing disabled.
	// = 1 – Antialiasing enabled.
	DataFlags uint8
	// Line Width specifies the width in pixels that should be used for rendering polylines.
	// The value of this field shall be greater than 0.0.
	LineWidth float32
}

func (n LinestyleAttribute) GUID() model.GUID {
	return model.LinestyleAttributeElement
}

func (n *LinestyleAttribute) Read(c *model.Context) error {
	c.LogGroup("LinestyleAttribute")
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

	n.DataFlags = c.Data.UInt8()
	c.Log("DataFlags: %d", n.DataFlags)

	n.LineWidth = c.Data.Float32()
	c.Log("LineWidth: %d", n.LineWidth)

	return c.Data.GetError()
}

func (n *LinestyleAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *LinestyleAttribute) BaseElement() *JTElement {
	return &n.JTElement
}