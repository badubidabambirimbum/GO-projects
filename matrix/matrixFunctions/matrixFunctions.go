// Package matrix Функции, применимые к типу Matrix
package matrix

import (
	"fmt"
	"math"
	dopf "matrix/dopFunctions"
	matrix "matrix/matrixStruct"
	inter "matrix/numInterface"
	"reflect"
)

// Sum Сумма матриц
func Sum[T inter.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) (matrix.Matrix[T], error) {
	var result matrix.Matrix[T]
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		return result, fmt.Errorf("Матрицы неравны! %d x %d != %d x %d\n", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	result.SetMatrix(dopf.CreateEmptyMatrix[T](m1.Row(), m1.Column()))
	result.SetRow(m1.Row())
	result.SetColumn(m1.Column())
	result.SetMax(T(math.MinInt))
	result.SetMin(T(math.MaxInt))
	for i := 0; i < m1.Row(); i++ {
		for j := 0; j < m1.Column(); j++ {
			result.Matrix()[i][j] = m1.Matrix()[i][j] + m2.Matrix()[i][j]
			if result.Matrix()[i][j] > result.Max() {
				result.SetMax(result.Matrix()[i][j])
			}
			if result.Matrix()[i][j] < result.Min() {
				result.SetMin(result.Matrix()[i][j])
			}
		}
	}
	if m1.Row() == m1.Column() {
		result.SetDeterminant()
	}
	return result, nil
}

// Difference Разность матриц
func Difference[T inter.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) (matrix.Matrix[T], error) {
	var result matrix.Matrix[T]
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		return result, fmt.Errorf("матрицы неравны! %d x %d != %d x %d", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	result.SetMatrix(dopf.CreateEmptyMatrix[T](m1.Row(), m1.Column()))
	result.SetRow(m1.Row())
	result.SetColumn(m1.Column())
	result.SetMax(T(math.MinInt))
	result.SetMin(T(math.MaxInt))
	for i := 0; i < m1.Row(); i++ {
		for j := 0; j < m1.Column(); j++ {
			result.Matrix()[i][j] = m1.Matrix()[i][j] - m2.Matrix()[i][j]
			if result.Matrix()[i][j] > result.Max() {
				result.SetMax(result.Matrix()[i][j])
			}
			if result.Matrix()[i][j] < result.Min() {
				result.SetMin(result.Matrix()[i][j])
			}
		}
	}
	if m1.Row() == m1.Column() {
		result.SetDeterminant()
	}
	return result, nil
}

// Product Произведение матриц
func Product[T inter.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) (matrix.Matrix[T], error) {
	var result matrix.Matrix[T]
	if m1.Column() != m2.Row() {
		return result, fmt.Errorf("кол-во столбцов первой матрицы не равно кол-ву строк второй")
	}
	result.SetMatrix(dopf.CreateEmptyMatrix[T](m1.Row(), m2.Column()))
	result.SetRow(m1.Row())
	result.SetColumn(m2.Column())
	result.SetMax(T(math.MinInt))
	result.SetMin(T(math.MaxInt))
	for i := 0; i < m1.Row(); i++ {
		for j := 0; j < m2.Column(); j++ {
			var val T
			for k := 0; k < m2.Row(); k++ {
				val += m1.Matrix()[i][k] * m2.Matrix()[k][j]
			}
			result.Matrix()[i][j] = val
			if result.Matrix()[i][j] > result.Max() {
				result.SetMax(result.Matrix()[i][j])
			}
			if result.Matrix()[i][j] < result.Min() {
				result.SetMin(result.Matrix()[i][j])
			}
		}
	}
	if result.Row() == result.Column() {
		result.SetDeterminant()
	}
	return result, nil
}

// MultiplicationByNumber Произведение матрицы на число
func MultiplicationByNumber[T inter.Number](m matrix.Matrix[T], val T) (matrix.Matrix[T], error) {
	var result matrix.Matrix[T]
	result.SetMatrix(dopf.CreateEmptyMatrix[T](m.Row(), m.Column()))
	result.SetRow(m.Row())
	result.SetColumn(m.Column())
	result.SetMax(T(math.MinInt))
	result.SetMin(T(math.MaxInt))
	for i := 0; i < m.Row(); i++ {
		for j := 0; j < m.Column(); j++ {
			result.Matrix()[i][j] = m.Matrix()[i][j] * val
			if result.Matrix()[i][j] > result.Max() {
				result.SetMax(result.Matrix()[i][j])
			}
			if result.Matrix()[i][j] < result.Min() {
				result.SetMin(result.Matrix()[i][j])
			}
		}
	}
	if result.Row() == result.Column() {
		result.SetDeterminant()
	}
	return result, nil
}

// Transposition Транспозиция матрицы
func Transposition[T inter.Number](m matrix.Matrix[T]) (matrix.Matrix[T], error) {
	if m.Column() != m.Row() {
		return m, fmt.Errorf("матрица должна быть квадратной")
	}
	for i := 0; i < m.Row(); i++ {
		for j := i + 1; j < m.Column(); j++ {
			m.Matrix()[i][j], m.Matrix()[j][i] = m.Matrix()[j][i], m.Matrix()[i][j]
		}
	}
	return m, nil
}

// Inverse Обратная матрица
func Inverse[T inter.Number](m matrix.Matrix[T]) (matrix.Matrix[T], error) {
	var result matrix.Matrix[T]
	if m.Determinant() == 0 {
		return result, fmt.Errorf("определитель равен нулю")
	}
	if reflect.TypeOf(*new(T)).Kind() == reflect.Int {
		return result, fmt.Errorf("обратная матрица для типа int находиться некорректно")
	}
	result = CreateMatrix(dopf.SearchMatrixAlgAdditions(m.Matrix()))
	result, err := Transposition(result)
	if err != nil {
		return result, err
	}
	result, err = MultiplicationByNumber(result, 1/T(math.Abs(float64(m.Determinant()))))
	if err != nil {
		return result, err
	}
	return result, nil
}

// CreateMatrix Создание матрицы
func CreateMatrix[T inter.Number](values [][]T) matrix.Matrix[T] {
	result := matrix.Matrix[T]{}
	result.SetMatrix(values)
	result.SetRow(len(values))
	result.SetColumn(len(values[0]))
	result.SetDeterminant()
	result.SetMax()
	result.SetMin()

	return result
}
