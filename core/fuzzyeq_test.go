package core

import (
	"math"
	"testing"
)

func TestFuzzyEqf64(t *testing.T) {
	var a, b, c, d, e float64

	a = 1.0
	b = 2.0
	c = 2.0 *1. / 2.0
	d = a + EPSILON
	e = a + 2*EPSILON

	if FuzzyEqf64(a, b) {
		t.Errorf("expected a!=b %v %v", a, b)
	}

	if !FuzzyEqf64(a, c) {
		t.Error("expected a == c")
	}

	if !FuzzyEqf64(a, d) {
		t.Errorf("expected %v == %v ", a, d)
	}

	if FuzzyEqf64(a, e) {
		t.Errorf("expected %v != %v, delta is %v", a, e, math.Abs(a-e))
	}
}
