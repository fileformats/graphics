package elements

import "github.com/fileformats/graphics/jt/model"

type PointstyleAttribute struct {
	JTElement
}

func (n PointstyleAttribute) GUID() model.GUID {
	return model.PointstyleAttributeElement
}

func (n *PointstyleAttribute) Read(c *model.Context) error {
	c.LogGroup("PointstyleAttribute")
	defer c.LogGroupEnd()

	c.Log("Length: %d", n.JTElement.Length)
	c.Log("GUID: %s (%s)", n.JTElement.Id, n.JTElement.Id.Name())
	c.Log("Base Type: %d (%s)", n.JTElement.Type, n.JTElement.Type)

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *PointstyleAttribute) BaseElement() *JTElement {
	return &n.JTElement
}