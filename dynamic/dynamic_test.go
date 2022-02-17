package dynamic

import (
	"testing"

	"github.com/stephstephg/rt-go/core"
)

func TestTickVelocityOnly(t *testing.T) {
	start := core.NewPoint(0.0, 0.0, 0.0)
	velocity := core.NewVector(1.0, 0.0, 0.0)

	projectile := Projectile{Position: start, Velocity: velocity}

	gravity := core.NewVector(0.0, 0.0, 0.0)
	wind := core.NewVector(0.0, 0.0, 0.0)
	env := NewEnvironment(gravity, wind)

	var old, current Projectile

	old = projectile
	current = Tick(env, &old)

	if !core.FuzzyEqf64(current.Position.X, 1.0) {
		t.Errorf("expected X==1.0, get %v", current.Position.X)
	}

	old = current
	current = Tick(env, &old)

	if !core.FuzzyEqf64(current.Position.X, 2.0) {
		t.Errorf("expected X==1.0, get %v", current.Position.X)
	}
}
