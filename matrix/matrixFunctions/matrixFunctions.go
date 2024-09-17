// Функции, применимые к типу Matrix
package matrix

import (
	"log"
	"math"
	dopf "matrix/dopFunctions"
	matrix "matrix/matrixStruct"
	inter "matrix/numInterface"
)

// Сумма матриц
func Sum[T inter.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) matrix.Matrix[T] {
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		log.Fatalf("Матрицы неравны! %d x %d != %d x %d", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	var result matrix.Matrix[T]
	result.SetMatrix(dopf.CreateEmptyMatrix[T](m1.Row(), m1.Column()))
	result.SetRow(m1.Row())
	result.SetColumn(m1.Column())
	result.SetMax(T(math.Inf(-1)))
	result.SetMin(T(math.Inf(1)))
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
	return result
}

// Разность матриц
func Difference[T inter.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) matrix.Matrix[T] {
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		log.Fatalf("Матрицы неравны! %d x %d != %d x %d", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	var result matrix.Matrix[T]
	result.SetMatrix(dopf.CreateEmptyMatrix[T](m1.Row(), m1.Column()))
	result.SetRow(m1.Row())
	result.SetColumn(m1.Column())
	result.SetMax(T(math.Inf(-1)))
	result.SetMin(T(math.Inf(1)))
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
	return result
}

// Произведение матриц
func Product[T inter.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) matrix.Matrix[T] {
	if m1.Column() != m2.Row() {
		log.Fatal("Кол-во столбцов первой матрицы не равно кол-ву строк второй")
	}
	var result matrix.Matrix[T]
	result.SetMatrix(dopf.CreateEmptyMatrix[T](m1.Row(), m2.Column()))
	result.SetRow(m1.Row())
	result.SetColumn(m2.Column())
	result.SetMax(T(math.Inf(-1)))
	result.SetMin(T(math.Inf(1)))
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
	return result
}

// Произведение матрицы на число
func MultiplicationByNumber[T inter.Number](m matrix.Matrix[T], val T) matrix.Matrix[T] {
	var result matrix.Matrix[T]
	result.SetMatrix(dopf.CreateEmptyMatrix[T](m.Row(), m.Column()))
	result.SetRow(m.Row())
	result.SetColumn(m.Column())
	result.SetMax(T(math.Inf(-1)))
	result.SetMin(T(math.Inf(1)))
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
	return result
}

// Транспозиция матрицы
func Transposition[T inter.Number](m matrix.Matrix[T]) matrix.Matrix[T] {
	if m.Column() != m.Row() {
		log.Fatal("Матрица должна быть квадратной!")
	}
	for i := 0; i < m.Row(); i++ {
		for j := i + 1; j < m.Column(); j++ {
			m.Matrix()[i][j], m.Matrix()[j][i] = m.Matrix()[j][i], m.Matrix()[i][j]
		}
	}
	return m
}

// Обратная матрица
func Inverse[T inter.Number](m matrix.Matrix[T]) matrix.Matrix[T] {
	if m.Column() != m.Row() {
		log.Fatal("Обратную матрицу можно найти только для квадрдатной матрицы")
	}
	if m.Determinant() == 0 {
		log.Fatal("Определитель равен нулю")
	}
	resMatrix := CreateMatrix(dopf.SearchMatrixAlgAdditions(m.Matrix()))
	return MultiplicationByNumber(Transposition(resMatrix), 1/T(math.Abs(float64(m.Determinant()))))
}

// Создание матрицы
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
