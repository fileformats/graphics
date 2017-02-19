package elements

import "github.com/fileformats/graphics/jt/model"

type BaseLight struct {
	BaseAttribute
	// Version number is the version identifier for this element
	VersionNumber uint8
	// Ambient Color specifies the ambient red, green, blue, alpha color values of the light.
	AmbientColor model.RGBA
	// Diffuse Color specifies the diffuse red, green, blue, alpha color values of the light
	DiffuseColor model.RGBA
	// Specular Color specifies the specular red, green, blue, alpha color values of the light.
	SpecularColor model.RGBA
	// Brightness specifies the Light brightness.  The Brightness value must be greater than or equal to “-1”.
	Brightness float32
	// Coord System specifies the coordinate space in which Light source is defined.  Valid values include the following:
	// = 1 Viewpoint Coordinate System. Light source is to move together with the viewpoint
	// = 2 Model Coordinate System. Light source is affected by whatever model transforms that are current when
	//     the light source is encountered in LSG.
	// = 3 World Coordinate system. Light source is not affected by model transforms in the LSG.
	CoordSystem int32
	// Shadow Caster Flag is a flag that indicates whether the light is a shadow caster or not.
	// = 0 Light source is not a shadow caster.
	// = 1 Light source is a shadow caster
	ShadowCasterFlag uint8
	// Shadow Opacity specifies the shadow opacity factor on Light source. Value must be within range [0.0, 1.0]
	// inclusive.  Shadow Opacity is intended to convey how dark a shadow cast by this light source are to be rendered.
	// A value of 1.0 means that no light from this light source reaches a shadowed surface, resulting in a black shadow
	ShadowOpacity float32
	// Non-shadow Alpha Factor is one of a matched pair of fields intended to govern how a shadowing light source (one
	// whose Shadow Caster Flag is set) casts "alpha light" into areas that it directly illuminates (i.e. are not in
	// shadow).  Those fragments directly lit by this light source will have their alpha values scaled by Non-shadow
	// Alpha Factor.  Non-shadow Alpha Factor value shall lie on the range [0.0, 1.0] inclusive
	NonShadowAlphaFactor float32
	// Shadow Alpha Factor is one of a matched pair of fields intended to govern how a shadowing light source (one
	// whose Shadow Caster Flag is set) casts "alpha light" into areas that it does not illuminate (i.e. are in
	// shadow).  Those fragments in shadow from this light source will have their alpha values scaled by Shadow
	// Alpha Factor.  Shadow Alpha Factor value shall lie on the range [0.0, 1.0] inclusive
	ShadowAlphaFactor float32
}

func (n *BaseLight) Read(c *model.Context) error {
	if err := (&n.BaseAttribute).ReadLight(c); err != nil {
		return err
	}

	if c.Version.Equal(model.V9) {
		n.VersionNumber = uint8(c.Data.Int16())
	}

	if c.Version.GreaterEqThan(model.V10) {
		n.VersionNumber = uint8(c.Data.Int16())
	}

	c.Data.Unpack(&n.AmbientColor)
	c.Log("AmbientColor: %s", n.AmbientColor)
	c.Data.Unpack(&n.DiffuseColor)
	c.Log("DiffuseColor: %s", n.DiffuseColor)
	c.Data.Unpack(&n.SpecularColor)
	c.Log("SpecularColor: %s", n.SpecularColor)
	n.Brightness = c.Data.Float32()
	c.Log("Brightness: %f", n.Brightness)

	if c.Version.GreaterEqThan(model.V9) {
		n.CoordSystem = c.Data.Int32()
		c.Log("CoordSystem: %d", n.CoordSystem)
		n.ShadowCasterFlag = c.Data.UInt8()
		c.Log("ShadowCasterFlag: %d", n.ShadowCasterFlag)
		n.ShadowOpacity = c.Data.Float32()
		c.Log("ShadowOpacity: %f", n.ShadowOpacity)
	}

	if c.Version.GreaterEqThan(model.V10) {
		n.NonShadowAlphaFactor = c.Data.Float32()
		c.Log("NonShaodwAlphaFactor: %f", n.NonShadowAlphaFactor)
		n.ShadowAlphaFactor = c.Data.Float32()
		c.Log("ShadowAlphaFactor: %f", n.ShadowAlphaFactor)
	}

	// TODO: verify this

	return c.Data.GetError()
}

func (n *BaseLight) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *BaseLight) BaseElement() *JTElement {
	return &n.JTElement
}
