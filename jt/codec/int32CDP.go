package codec

import (
	"github.com/fileformats/graphics/jt/model"
	"fmt"
)

// Int32 Compressed Data Packet
type Int32CDP struct {
	CodecType CodecType
	// Data to decompress
	Data               []byte
	// Bits to interpret from code text bytes
	CodeTextLen        int32
	// Value element count
	ElementCount       int32
	// Symbol count
	SymbolCount        int32
	// Out of band values
	OutOfBandValues    []int
	// Probability contexts
	ProbabilityContext *ProbContext
	codec codec
}

func (i *Int32CDP) ReadVecI32(c *model.Context) (decoded []int32, err error) {
	i.ProbabilityContext = &ProbContext{}
	i.CodecType = CodecType(c.Data.UInt8())
	c.Log("CodecType: %d", i.CodecType)

	if i.CodecType < 0 || i.CodecType > 5 {
		return []int32{}, fmt.Errorf("Invalid codec type: %d", i.CodecType)
	}
	switch i.CodecType {
	case Null:
		i.codec = NullCodec{}
	case Bitlength:
		i.codec = BitlengtCodec{}
	case Chopper:
		i.codec = ChopperCodec{}
	case Arithmetic:
		i.codec = ArithmeticCodec{}
	case Huffman:
		i.codec = HuffmanCodec{}
	case MoveToFront:
		i.codec = MoveToFrontCodec{}
	}

	if i.CodecType == Huffman || i.CodecType == Arithmetic {
		i.ProbabilityContext.read(c, i.codec)
	}

	if i.CodecType != Null {
		i.CodeTextLen = c.Data.Int32()
		c.Log("CodeTextLength: %d", i.CodeTextLen)
		i.ElementCount = c.Data.Int32()
		c.Log("ElementCount: %d", i.ElementCount)
	}

	if i.ProbabilityContext.TableCount == 1 || i.ProbabilityContext.TableCount == 0 {
		i.SymbolCount = i.ElementCount
	} else if i.ProbabilityContext.TableCount == 2 {
		i.SymbolCount = c.Data.Int32()
	}
	c.Log("SymbolCount: %d", i.SymbolCount)

	residuals, err := i.codec.Decode(c, i)
	return residuals, err
}

func (i *Int32CDP) ReadVecUI32(c *model.Context) (decoded []uint32, err error) {
	data, err := i.ReadVecI32(c)
	if err != nil {
		return
	}
	for _, val := range data {
		decoded = append(decoded, uint32(val) & 0xFFFFFFFF)
	}
	return
}