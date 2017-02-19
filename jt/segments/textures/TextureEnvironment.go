package textures

import "github.com/fileformats/graphics/jt/model"

// The Texture Environment is a collection of data defining various aspects of how a texture image is
// to be mapped/applied to a surface.
type TextureEnvironment struct {
	TextureEnvVersion uint8
	// Border Mode specifies the texture border mode.
	// = 0 No border.
	// = 1 Constant Border Color. Indicates that the texture has a constant border color whose value is
	//     defined in data field Border Color.
	// = 2 Explicit. Indicates that a border texel ring is present in the texture image definition.
	BorderMode int32
	// Mipmap Magnification Filter specifies the texture filtering method to apply when a single pixel
	// on screen maps to a tiny portion of a texel
	// = 0 None.
	// = 1 Nearest. Texel with coordinates nearest the centre of the pixel is used.
	// = 2 Linear. A weighted linear average of the 2 x 2 array of texels nearest to the centre of the pixel is used.
	//     For one-dimensional texture is average of 2 texels. For three dimensional texel is 2 x 2 x 2 array.
	MipmapMagnificationFilter int32
	// Mipmap Minification Filter specifies the texture filtering method to apply when a single pixel on
	// screen maps to a large collection of texels.
	// = 0 None.
	// = 1 Nearest. Texel with coordinates nearest the center of the pixel is used.
	// = 2 Linear. A weighted linear average of the 2 x 2 array of texels nearest to the center of the pixel is used.
	//     For one-dimensional texture is average of 2 texels. For three-dimensional texture is 2 x 2 x 2 array.
	// = 3 Nearest in Mipmap. Within an individual mipmap, the texel with coordinates nearest the center of the pixel is used.
	// = 4 Linear in Mipmap. Within an individual mipmap, a weighted linear average of the 2 x 2 array of texels nearest
	//     to the center of the pixel is used. For one-dimensional texture is average of 2 texels.
	//     For three-dimensional texture is 2 x 2 x 2 array
	// = 5 Nearest between Mipmaps. Within each of the adjacent two mipmaps, selects the texel with coordinates nearest
	//     the center of the pixel and then interpolates linearly between these two selected mipmap values.
	// = 6 Linear between Mipmaps. Within each of the two adjacent mipmaps, computes value based on a weighted linear
	//     average of the 2 x 2 array of texels nearest to the center of the pixel and then interpolates linearly
	//     between these two computed mipmap values
	MipmapMinificationFilter int32
	// SDimensionWrapMode specifies the mode for handling texture coordinates S-Dimension values outside the range [0, 1].
	// = 0 None.
	// = 1 Clamp. Any values greater than 1.0 are set to 1.0; any values less than 0.0 are set to 0.0
	// = 2 Repeat Integer parts of the texture coordinates are ignored (i.e. retains only the fractional component
	//     o texture coordinates greater than 1.0 and only one-minus the fractional component of values less than zero).
	//     Resulting in copies of the texture map tiling the surface
	// = 3 Mirror Repeat. Like Repeat, except the surface tiles “flip-flop” resulting in an alternating mirror pattern
	//     of surface tiles.
	// = 4 Clamp to Edge. Border is always ignored and instead texel at or near the edge is chosen for coordinates
	//     outside the range [0, 1]. Whether the exact nearest edge texel or some average of the nearest edge texels
	//     is used is dependent upon the mipmap filtering value.
	// = 5 Clamp to Border. Nearest border texel is chosen for coordinates outside the range [0, 1]. Whether the exact
	//     nearest border texel or some average of the nearest border texels is used is dependent upon the
	//     mipmap filtering value.
	SDimensionWrapMode int32
	// TDimensionWrapMode specifies the mode for handling texture coordinates T-Dimension values outside the range [0, 1].
	// Same mode values as documented for SDimensionWrapMode
	TDimensionWrapMode int32
	// R-Dimen Wrap Mode specifies the mode for handling texture coordinates R-Dimension values outside the range [0, 1].
	// Same mode values as documented for SDimensionWrapMode
	RdimentsionWrapMode int32
	// Texture Function Data contains information indicating how the values in the texture map are to be
	// modulated/combined/blended with the original color of the surface or some other alternative color to compute
	// the final color to be painted on the surface
	// Bits 0 - 2 Texture Environment Mode
	//            = 0 − None.
	//            = 1 − Decal. Interpret same as OpenGL GL_DECAL environment mode.
	//            = 2 − Modulate. Interpret same as OpenGL GL_MODULATE environment mode.
	//            = 3 − Replace. Interpret same as OpenGL GL_REPLACE environment mode.
	//            = 4 − Blend. Interpret same as OpenGL GL_BLEND environment mode.
	//            = 5 − Add. Interpret same as OpenGL GL_ADD environment mode.
	//            = 6 − Combine. Interpret same as OpenGL GL_COMBINE environment mode
	// Bit 3 Environment Mapping Flag
	// Bits 4 - 31 Reserved for future use
	TextureFunctionData int32
	// Blend Type contains information indicating how the values in the texture map are to be modulated/combined/blended
	// with the original color of the surface or some other alternative color to compute the final color to be
	// painted on the surface.
	// = 0 None.
	// = 1 Decal. Interpret same as OpenGL GL_DECAL environment mode.
	// = 2 Modulate. Interpret same as OpenGL GL_MODULATE environment mode.
	// = 3 Replace. Interpret same as OpenGL GL_REPLACE environment mode.
	// = 4 Blend. Interpret same as OpenGL GL_BLEND environment mode.
	// = 5 Add. Interpret same as OpenGL GL_ADD environment mode.
	// = 6 Combine. Interpret same as OpenGL GL_COMBINE environment mode.
	BlendType int32
	// Internal Compression Level specifies a data compression hint/recommendation that a JT file loader is free to
	// follow for internally (in memory) storing texel data.
	// = 0 None. No compression of texel data.
	// = 1 Conservative. Lossless compression of texel data.
	// = 2 Moderate. Texel components truncated to 8-bits each.
	// = 3 Aggressive. Texel components truncates to 4-bits each (or 5 bits for RGB images).
	InternalCompressionLevel int32
	// Blend Color specifies the color to be used for the “Blend” mode of Blend Type operations
	BlendColor model.RGBA
	// Border Color specifies the constant border color to use for “Clamp to Border” style wrap modes when the
	// texture itself does not have a border
	BorderColor model.RGBA
	// Texture Transform defines the texture coordinate transformation matrix
	TextureTransform model.Matrix4F32
}

