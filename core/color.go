package core

import (
	"math"
)

type Color struct {
	R float64
	G float64
	B float64
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c Color) AddColor(other Color) Color {
	return Color{c.R + other.R, c.G + other.G, c.B + other.B}
}

func (c Color) SubColor(other Color) Color {
	return Color{c.R - other.R, c.G - other.G, c.B - other.B}
}

func (c Color) MulColor(other Color) Color {
	return Color{c.R * other.R, c.G * other.G, c.B * other.B}
}

func (c Color) MulColorF64(other float64) Color {
	return Color{c.R * other, c.G * other, c.B * other}
}

func (c Color) FuzzyEqColor(other Color) bool {
	return FuzzyEqf64(c.R, other.R) && FuzzyEqf64(c.G, other.G) && FuzzyEqf64(c.B, other.B)
}

func (c Color) ToRgb() []byte {
	var rgb []byte
	var r, g, b byte

	r = (byte)(math.Round(math.Max(math.Min(c.R, 1.0), 0.0) * 255.0))
	g = (byte)(math.Round(math.Max(math.Min(c.G, 1.0), 0.0) * 255.0))
	b = (byte)(math.Round(math.Max(math.Min(c.B, 1.0), 0.0) * 255.0))
	rgb = append(rgb, r)
	rgb = append(rgb, g)
	rgb = append(rgb, b)

	return rgb
}
