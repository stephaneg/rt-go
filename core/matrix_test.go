package core

import (
	"testing"
)

func TestIdentity4(t *testing.T) {
	m := Identity4()

	if !(m[0][0] == 1.0) {
		t.Errorf("expecting 1.0, get %v", m[0][0])
	}
	if !(m[0][1] == 0.0) {
		t.Errorf("expecting 0.0, get %v", m[0][0])
	}
	if !(m[0][2] == 0.0) {
		t.Errorf("expecting 0.0, get %v", m[0][0])
	}
	if !(m[3][3] == 1.0) {
		t.Errorf("expecting 1.0, get %v", m[0][0])
	}
}

func TestEqualMatrix4(t *testing.T) {
	m1 := NewMatrix4()
	m2 := NewMatrix4()

	if !m1.EqualMatrix4(m2) {
		t.Error("expected empty matrices to be equal")
	}
	m1[0][0] = 3.0

	if m1.EqualMatrix4(m2) {
		t.Error("m1 has changed and is no more equal to m2")
	}

	m1 = Identity4()
	m2 = Identity4()
	if !m1.EqualMatrix4(m2) {
		t.Error("Identify are equals")
	}

}

func TestDotlMatrix4(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{1.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0, 8.0},
		{9.0, 8.0, 7.0, 6.0},
		{5.0, 4.0, 3.0, 2.0},
	})

	b := NewMatrix4From([4][4]float64{
		{-2.0, 1.0, 2.0, 3.0},
		{3.0, 2.0, 1.0, -1.0},
		{4.0, 3.0, 6.0, 5.0},
		{1.0, 2.0, 7.0, 8.0},
	})

	expected := NewMatrix4From([4][4]float64{
		{20.0, 22.0, 50.0, 48.0},
		{44.0, 54.0, 114.0, 108.0},
		{40.0, 58.0, 110.0, 102.0},
		{16.0, 26.0, 46.0, 42.0},
	})

	res := a.DotMatrix4(b)

	if !res.EqualMatrix4(expected) {
		t.Error("error in matrix dot matrix ")
	}
}

func TestMatrix4xTuple(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{1.0, 2.0, 3.0, 4.0},
		{2.0, 4.0, 4.0, 2.0},
		{8.0, 6.0, 4.0, 1.0},
		{0.0, 0.0, 0.0, 1.0},
	})

	tuple := NewTuple(1.0, 2.0, 3.0, 1.0)
	res := a.Matrix4xTuple(tuple)

	expected := NewTuple(18.0, 24.0, 33.0, 1.0)

	if !res.EqualTuple(expected) {
		t.Errorf("error in matrix dot tuple : %v ", res)
	}

}

func TestTransposeMatrix4(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{0.0, 9.0, 3.0, 0.0},
		{9.0, 8.0, 0.0, 8.0},
		{1.0, 8.0, 5.0, 3.0},
		{0.0, 0.0, 5.0, 8.0},
	})

	expected := NewMatrix4From([4][4]float64{
		{0.0, 9.0, 1.0, 0.0},
		{9.0, 8.0, 8.0, 0.0},
		{3.0, 0.0, 5.0, 5.0},
		{0.0, 8.0, 3.0, 8.0},
	})

	res := a.Transpose()

	if !res.EqualMatrix4(expected) {
		t.Errorf("error in matrix transposition")
	}

}
