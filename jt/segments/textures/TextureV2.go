package textures

import "github.com/fileformats/graphics/jt/model"

type TextureV2 struct {
}

func (n *TextureV2) Read(c *model.Context) error {
	return c.Data.GetError()
}