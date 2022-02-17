package core

type Matrix4 [4][4]float64

func NewMatrix4() Matrix4 {
	m := [4][4]float64{
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
	}
	return m
}

func Identity4() Matrix4 {
	m := [4][4]float64{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	}
	return m
}
