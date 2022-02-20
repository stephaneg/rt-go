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

func TestEqualMatrix4(t *testing.T) {
	m1 := NewMatrix4()
	m2 := NewMatrix4()

	if !m1.EqualMatrix(m2) {
		t.Error("expected empty matrices to be equal")
	}
	m1[0][0] = 3.0

	if m1.EqualMatrix(m2) {
		t.Error("m1 has changed and is no more equal to m2")
	}

	m1 = Identity4()
	m2 = Identity4()
	if !m1.EqualMatrix(m2) {
		t.Error("Identify are equals")
	}

}

func TestDotlMatrix4(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{1.0, 2.0, 3.0, 4.0},
		{5.0, 6.0, 7.0, 8.0},
		{9.0, 8.0, 7.0, 6.0},
		{5.0, 4.0, 3.0, 2.0},
	})

	b := NewMatrix4From([4][4]float64{
		{-2.0, 1.0, 2.0, 3.0},
		{3.0, 2.0, 1.0, -1.0},
		{4.0, 3.0, 6.0, 5.0},
		{1.0, 2.0, 7.0, 8.0},
	})

	expected := NewMatrix4From([4][4]float64{
		{20.0, 22.0, 50.0, 48.0},
		{44.0, 54.0, 114.0, 108.0},
		{40.0, 58.0, 110.0, 102.0},
		{16.0, 26.0, 46.0, 42.0},
	})

	res := a.DotMatrix(b)

	if !res.EqualMatrix(expected) {
		t.Error("error in matrix dot matrix ")
	}
}

func TestMatrixDotTuple(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{1.0, 2.0, 3.0, 4.0},
		{2.0, 4.0, 4.0, 2.0},
		{8.0, 6.0, 4.0, 1.0},
		{0.0, 0.0, 0.0, 1.0},
	})

	tuple := NewTuple(1.0, 2.0, 3.0, 1.0)
	res := a.DotTuple(tuple)

	expected := NewTuple(18.0, 24.0, 33.0, 1.0)

	if !res.EqualTuple(expected) {
		t.Errorf("error in matrix dot tuple : %v ", res)
	}

}

func TestTransposeMatrix4(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{0.0, 9.0, 3.0, 0.0},
		{9.0, 8.0, 0.0, 8.0},
		{1.0, 8.0, 5.0, 3.0},
		{0.0, 0.0, 5.0, 8.0},
	})

	expected := NewMatrix4From([4][4]float64{
		{0.0, 9.0, 1.0, 0.0},
		{9.0, 8.0, 8.0, 0.0},
		{3.0, 0.0, 5.0, 5.0},
		{0.0, 8.0, 3.0, 8.0},
	})

	res := a.Transpose()

	if !res.EqualMatrix(expected) {
		t.Errorf("error in matrix transposition")
	}

}

func TestDeterminantMatrix2(t *testing.T) {
	m := NewMatrix2From([2][2]float64{
		{1.0, 5.0},
		{-3.0, 2.0},
	})

	result := m.Determinant()
	expected := 17.0

	if !FuzzyEqf64(result, expected) {
		t.Errorf("error in computing 2x2 determinant, expecting %v, get %v", expected, result)
	}

}

func TestSubMatrix3(t *testing.T) {
	m := NewMatrix3From([3][3]float64{
		{1.0, 5.0, 0.0},
		{-3.0, 2.0, 7.0},
		{0.0, 6.0, -3.0},
	})

	expected := NewMatrix2From([2][2]float64{
		{-3.0, 2.0},
		{0.0, 6.0},
	})

	result := m.SubMatrix(0, 2)

	if !expected.EqualMatrix(result) {
		t.Error("error in extracting the submatrix from a 3x3 matrix")
	}

}

func TestMinorMatrix3(t *testing.T) {

	m := NewMatrix3From([3][3]float64{
		{3.0, 5.0, 0.0},
		{2.0, -1.0, 7.0},
		{6.0, -1.0, 5.0},
	})

	result := m.Minor(1, 0)
	expected := 25.0

	if !FuzzyEqf64(result, expected) {
		t.Errorf("error in computing minor, expecting %v, get %v", expected, result)
	}

}

func TestCofactorMatrix3(t *testing.T) {

	m := NewMatrix3From([3][3]float64{
		{3.0, 5.0, 0.0},
		{2.0, -1.0, -7.0},
		{6.0, -1.0, 5.0},
	})
	var res float64
	res = m.Cofactor(0, 0)

	if !FuzzyEqf64(res, -12.0) {
		t.Errorf("error in computing the cofactor for 0,0, expected -12, get %v", res)
	}

	res = m.Cofactor(1, 0)

	if !FuzzyEqf64(res, -25.0) {
		t.Errorf("error in computing the cofactor for 1,0, expected -25, get %v", res)
	}
}

