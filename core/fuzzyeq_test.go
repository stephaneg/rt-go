package core

import (
	"testing"
)

func TestFuzzyEqf64(t *testing.T) {
	var a, b, c, d, e float64

	a = 1.0
	b = 2.0
	c = 2.0 / 2.0
	d = a + EPSILON
	e = a - EPSILON

	if FuzzyEqf64(a, b) {
		t.Error("expected a!=b")
	}

	if !FuzzyEqf64(a, c) {
		t.Error("expected a == c")
	}

	if FuzzyEqf64(a, d) {
		t.Error("expected a != d")
	}

	if !FuzzyEqf64(a, e) {
		t.Error("expected a == e")
	}
}
