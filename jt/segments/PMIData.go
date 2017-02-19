package segments

import (
	"github.com/fileformats/graphics/jt/model"
)

type PMIDataSegment struct {
	*SegmentData
}

func (l *PMIDataSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("PMIDataSegment")
	defer context.LogGroupEnd()
	context.Log("NOT IMPLEMENTED")
	return context.Data.GetError()
}

func (l *PMIDataSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}