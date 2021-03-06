package core

import (
	"testing"
)

func TestToPPM(t *testing.T) {

	c := NewCanvas(10, 2)
	var i, j uint

	for i = 0; i < 10; i++ {
		for j = 0; j < 2; j++ {
			c.Write(i, j, NewColor(1.0, 0.8, 0.6))
		}
	}

	c.ToPPM()

}
