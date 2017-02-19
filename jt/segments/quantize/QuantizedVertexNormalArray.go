package quantize

import (
	"github.com/fileformats/graphics/jt/codec"
	"github.com/fileformats/graphics/jt/model"
	"fmt"
)

// The Quantized Vertex Normal Array data collection contains the quantization data/representation for a set of
// vertex normals. Quantized Vertex Normal Array data collection is only present if previously read Normal Binding
// value is not equal to zero
type QuantizedVertexNormalArray struct {
	// Number of Bits specifies the quantized size (i.e. the number of bits of precision) for the Theta and PSI angles.
	// This value must satisfy the following condition:  “0 <= Number of Bits <= 13”
	NumberOfBits uint8
	// Normal Count specifies the count (number of unique) Normal Codes
	NormalCount int32
	// Sextant Codes is a vector of “codes” (one per normal) for a set of normals identifying which Sextant of the
	// corresponding sphere Octant each normal is located in
	SextantCodes []int32
	// Octant Codes is a vector of “codes” (one per normal) for a set of normals identifying which sphere Octant
	// each normal is located in
	OctantCodes []int32
	// Theta Codes is a vector of “codes” (one per normal) for a set of normals representing in Sextant coordinates
	// the quantized theta angle for each normal’s location on the unit radius sphere; where theta angle is defined
	// as the angle in spherical coordinates about the Y-axis on a unit radius sphere
	ThetaCodes []int32
	// Psi Codes is a vector of “codes” (one per normal) for a set of normals representing in Sextant coordinates the
	// quantized Psi angle for each normal’s location on the unit radius sphere; where Psi angle is defined as the
	// longitudinal angle in spherical coordinates from the y = 0 plane on the unit radius sphere
	PsiCodes []int32

	Normals []model.Vector3D
}

func (n *QuantizedVertexNormalArray) Read(c *model.Context) (err error) {
	c.Log("QuantizedVertexNormalArray")
	n.NumberOfBits = c.Data.UInt8()
	c.Log("NumberOfBits: %d", n.NumberOfBits)
	n.NormalCount = c.Data.Int32()
	c.Log("NormalCount: %d", n.NormalCount)

	if n.SextantCodes, err = (&codec.Int32CDP{}).ReadVecI32(c); err != nil {
		return
	}
	codec.UnpackResidual(n.SextantCodes, codec.Lag1)
	c.Log("SextantCodes: %v", n.SextantCodes)

	if n.OctantCodes, err = (&codec.Int32CDP{}).ReadVecI32(c); err != nil {
		return
	}
	codec.UnpackResidual(n.OctantCodes, codec.Lag1)
	c.Log("OctantCodes: %v", n.OctantCodes)

	if n.ThetaCodes , err = (&codec.Int32CDP{}).ReadVecI32(c); err != nil {
		return
	}
	codec.UnpackResidual(n.ThetaCodes , codec.Lag1)
	c.Log("ThetaCodes: %v", n.ThetaCodes)

	if n.PsiCodes, err = (&codec.Int32CDP{}).ReadVecI32(c); err != nil {
		return
	}
	codec.UnpackResidual(n.PsiCodes, codec.Lag1)
	c.Log("PsiCodes: %v", n.PsiCodes)

	codec := codec.NewDeeringCodec(int(n.NumberOfBits))

	n.Normals = make([]model.Vector3D, 0)
	for i := 0; i < len(n.PsiCodes); i++ {
		vect := codec.ToVector3D(
			uint32(n.SextantCodes[i]) & 0xFFFFFFFF,
			uint32(n.OctantCodes[i]) & 0xFFFFFFFF,
			uint32(n.ThetaCodes[i]) & 0xFFFFFFFF,
			uint32(n.PsiCodes[i]) & 0xFFFFFFFF,
		)
		n.Normals = append(n.Normals, vect)
		fmt.Printf(" %+v ", vect)
	}

	return c.Data.GetError()
}
