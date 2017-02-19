package elements

import "github.com/fileformats/graphics/jt/model"

type EndOfElements struct {
	JTElement
}

func (n EndOfElements) GUID() model.GUID {
	return model.EndOfElements
}

func (n *EndOfElements) Read(c *model.Context) error {
	c.LogGroup("EndOfElements")
	defer c.LogGroupEnd()

	return c.Data.GetError()
}
func (n *EndOfElements) BaseElement() *JTElement {
	return &n.JTElement
}