func TestDeterminantMatrix3(t *testing.T) {
	m := NewMatrix3From([3][3]float64{
		{1.0, 2.0, 6.0},
		{-5.0, 8.0, -4.0},
		{2.0, 6.0, 4.0},
	})

	expectedDeterminant := -196.0
	expectedCofactor00 := 56.0
	expectedCofactor01 := 12.0
	expectedCofactor02 := -46.0

	resDeterminant := m.Determinant()
	resCofactor00 := m.Cofactor(0, 0)
	resCofactor01 := m.Cofactor(0, 1)
	resCofactor02 := m.Cofactor(0, 2)

	if !FuzzyEqf64(expectedCofactor00, resCofactor00) {
		t.Errorf("error in computing cofactor in 0 0, expecting %v, get %v", expectedCofactor00, resCofactor00)
	}

	if !FuzzyEqf64(expectedCofactor01, resCofactor01) {
		t.Errorf("error in computing cofactor in 0 1, expecting %v, get %v", expectedCofactor01, resCofactor01)
	}

	if !FuzzyEqf64(expectedCofactor02, resCofactor02) {
		t.Errorf("error in computing cofactor in 0 2, expecting %v, get %v", expectedCofactor02, resCofactor02)
	}

	if !FuzzyEqf64(expectedDeterminant, resDeterminant) {
		t.Errorf("error in computing 3x3 determinant, expecting %v, get %v", expectedDeterminant, resDeterminant)
	}

}

func TestDeterminantMatrix4(t *testing.T) {
	m := NewMatrix4From([4][4]float64{
		{-2.0, -8.0, 3.0, 5.0},
		{-3.0, 1.0, 7.0, 3.0},
		{1.0, 2.0, -9.0, 6.0},
		{-6.0, 7.0, 7.0, -9.0},
	})

	expectedDeterminant := -4071.0
	expectedCofactor00 := 690.0
	expectedCofactor01 := 447.0
	expectedCofactor02 := 210.0
	expectedCofactor03 := 51.0

	resDeterminant := m.Determinant()
	resCofactor00 := m.Cofactor(0, 0)
	resCofactor01 := m.Cofactor(0, 1)
	resCofactor02 := m.Cofactor(0, 2)
	resCofactor03 := m.Cofactor(0, 3)

	if !FuzzyEqf64(expectedCofactor00, resCofactor00) {
		t.Errorf("error in computing cofactor in 0 0, expecting %v, get %v", expectedCofactor00, resCofactor00)
	}

	if !FuzzyEqf64(expectedCofactor01, resCofactor01) {
		t.Errorf("error in computing cofactor in 0 1 , expecting %v, get %v", expectedCofactor01, resCofactor01)
	}

	if !FuzzyEqf64(expectedCofactor02, resCofactor02) {
		t.Errorf("error in computing cofactor in 0 2, expecting %v, get %v", expectedCofactor02, resCofactor02)
	}

	if !FuzzyEqf64(expectedCofactor03, resCofactor03) {
		t.Errorf("error in computing cofactor in 0 3, expecting %v, get %v", expectedCofactor03, resCofactor03)
	}

	if !FuzzyEqf64(expectedDeterminant, resDeterminant) {
		t.Errorf("error in computing 4x4 determinant, expecting %v, get %v", expectedDeterminant, resDeterminant)
	}

}

func TestInvertibleMatrix4(t *testing.T) {
	m := NewMatrix4From([4][4]float64{
		{6.0, 4.0, 4.0, 4.0},
		{5.0, 5.0, 7.0, 6.0},
		{4.0, -9.0, 3.0, -7.0},
		{9.0, 1.0, 7.0, -6.0},
	})

	expectedDeterminant := -2120.0
	expectedIsInvertible := true

	resDeterminant := m.Determinant()
	resIsInvertible := m.IsInvertible()

	if !FuzzyEqf64(expectedDeterminant, resDeterminant) {
		t.Errorf("error in computing determinant, expecting %v, get %v", expectedDeterminant, resDeterminant)
	}

	if resIsInvertible != expectedIsInvertible {
		t.Errorf("error in computing IsInvertible expecting %v, get %v", expectedIsInvertible, resIsInvertible)
	}

}

func TestNonInvertibleMatrix4(t *testing.T) {
	m := NewMatrix4From([4][4]float64{
		{-4.0, 2.0, -2.0, -3.0},
		{9.0, 6.0, 2.0, 6.0},
		{0.0, -5.0, 1.0, -5.0},
		{0.0, 0.0, 0.0, 0.0},
	})

	expectedDeterminant := 0.0
	expectedIsInvertible := false

	resDeterminant := m.Determinant()
	resIsInvertible := m.IsInvertible()

	if !FuzzyEqf64(expectedDeterminant, resDeterminant) {
		t.Errorf("error in computing determinant, expecting %v, get %v", expectedDeterminant, resDeterminant)
	}

	if resIsInvertible != expectedIsInvertible {
		t.Errorf("error in computing IsInvertible expecting %v, get %v", expectedIsInvertible, resIsInvertible)
	}

}

