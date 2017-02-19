package textures

// The Vers-0 Image Format Description is a collection of data defining the pixel format, data type, size, and other
// miscellaneous characteristics of the monolithic block of image data
type ImageV0 struct {
	// Pixel format specifies the format of the texture image pixel data. Depending on the format, anywhere from one
	// to four elements of data exists per texel.
	// = 0 − No format specified.  Texture mapping is not applied.
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
	// (e.g. a Pixel Format Type of “1” with Pixel Data Type of “7” would require 12 bytes of storage for each texel).
	// = 0 − No type specified. Texture mapping is not applied.
	// = 1 − Signed 8-bit integer
	// = 2 − Single-precision 32-bit floating point
	// = 3 − Unsigned 8-bit integer
	// = 4 − Single bits in unsigned 8-bit integers
	// = 5 − Unsigned 16-bit integer
	// = 6 − Signed 16-bit integer
	// = 7 − Unsigned 32-bit integer
	// = 8 − Signed 32-bit integer
	// = 9 − 16-bit floating point according to IEEE-754 format (i.e. 1 sign bit, 5 exponent bits, 10 mantissa bits)
	PixelDataType uint32
	// Dimensionality specifies the number of dimensions the texture image has.  Valid values include:
	// = 1 − One-dimensional texture
	// = 2 − Two-dimensional texture
	Dimensionality uint16
	// Width specifies the width dimension (number of texel columns) of the texture image in number of pixels
	Width int32
	// Height specifies the height dimension (number of texel rows) of the texture image in number of pixels.
	// Height is “1” for one-dimensional images.
	Height int32
	// Mipmaps Flag is a flag indicating whether the texture image has mipmaps.
	// = 0 − No mipmaps
	// = 1 − Yes has mipmaps. Image Texel Data is assumed to contain multiple textures, each a mipmap of the base texture
	MipmapsFlag uint32
	// Shared Image Flag is a flag indicating whether this texture image is shareable with other Texture Image Element attributes.
	// = 0 − Image is not shareable with other Texture Image Elements.
	// = 1 − Image is shareable with other Texture Image Elements.
	SharedImageFlag uint32
}
