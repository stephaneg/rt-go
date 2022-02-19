package core

import (
	"math"
)

const EPSILON float64 = 0.000001

// FuzzyEqf64 is a fuzzy equal operator for float64
func FuzzyEqf64(a, b float64) bool {
	return (math.Abs(a-b) < EPSILON)
}
