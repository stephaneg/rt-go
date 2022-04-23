package rays

import (
	"math"
	"sort"

	"github.com/stephstephg/rt-go/core"
)

type Intersection struct {
	T float64
	S Shape
}

// Intersections implements sort.Interface based on the T field
type Intersections []Intersection

func (ixs Intersections) Len() int           { return len(ixs) }
func (ixs Intersections) Less(i, j int) bool { return ixs[i].T < ixs[j].T }
func (ixs Intersections) Swap(i, j int)      { ixs[i], ixs[j] = ixs[j], ixs[i] }

// NewInterSection creates an Intersection at a point t on a Ray with a Shape shape
func NewIntersection(t float64, shape Shape) Intersection {
	return Intersection{t, shape}
}

// Intersect intersect a Sphere and a Ray
func Intersect(sphere *Sphere, ray *Ray) Intersections {
	var ret Intersections
	itxs := make([]Intersection, 0)

	sphere_to_ray := ray.Origin.SubTuple(core.NewPoint(0.0, 0.0, 0.0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) - 1

	discriminant := b*b - 4.0*a*c

	if discriminant < 0 {
		return Intersections(itxs)
	} else {
		t1 := (-1.0*b - math.Sqrt(discriminant)) / (2.0 * a)
		t2 := (-1.0*b + math.Sqrt(discriminant)) / (2.0 * a)
		itxs = append(itxs, NewIntersection(t1, sphere))
		itxs = append(itxs, NewIntersection(t2, sphere))
	}

	ret = Intersections(itxs)
	sort.Sort(ret)
	return ret
}

func Hit(itxs Intersections) Intersection {
	var res Intersection = NewIntersection(0.0, NewVoidShape())


	sort.Sort(itxs)
	for _, v := range itxs {
		if v.T > 0 {
			res = v
			break
		}
	}
	return res
}
