package elements

import "github.com/fileformats/graphics/jt/model"

type WireHarnessSetShape struct {
	JTElement
}

func (n WireHarnessSetShape) GUID() model.GUID {
	return model.WireHarnessSetShapeElement
}

func (n *WireHarnessSetShape) Read(c *model.Context) error {
	c.LogGroup("WireHarnessSetShape")
	defer c.LogGroupEnd()

	c.Log("Length: %d", n.JTElement.Length)
	c.Log("GUID: %s (%s)", n.JTElement.Id, n.JTElement.Id.Name())
	c.Log("Base Type: %d (%s)", n.JTElement.Type, n.JTElement.Type)

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *WireHarnessSetShape) BaseElement() *JTElement {
	return &n.JTElement
}