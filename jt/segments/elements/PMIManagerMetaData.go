package elements

import "github.com/fileformats/graphics/jt/model"

type PMIManagerMetaData struct {
	JTElement
}

func (n PMIManagerMetaData) GUID() model.GUID {
	return model.PMIManagerMetaDataElement
}

func (n *PMIManagerMetaData) Read(c *model.Context) error {
	c.LogGroup("PMIManagerMetaData")
	defer c.LogGroupEnd()

	c.Log("Length: %d", n.JTElement.Length)
	c.Log("GUID: %s (%s)", n.JTElement.Id, n.JTElement.Id.Name())
	c.Log("Base Type: %d (%s)", n.JTElement.Type, n.JTElement.Type)

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *PMIManagerMetaData) BaseElement() *JTElement {
	return &n.JTElement
}