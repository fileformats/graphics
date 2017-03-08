package jt

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/cthackers/go/version"
	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments"
)

type LevelOfDetail int

const (
	LOD0 LevelOfDetail = iota
	LOD1
	LOD2
	LOD3
	LOD4
	LOD5
	LOD6
	LOD7
	LOD8
	LOD9
)

// Jupiter Tesselation file
// Version 10.0 Rev-B
// https://www.plm.automation.siemens.com/en_us/Images/JT-v10-file-format-reference-rev-B_tcm1023-233786.pdf
type JTFile struct {
	// JT file size
	FileSize uint64
	// Path of root JT file
	FilePath string
	// JT File Header
	Header JTHeader
	// List of file segments
	Segments []segments.Segment
}

type JtJSON struct {
	BBox model.BoundingBox `json:"bbox"`
	// Area is the total surface area for this node and all of its descendents.
	Area float32 `json:"area"`
	// JT File version
	Version string `json:"version,omitempty"`
	// JT Internal filename
	Filename string `json:"filename,omitempty"`
	// Properties table
	Properties map[string]interface{} `json:"properties,omitempty"`
	// Level of detail for this geometry
	LevelOfDetail LevelOfDetail `json:"levelOfDetail"`
	// Available Level of Detail
	AvailableLOD []LevelOfDetail `json:"availableLOD"`
	// Geometry data
	Vertices []model.Vector3D `json:"vertices,omitempty"`
	Normals  []model.Vector3D `json:"normals,omitempty"`
	Faces    []int32          `json:"faces,omitempty"`
	Colors   []string         `json:"colors,omitempty"`
	// Child files
	Children []string `json:"children,omitempty"`
}

func (jt *JTFile) GetLSG() *segments.LSGSegment {
	for _, s := range jt.Segments {
		if lsg, ok := s.(*segments.LSGSegment); ok {
			return lsg
		}
	}
	return nil
}

func (jt *JTFile) GetLOD(lod LevelOfDetail) *segments.ShapeLODSegment {
	var lodSegment *segments.ShapeLODSegment

	for _, s := range jt.Segments {
		if seg, ok := s.(*segments.ShapeLODSegment); ok {
			if seg.Level <= int(lod) {
				lodSegment = seg
			}
		}
	}

	return lodSegment
}

func (jt *JTFile) AvailableLOD() []LevelOfDetail {
	var res = []LevelOfDetail{}
	for _, s := range jt.Segments {
		if lod, ok := s.(*segments.ShapeLODSegment); ok {
			res = append(res, LevelOfDetail(lod.Level))
		}
	}
	sort.Slice(res, func(a, b int) bool {
		if res[a] < res[b] {
			return true
		}
		return false
	})

	return res
}

func (jt *JTFile) InternalFilename() string {
	lsg := jt.GetLSG()
	if lsg == nil {
		return ""
	}
	return lsg.PartitionNode.FileName.String()
}

func (jt *JTFile) JSONStruct(lod LevelOfDetail) (*JtJSON, error) {
	js := &JtJSON{
		Version:      jt.Header.Version.String(),
		AvailableLOD: jt.AvailableLOD(),
	}
	lsg := jt.GetLSG()
	if lsg == nil {
		return nil, fmt.Errorf("Could not read JT file. No LSG Segment found")
	}
	if lsg.PartitionNode == nil {
		return nil, fmt.Errorf("Could not read JT file. No PartitionNode found")
	}
	js.BBox = lsg.PartitionNode.UntransformedBBox
	js.Area = lsg.PartitionNode.Area
	js.Filename = lsg.PartitionNode.FileName.String()

	lodSegment := jt.GetLOD(lod)
	if lodSegment == nil {
		return nil, fmt.Errorf("Could not read JT file. No LOD Segment found")
	}
	js.LevelOfDetail = LevelOfDetail(lodSegment.Level)

	geo, err := lodSegment.GetGeometry()
	if err != nil {
		return nil, err
	}
	js.Normals = geo.Normals
	js.Vertices = geo.Vertex
	js.Faces = geo.Faces

	for _, prop := range lsg.Properties {
		_ = prop
		//	fmt.Println(prop.Value())
	}
	return js, nil
}

// ToJSON returns file geometry with the given level of detail (0 - 9).
// If the requested level of detail does not exist, the previous lower value LOD will be returned
func (jt *JTFile) ToJSON(lod LevelOfDetail) ([]byte, error) {
	js, err := jt.JSONStruct(lod)
	if err != nil {
		return nil, err
	}
	return json.Marshal(js)
}

