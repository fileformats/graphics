package textures

// The Vers-2 Image Format Description is a collection of data defining the pixel format, data type, size, and
// other miscellaneous characteristics of the texel image data
type ImageV2 struct {
	// Pixel format specifies the format of the texture image pixel data. Depending on the format, anywhere from one
	// to four elements of data exists per texel.
	//  = 0 − No format specified.  Texture mapping is not applied.
	// = 1 − A red color component followed by green and blue color components
    // = 2 − A red color component followed by green, blue, and alpha color components
	// = 3 − A single luminance component
	// = 4 − A luminance component followed by an alpha color component.
	// = 5 − A single stencil index.
	// = 6 − A single depth component
	// = 7 − A single red color component
	// = 8 − A single green color component
	// = 9 − A single blue color component
	// = 10 − A single alpha color component
	// = 11 − A blue color component, followed by green and red color components
	// = 12 − A blue color component, followed by green, red, and alpha color components
	PixelFormat uint32
	// Pixel Data Type specifies the data type used to store the per texel data. If the Pixel Format represents a multi
	// component value (e.g. red, green, blue) then each value requires the Pixel Data Type number of bytes of storage
	// (e.g. a Pixel Format Type of “1” with Pixel Data Type of “3” would require 3 bytes of storage for each texel).
	// = 3 − Unsigned 8-bit integer
    PixelDataType uint32
	// Dimensionality specifies the number of dimensions the texture image has.  Valid values include:
	// = 1 − One-dimensional texture
	// = 2 − Two-dimensional texture
	Dimensionality int16
	// Row Alignment specifies the byte alignment for image data rows.  This data field must have a value of 1, 2, 4, or 8
	// BytesPerRow = (numBytesPerPixel * ImageWidth + RowAlignmnet – 1)  &  ~(RowAlignment – 1)
	RowAlignment int16
	// Width specifies the width dimension (number of texel columns) of the texture image in number of pixels
	Width int16
	// Height specifies the height dimension (number of texel rows) of the texture image in number of pixels.
	// Height is “1” for one-dimensional images
	Height int16
	// Depth specifies the depth dimension (number of texel slices) of the texture image in number of pixels. Depth
	// is “1” for one-dimensional and two-dimensional images
	Depth int16
	// Number Border Texels specifies the number of border texels in the texture image definition. Valid values are 0 or 1
	NumberBorderTexels int16
	// Shared Image Flag is a flag indicating whether this texture image is shareable with other Texture Image Element attributes.
	// = 0 − Image is not shareable with other Texture Image Elements.
	// = 1 − Image is shareable with other Texture Image Elements
	ShareImageFlag uint32
	// Mipmaps Count specifies the number of mipmap images.  A value of “1” indicates that no mipmaps are used.
	// A value greater than “1” indicates that mipmaps are present all the way down to a 1-by-1 texel
	MipmapsCount int16
}
