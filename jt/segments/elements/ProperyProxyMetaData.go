package elements

import (
	"github.com/fileformats/graphics/jt/model"
	"time"
)

// A Property Proxy Meta Data Element  serves as a ―proxy‖ for all meta-data properties
// associated with a particular Meta Data Node Element
type PropertyProxyMetaData struct {
	JTElement
	// Version Number is the version identifier for this element
	VersionNumber uint8
	Properties map[string]interface{}
}

func (n PropertyProxyMetaData) GUID() model.GUID {
	return model.PropertyProxyMetaDataElement
}

func (n *PropertyProxyMetaData) Read(c *model.Context) error {
	c.LogGroup("PropertyProxyMetaData")
	defer c.LogGroupEnd()

	if c.Version.GreaterEqThan(model.V9) {
		if c.Version.GreaterEqThan(model.V10) {
			n.VersionNumber = c.Data.UInt8()
		} else {
			n.VersionNumber = uint8(c.Data.Int16())
		}
	}

	n.Properties = map[string]interface{}{}

	for {
		var key = model.MbString{}
		if err := (&key).Read(c); err != nil {
			return err
		}

		if key.String() == "" {
			return nil
		}

		propType := c.Data.UInt8()
		switch propType {
		case 1:
			tmp := model.MbString{}
			if err := (&tmp).Read(c); err != nil {
				return err
			}
			n.Properties[key.String()] = tmp.String()
		case 2:
			n.Properties[key.String()] = c.Data.Int32()
		case 3:
			n.Properties[key.String()] = c.Data.Float32()
		case 4:
			year := c.Data.Int16()
			month := c.Data.Int16()
			day := c.Data.Int16()
			hour := c.Data.Int16()
			minute := c.Data.Int16()
			second := c.Data.Int16()

			n.Properties[key.String()] = time.Date(
				int(year), time.Month(month), int(day),
				int(hour), int(minute), int(second),
				0,
				time.UTC,
			)
		}
	}

	return c.Data.GetError()
}

func (n *PropertyProxyMetaData) BaseElement() *JTElement {
	return &n.JTElement
}
