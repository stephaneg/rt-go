package rays

import "math/rand"

type Sphere struct {
	Id    int64
	Label string
}

func NewSphere() *Sphere {
	return &Sphere{
		Id:    rand.Int63(),
		Label: "sphere",
	}
}

func (sphere Sphere) ID() int64 {
	return sphere.Id
}

func (sphere Sphere) Name() string {
	return sphere.Label
}
