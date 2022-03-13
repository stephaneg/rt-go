package rays

import (
	"github.com/stephstephg/rt-go/core"
)

type Ray struct {
	Origin    core.Tuple
	Direction core.Tuple
}

func (r Ray) Position(t float64) core.Tuple {
	return r.Origin.AddTuple(r.Direction.MulTuple(t))
}

func Intersect(sphere *Sphere, ray *Ray) []float64 {
	ret := make([]float64, 0)
	return ret
}
