package elements

import (
	"errors"
	"github.com/fileformats/graphics/jt/model"
)

// Material Attribute Element defines the reflective characteristics of a material.
// JT format LSG traversal semantics dictate that material attributes accumulate down the LSG by replacement
// The Field Inhibit flag  from the Base Attribute bit assignments for the Material Attribute Element data
// fields, are as follow:
// Version JT ^v10
// - 0 Ambient Common RGB Value, Ambient Colour
// - 1 Specular Common RGB Value, Specular Colour
// - 2 Emission Common RGB Value, Emission Colour
// - 3 Blending Flag, Source Blending Factor, Destination Blending Factor
// - 4 Override Vertex Colour Flag
// - 5 Material Reflectivity
// - 6 Diffuse Colour
// - 7 Diffuse Alpha
//
// Version JT ^v9
// - 0 Ambient Common RGB Value, Ambient Color
// - 1 Diffuse Color and Alpha (Legacy)
// - 2 Specular Common RGB Value, Specular Color
// - 3 Emission Common RGB Value, Emission Color
// - 4 Blending Flag, Source Blending Factor, Destination Blending Factor
// - 5 Override Vertex Color Flag
// - 6 Material Reflectivity
// - 7 Diffuse Color
// - 8 Diffuse Alpha
//
// Version JT ^v8
// - 0 Ambient Common RGB Value, Ambient Color
// - 1 Diffuse Color
// - 2 Specular Common RGB Value, Specular Color
// - 3 Emission Common RGB Value, Emission Color
// - 4 Blending Flag, Source Blending Factor, Destination Blending Factor
// - 5 Override Vertex Color Flag
type MaterialAttribute struct {
	BaseAttribute
	// Version Number is the version identifier for this element
	VersionNumber uint8
	// Data Flags is a collection of flags and factor data. The flags and factor data are combined
	// using the binary OR operator.  The flags store information to be used for interpreting
	// how to read subsequent Material data fields
	//      0x0010 - Blending Flag
	//      0x0020 - Override Vertex Colours Flag
	//      0x07C0 - Source Blend Factor (stored in bits 6-10). If Blending Flag enabled, this value
	//               indicates how the incoming fragment‘s  RGBA colour values are to be used to blend with
	//               the current framebuffer‘s
	//               = 0 – Interpret same as OpenGL GL_ZERO Blending Factor
	//               = 1 – Interpret same as OpenGL GL_ONE Blending Factor
	//               = 2 – Interpret same as OpenGL GL_DST_COLOUR Blending Factor
	//               = 3 – Interpret same as OpenGL GL_SRC_COLOUR Blending Factor
	//               = 4 – Interpret same as OpenGL GL_ONE_MINUS_DST_COLOUR Blending Factor
	//               = 5 – Interpret same as OpenGL GL_ONE_MINUS_SRC_COLOUR Blending Factor
	//               = 6 – Interpret same as OpenGL GL_SRC_ALPHA Blending Factor
	//               = 7 – Interpret same as OpenGL GL_ONE_MINUS_SRC_ALPHA Blending Factor
	//               = 8 – Interpret same as OpenGL GL_DST_ALPHA Blending Factor
	//               = 9 – Interpret same as OpenGL GL_ONE_MINUS_DST_ALPHA Blending Factor
	//               = 10 – Interpret same as OpenGL GL_SRC_ALPHA_SATURATE Blending Factor
	//     0xF800 - Destination Blend Factor (stored in bits 11-15). If Blending Flag enabled, this value indicates how
	//              the current framebuffer‘s (the destination) RGBA colour values are to be used to blend
	//              with the incoming fragment‘s (the source) RGBA colour values.
	//				= 0 – Interpret same as OpenGL GL_ZERO Blending Factor
	//              = 1 – Interpret same as OpenGL GL_ONE Blending Factor
	//              = 2 – Interpret same as OpenGL GL_DST_COLOUR Blending Factor
	//              = 3 – Interpret same as OpenGL GL_SRC_COLOUR Blending Factor
	//              = 4 – Interpret same as OpenGL GL_ONE_MINUS_DST_COLOUR Blending Factor
	//              = 5 – Interpret same as OpenGL GL_ONE_MINUS_SRC_COLOUR Blending Factor
	//              = 6 – Interpret same as OpenGL GL_SRC_ALPHA Blending Factor
	//              = 7 – Interpret same as OpenGL GL_ONE_MINUS_SRC_ALPHA Blending Factor
	//              = 8 – Interpret same as OpenGL GL_DST_ALPHA Blending Factor
	//              = 9 – Interpret same as OpenGL GL_ONE_MINUS_DST_ALPHA Blending Factor
	//              = 10 – Interpret same as OpenGL GL_SRC_ALPHA_SATURATE Blending Factor
	DataFlags uint16
	// Ambient Colour specifies the ambient red, green, blue, alpha colour values of the material
	AmbientColor model.RGBA
	// Diffuse Colour and Alpha specify the diffuse red, green, blue colour components, and alpha value of the material.
	DiffuseColor model.RGBA
	// Specular Colour specifies the specular red, green, blue, alpha colour values of the material
	SpecularColor model.RGBA
	// Emission Colour specifies the emissive red, green, blue, alpha colour values of the material
	EmissionColor model.RGBA
	// Shininess is the exponent associated with specular reflection and highlighting of the Phong specular lighting
	// model. Shininess controls the degree with which the specular highlight decays.
	// Only values in the range [1,128] are valid
	Shininess float32
	// Reflectivity specifies the material reflectivity of the material. It represents the fraction of
	// light reflected in the mirror direction by the material.
	// Only values in the range [0.0, 1.0] are valid
	Reflectivity float32
	// Bumpiness is used to control bump mapping, and specifies the degree to which bump mapping modifies the local normal vector.
	// A value of 1.0 is the default.
	// Values larger than 1.0 are intended to make the shaded object look as if it is more highly embossed;
	// values between 0.0 and 1.0 make it look less so.
	// Negative values are legal and make the object appear to be engraved rather than embossed
	Bumpiness    float32
}

