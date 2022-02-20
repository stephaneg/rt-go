package core

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
