package rays

import (
	"reflect"
	"testing"

	"github.com/stephstephg/rt-go/core"
)

func TestIntersection(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, s)

	if !core.FuzzyEqf64(i.T, 3.5) {
		t.Errorf("error in Intersection.T, expecting 3.5, get %v", i.T)
	}

	if i.S != s {
		t.Errorf("error in Intersection.S, expecting %v, get %v", s, i.T)
	}

}

func TestHitWithAllPositiveIntersections(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1.0, s)
	i2 := NewIntersection(2.0, s)
	ixs := []Intersection{i1, i2}

	res := Hit(ixs)

	if res != i1 {
		t.Errorf("error in Hit, expecting %+v, get %+v", i1, res)
	}

}

func TestHitWithSomeNegativeIntersections(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-1.0, s)
	i2 := NewIntersection(1.0, s)
	ixs := []Intersection{i1, i2}

	res := Hit(ixs)

	if res != i2 {
		t.Errorf("error in Hit, expecting %+v, get %+v", i2, res)
	}
}

func TestHitWithAllNegativeIntersections(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-1.0, s)
	i2 := NewIntersection(-2.0, s)
	ixs := []Intersection{i1, i2}

	res := Hit(ixs)

	if reflect.DeepEqual(res, NewVoidShape()) {
		t.Errorf("error in Hit, expecting nil, get %+v", res)
	}

}

func TestHitLowestNonNegativeIntersections(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(5.0, s)
	i2 := NewIntersection(7.0, s)
	i3 := NewIntersection(-3.0, s)
	i4 := NewIntersection(2.0, s)
	itxs := []Intersection{i1, i2, i3, i4}

	res := Hit(itxs)

	if !reflect.DeepEqual(res, i4) {
		t.Errorf("error in TestHitLowestNonNegativeIntersections, expecting %+v, get %+v", i4, res)
	}
}
