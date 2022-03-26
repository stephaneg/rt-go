package core

import (
	"testing"
)

func TestAddTuple(t *testing.T) {
	t1 := Tuple{1.0, 2.0, 3.0, 1.0}
	t2 := Tuple{4.0, -5.0, -6.0, 2.0}
	t3 := t1.AddTuple(t2)

	if !FuzzyEqf64(t3.X(), 5.0) {
		t.Errorf("expected 5.0, get %v", t3.X())
	}

	if !FuzzyEqf64(t3.Y(), -3.0) {
		t.Errorf("expected -3.0, get %v", t3.Y())
	}

	if !FuzzyEqf64(t3.Z(), -3.0) {
		t.Errorf("expected -3.0, get %v", t3.Z())
	}

	if !FuzzyEqf64(t3.W(), 3.0) {
		t.Errorf("expected 3.0, get %v", t3.W())
	}

}

func TestDot(t *testing.T) {
	a := NewVector(1.0, 2.0, 3.0)
	b := NewVector(2.0, 3.0, 4.0)

	res := a.Dot(b)

	if !FuzzyEqf64(res, 20.0) {
		t.Errorf("expected 20.0, get %v", res)
	}

}
