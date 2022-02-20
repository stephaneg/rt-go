package core

import (
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
