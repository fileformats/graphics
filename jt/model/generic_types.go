package model

import "fmt"

type Vector3D struct {
	X, Y, Z float32
}

func (v Vector3D) String() string {
	return fmt.Sprintf("{X:%f Y:%f Z:%f}", v.X, v.Y, v.Z)
}

// The BoundingBox type defines a bounding box using two Vector3D types to store the XYZ coordinates
// for the bounding box minimum and maximum corner points.
type BoundingBox struct {
	Min, Max Vector3D
}

func (b BoundingBox) String() string {
	return fmt.Sprintf("Min: %s Max: %s", b.Min, b.Max)
}

type HCoordF32 struct {
	X, Y, Z, W float32
}

func (v HCoordF32) String() string {
	return fmt.Sprintf("{X:%f Y:%f Z:%f W:%f}", v.X, v.Y, v.Z, v.W)
}

type HCoordF64 struct {
	X, Y, Z, W float64
}

func (v HCoordF64) String() string {
	return fmt.Sprintf("{X:%f Y:%f Z:%f W:%f}", v.X, v.Y, v.Z, v.W)
}

type String []byte

func (m *String)  Read(c *Context) error {
	var length = c.Data.Int32()
	if length == 0 {
		return nil
	}
	for i := 0; i < int(length); i++ {
		*m = append(*m, c.Data.UInt8())
	}
	return c.Data.GetError()
}

func (m String) String() string {
	return string(m)
}

type MbString []uint16

func (m *MbString) Read(c *Context) error {
	var length = c.Data.Int32()
	if length == 0 {
		return nil
	}
	for i := 0; i < int(length); i++ {
		*m = append(*m, c.Data.UInt16())
	}
	return c.Data.GetError()
}

func (m MbString) String() string {
	var tmp = "";
	for _, i := range m {
		tmp += string(i)
	}
	return tmp
}

// Defines a 4-by-4 matrix of F32 values for a total of 16 F32 values.
// The values are stored in row major order (right most subscript, column varies fastest),
// that is, the first 4 elements form the first row of the matrix
type Matrix4F32 [16]float32

// Defines a 4-by-4 matrix of F64 values for a total of 16 F64 values.
// The values are stored in row major order (right most subscript, column varies fastest),
// that is, the first 4 elements form the first row of the matrix.
type Matrix4F64 [16]float64

// The Plane type defines a geometric Plane using the General Form of the plane equation
// (Ax + By + Cz + D = 0). The PlaneF32 type is made up of four F32 base types where the
// first three F32 define the plane unit normal vector (A, B, C) and the last F32 defines
// the negated perpendicular distance (D), along normal vector, from the origin to the plane.
type Plane struct {
	A, B, C float32
	D       float32
}

// The Quaternion type defines a 3-dimensional orientation (no translation) in quaternion
// linear combination form (a + bi + cj + dk) where the four scalar values (a, b, c, d) are
// associated with the 4 dimensions of a quaternion (1 real dimension, and 3 imaginary
// dimensions).  So the Quaternion type is made up of four F32 base types.
type Quaternion struct {
	A, B, C, D float32
}

// The RGB type defines a colour composed of Red, Green, Blue components, each of which is a F32.
// So a RGB type is made up of three F32 base types.  The Red, Green, Blue colour values
// typically range from 0.0 to 1.0.
type RGB struct {
	R, G, B float32
}
func (c RGB) String() string {
	return fmt.Sprintf("RGB{R:%f, G:%f, B:%f}", c.R, c.G, c.B)
}

// The RGBA type defines a colour composed of Red, Green, Blue, Alpha components, each of which
// is a F32.  So a RGBA type is made up of four F32 base types.  The Red, Green, Blue colour values
// typically range from 0.0 to 1.0.  The Alpha value ranges from 0.0 to 1.0 where 1.0
// indicates completely opaque.
type RGBA struct {
	R, G, B, A float32
}

func (c RGBA) String() string {
	return fmt.Sprintf("RGBA{R:%f, G:%f, B:%f, A:%f}", c.R, c.G, c.B, c.A)
}

type Int32Range struct {
	Min int32
	Max int32
}

type Float32Range struct {
	Min float32
	Max float32
}

type VectorF32 []float32

func (m *VectorF32) Read(c *Context) error {
	var length = c.Data.Int32()
	if length == 0 {
		return nil
	}
	for i := 0; i < int(length); i++ {
		*m = append(*m, c.Data.Float32())
	}
	return c.Data.GetError()
}

type VectorF64 []float64

func (m *VectorF64) Read(c *Context) error {
	var length = c.Data.Int32()
	if length == 0 {
		return nil
	}
	for i := 0; i < int(length); i++ {
		*m = append(*m, c.Data.Float64())
	}
	return c.Data.GetError()
}

type VectorI16 []int16

func (m *VectorI16) Read(c *Context) error {
	var length = c.Data.Int32()
	if length == 0 {
		return nil
	}
	for i := 0; i < int(length); i++ {
		*m = append(*m, c.Data.Int16())
	}
	return c.Data.GetError()
}

type VectorU16 []uint16

func (m *VectorU16) Read(c *Context) error {
	var length = c.Data.Int32()
	if length == 0 {
		return nil
	}
	for i := 0; i < int(length); i++ {
		*m = append(*m, c.Data.UInt16())
	}
	return c.Data.GetError()
}

type VectorI32 []int32

func (m *VectorI32) Read(c *Context) error {
	var length = c.Data.Int32()
	if length == 0 {
		return nil
	}
	for i := 0; i < int(length); i++ {
		*m = append(*m, c.Data.Int32())
	}
	return c.Data.GetError()
}

type VectorU32 []uint32

func (m *VectorU32) Read(c *Context) error {
	var length = c.Data.Int32()
	if length == 0 {
		return nil
	}
	for i := 0; i < int(length); i++ {
		*m = append(*m, c.Data.UInt32())
	}
	return c.Data.GetError()
}


