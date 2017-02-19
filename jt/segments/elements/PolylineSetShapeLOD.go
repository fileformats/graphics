package elements

import "github.com/fileformats/graphics/jt/model"

type PolylineSetShapeLOD struct {
	JTElement
}

func (n PolylineSetShapeLOD) GUID() model.GUID {
	return model.PolylineSetShapeLODElement
}

func (n *PolylineSetShapeLOD) Read(c *model.Context) error {
	c.LogGroup("PolylineSetShapeLOD")
	defer c.LogGroupEnd()

	c.Log("Length: %d", n.JTElement.Length)
	c.Log("GUID: %s (%s)", n.JTElement.Id, n.JTElement.Id.Name())
	c.Log("Base Type: %d (%s)", n.JTElement.Type, n.JTElement.Type)

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *PolylineSetShapeLOD) BaseElement() *JTElement {
	return &n.JTElement
}