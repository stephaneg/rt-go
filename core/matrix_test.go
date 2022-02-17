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
