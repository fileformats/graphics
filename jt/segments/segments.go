package segments

import (
	"compress/zlib"
	"errors"
	"io"

	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments/elements"
	"github.com/lxq/lzma"
)

type Segment interface {
	model.FileObjectReader
	ReadSegmentData(*model.Context) error
	GUID() model.GUID
}

type SegmentData struct {
	// Segment ID is the globally unique identifier for the segment.
	GUID model.GUID
	// Segment Offset defines the byte offset from the top of the file to start of the segment.
	Offset uint64
	// Segment Length is the total size of the segment in bytes.
	Length uint32
	// Segment Attributes is a collection of segment information encoded within a single U32
	// bit 0-23 - reserved for future use
	// bit 24-31 - Segment type
	Attr uint32
	// Segment Type defines a broad classification of the segment contents
	Type SegmentType
	// Compression flag
	CompressionFlag uint32
	// Compressed data length
	CompressedLength int32
	// Compression algorithm
	CompressionAlgo uint8
	// Compression indicator
	CompressionIndicator uint8
	// Element Length is the total length in bytes of the element Object Data
	ElementLength int32
	// Segment elements
	Elements []*elements.Element
	// Segment context
	Context *model.Context
}

func (s *SegmentData) ReadSegmentData(context *model.Context) error {
	context.LogGroup("Data Segment")
	defer context.LogGroupEnd()

	context.Data.Seek(int64(s.Offset), 0)

	// read segment header
	var id model.GUID
	context.Data.Unpack(&id)
	if !id.Equals(s.GUID) {
		return errors.New("Segment GUID mismatch")
	}
	context.Log("Id: %s (%s)", s.GUID, s.GUID.Name())

	typ := SegmentType(context.Data.Int32())
	if typ != s.Type {
		return errors.New("Segment type mismatch")
	}
	context.Log("Type: %d (%s)", s.Type, s.Type.Name())

	length := uint32(context.Data.Int32())
	if s.Length != length {
		return errors.New("Segment length mismatch")
	}
	context.Log("Length: %d", s.Length)

	s.Context = context.Clone()

	if !typ.Compressed() {
		return context.Data.GetError()
	}

	s.CompressionIndicator = 3
	if context.Version.GreaterEqThan(model.V10) {
		s.CompressionFlag = context.Data.UInt32()
	} else {
		s.CompressionFlag = uint32(context.Data.Int32())
		s.CompressionIndicator = 2
	}
	s.CompressedLength = context.Data.Int32() - 1 // compressionAlgo flag is also included here so we must remove it
	s.CompressionAlgo = context.Data.UInt8()

	context.Log("Compressed Flag: %d", s.CompressionFlag)
	context.Log("Compressed Length: %d", s.CompressedLength)
	context.Log("Compression Algo: %d", s.CompressionAlgo)

	if s.CompressionFlag == uint32(s.CompressionIndicator) || s.CompressionAlgo == uint8(s.CompressionIndicator) {
		//var buf = make([]byte, s.CompressedLength)
		//context.Data.Unpack(buf)
		var uncompressedReader io.Reader
		if context.Version.GreaterEqThan(model.V10) { // LZMA compression for JT files version >= 10.0
			uncompressedReader = lzma.NewReader(context.File)
		} else { // ZLIB compression for JT files version < 10.0
			uncompressedReader, _ = zlib.NewReader(context.File)
		}
		//data, _ := ioutil.ReadAll(uncompressedReader)
		s.Context.SetReader(uncompressedReader)
	}

	return context.Data.GetError()
}

