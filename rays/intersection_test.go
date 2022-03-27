package rays

import (
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


