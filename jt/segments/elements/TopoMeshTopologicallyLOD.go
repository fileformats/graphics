package elements

import (
	"errors"

	"github.com/fileformats/graphics/jt/model"
)

// TopoMesh Compressed LOD Data collection contains the common items to all TopoMesh Compressed LOD data elements.
type TopoMeshTopologicallyCompressedLOD struct {
	TopoMeshLOD
	VersionNumber uint8
}

func (n *TopoMeshTopologicallyCompressedLOD) Read(c *model.Context) error {
	c.LogGroup("TopoMeshTopologicallyCompressedLOD")

	if err := (&n.TopoMeshLOD).Read(c); err != nil {
		return err
	}

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
	} else if c.Version.GreaterEqThan(model.V9) {
		n.VersionNumber = uint8(c.Data.Int16())
	}
	if n.VersionNumber != 1 && n.VersionNumber != 2 {
		return errors.New("Invalid TopoMeshTopologicallyCompresseedLODData version number")
	}

	return c.Data.GetError()
}

func (n *TopoMeshTopologicallyCompressedLOD) readCompressedRepData(c *model.Context) error {
	//	var FaceDegrees = make([]codec.Int32CDP2, 8)
	// @TODO: implement this
	return c.Data.GetError()
}
