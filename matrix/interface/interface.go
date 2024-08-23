package matrix

import (
	"fmt"
	"log"
	matrixFloat "matrix/matrixFloat"
	matrixInt "matrix/matrixInt"
)

type Matrix interface {
	Create()
	SetMatrix(i, j int, val any)
	SetDeterminant()
	Determinant() any
	SetRow(int)
	Row() int
	SetColumn(int)
	Column() int
}

// Сумма матриц
func Sum(m1 Matrix, m2 Matrix) Matrix {
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		log.Fatalf("Матрицы неравны! %d x %d != %d x %d", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	var m3 Matrix

	return m3
}

// Разность матриц
func Difference(m1 Matrix, m2 Matrix) Matrix {
	if m1.Row() != m2.Row() || m1.Column() != m2.Column() {
		log.Fatalf("Матрицы неравны! %d x %d != %d x %d", m1.Row(), m1.Column(), m2.Row(), m2.Column())
	}
	var m3 Matrix

	return m3
}

// Произведение матриц
func ProductMatrix(m1 Matrix, m2 Matrix) Matrix {
	if m1.Column() != m2.Row() {
		log.Fatal("Кол-во столбцов первой матрицы не равно кол-ву строк второй")
	}
	var m3 Matrix

	return m3
}

// Произведение матрицы на число
func MultiplicationByNumber(m Matrix, val any, result Matrix) {
	if mInt, ok := m.(*(matrixInt.MatrixInt)); ok {
		if valInt, ok := val.(int); ok {
			if resInt, ok := result.(*(matrixInt.MatrixInt)); ok {
				if resInt.Len() == 0 {
					resInt.CreateEmptyMatrix(m.Row(), m.Column())
					resInt.SetRow(m.Row())
					resInt.SetColumn(m.Column())
				}
				for i := 0; i < m.Row(); i++ {
					for j := 0; j < m.Column(); j++ {
						resInt.Matrix()[i][j] = mInt.Matrix()[i][j] * int(valInt)
					}
				}
			}
		} else {
			log.Fatalf("Значение типа %T не подходит!", val)
		}
	} else if mFloat, ok := m.(*(matrixFloat.MatrixFloat)); ok {
		if valFloat, ok := val.(float64); ok {
			if resFloat, ok := result.(*(matrixFloat.MatrixFloat)); ok {
				if resFloat.Len() == 0 {
					resFloat.CreateEmptyMatrix(m.Row(), m.Column())
					resFloat.SetRow(m.Row())
					resFloat.SetColumn(m.Column())
				}
				for i := 0; i < m.Row(); i++ {
					for j := 0; j < m.Column(); j++ {
						resFloat.Matrix()[i][j] = mFloat.Matrix()[i][j] * float64(valFloat)
					}
				}
			}
		} else {
			log.Fatalf("Значение типа %T не подходит!", val)
		}
	} else {
		log.Fatal("Матрица не подходит!")
	}
}

// Обратная матрица
func Inverse(m Matrix) Matrix {
	var m3 Matrix

	return m3
}

// Печать матрицы
func PrintMatrix(m Matrix) {
	if mInt, ok := m.(*(matrixInt.MatrixInt)); ok {
		for i := 0; i < m.Row(); i++ {
			fmt.Println(mInt.Matrix()[i])
		}
	} else if mFloat, ok := m.(*(matrixFloat.MatrixFloat)); ok {
		for i := 0; i < m.Row(); i++ {
			fmt.Println(mFloat.Matrix()[i])
		}
	}
}
