package elements

import "github.com/fileformats/graphics/jt/model"

// Point Light Attribute Element specifies a light source emitting light from a specified position, along a
// specified direction, and with a specified spread angle
type PointLightAttribute struct {
	BaseLight
	// Version Number is the version identifier for this element
	VersionNumber uint8
	// Position specifies the light position in homogeneous coordinates
	Position model.HCoordF32
	// Constant Attenuation specifies the constant coefficient for how light intensity decreases with distance.
	// Value shall be greater than or equal to 0
	ConstantAttenuation float32
	// Linear Attenuation specifies the linear coefficient for how light intensity decreases with distance.
	// Value shall be greater than or equal to 0
	LinearAttenuation float32
	// Quadratic Attenuation specifies the quadratic coefficient for how light intensity decreases with distance.
	// Value shall be greater than or equal to 0
	QuadraticAttenuation float32
	// Spread Angle value with respect to the light cone below, specifies in degrees the half angle of the light cone.
	// Valid Spread Angle values are clamped and interpreted as follows:
	// angle == 180.0        Simple point light
	// 0.0 >= angle <= 90.0  Spot Light
	SpreadAngle float32
	// Spot Direction specifies the direction the spot light is pointing in
	SpotDirection model.Vector3D
	// Spot Intensity specifies the intensity distribution of the light within the spot light cone.
	// Spot Intensity is really a spot exponent in a lighting equation and indicates how focused the light is at the centre.
	// The larger the value, the more focused the light source.  Only non-negative Spot intensity values are valid
	SpotIntensity int32
}

func (n PointLightAttribute) GUID() model.GUID {
	return model.PointLightAttributeElement
}

func (n *PointLightAttribute) Read(c *model.Context) error {
	c.LogGroup("PointLightAttribute")
	defer c.LogGroupEnd()

	if err := (&n.BaseLight).Read(c); err != nil {
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

	c.Data.Unpack(&n.Position)
	c.Log("Position: %s", n.Position)

	n.ConstantAttenuation = c.Data.Float32()
	n.LinearAttenuation = c.Data.Float32()
	n.QuadraticAttenuation = c.Data.Float32()

	n.SpreadAngle = c.Data.Float32()

	c.Data.Unpack(&n.SpotDirection)

	n.SpotIntensity = c.Data.Int32()

	if c.Version.Equal(model.V9) && n.VersionNumber == 2 {
		n.ShadowOpacity = c.Data.Float32()
	}

	if c.Version.Equal(model.V8) && c.Data.Int16() == 1 {
		n.CoordSystem = c.Data.Int32()
	}

	return c.Data.GetError()
}

func (n *PointLightAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}


func (n *PointLightAttribute) BaseElement() *JTElement {
	return &n.JTElement
}