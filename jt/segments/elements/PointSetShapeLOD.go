package elements

import "github.com/fileformats/graphics/jt/model"

type PointSetShapeLOD struct {
	JTElement
}

func (n PointSetShapeLOD) GUID() model.GUID {
	return model.PointSetShapeLODElement
}

func (n *PointSetShapeLOD) Read(c *model.Context) error {
	c.LogGroup("PointSetShapeLOD")
	defer c.LogGroupEnd()

	c.Log("Length: %d", n.JTElement.Length)
	c.Log("GUID: %s (%s)", n.JTElement.Id, n.JTElement.Id.Name())
	c.Log("Base Type: %d (%s)", n.JTElement.Type, n.JTElement.Type)

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *PointSetShapeLOD) BaseElement() *JTElement {
	return &n.JTElement
}