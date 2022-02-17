package dynamic

import (
	"github.com/stephstephg/rt-go/core"
)

type Projectile struct {
	Position core.Tuple
	Velocity core.Tuple
}

type Environment struct {
	Gravity core.Tuple
	Wind    core.Tuple
}

func NewEnvironment(gravity core.Tuple, wind core.Tuple) Environment {
	return Environment{gravity, wind}
}

func Tick(env Environment, proj *Projectile) Projectile {
	p := Projectile{proj.Position.AddTuple(proj.Velocity), proj.Velocity.AddTuple(env.Gravity).AddTuple(env.Wind)}
	return p
}
