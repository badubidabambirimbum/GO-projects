package matrix

import (
	"fmt"
	"log"
	matrix "matrix/matrixStruct"
)

// Сумма матриц
func Sum[T matrix.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) matrix.Matrix[T] {
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		log.Fatalf("Матрицы неравны! %d x %d != %d x %d", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	var result matrix.Matrix[T]
	result.CreateEmptyMatrix(m1.Row(), m1.Column())
	result.SetRow(m1.Row())
	result.SetColumn(m1.Column())
	for i := 0; i < m1.Row(); i++ {
		for j := 0; j < m1.Column(); j++ {
			result.Matrix()[i][j] = m1.Matrix()[i][j] + m2.Matrix()[i][j]
		}
	}
	return result
}

// Разность матриц
func Difference[T matrix.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) matrix.Matrix[T] {
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		log.Fatalf("Матрицы неравны! %d x %d != %d x %d", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	var result matrix.Matrix[T]
	result.CreateEmptyMatrix(m1.Row(), m1.Column())
	result.SetRow(m1.Row())
	result.SetColumn(m1.Column())
	for i := 0; i < m1.Row(); i++ {
		for j := 0; j < m1.Column(); j++ {
			result.Matrix()[i][j] = m1.Matrix()[i][j] - m2.Matrix()[i][j]
		}
	}
	return result
}

// Произведение матриц
func ProductMatrix[T matrix.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) matrix.Matrix[T] {
	if m1.Column() != m2.Row() {
		log.Fatal("Кол-во столбцов первой матрицы не равно кол-ву строк второй")
	}
	var result matrix.Matrix[T]
	result.CreateEmptyMatrix(m1.Row(), m2.Column())
	result.SetRow(m1.Row())
	result.SetColumn(m2.Column())
	for i := 0; i < m1.Row(); i++ {
		for j := 0; j < m2.Column(); j++ {
			var val T
			for k := 0; k < m2.Row(); k++ {
				val += m1.Matrix()[i][k] * m2.Matrix()[k][j]
			}
			result.Matrix()[i][j] = val
		}
	}
	return result
}

// Произведение матрицы на число
func MultiplicationByNumber[T matrix.Number](m matrix.Matrix[T], val T) matrix.Matrix[T] {
	var result matrix.Matrix[T]
	result.CreateEmptyMatrix(m.Row(), m.Column())
	result.SetRow(m.Row())
	result.SetColumn(m.Column())
	for i := 0; i < m.Row(); i++ {
		for j := 0; j < m.Column(); j++ {
			result.Matrix()[i][j] = m.Matrix()[i][j] * val
		}
	}
	return result
}

// Транспозиция матрицы
func Transposition[T matrix.Number](m matrix.Matrix[T]) matrix.Matrix[T] {
	if m.Column() != m.Row() {
		log.Fatal("Матрица должна быть квадратной!")
	}
	var result matrix.Matrix[T]
	for i := 0; i < m.Row(); i++ {
		for j := i + 1; j < m.Column(); j++ {
			m.Matrix()[i][j], m.Matrix()[j][i] = m.Matrix()[j][i], m.Matrix()[i][j]
		}
	}
	return result
}

// Обратная матрица
func Inverse[T matrix.Number](m matrix.Matrix[T]) matrix.Matrix[T] {
	var m3 matrix.Matrix[T]

	return m3
}

// Печать матрицы
func PrintMatrix[T matrix.Number](m matrix.Matrix[T]) {
	for i := 0; i < m.Row(); i++ {
		fmt.Println(m.Matrix()[i])
	}
}

func CreateMatrix[T matrix.Number](values [][]T) matrix.Matrix[T] {
	result := matrix.Matrix[T]{}
	result.SetMatrix(values)
	result.SetRow(len(values))
	result.SetColumn(len(values[0]))
	result.SetDeterminant()

	return result
}
