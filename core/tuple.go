package core

import (
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// isPoint return true if the tuple is a point : w == 1
func (t Tuple) IsPoint() bool {
	return !t.IsVector()
}

// isVector return true if the tuple is a point : w == 1
func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func NewTuple(x, y, z, w float64) Tuple {
	t := Tuple{x, y, z, w}
	return t
}

func NewPoint(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 1.0)
}

func NewVector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 0.0)
}

func FuzzyEqTuple(t1, t2 Tuple) bool {
	return FuzzyEqf64(t1.X, t2.X) && FuzzyEqf64(t1.Y, t2.Y) && FuzzyEqf64(t1.Z, t2.Z) && FuzzyEqf64(t1.W, t2.W)
}

func (t Tuple) AddTuple(other Tuple) Tuple {
	return Tuple{t.X + other.X, t.Y + other.Y, t.Z + other.Z, t.W + other.W}
}

func (t Tuple) SubTuple(other Tuple) Tuple {
	return Tuple{t.X - other.X, t.Y - other.Y, t.Z - other.Z, t.W - other.W}
}

func (t Tuple) NegTuple() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t Tuple) MulTuple(other float64) Tuple {
	return Tuple{t.X * other, t.Y * other, t.Z * other, t.W * other}
}

func (t Tuple) DivTuple(other float64) Tuple {
	return Tuple{t.X / other, t.Y / other, t.Z / other, t.W / other}
}

func (t Tuple) DotTuple(other Tuple) Tuple {
	return Tuple{t.X * other.X, t.Y * other.Y, t.Z * other.Z, t.W * other.W}
}

func (t Tuple) Magintude() float64 {
	r := t.X*t.X + t.Y*t.Y + t.Z*t.Z
	return math.Sqrt(r)
}

func (t Tuple) Normalize() Tuple {
	return t.DivTuple(t.Magintude())
}
