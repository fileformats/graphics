package elements

import "github.com/fileformats/graphics/jt/model"

// A Tri-Strip Set Shape LOD Element contains the geometric shape definition data (e.g. vertices, polygons, normals, etc.)
// for a single LOD of a collection of independent and unconnected triangle strips. Each strip constitutes one primitive
// of the set and  the ordering of the vertices in forming triangles, is the same as OpenGL‘s triangle strip definition
type TriStripSetShapeLOD struct {
	VertexShapeLOD
	// Version Number is the version identifier for this Tri-Strip Set Shape LOD
	VersionNumber uint8
}

func (n TriStripSetShapeLOD) GUID() model.GUID {
	return model.TriStripSetShapeLODElement
}

func (n *TriStripSetShapeLOD) Read(c *model.Context) error {
	c.LogGroup("TriStripSetShapeLOD")
	defer c.LogGroupEnd()

	if err := (&n.VertexShapeLOD).Read(c); err != nil {
		return err
	}

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
	} else {
		n.VersionNumber = uint8(c.Data.Int16())
	}
	c.Log("VersionNumber: %d", n.VersionNumber)

	if c.Version.Equal(model.V8) {
		if err := (&n.VertexShapeLOD).ReadCompressedData(c); err != nil {
			return err
		}
	} else {
		if err := (&n.VertexShapeLOD).ReadData(c, true); err != nil {
			return err
		}
	}
	return c.Data.GetError()
}

func (n *TriStripSetShapeLOD) BaseElement() *JTElement {
	return &n.JTElement
}