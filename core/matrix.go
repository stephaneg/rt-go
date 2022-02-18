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

func NewMatrix4From(data [4][4]float64) Matrix4 {
	m := Matrix4(data)
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

func (m Matrix4) EqualMatrix4(other Matrix4) bool {
	for i, row := range m {
		for j, val := range row {
			if !FuzzyEqf64(val, other[i][j]) {
				return false
			}
		}
	}
	return true
}

func (m Matrix4) DotMatrix4(o Matrix4) Matrix4 {
	ret := NewMatrix4()
	for i, row := range ret {
		for j := range row {
			ret[i][j] = m[i][0]*o[0][j] + m[i][1]*o[1][j] + m[i][2]*o[2][j] + m[i][3]*o[3][j]
		}
	}
	return ret
}

func (m Matrix4) Matrix4xTuple(o Tuple) Tuple {
	ret := NewTuple(0.0, 0.0, 0.0, 0.0)
	for i, row := range m {
		for j, val := range row {
			ret[i] += val * o[j]
		}
	}
	return ret
}

func (m Matrix4) Transpose() Matrix4 {
	ret := NewMatrix4()
	return ret
}
