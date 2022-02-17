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
	x = uint(current.Position.X)
	y = c.Height - uint(current.Position.Y)

	c.Write(x, y, color)

	for current.Position.Y > 0.0 {
		fmt.Printf("position : %v", current.Position.Y)
		current = dynamic.Tick(env, &projectile)

		x = uint(current.Position.X)
		y = c.Height - uint(current.Position.Y)
		if current.Position.Y > 0.0 {
			c.Write(x, y, color)
		}

	}
	s := c.ToPPM()
	fmt.Print(s)

}
