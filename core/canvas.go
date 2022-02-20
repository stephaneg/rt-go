package core

import (
	"fmt"
	"strings"
)

type Canvas struct {
	Width  uint
	Height uint
	Pixels []Color
}

func NewCanvas(width, height uint) *Canvas {
	return &Canvas{width, height, make([]Color, width*height)}
}

func (c *Canvas) Write(width, height uint, color Color) {
	index := c.getIndex(width, height)
	c.Pixels[index] = color
}

func (c *Canvas) Read(width, height uint) Color {
	index := c.getIndex(width, height)
	return c.Pixels[index]
}

func (c *Canvas) getIndex(x, y uint) uint {
	return y*c.Width + x
}

func (c *Canvas) ToPPM() string {
	var b strings.Builder
	b.Grow(32)

	fmt.Fprintf(&b, "P3\n")
	fmt.Fprintf(&b, "%d %d\n", c.Width, c.Height)
	fmt.Fprintf(&b, "%d\n", 255)

	var current_col int = 0
	var row_size int = 0

	for _, p := range c.Pixels {
		colors := p.ToRgb()

		for _, color := range colors {
			str := fmt.Sprintf(" %d", color)
			fmt.Fprint(&b, str)
			row_size += len(str)

			if row_size > 70 {
				fmt.Fprintf(&b, "\n")
				row_size = 0
			}
		}

		// saut de ligne pour chaque ligne de la matrice de pixems
		if current_col < int(c.Width)-1 {
			current_col += 1
		} else {
			current_col = 0
			row_size = 0
			fmt.Fprintf(&b, "\n")

		}

	}

	return b.String()
}
