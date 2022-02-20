package core

import (
	"testing"
)

func TestAddColor(t *testing.T) {
	c1 := Color{1.0, 2.0, 3.0}
	c2 := Color{4.0, 5.0, 6.0}
	c3 := c1.AddColor(c2)

	if !FuzzyEqf64(c3.R, 5.0) {
		t.Errorf("expected 5.0, get %v", c3.R)
	}

	if !FuzzyEqf64(c3.G, 7.0) {
		t.Errorf("expected 7.0, get %v", c3.G)
	}

	if !FuzzyEqf64(c3.B, 9.0) {
		t.Errorf("expected 9.0, get %v", c3.B)
	}
}

func TestFuzzyEqColor(t *testing.T) {
	c1 := NewColor(1.0, 2.0, 3.0)
	c2 := NewColor(1.0/(1.0*1.0), 4.0/2.0*1.0, 6.0/2.0)
	c3 := NewColor(1.0, 2.0, 3.0+EPSILON)

	if !c1.FuzzyEqColor(c2) {
		t.Error("expecting c1 == c2")
	}

	if c1.FuzzyEqColor(c3) {
		t.Error("expecting c1 =! c3")
	}

}
