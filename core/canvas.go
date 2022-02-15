package core

type Canvas struct {
	Width  uint
	Height uint
	Pixels []Color
}

func NewCanvas(width, height uint) *Canvas {
	return &Canvas{width, height, make([]Color, width*height)}
}

func (c *Canvas) write(width, height uint, color Color) {
	index := c.getIndex(width, height)
	c.Pixels[index] = color
}

func (c *Canvas) read(width, height uint) Color {
	index := c.getIndex(width, height)
	return c.Pixels[index]
}

func (c *Canvas) getIndex(x, y uint) uint {
	return y*c.Width + x
}
