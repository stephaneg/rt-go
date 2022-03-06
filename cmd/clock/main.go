package main

import (
	"fmt"
	"math"

	"github.com/stephstephg/rt-go/core"
)

func main() {
	//origin := core.NewPoint(0.0, 0.0, 0.0)
	r := core.RotationY(3.0 * math.Pi / 6.0)

	twelve := core.NewPoint(0.0, 0.0, 1.0)

	three := r.DotTuple(twelve)

	c := core.NewCanvas(800, 800)
	color := core.NewColor(1.0, 0.0, 0.0)

	radius := 3.0 / 8.0 * 800.0

	var x, y uint
	x = uint(twelve.X() + radius)
	y = uint(twelve.Z() + radius)
	c.Write(x, y, color)

	x = uint(three.X() + radius)
	y = uint(three.Z() + radius)
	c.Write(x, y, color)

	s := c.ToPPM()
	fmt.Print(s)
}
