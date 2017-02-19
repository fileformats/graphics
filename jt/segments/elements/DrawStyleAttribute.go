package elements

import "github.com/fileformats/graphics/jt/model"

// Draw Style Attribute Element contains information defining various aspects of the graphics state/style that
// should be used for rendering associated geometry.
// JT format LSG traversal semantics state that draw style attributes accumulate down the LSG by replacement
type DrawStyleAttribute struct {
	BaseAttribute
	// Version Number is the version identifier for this node.
	VersionNumber uint8
	// Data Flags is a collection of flags.  The flags are combined using the binary OR operator and
	// store various state settings for Draw Style Attribute Elements
	//      0x01 Back-face Culling Flag
	//      0x02 Two Sided Lighting Flag.
	//      0x04 Outlined Polygons Flag
	//      0x08 Lighting Enabled Flag
	//      0x10 Flat Shading Flag
	//      0x20 Separate Specular Flag
	DataFlags uint8
	// DataFlags1 field is only present if Version Number equals “1” and for JT files ^v8
	// The DataFlags1 includes the DataFlags data along with some additional flags.
	DataFlags1 uint8
}

func (n DrawStyleAttribute) GUID() model.GUID {
	return model.DrawStyleAttributeElement
}

func (n *DrawStyleAttribute) Read(c *model.Context) error {
	c.LogGroup("DrawStyleAttribute")
	defer c.LogGroupEnd()

	if err := (&n.BaseAttribute).Read(c); err != nil {
		return err
	}

	switch {
	case c.Version.Equal(model.V8):
		n.DataFlags = c.Data.UInt8()
		n.VersionNumber = uint8(c.Data.UInt16())
		if n.VersionNumber != 0 {
			c.Data.UInt8()
		}
	case c.Version.Equal(model.V9):
		n.VersionNumber = uint8(c.Data.UInt16())
		n.DataFlags = c.Data.UInt8()
	case c.Version.Equal(model.V10):
		n.VersionNumber = c.Data.UInt8()
		n.DataFlags = c.Data.UInt8()
	}

	return c.Data.GetError()
}

func (n *DrawStyleAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *DrawStyleAttribute) BaseElement() *JTElement {
	return &n.JTElement
}