package segments

import (
	"github.com/fileformats/graphics/jt/model"
)

type ShapeSegment struct {
	*SegmentData
	Level int
}

func (l *ShapeSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("ShapeSegment")
	defer context.LogGroupEnd()
	context.Log("NOT IMPLEMENTED")
	return context.Data.GetError()
}

func (l *ShapeSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}
