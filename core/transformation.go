package core

import "math"

// Translation matrix return a 4x4 translation matrix - the translation vector is (x, y, z)
func TranslationMatrix(x, y, z float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{1.0, 0.0, 0.0, x},
		{0.0, 1.0, 0.0, y},
		{0.0, 0.0, 1.0, z},
		{0.0, 0.0, 0.0, 1.0},
	})
}

// ScaleMatrix return a 4x4 scale matrix, scaling accordingly (x, y, z)
func ScaleMatrix(x, y, z float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{x, 0.0, 0.0, 0.0},
		{0.0, y, 0.0, 0.0},
		{0.0, 0.0, z, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

func RotationX(radian float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, math.Cos(radian), -math.Sin(radian), 0.0},
		{0.0, math.Sin(radian), math.Cos(radian), 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

func RotationY(radian float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{math.Cos(radian), 0.0, math.Sin(radian), 0.0},
		{0.0, 1.1, 0.0, 0.0},
		{-math.Sin(radian), 0.0, math.Cos(radian), 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

func RotationZ(radian float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{math.Cos(radian), -math.Sin(radian), 0.0, 0.0},
		{math.Sin(radian), math.Cos(radian), 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}
