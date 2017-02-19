package textures

import "github.com/fileformats/graphics/jt/model"

type TextureV3 struct {
}

func (n *TextureV3) Read(c *model.Context) error {
	return c.Data.GetError()
}
