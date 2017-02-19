package quantize

import "github.com/fileformats/graphics/jt/model"

// The Uniform Quantizer Data collection contains information that defines a scalar quantizer/dequantizer
// (encoder/decoder) whose range is divided into levels of equal spacing.
type UniformQuantizer struct {
	// Min specifies the minimum of the quantized range
	Min float64
	// Max specifies the maximum of the quantized range
	Max float64
	// Number of Bits specifies the quantized size (i.e. the number of bits of precision).
	// In general, this value must satisfy the following condition: “0 <= Number Of Bits <= 32”
	NumberOfBits uint8
}

func (n *UniformQuantizer) Read(c *model.Context) error {
	c.Log("UniformQuantizer")
	n.Min = float64(c.Data.Float32())
	c.Log("Min: %f", n.Min)
	n.Max = float64(c.Data.Float32())
	c.Log("Max: %f", n.Max)
	n.NumberOfBits = c.Data.UInt8()
	c.Log("NumberOfBits: %d", n.NumberOfBits)
	return c.Data.GetError()
}

func (n *UniformQuantizer) Dequantize(values []int32) []float32 {
	res := []float32{}
	max := 0xffffffff
	if n.NumberOfBits < 32 {
		max = 1 << n.NumberOfBits
	}
	encode := float64(max) / (n.Max - n.Min)
	for _, val := range values {
		res = append(res, float32((float64(val) - 0.5) / encode + n.Min))
	}
	return res
}