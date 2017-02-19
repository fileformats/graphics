package elements

import "github.com/fileformats/graphics/jt/model"

type PolygonSetLOD struct {
	JTElement
}

func (n PolygonSetLOD) GUID() model.GUID {
	return model.PolygonSetLODElement
}

func (n *PolygonSetLOD) Read(c *model.Context) error {
	c.LogGroup("PolygonSetLOD")
	defer c.LogGroupEnd()

	c.Log("Length: %d", n.JTElement.Length)
	c.Log("GUID: %s (%s)", n.JTElement.Id, n.JTElement.Id.Name())
	c.Log("Base Type: %d (%s)", n.JTElement.Type, n.JTElement.Type)

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *PolygonSetLOD) BaseElement() *JTElement {
	return &n.JTElement
}