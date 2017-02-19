package elements

import "github.com/fileformats/graphics/jt/model"

// TopoMesh LOD Data collection contains the common items to all TopoMesh LOD elements.
type TopoMeshLOD struct {
	// Version Number is the version identifier for this TopoMesh LOD Data
	VersionNumber uint8
	// Vertex Records Object ID is the identifier for the vertex records associated with this Object.
	// Other objects referencing these vertex records will do so using this Object ID
	VertexRecordsObjectId uint32
}

func (n *TopoMeshLOD) Read(c *model.Context) error {
	c.LogGroup("TopoMeshLOD")

	if c.Version.Equal(model.V8) {
		n.VersionNumber = uint8(c.Data.Int16())
	} else if c.Version.Equal(model.V9) {
		n.VersionNumber = uint8(c.Data.Int16())
		n.VertexRecordsObjectId = uint32(c.Data.Int32())
	} else if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
		n.VertexRecordsObjectId = c.Data.UInt32()
	}


	c.Log("VersionNumber: %d", n.VersionNumber)

	// @TODO: implement this
	return c.Data.GetError()
}