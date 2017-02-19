package elements

import "github.com/fileformats/graphics/jt/model"

type TextureCoordinateGeneratorAttribute struct {
	BaseAttribute
}

func (n TextureCoordinateGeneratorAttribute) GUID() model.GUID {
	return model.TextureCoordinateGeneratorAttributeElement
}

func (n *TextureCoordinateGeneratorAttribute) Read(c *model.Context) error {
	c.LogGroup("TextureCoordinateGeneratorAttribute")
	defer c.LogGroupEnd()

	c.Log("Length: %d", n.JTElement.Length)
	c.Log("GUID: %s (%s)", n.JTElement.Id, n.JTElement.Id.Name())
	c.Log("Base Type: %d (%s)", n.JTElement.Type, n.JTElement.Type)

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *TextureCoordinateGeneratorAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *TextureCoordinateGeneratorAttribute) BaseElement() *JTElement {
	return &n.JTElement
}