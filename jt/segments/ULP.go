package segments

import (
	"github.com/fileformats/graphics/jt/model"
)

type ULPSegment struct {
	*SegmentData
}

func (l *ULPSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("ULPSegment")
	defer context.LogGroupEnd()
	context.Log("NOT IMPLEMENTED")
	return context.Data.GetError()
}

func (l *ULPSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}