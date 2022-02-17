package main

import (
	"fmt"

	"github.com/stephstephg/rt-go/core"
	"github.com/stephstephg/rt-go/dynamic"
)

func main() {
	start := core.NewPoint(0.0, 1.0, 0.0)
	velocity := core.NewVector(1.0, 1.8, 0.0).Normalize().MulTuple(11.25)

	projectile := dynamic.Projectile{Position: start, Velocity: velocity}

	gravity := core.NewVector(0.0, -0.1, 0.0)
	wind := core.NewVector(-0.01, 0.0, 0.0)
	env := dynamic.NewEnvironment(gravity, wind)

	c := core.NewCanvas(900, 600)
	color := core.NewColor(1.0, 0.0, 0.0)

	var current dynamic.Projectile = projectile
	var x, y uint
	var current_y float64
	x = uint(current.Position.X())
	y = c.Height - uint(current.Position.Y())

	c.Write(x, y, color)

	current_y = current.Position.Y()

	for current_y > 0.0 {
		current = dynamic.Tick(env, &current)
		current_y = current.Position.Y()

		x = uint(current.Position.X())
		y = c.Height - uint(current.Position.Y())
		if current.Position.Y() > 0.0 {
			c.Write(x, y, color)
		}

	}
	s := c.ToPPM()
	fmt.Print(s)

}
