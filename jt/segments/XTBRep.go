package segments

import (
	"github.com/fileformats/graphics/jt/model"
)

type XTBrepSegment struct {
	*SegmentData
}

func (l *XTBrepSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("XTBrepSegment")
	defer context.LogGroupEnd()
	context.Log("NOT IMPLEMENTED")
	return context.Data.GetError()
}

func (l *XTBrepSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}
