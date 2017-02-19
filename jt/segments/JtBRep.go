package segments

import (
	"github.com/fileformats/graphics/jt/model"
)

type JTBRepSegment struct {
	*SegmentData
}

func (l *JTBRepSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("JTBRepSegment")
	defer context.LogGroupEnd()
	context.Log("NOT IMPLEMENTED")
	return context.Data.GetError()
}

func (l *JTBRepSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}
