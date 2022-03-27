package rays

import (
	"math"

	"github.com/stephstephg/rt-go/core"
)

type Intersection struct {
	T float64
	S Shape
}

func NewIntersection(t float64, shape Shape) Intersection {
	return Intersection{t, shape}
}

func Intersect(sphere *Sphere, ray *Ray) []Intersection {
	ret := make([]Intersection, 0)

	sphere_to_ray := ray.Origin.SubTuple(core.NewPoint(0.0, 0.0, 0.0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) - 1

	discriminant := b*b - 4.0*a*c

	if discriminant < 0 {
		return ret
	} else {
		t1 := (-1.0*b - math.Sqrt(discriminant)) / (2.0 * a)
		t2 := (-1.0*b + math.Sqrt(discriminant)) / (2.0 * a)
		ret = append(ret, NewIntersection(t1, sphere))
		ret = append(ret, NewIntersection(t2, sphere))
	}

	return ret
}
