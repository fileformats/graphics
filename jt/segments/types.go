package segments

type SegmentType int32

const (
	logicalSceneGraphSegment       SegmentType = 1
	jtBRepSegment                  SegmentType = 2
	pmiDataSegment                 SegmentType = 3
	metaDataSegment                SegmentType = 4
	shapeSegment                   SegmentType = 6
	shapeLOD0Segment               SegmentType = 7
	shapeLOD1Segment               SegmentType = 8
	shapeLOD2Segment               SegmentType = 9
	shapeLOD3Segment               SegmentType = 10
	shapeLOD4Segment               SegmentType = 11
	shapeLOD5Segment               SegmentType = 12
	shapeLOD6Segment               SegmentType = 13
	shapeLOD7Segment               SegmentType = 14
	shapeLOD8Segment               SegmentType = 15
	shapeLOD9Segment               SegmentType = 16
	xtBRepSegment                  SegmentType = 17
	wireframeRepresentationSegment SegmentType = 18
	ulpSegment                     SegmentType = 20
	lwpaSegment                    SegmentType = 24
	xtbRep1Segment                 SegmentType = 30
)

func (s SegmentType) Compressed() bool {
	if int(s) < 5 || int(s) >= 17 {
		return true
	}
	return false
}

func (s SegmentType) Name() string {
	switch s {
	case 1:
		return "logicalSceneGraphSegment"
	case 2:
		return "jtBRepSegment"
	case 3:
		return "pmiDataSegment"
	case 4:
		return "metaDataSegment"
	case 6:
		return "shapeSegment"
	case 7:
		return "shapeLOD0Segment"
	case 8:
		return "shapeLOD1Segment"
	case 9:
		return "shapeLOD2Segment"
	case 10:
		return "shapeLOD3Segment"
	case 11:
		return "shapeLOD4Segment"
	case 12:
		return "shapeLOD5Segment"
	case 13:
		return "shapeLOD6Segment"
	case 14:
		return "shapeLOD7Segment"
	case 15:
		return "shapeLOD8Segment"
	case 16:
		return "shapeLOD9Segment"
	case 17:
		return "xtBRepSegment"
	case 18:
		return "wireframeRepresentationSegment"
	case 20:
		return "ulpSegment"
	case 24:
		return "lwpaSegment"
	case 30:
		return "xtbRep1Segment"
	}
	return ""
}