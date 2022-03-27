package rays

import (
	"testing"

	"github.com/stephstephg/rt-go/core"
)

func TestIntersectSphereAtTwoPoints(t *testing.T) {

	ray := Ray{core.NewPoint(0.0, 0.0, -5.0), core.NewVector(0.0, 0.0, 1.0)}

	s := Sphere{}

	var itxs []Intersection = Intersect(&s, &ray)

	if len(itxs) != 2 {
		t.Errorf("error in the number of intersection points, expected 2, get %v", len(itxs))
	}

	if !core.FuzzyEqf64(itxs[0].T, 4.0) {
		t.Errorf("expected 4.0 for the first point, get %v", itxs[0].T)
	}

	if !core.FuzzyEqf64(itxs[1].T, 6.0) {
		t.Errorf("expected 6.0 for the first point, get %v", itxs[1].T)
	}
}

func TestIntersectSphereAtTangent(t *testing.T) {

	ray := Ray{core.NewPoint(0.0, 1.0, -5.0), core.NewVector(0.0, 0.0, 1.0)}

	s := Sphere{}

	var itxs []Intersection = Intersect(&s, &ray)

	if len(itxs) != 2 {
		t.Errorf("error in the number of intersection points, expected 2, get %v", len(itxs))
	}

	if !core.FuzzyEqf64(itxs[0].T, 5.0) {
		t.Errorf("expected 5.0 for the first point, get %v", itxs[0].T)
	}

	if !core.FuzzyEqf64(itxs[1].T, 5.0) {
		t.Errorf("expected 5.0 for the first point, get %v", itxs[1].T)
	}
}

func TestIntersectRayMissSphere(t *testing.T) {

	ray := Ray{core.NewPoint(0.0, 2.0, -5.0), core.NewVector(0.0, 0.0, 1.0)}

	s := Sphere{}

	var itxs []Intersection = Intersect(&s, &ray)

	if len(itxs) != 0 {
		t.Errorf("error in the number of intersection points, expected 0, get %v", len(itxs))
	}
}

func TestIntersectRayInsideASphere(t *testing.T) {

	ray := Ray{core.NewPoint(0.0, 0.0, 0.0), core.NewVector(0.0, 0.0, 1.0)}

	s := Sphere{}

	var itxs []Intersection = Intersect(&s, &ray)

	if len(itxs) != 2 {
		t.Errorf("error in the number of intersection points, expected 2, get %v", len(itxs))
	}

	if !core.FuzzyEqf64(itxs[0].T, -1.0) {
		t.Errorf("expected -1.0 for the first point, get %v", itxs[0].T)
	}

	if !core.FuzzyEqf64(itxs[1].T, 1.0) {
		t.Errorf("expected 1.0 for the first point, get %v", itxs[1].T)
	}
}

func TestIntersectSphereBehindRay(t *testing.T) {

	ray := Ray{core.NewPoint(0.0, 0.0, 5.0), core.NewVector(0.0, 0.0, 1.0)}

	s := Sphere{}

	var itxs []Intersection = Intersect(&s, &ray)

	if len(itxs) != 2 {
		t.Errorf("error in the number of intersection points, expected 2, get %v", len(itxs))
	}

	if !core.FuzzyEqf64(itxs[0].T, -6.0) {
		t.Errorf("expected -6.0 for the first point, get %v", itxs[0].T)
	}

	if !core.FuzzyEqf64(itxs[1].T, -4.0) {
		t.Errorf("expected -4.0 for the first point, get %v", itxs[1].T)
	}
}
