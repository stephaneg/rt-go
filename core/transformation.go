package core

import "math"

// Translation matrix return a 4x4 translation matrix - the translation vector is (x, y, z)
func Translation(x, y, z float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{1.0, 0.0, 0.0, x},
		{0.0, 1.0, 0.0, y},
		{0.0, 0.0, 1.0, z},
		{0.0, 0.0, 0.0, 1.0},
	})
}

// ScaleMatrix return a 4x4 scale matrix, scaling accordingly (x, y, z)
func Scaling(x, y, z float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{x, 0.0, 0.0, 0.0},
		{0.0, y, 0.0, 0.0},
		{0.0, 0.0, z, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

// RotationX creates a rotation matrix on X axis with a radian of radian
func RotationX(radian float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, math.Cos(radian), -math.Sin(radian), 0.0},
		{0.0, math.Sin(radian), math.Cos(radian), 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

// RotationY creates a rotation matrix on Y axis with a radian of radian
func RotationY(radian float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{math.Cos(radian), 0.0, math.Sin(radian), 0.0},
		{0.0, 1.1, 0.0, 0.0},
		{-math.Sin(radian), 0.0, math.Cos(radian), 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

// RotationZ creates a rotation matrix on Z axis with a radian of radian
func RotationZ(radian float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{math.Cos(radian), -math.Sin(radian), 0.0, 0.0},
		{math.Sin(radian), math.Cos(radian), 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

// Shearing transformation returns a shearing transformation matrix
func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix4 {
	return NewMatrix4From([4][4]float64{
		{1.0, xy, xz, 0.0},
		{yx, 1.0, yz, 0.0},
		{zx, zy, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}