type JTHeader struct {
	// File version
	Version *version.Version
	// File comment
	Comment string
	// File byte order
	ByteOrder binary.ByteOrder
	// Reserved field
	Reserved int32
	// TOC offset defines the byte offset from the top of the file to the start of the TOC Segment
	TOCOffset uint64
	// LSG Segment ID specifies the globally unique identifier for the Logical Scene Graph Data Segment in the file
	LSGSegmentId model.GUID
}

func (h *JTHeader) Read(context *model.Context) error {
	context.LogGroup("JTHeader")
	defer context.LogGroupEnd()

	h.ByteOrder = binary.BigEndian

	// read file version
	var buf = make([]byte, 80)
	context.Data.Unpack(buf)

	match := versionPattern.FindSubmatch(buf)
	if match == nil {
		return errors.New("Invalid file signature")
	}
	var err error
	if h.Version, err = version.New(string(match[1])); err != nil {
		return errors.New("Invalid file version")
	}
	context.Version = h.Version
	context.Log(" Version: %s", h.Version)

	h.Comment = string(bytes.TrimSpace(match[2]))
	context.Log(" Comment: %s", h.Comment)

	// read byte order
	if context.Data.Byte() == 0 {
		context.Data.ByteOrder(binary.LittleEndian)
		context.ByteOrder = binary.LittleEndian
		h.ByteOrder = binary.LittleEndian
	}

	context.Log(" ByteOrder: %s", h.ByteOrder)

	// empty field. File attributes for version ^8.0
	h.Reserved = context.Data.Int32()
	context.Log(" Reserved: %d", h.Reserved)

	// TOC Offset
	switch {
	case h.Version.Equal(model.V8) || h.Version.Equal(model.V9):
		h.TOCOffset = uint64(context.Data.Int32())

	case h.Version.Equal(model.V10):
		h.TOCOffset = context.Data.UInt64()
	}
	context.Log(" TOC Offset: %d", h.TOCOffset)

	// LSG Segment ID
	context.Data.Unpack(&h.LSGSegmentId)
	if h.Reserved != 0 {
		h.LSGSegmentId = model.GUID{}
	}
	context.Log(" LSG Segment Id: %s (%s)", h.LSGSegmentId, h.LSGSegmentId.Name())

	return nil
}

func readTOCEntry(context *model.Context) (segments.Segment, error) {
	context.LogGroup("TOC Entry")
	defer context.LogGroupEnd()

	data := &segments.SegmentData{}

	context.Data.Unpack(&data.GUID)
	context.Log("Id: %s (%s)", data.GUID, data.GUID.Name())

	switch {
	case context.Version.Equal(model.V8) || context.Version.Equal(model.V9):
		data.Offset = uint64(context.Data.Int32())
		data.Length = uint32(context.Data.Int32())

	case context.Version.Equal(model.V10):
		data.Offset = context.Data.UInt64()
		data.Length = context.Data.UInt32()
	}
	context.Log("Offset: %d", data.Offset)
	context.Log("Length: %d", data.Length)

	data.Attr = context.Data.UInt32()
	context.Log("Attr: %d", data.Attr)
	data.Type = segments.SegmentType(int(data.Attr) >> 24)
	context.Log("Type: %d (%s)", data.Type, data.Type.Name())

	var segment segments.Segment

	switch int(data.Type) {
	case 1:
		segment = &segments.LSGSegment{
			SegmentData: data,
		}
	case 2:
		segment = &segments.JTBRepSegment{
			SegmentData: data,
		}
	case 3:
		segment = &segments.PMIDataSegment{
			SegmentData: data,
		}
	case 4:
		segment = &segments.MetaDataSegment{
			SegmentData: data,
		}
	case 6:
		segment = &segments.ShapeLODSegment{
			SegmentData: data,
			Level:       -1,
		}
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		fallthrough
	case 10:
		fallthrough
	case 11:
		fallthrough
	case 12:
		fallthrough
	case 13:
		fallthrough
	case 14:
		fallthrough
	case 15:
		fallthrough
	case 16:
		segment = &segments.ShapeLODSegment{
			SegmentData: data,
			Level:       int(data.Type) - 7,
		}
	case 17:
		segment = &segments.XTBrepSegment{
			SegmentData: data,
		}
	case 18:
		segment = &segments.WireframeSegment{
			SegmentData: data,
		}
	case 20:
		segment = &segments.ULPSegment{
			SegmentData: data,
		}
	case 24:
		segment = &segments.LWPASegment{
			SegmentData: data,
		}
	case 30:
		segment = &segments.XTBrepSegment{
			SegmentData: data,
		}
	default:
		return nil, fmt.Errorf("Unknown segment type %d", data.Type)
	}

	return segment, context.Data.GetError()
}