func (n MaterialAttribute) GUID() model.GUID {
	return model.MaterialAttributeElement
}

func (n *MaterialAttribute) Read(c *model.Context) error {
	c.LogGroup("MaterialAttribute")
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
		if n.VersionNumber != 1 && n.VersionNumber != 2 {
			return errors.New("Invalid version number")
		}
	}

	n.DataFlags = c.Data.UInt16()
	c.Log("DataFlags: %d", n.DataFlags)

	usePattern := n.DataFlags & 1 != 0
	useAmbientPattern := n.DataFlags & 2 != 0
	useSpecularPattern := n.DataFlags & 4 != 0
	useEmissionPattern := n.DataFlags & 8 != 0


	if usePattern && useAmbientPattern {
		color := c.Data.Float32()
		n.AmbientColor = model.RGBA{R:color, G:color, B:color, A:1}
	} else {
		c.Data.Unpack(&n.AmbientColor)
	}
	c.Log("AmbientColor: %s", n.AmbientColor)

	c.Data.Unpack(&n.DiffuseColor)
	c.Log("DiffuseColor: %s", n.DiffuseColor)

	if usePattern && useSpecularPattern {
		color := c.Data.Float32()
		n.SpecularColor = model.RGBA{R:color, G:color, B:color, A:1}
	} else {
		c.Data.Unpack(&n.SpecularColor)
	}
	c.Log("SpecularColor: %f", n.SpecularColor)

	if usePattern && useEmissionPattern {
		color := c.Data.Float32()
		n.EmissionColor = model.RGBA{R:color, G:color, B:color, A:1}
	} else {
		c.Data.Unpack(&n.EmissionColor)
	}
	c.Log("EmissionColor: %f", n.EmissionColor)

	n.Shininess = c.Data.Float32()
	c.Log("Shininess: %f", n.Shininess)

	if n.VersionNumber == 2 {
		n.Reflectivity = c.Data.Float32()
		c.Log("Reflectivity: %f", n.Reflectivity)
	}

	return c.Data.GetError()
}

func (n *MaterialAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *MaterialAttribute) BaseElement() *JTElement {
	return &n.JTElement
}
