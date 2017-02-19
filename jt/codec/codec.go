package codec

import (
	"encoding/binary"

	"github.com/fileformats/graphics/jt/model"
)

type CodecType uint8

const (
	Null        CodecType = 0
	Bitlength   CodecType = 1
	Huffman     CodecType = 2
	Arithmetic  CodecType = 3
	Chopper     CodecType = 4
	MoveToFront CodecType = 5
)

type codec interface {
	Decode(context *model.Context, cdp *Int32CDP) ([]int32, error)
}

func getCodeText(c *model.Context) []byte {
	count := c.Data.Int32()
	codeText := make([]byte, count*4)
	c.Data.Unpack(&codeText)

	// must reorder
	if c.ByteOrder == binary.LittleEndian {
		for i := 0; i < int(count); i++ {
			codeText[0+i*4], codeText[3+i*4] = codeText[3+i*4], codeText[0+i*4]
			codeText[1+i*4], codeText[2+i*4] = codeText[2+i*4], codeText[1+i*4]
		}
	}

	return codeText
}
