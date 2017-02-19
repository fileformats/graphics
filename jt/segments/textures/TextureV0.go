package textures

import "github.com/fileformats/graphics/jt/model"

// Texture Vers-0 format is stored in JT file if the Texture Image Element is a vanilla/basic texture image
type TextureV0 struct {
	// Number of Bytes specifies the length, in bytes, of the on-disk representation of the texture image.
	// The texture image in a JT file is a single monolithic/contiguous block of data beginning with the
	// highest-level mip image, and processing through the mipmaps down to a one-by-one texel image
	NumberOfBytes int32
	// Image Format description
	ImageFormatDescription ImageV0
	// Image Texel Data is the single monolithic/contiguous block of image data.
	// The length of this field in bytes is specified by the value of data field Number of Bytes
	ImageTexelData []byte
	// Texture environment details
	TextureEnvironment TextureEnvironment
}

func (n TextureV0) Read(c *model.Context) error {
	n.NumberOfBytes = c.Data.Int32()
	if n.NumberOfBytes == 0 {
		return c.Data.GetError()
	}
	c.Data.Unpack(&n.ImageFormatDescription)

	n.ImageTexelData = make([]byte, int(n.NumberOfBytes))
	c.Data.Unpack(n.ImageTexelData)

	return (&n.TextureEnvironment).Read(c, 1)
}

