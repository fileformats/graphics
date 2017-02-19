package segments

import (
	"github.com/fileformats/graphics/jt/model"
	"github.com/fileformats/graphics/jt/segments/elements"
)

type MetaDataSegment struct {
	*SegmentData
}

func (l *MetaDataSegment) Read(context *model.Context) error {
	l.SegmentData.ReadSegmentData(context)
	// switch to segment context
	context = l.Context

	context.LogGroup("MetaDataSegment")
	defer context.LogGroupEnd()

	if element := elements.New(context); element != nil {
		element.Read(context)
	}

	return context.Data.GetError()
}

func (l *MetaDataSegment) GUID() model.GUID {
	return l.SegmentData.GUID
}