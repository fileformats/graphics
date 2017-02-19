package elements

import "github.com/fileformats/graphics/jt/model"

type PrimitiveSetShape struct {
	JTElement
}

func (n PrimitiveSetShape) GUID() model.GUID {
	return model.PrimitiveSetShapeElement
}

func (n *PrimitiveSetShape) Read(c *model.Context) error {
	c.LogGroup("PrimitiveSetShape")
	defer c.LogGroupEnd()

	c.Log("Length: %d", n.JTElement.Length)
	c.Log("GUID: %s (%s)", n.JTElement.Id, n.JTElement.Id.Name())
	c.Log("Base Type: %d (%s)", n.JTElement.Type, n.JTElement.Type)

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *PrimitiveSetShape) BaseElement() *JTElement {
	return &n.JTElement
}