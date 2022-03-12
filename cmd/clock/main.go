package main

import (
	"fmt"
	"math"

	"github.com/stephstephg/rt-go/core"
)

func main() {
	r := core.RotationY(2.0 * math.Pi / 12.0)
	var x, y uint
	current := core.NewPoint(0.0, 0.0, 1.0)

	c := core.NewCanvas(800, 800)
	color := core.NewColor(1.0, 0.0, 0.0)
	radius := 3.0 / 8.0 * 800.0

	for i := 0; i < 12; i++ {
		x = uint(current.X()*radius + 400)
		y = uint(current.Z()*radius + 400)
		c.Write(x, y, color)

		current = r.DotTuple(current)
	}

	s := c.ToPPM()
	fmt.Print(s)

}
