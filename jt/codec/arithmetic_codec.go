package codec

import (
	"github.com/fileformats/graphics/jt/model"
)

// Arithmetic encoding is a lossless compression algorithm that replaces an input stream of symbols
// or bytes with a single fixed point output number. The total number of bits needed in the output number is dependent
// upon the length/complexity of the input message.  This single fixed point number output from an arithmetic encoding
// process must be uniquely decodable to create the exact stream of input symbols that were used to create it
type ArithmeticCodec struct {
}

func (n ArithmeticCodec) Decode(c *model.Context, cdp *Int32CDP) ([]int32, error) {
	if c.Version.Equal(model.V8) {

	}
	return []int32{}, c.Data.GetError()
}