func TestInverseMatrix4(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{-5.0, 2.0, 6.0, -8.0},
		{1.0, -5.0, 1.0, 8.0},
		{7.0, 7.0, -6.0, -7.0},
		{1.0, -3.0, 7.0, 4.0},
	})

	expectedDeterminant := 532.0
	expectedCofactorA23 := -160.0
	expectedCofactorA32 := 105.0
	expectedB32 := -160.0 / 532.0
	expectedB23 := 105.0 / 532.0

	b := NewMatrix4From([4][4]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	})

	resDeterminant := a.Determinant()
	resCofactorA23 := a.Cofactor(2, 3)
	resCofactorA32 := a.Cofactor(3, 2)

	res, err := a.Inverse()
	resB32 := res[3][2]
	resB23 := res[2][3]

	if err != nil {
		t.Error("this matrix should be invertible")
	}

	if !FuzzyEqf64(expectedDeterminant, resDeterminant) {
		t.Errorf("error in computing determinant, expecting %v, get %v", expectedDeterminant, resDeterminant)
	}

	if !FuzzyEqf64(expectedCofactorA23, resCofactorA23) {
		t.Errorf("error in computing cofactor a in 2 3, expecting %v, get %v", expectedCofactorA23, resCofactorA23)
	}

	if !FuzzyEqf64(expectedCofactorA32, resCofactorA32) {
		t.Errorf("error in computing cofactor a in 3 2, expecting %v, get %v", expectedCofactorA32, resCofactorA32)
	}

	if !FuzzyEqf64(expectedB32, resB32) {
		t.Errorf("error in comparing b 3 2, expecting %v, get %v", expectedB32, resB32)
	}

	if !FuzzyEqf64(expectedB23, resB23) {
		t.Errorf("error in comparing b 2 3, expecting %v, get %v", expectedB23, resB23)
	}

	if b.EqualMatrix(res) {
		t.Errorf("inverse matrice %v is not equal to the expected one %v ", res, b)
	}

}

func TestInverseAnotherMatrix4(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{8.0, -5.0, 9.0, 2.0},
		{7.0, 5.0, 6.0, 1.0},
		{-6.0, 0.0, 9.0, 6.0},
		{-3.0, 0.0, -9.0, -4.0},
	})

	b := NewMatrix4From([4][4]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	})

	res, err := a.Inverse()

	if err != nil {
		t.Error("this matrix should be invertible")
	}

	if b.EqualMatrix(res) {
		t.Errorf("inverse matrice %v is not equal to the expected one %v ", res, b)
	}

}

func TestInverseThirdMatrix4(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{9.0, 3.0, 0.0, 9.0},
		{-5.0, -2.0, -6.0, -3.0},
		{-4.0, 9.0, 6.0, 4.0},
		{-7.0, 6.0, 6.0, 2.0},
	})

	b := NewMatrix4From([4][4]float64{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	})

	res, err := a.Inverse()

	if err != nil {
		t.Error("this matrix should be invertible")
	}

	if b.EqualMatrix(res) {
		t.Errorf("inverse matrice %v is not equal to the expected one %v ", res, b)
	}

}

func TestMultiplyMatrix4ByInverse(t *testing.T) {
	a := NewMatrix4From([4][4]float64{
		{3.0, -9.0, 7.0, 3.0},
		{3.0, -8.0, 2.0, -9.0},
		{-4.0, 4.0, 4.0, 1.0},
		{-6.0, 5.0, -1.0, 1.0},
	})

	b := NewMatrix4From([4][4]float64{
		{8.0, 2.0, 2.0, 2.0},
		{3.0, -1.0, 7.0, 0.0},
		{7.0, 0.0, 5.0, 4.0},
		{6.0, -2.0, 0.0, 5.0},
	})

	c := a.DotMatrix(b)
	inv, _ := b.Inverse()
	d := c.DotMatrix(inv)

	if !a.EqualMatrix(d) {
		t.Error("A * B * inverse(B) should be equal to A")
	}
}

func TestMultiplyByTranslationMatrix(t *testing.T) {
	m := Translation(5.0, -3.0, 2.0)

	v := NewVector(-3, 4, 5)

	res := m.DotTuple(v)

	if !res.EqualTuple(v) {
		t.Error("translation matrix x vector should be equals to vector")
	}
}

func TestMultiplyPointByTranslationMatrix(t *testing.T) {
	m := Translation(5.0, -3.0, 2.0)
	v := NewPoint(-3, 4, 5)
	expected := NewPoint(2, 1, 7)

	res := m.DotTuple(v)

	if !res.EqualTuple(expected) {
		t.Error("translation matrix x point should be equals to a point")
	}
}

func TestMultiplyPointByInvTranslationMatrix(t *testing.T) {

	m := Translation(5.0, -3.0, 2.0)
	inv, _ := m.Inverse()

	v := NewPoint(-3, 4, 5)
	expected := NewPoint(-8, 7, 3)

	res := inv.DotTuple(v)

	if !res.EqualTuple(expected) {
		t.Error("translation matrix x point should be equals to a point")
	}
}
