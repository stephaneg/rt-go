package core

import (
	"math"
)

type Tuple [4]float64

func (t Tuple) X() float64 {
	return t[0]
}

func (t Tuple) Y() float64 {
	return t[1]
}

func (t Tuple) Z() float64 {
	return t[2]
}

func (t Tuple) W() float64 {
	return t[3]
}

// isPoint return true if the tuple is a point : w == 1
func (t Tuple) IsPoint() bool {
	return !t.IsVector()
}

// isVector return true if the tuple is a point : w == 1
func (t Tuple) IsVector() bool {
	return t.W() == 0.0
}

func NewTuple(x, y, z, w float64) Tuple {
	t := [4]float64{x, y, z, w}
	return t
}

func NewTupleFrom(data [4]float64) Tuple {
	return Tuple(data)
}

func NewPoint(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 1.0)
}

func NewVector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 0.0)
}

func (t Tuple) EqualTuple(other Tuple) bool {
	return FuzzyEqf64(t.X(), other.X()) && FuzzyEqf64(t.Y(), other.Y()) && FuzzyEqf64(t.Z(), other.Z()) && FuzzyEqf64(t.W(), other.W())
}

func (t Tuple) AddTuple(other Tuple) Tuple {
	return Tuple{t.X() + other.X(), t.Y() + other.Y(), t.Z() + other.Z(), t.W() + other.W()}
}

func (t Tuple) SubTuple(other Tuple) Tuple {
	return Tuple{t.X() - other.X(), t.Y() - other.Y(), t.Z() - other.Z(), t.W() - other.W()}
}

func (t Tuple) NegTuple() Tuple {
	return Tuple{-t.X(), -t.Y(), -t.Z(), -t.W()}
}

func (t Tuple) MulTuple(other float64) Tuple {
	return Tuple{t.X() * other, t.Y() * other, t.Z() * other, t.W() * other}
}

func (t Tuple) DivTuple(other float64) Tuple {
	return Tuple{t.X() / other, t.Y() / other, t.Z() / other, t.W() / other}
}

func (t Tuple) DotTuple(other Tuple) Tuple {
	return Tuple{t.X() * other.X(), t.Y() * other.Y(), t.Z() * other.Z(), t.W() * other.W()}
}

func (t Tuple) Magintude() float64 {
	r := t.X()*t.X() + t.Y()*t.Y() + t.Z()*t.Z()
	return math.Sqrt(r)
}

func (t Tuple) Normalize() Tuple {
	return t.DivTuple(t.Magintude())
}
