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

func TestMatrix4xTuple(t *testing.T) {
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
