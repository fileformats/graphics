package codec

import (
	"math"
	"github.com/fileformats/graphics/jt/model"
)

type DeeringCodec struct {
	lookupTable *deeringLookupTable
	numBits float64
}

func NewDeeringCodec(numBits int) *DeeringCodec {
	return &DeeringCodec{
		numBits: float64(numBits),
		lookupTable:  newDeeringLookupTable(),
	}
}

type deeringCode struct {
	sextant int64
	octant int64
	theta int64
	psi int64
}

type deeringLookupTable struct {
	nBits float64
	cosTheta []float64
	sinTheta []float64
	cosPsi []float64
	sinPsi []float64
}

func (c *DeeringCodec) ToVector3D(sextant, octant, theta, psi uint32) model.Vector3D {
	if c.lookupTable == nil {
		c.lookupTable = newDeeringLookupTable()
	}
	if c.numBits == 0 {
		c.numBits = 6
	}
	theta += sextant & 1
	cosTheta, sinTheta, cosPsi, sinPsi := c.lookupTable.lookupThetaPsi(float64(theta), float64(psi), c.numBits)

	vector := model.Vector3D{
		X: float32(cosTheta * cosPsi),
		Y: float32(sinPsi),
		Z: float32(sinTheta * cosPsi),
	}

	switch sextant {
	case 0:
	case 1:
		vector.Z, vector.X = vector.X, vector.Z
	case 2:
		vector.Z, vector.X, vector.Y = vector.X, vector.Y, vector.Z
	case 3:
		vector.Y, vector.X = vector.X, vector.Y
	case 4:
		vector.Y, vector.Z, vector.X = vector.X, vector.Y, vector.Z
	case 5:
		vector.Z, vector.Y = vector.Y, vector.Z
	}

	if octant & 0x4 == 0 {
		vector.X = -vector.X
	}
	if octant & 0x2 == 0 {
		vector.Y = -vector.Y
	}
	if octant & 0x1 == 0 {
		vector.Z = -vector.Z
	}
	return vector
}

func (tbl *deeringLookupTable) lookupThetaPsi(theta, psi, count float64) (cosTheta float64, sinTheta float64, cosPsi float64, sinPsi float64) {
	offset := uint(tbl.nBits - count)
	offTheta := (int(theta) << offset) & 0xFFFFFFFF
	offPsi := (int(psi) << offset) & 0xFFFFFFFF

	return tbl.cosTheta[offTheta], tbl.sinTheta[offTheta], tbl.cosPsi[offPsi], tbl.sinPsi[offPsi]
}

func newDeeringLookupTable() *deeringLookupTable {
	tbl := &deeringLookupTable{
		nBits: 8,
		cosTheta: []float64{},
		sinTheta: []float64{},
		cosPsi: []float64{},
		sinPsi: []float64{},
	}
	var tblSize float64 = 256
	psiMax := 0.615479709

	for i := 0; i <= int(tblSize); i++ {
		theta := math.Asin(math.Tan(psiMax * (tblSize - float64(i)) / tblSize))
		psi := psiMax * (float64(i) / tblSize)
		tbl.cosTheta = append(tbl.cosTheta, math.Cos(theta))
		tbl.sinTheta = append(tbl.sinTheta, math.Sin(theta))
		tbl.cosPsi = append(tbl.cosPsi, math.Cos(psi))
		tbl.sinPsi = append(tbl.sinPsi, math.Sin(psi))
	}
	return tbl
}