func (n *TextureEnvironment) Read(c *model.Context, version uint8) error {
	n.TextureEnvVersion = version

	if (c.Version.Equal(model.V8) && version == 2) || c.Version.GreaterEqThan(model.V9) {
		n.BorderMode = c.Data.Int32()
	}

	n.MipmapMagnificationFilter = c.Data.Int32()
	n.MipmapMinificationFilter = c.Data.Int32()
	n.SDimensionWrapMode = c.Data.Int32()
	n.TDimensionWrapMode = c.Data.Int32()

	if (c.Version.Equal(model.V8) && version == 2) || c.Version.GreaterEqThan(model.V9) {
		n.RdimentsionWrapMode = c.Data.Int32()
		n.BlendType = c.Data.Int32()
		n.InternalCompressionLevel = c.Data.Int32()
	}

	if c.Version.Equal(model.V8) && version == 1 {
		n.TextureFunctionData = c.Data.Int32()
	}

	c.Data.Unpack(&n.BlendColor)

	if (c.Version.Equal(model.V8) && version == 2) || c.Version.GreaterEqThan(model.V9) {
		c.Data.Unpack(&n.BorderColor)
	}
	n.TextureTransform = model.Matrix4F32{
	// TODO: read matrix
	}

	return c.Data.GetError()
}
