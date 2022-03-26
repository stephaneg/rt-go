package rays

import (
	"math"

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

	sphere_to_ray := ray.Origin.SubTuple(core.NewPoint(0.0, 0.0, 0.0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) -1

	discriminant := b*b - 4.0*a*c

	if discriminant < 0 {
		return ret
	} else {
		t1 := (-1.0*b - math.Sqrt(discriminant)) / (2.0 * a)
		t2 := (-1.0*b + math.Sqrt(discriminant)) / (2.0 * a)
		ret = append(ret, t1)
		ret = append(ret, t2)
	}

	return ret
}
