package segments

import (
	"github.com/fileformats/graphics/jt/model"
)

type WireframeSegment struct {
	*SegmentData
}

func (l *WireframeSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("WireframeSegment")
	defer context.LogGroupEnd()
	context.Log("NOT IMPLEMENTED")
	return context.Data.GetError()
}

func (l *WireframeSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}
