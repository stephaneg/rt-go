package core

import (
	"math"
	"testing"
)

func TestScaleMatrixToPoint(t *testing.T) {
	m := ScaleMatrix(2.0, 3.0, 4.0)
	p := NewPoint(-4.0, 6.0, 8.0)

	res := m.DotTuple(p)
	expected := NewPoint(-8.0, 18.0, 32.0)

	if !res.EqualTuple(expected) {
		t.Errorf("scale matrix : expected %v, get %v", expected, res)
	}
}

func TestScaleMatrixToVector(t *testing.T) {
	m := ScaleMatrix(2.0, 3.0, 4.0)
	p := NewVector(-4.0, 6.0, 8.0)

	res := m.DotTuple(p)
	expected := NewVector(-8.0, 18.0, 32.0)

	if !res.EqualTuple(expected) {
		t.Errorf("scale matrix : expected %v, get %v", expected, res)
	}
}

func TestInversedScaleMatrixToVector(t *testing.T) {
	m := ScaleMatrix(2.0, 3.0, 4.0)
	inv, _ := m.Inverse()
	p := NewVector(-4.0, 6.0, 8.0)

	res := inv.DotTuple(p)
	expected := NewVector(-2.0, 2.0, 2.0)

	if !res.EqualTuple(expected) {
		t.Errorf("scale matrix : expected %v, get %v", expected, res)
	}
}

func TestReflectiongMatrix(t *testing.T) {
	m := ScaleMatrix(-1.0, 1.0, 1.0)
	p := NewVector(2.0, 3.0, 4.0)

	res := m.DotTuple(p)
	expected := NewVector(-2.0, 3.0, 4.0)

	if !res.EqualTuple(expected) {
		t.Errorf("scale matrix : expected %v, get %v", expected, res)
	}
}

func TestRotationAroundX(t *testing.T) {
	p := NewPoint(0.0, 1.0, 0.0)
	half_quarter := RotationX(math.Pi / 4)
	full_quarter := RotationX(math.Pi / 2)

	res1 := half_quarter.DotTuple(p)
	res2 := full_quarter.DotTuple(p)

	expected1 := NewPoint(0.0, math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0)
	expected2 := NewPoint(0.0, 0.0, 1.0)

	if !res1.EqualTuple(expected1) {
		t.Errorf("expected tupe is %v, result is %v", expected1, res1)
	}

	if !res2.EqualTuple(expected2) {
		t.Errorf("expected tupe is %v, result is %v", expected2, res2)
	}
}

func TestRotationAroundY(t *testing.T) {
	p := NewPoint(0.0, 0.0, 1.0)
	half_quarter := RotationY(math.Pi / 4)
	full_quarter := RotationY(math.Pi / 2)

	res1 := half_quarter.DotTuple(p)
	res2 := full_quarter.DotTuple(p)

	expected1 := NewPoint(math.Sqrt(2.0)/2.0, 0.0, math.Sqrt(2.0)/2.0)
	expected2 := NewPoint(1.0, 0.0, 0.0)

	if !res1.EqualTuple(expected1) {
		t.Errorf("expected tupe is %v, result is %v", expected1, res1)
	}

	if !res2.EqualTuple(expected2) {
		t.Errorf("expected tupe is %v, result is %v", expected2, res2)
	}
}

func TestRotationAroundZ(t *testing.T) {
	p := NewPoint(0.0, 1.0, 0.0)
	half_quarter := RotationZ(math.Pi / 4)
	full_quarter := RotationZ(math.Pi / 2)

	res1 := half_quarter.DotTuple(p)
	res2 := full_quarter.DotTuple(p)

	expected1 := NewPoint(-math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0, 0.0)
	expected2 := NewPoint(-1.0, 0.0, 0.0)

	if !res1.EqualTuple(expected1) {
		t.Errorf("expected tupe is %v, result is %v", expected1, res1)
	}

	if !res2.EqualTuple(expected2) {
		t.Errorf("expected tupe is %v, result is %v", expected2, res2)
	}
}
