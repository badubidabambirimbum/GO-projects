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
	var m3 matrix.Matrix[T]

	return m3
}

// Разность матриц
func Difference[T matrix.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) matrix.Matrix[T] {
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		log.Fatalf("Матрицы неравны! %d x %d != %d x %d", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	var m3 matrix.Matrix[T]

	return m3
}

// Произведение матриц
func ProductMatrix[T matrix.Number](m1 matrix.Matrix[T], m2 matrix.Matrix[T]) matrix.Matrix[T] {
	if m1.Column() != m2.Row() {
		log.Fatal("Кол-во столбцов первой матрицы не равно кол-ву строк второй")
	}
	var m3 matrix.Matrix[T]

	return m3
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
