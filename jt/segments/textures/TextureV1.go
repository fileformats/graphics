package textures

import "github.com/fileformats/graphics/jt/model"

// Texture Vers-1 Data format is stored in JT file if the Texture Image Element is a vanilla/basic texture image
type TextureV1 struct {
	// Texture Type specifies the type of texture
	// = 0 None.
	// = 1 One-Dimensional. A one-dimensional texture has a height (T-Dimension) and depth (R-Dimension)
	//     equal to “1” and no top or bottom border.
	// = 2 Two-Dimensional. A two-dimensional texture has a depth (R-Dimension) equal to “1.”
	// = 3 Three-Dimensional. A three-dimensional texture can be thought of as layers of two-dimensional
	//     sub image rectangles arranged in a sequence.
	// = 4 Bump Map. A bump map texture is a texture where the image texel data (e.g. RGB color values) represents
	//     surface normal XYZ components.
	// = 5 Cube Map. A cube map texture is a texture cube centered at the origin and formed by a set of six
	//     two-dimensional texture images.
	// = 6 Depth Map. A depth map texture is a texture where the image texel data represents depth values.
	TextureType int32

	TextureEnvironment TextureEnvironment
	// Texture Channel specifies the texture channel number for the Texture Image Element. For purposes of multi-texturing,
	// the JT concept of a texture channel corresponds to the OpenGL concept of a “texture unit.” The Texture Channel
	// value must be between 0 and 31 inclusive
	TextureChannel uint32
	// Reserved Field is a data field reserved for future JT format expansion
	Reserved uint32
	// Inline Image Storage Flag is a flag that indicates whether the texture image is stored within the JT File
	// (i.e. inline) or in some other external file.
	// = 0 Texture image stored in an external file.
	// = 1 Texture image stored inline in this JT file
	InlineImageStorageFlag uint8
	// Image Count specifies the number of texture images
	ImageCount int32
}

func (n *TextureV1) Read(c *model.Context) error {
	if c.Version.Equal(model.V8) {

	}
	if c.Version.Equal(model.V9) {

	}
	return c.Data.GetError()
}