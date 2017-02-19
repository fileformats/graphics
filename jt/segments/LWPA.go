package segments

import (
	"github.com/fileformats/graphics/jt/model"
)

type LWPASegment struct {
	*SegmentData
}

func (l *LWPASegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("LWPASegment")
	defer context.LogGroupEnd()
	context.Log("NOT IMPLEMENTED")
	return context.Data.GetError()
}

func (l *LWPASegment) GUID() model.GUID {
	return l.SegmentData.GUID
}
