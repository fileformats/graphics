package elements

import "github.com/fileformats/graphics/jt/model"

type MappingCylinder struct {
	JTElement
	// Version Number is the version identifier for this element
	VersionNumber uint8
	// Mapping Plane Matrix specifies the transformation matrix and mapping parameters for the mapping plane
	MappingPlaneMatrix model.Matrix4F64
	// Coordinate system specifies the coordinate space in which mapping plane is defined. Valid values include the following
	// = 0 Undefined Coordinate System.
	// = 1 Viewpoint Coordinate System. Mapping plane is to move together with the viewpoint.
	// = 2 Model Coordinate System. Mapping plane is affected by whatever model transforms that are current when
	//     the mapping plane is encountered in LSG.
	// = 3 World Coordinate system. Mapping plane is not affected by model transforms in the LSG
	CoordinateSystem int32
}

func (n MappingCylinder) GUID() model.GUID {
	return model.MappingCylinderElement
}

func (n *MappingCylinder) Read(c *model.Context) error {
	c.LogGroup("MappingCylinder")
	defer c.LogGroupEnd()

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
	} else if c.Version.GreaterEqThan(model.V9) {
		n.VersionNumber = uint8(c.Data.Int16())
	}
	n.MappingPlaneMatrix = model.Matrix4F64{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
	for i := 0; i < 16; i++ {
		n.MappingPlaneMatrix[i] = c.Data.Float64()
	}
	n.CoordinateSystem = c.Data.Int32()
	return c.Data.GetError()
}

func (n *MappingCylinder) BaseElement() *JTElement {
	return &n.JTElement
}