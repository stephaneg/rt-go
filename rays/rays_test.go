package rays

import (
	"testing"

	"github.com/stephstephg/rt-go/core"
)

func TestPosition(t *testing.T) {
	ray := Ray{core.NewPoint(2.0, 3.0, 4.0), core.NewVector(1.0, 0.0, 0.0)}

	var result core.Tuple
	var expected core.Tuple

	result = ray.Position(0.0)
	expected = core.NewPoint(2.0, 3.0, 4.0)

	if !result.EqualTuple(expected) {
		t.Errorf("Error in Position, expected %v, get %v", expected, result)
	}

	result = ray.Position(1.0)
	expected = core.NewPoint(3.0, 3.0, 4.0)

	if !result.EqualTuple(expected) {
		t.Errorf("Error in Position, expected %v, get %v", expected, result)
	}

	result = ray.Position(-1.0)
	expected = core.NewPoint(1.0, 3.0, 4.0)

	if !result.EqualTuple(expected) {
		t.Errorf("Error in Position, expected %v, get %v", expected, result)
	}

	result = ray.Position(2.5)
	expected = core.NewPoint(4.5, 3.0, 4.0)

	if !result.EqualTuple(expected) {
		t.Errorf("Error in Position, expected %v, get %v", expected, result)
	}

}
