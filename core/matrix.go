package core

import (
	"errors"
)

type Matrix4 [4][4]float64
type Matrix3 [3][3]float64
type Matrix2 [2][2]float64

// NewMatrix4 creates a 4x4 matrix initialized with 0.0
func NewMatrix4() Matrix4 {
	m := [4][4]float64{
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0},
	}
	return m
}

// NewMatrix2 creates a 2x2 matrix initialized with 0.0
func NewMatrix2() Matrix2 {
	m := [2][2]float64{
		{0.0, 0.0},
		{0.0, 0.0},
	}
	return m
}

// NewMatrix3 creates a 3x3 matrix initialized with 0.0
func NewMatrix3() Matrix3 {
	m := [3][3]float64{
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0},
	}
	return m
}

// NewMatrix4From creates a 4x4 matrix initialized the [4][4]float64 array in param
func NewMatrix4From(data [4][4]float64) Matrix4 {
	m := Matrix4(data)
	return m
}

// NewMatrix3From creates a 3x3 matrix initialized the [3][3]float64 array in param
func NewMatrix3From(data [3][3]float64) Matrix3 {
	m := Matrix3(data)
	return m
}

// NewMatrix2From creates a 2x2 matrix initialized the [2][2]float64 array in param
func NewMatrix2From(data [2][2]float64) Matrix2 {
	m := Matrix2(data)
	return m
}

// Identity4 creates the 4x4 Identity matrix
func Identity4() Matrix4 {
	m := [4][4]float64{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	}
	return m
}

// EqualMatrix test the fuzzy equality of two 4x4 matrices of float64
func (m Matrix4) EqualMatrix(other Matrix4) bool {
	for i, row := range m {
		for j, val := range row {
			if !FuzzyEqf64(val, other[i][j]) {
				return false
			}
		}
	}
	return true
}

// EqualMatrix test the fuzzy equality of two 3x3 matrices of float64
func (m Matrix3) EqualMatrix(other Matrix3) bool {
	for i, row := range m {
		for j, val := range row {
			if !FuzzyEqf64(val, other[i][j]) {
				return false
			}
		}
	}
	return true
}

// EqualMatrix test the fuzzy equality of two 2x2 matrices of float64
func (m Matrix2) EqualMatrix(other Matrix2) bool {
	for i, row := range m {
		for j, val := range row {
			if !FuzzyEqf64(val, other[i][j]) {
				return false
			}
		}
	}
	return true
}

// DotMatrix4 produce the dot product of two 4x4 matrices. The result is a 4x4 matrix
func (m Matrix4) DotMatrix(o Matrix4) Matrix4 {
	ret := NewMatrix4()
	for i, row := range ret {
		for j := range row {
			ret[i][j] = m[i][0]*o[0][j] + m[i][1]*o[1][j] + m[i][2]*o[2][j] + m[i][3]*o[3][j]
		}
	}
	return ret
}

// Matrix4xTuple produce the dot product of a 4x4 matrix with a 4x1 tuple
func (m Matrix4) DotTuple(o Tuple) Tuple {
	ret := NewTuple(0.0, 0.0, 0.0, 0.0)
	for i, row := range m {
		for j, val := range row {
			ret[i] += val * o[j]
		}
	}
	return ret
}

// Transpose a 4x4 matrix
func (m Matrix4) Transpose() Matrix4 {
	ret := NewMatrix4()
	for i, row := range m {
		for j, val := range row {
			ret[j][i] = val
		}
	}
	return ret
}

// Determinant compute the Determinant for a 2x2 matrix
func (m Matrix2) Determinant() float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

// SubMatrix extract the 2x2 sub-matrix from a 3x3 matrix by removing row r and column c
func (m Matrix3) SubMatrix(r int, c int) Matrix2 {
	res := NewMatrix2()
	var x int = 0
	var y int = 0

	for i, row := range m {

		if i != r {
			y = 0
			for j, val := range row {
				if j != c {
					res[x][y] = val
					y += 1
				}
			}
			x += 1
		}
	}
	return res
}

// SubMatrix extract the 3x3 sub-matrix from a 4x4 matrix by removing row r and column c
func (m Matrix4) SubMatrix(r int, c int) Matrix3 {
	res := NewMatrix3()
	var x int = 0
	var y int = 0

	for i, row := range m {

		if i != r {
			y = 0
			for j, val := range row {
				if j != c {
					res[x][y] = val
					y += 1
				}
			}
			x += 1
		}
	}
	return res
}

// Minor compute the Minor of a 3x3 matrix for row i and col j (i.e compute the determinant of the i,j 2x2 sub matrix)
func (m Matrix3) Minor(i, j int) float64 {
	subMatrix := m.SubMatrix(i, j)
	return subMatrix.Determinant()
}

// Minor compute the Minor of a 4x4 matrix for row i and col j (i.e compute the determinant of the i,j 2x2 sub matrix)
func (m Matrix4) Minor(i, j int) float64 {
	subMatrix := m.SubMatrix(i, j)
	return subMatrix.Determinant()
}

// Cofactor computes the cofactor for a 3x3 matrix
func (m Matrix3) Cofactor(i, j int) float64 {
	res := m.Minor(i, j)

	// if i+j is even, return minor, else return - minor
	if (i+j)%2 != 0 {
		return -res
	} else {
		return res
	}

}

// Cofactor computes the cofactor for a 4x4 matrix
func (m Matrix4) Cofactor(i, j int) float64 {
	res := m.Minor(i, j)

	// if i+j is even, return minor, else return - minor
	if (i+j)%2 != 0 {
		return -res
	} else {
		return res
	}

}

// Determinant compute the Determinant for a 3x3 matrix
func (m Matrix3) Determinant() float64 {
	var res float64

	for j, val := range m[0] {
		res += val * m.Cofactor(0, j)
	}

	return res
}

// Determinant compute the Determinant for a 4x4 matrix
func (m Matrix4) Determinant() float64 {
	var res float64

	for j, val := range m[0] {
		res += val * m.Cofactor(0, j)
	}

	return res
}

// IsInvertible return true if the matrix is invertible
func (m Matrix4) IsInvertible() bool {
	return m.Determinant() != 0
}

// Inverse is inverting a 4x4 matrix
func (m Matrix4) Inverse() (Matrix4, error) {
	ret := NewMatrix4()
	determinant := m.Determinant()
	if !m.IsInvertible() {
		return ret, errors.New("matrix is not invertible")
	}

	for i, row := range m {
		for j := range row {
			c := m.Cofactor(i, j)
			ret[j][i] = c / determinant
		}
	}

	return ret, nil
}

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
