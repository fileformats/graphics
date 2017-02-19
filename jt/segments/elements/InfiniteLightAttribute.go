package elements

import "github.com/fileformats/graphics/jt/model"

// Infinite Light Attribute Element specifies a light source emitting unattenuated light in a single direction from
// every point on an infinite plane. The infinite location indicates that the rays of light can be considered parallel
// by the time they reach an object.
type InfiniteLightAttribute struct {
	BaseLight
	// Version Number is the version identifier for this element
	VersionNumber uint8
	// Direction specifies the direction the light is pointing in.
	Direction model.Vector3D
}

func (n InfiniteLightAttribute) GUID() model.GUID {
	return model.InfiniteLightAttributeElement
}

func (n *InfiniteLightAttribute) Read(c *model.Context) error {
	c.LogGroup("InfiniteLightAttribute")
	defer c.LogGroupEnd()

	if err := (&n.BaseLight).Read(c); err != nil {
		return err
	}

	if c.Version.Equal(model.V8) {
		c.Data.Unpack(&n.Direction)
		n.VersionNumber = uint8(c.Data.Int16())

		if n.VersionNumber == 1 {
			n.CoordSystem = c.Data.Int32()
			c.Log("CoordSystem: %d", n.CoordSystem)
		}
	}

	if c.Version.Equal(model.V9) {
		n.VersionNumber = uint8(c.Data.Int16())
		c.Data.Unpack(&n.Direction)

		if n.VersionNumber == 2 {
			n.ShadowOpacity = c.Data.Float32()
			c.Log("ShadowOpacity: %f", n.ShadowOpacity)
		}
	}

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = c.Data.UInt8()
		c.Data.Unpack(&n.Direction)
	}
	c.Log("VersionNumber: %d", n.VersionNumber)
	c.Log("Direction: %s", n.Direction)

	// TODO: verify this

	return c.Data.GetError()
}

func (n *InfiniteLightAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *InfiniteLightAttribute) BaseElement() *JTElement {
	return &n.JTElement
}