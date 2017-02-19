package elements

import "github.com/fileformats/graphics/jt/model"

// Texture Image Attribute Element defines a texture image and its mapping environment.  JT format LSG traversal
// semantics dictate that texture image attributes accumulate down the LSG by replacement
type TextureImageAttribute struct {
	BaseAttribute
	// Version Number is the version identifier for this element
	VersionNumber uint8
}

func (n TextureImageAttribute) GUID() model.GUID {
	return model.TextureImageAttributeElement
}

func (n *TextureImageAttribute) Read(c *model.Context) error {
	c.LogGroup("TextureImageAttribute")
	defer c.LogGroupEnd()

	if err := (&n.BaseAttribute).Read(c); err != nil {
		return err
	}

	c.Data.Skip(int(n.JTElement.Length))

	return c.Data.GetError()
}

func (n *TextureImageAttribute) GetBaseAttribute() *BaseAttribute {
	return &n.BaseAttribute
}

func (n *TextureImageAttribute) BaseElement() *JTElement {
	return &n.JTElement
}