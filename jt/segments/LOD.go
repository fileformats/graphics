package segments

import (
	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments/elements"
	"fmt"
)

// A Shape LOD Element is the holder/container of the geometric shape definition data (e.g. vertices, polygons, normals, etc.)
// for a single LOD. Much of the heavyweight data contained within a Shape LOD Element may be optionally compressed and/or encoded.
// The compression and/or encoding state is indicated through other data stored in each Shape LOD Element
type ShapeLODSegment struct {
	*SegmentData
	Level   int
	Element elements.Element
}

func (l *ShapeLODSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("ShapeLODSegment")
	defer context.LogGroupEnd()

	if element := elements.New(context); element != nil {
		element.Read(context)
		l.Element = element
	}
	return context.Data.GetError()
}

func (l *ShapeLODSegment) GetGeometry() (*elements.TriStripSetShapeLOD, error) {
	tri, ok := l.Element.(*elements.TriStripSetShapeLOD)
	if !ok {
		return nil, fmt.Errorf("Invalid Shape LOD element")
	}
	return tri, nil
}

func (l *ShapeLODSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}
