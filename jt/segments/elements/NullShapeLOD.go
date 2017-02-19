package elements

import "github.com/fileformats/graphics/jt/model"

// A Null Shape LOD Element represents the pseudo geometric shape definition data for a NULL Shape Node Element.
// Although a NULL Shape Node Element has no real geometric primitive representation (i.e. is empty), its usage as a
// proxy/placeholder node within the LSG still supports the concept of having a defined bounding box and thus the
// existence of this Null Shape LOD Element type
type NullShapeLOD struct {
	JTElement
	// Version Number is the version identifier for this Null Shape LOD Element
	VersionNumber uint8
	// The Untransformed BBox is an axis-aligned LCS bounding box and represents the untransformed extents
	// for this Null Shape LOD Element
	UntransformedBBox model.BoundingBox
}

func (n NullShapeLOD) GUID() model.GUID {
	return model.NullShapeLODElement
}

func (n *NullShapeLOD) Read(c *model.Context) error {
	c.LogGroup("NullShapeLOD")
	defer c.LogGroupEnd()

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
	} else {
		n.VersionNumber = uint8(c.Data.Int16())
	}
	c.Log("VersionNumber: %d", n.VersionNumber)

	c.Data.Unpack(&n.UntransformedBBox)
	c.Log("UntransformedBBox: %s", n.UntransformedBBox)

	return c.Data.GetError()
}

func (n *NullShapeLOD) BaseElement() *JTElement {
	return &n.JTElement
}