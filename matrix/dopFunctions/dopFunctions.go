package matrix

import (
	"math"
	inter "matrix/numInterface"
)

// RecursSearchDet Поиск определителя
func RecursSearchDet[T inter.Number](m [][]T) T {
	var det T
	det = 0

	if len(m) == 1 {
		det = m[0][0]
	} else if len(m) == 2 {
		det = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	} else {
		for i := 0; i < len(m); i++ {
			var mRec [][]T
			for j := 0; j < len(m); j++ {
				if i != j {
					mRec = append(mRec, m[j][1:])
				}
			}
			det += T(math.Pow(-1, float64(i))) * m[i][0] * RecursSearchDet(mRec)
		}
	}
	return det
}

// SearchMatrixAlgAdditions Составление матрицы алгебраических дополнений
func SearchMatrixAlgAdditions[T inter.Number](m [][]T) [][]T {
	resMatrix := CreateEmptyMatrix[T](len(m), len(m[0]))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			resMatrix[i][j] = T(math.Pow(-1, float64(i+j))) * RecursSearchDet(getMinor(m, i, j))
		}
	}
	return resMatrix
}

// Получение минора
func getMinor[T inter.Number](m [][]T, row, col int) [][]T {
	size := len(m)
	minor := make([][]T, size-1)

	for i := 0; i < size; i++ {
		if i == row {
			continue
		}
		var newRow []T
		for j := 0; j < size; j++ {
			if j == col {
				continue
			}
			newRow = append(newRow, m[i][j])
		}
		if i < row {
			minor[i] = newRow
		} else {
			minor[i-1] = newRow
		}
	}
	return minor
}

// CreateEmptyMatrix Создание матрицы из нулей
func CreateEmptyMatrix[T inter.Number](row, col int) [][]T {
	var m [][]T
	for i := 0; i < row; i++ {
		m = append(m, []T{})
		for j := 0; j < col; j++ {
			m[i] = append(m[i], 0)
		}
	}
	return m
}
