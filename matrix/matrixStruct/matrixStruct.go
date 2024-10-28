// Package matrix Тип Matrix и его методы
package matrix

import (
	"bufio"
	"fmt"
	"log"
	"math"
	dopf "matrix/dopFunctions"
	inter "matrix/numInterface"
	"os"
	"strconv"
	"strings"
)

// Matrix Структура Matrix, которая работает с типами из интерфейса Number
type Matrix[T inter.Number] struct {
	matrix      [][]T // матрица
	determinant T     // определитель
	row         int   // ряд
	column      int   // столбец
	max         T     // Максимальное значение
	min         T     // Минимальное значение
}

// Create Создание матрицы
func (m *Matrix[T]) Create() {
	fmt.Print("Введите количество строк и столбцов: ")
	fmt.Scanln(&(m.row), &(m.column))
	fmt.Print("Введите матрицу: ")

	// Создаем новый сканер
	scanner := bufio.NewScanner(os.Stdin)

	// Считываем строку с консоли
	if scanner.Scan() {

		// Получаем строку и разбиваем её на части
		input := scanner.Text()
		parts := strings.Split(input, " ")

		if len(parts) != m.row*m.column {
			log.Fatalf("Вы задали размер матрицы %d x %d, а ввели %d", m.row, m.column, len(parts))
		}
		m.max = T(math.MinInt)
		m.min = T(math.MaxInt)
		for i := 0; i < m.row; i++ {
			m.matrix = append(m.matrix, []T{})
			for j := 0; j < m.column; j++ {
				switch any(m.determinant).(type) {
				case int:
					num, err := strconv.Atoi(parts[i*m.column+j])
					if err != nil {
						fmt.Println("Ошибка преобразования:", err)
						return
					}
					m.matrix[i] = append(m.matrix[i], T(num))
					if T(num) > m.max {
						m.max = T(num)
					}
					if T(num) < m.min {
						m.min = T(num)
					}
				case float64:
					num, err := strconv.ParseFloat(parts[i*m.column+j], 64)
					if err != nil {
						fmt.Println("Ошибка преобразования:", err)
						return
					}
					m.matrix[i] = append(m.matrix[i], T(num))
					if T(num) > m.max {
						m.max = T(num)
					}
					if T(num) < m.min {
						m.min = T(num)
					}
				default:
					fmt.Println("Ошибка типа")
					return
				}
			}
		}
	}
	m.SetDeterminant()
}

// Len Длина матрицы (row * col)
func (m *Matrix[T]) Len() int {
	return m.row * m.column
}

// SetMatrix Изменение значения матрицы ?
func (m *Matrix[T]) SetMatrix(values [][]T) {
	m.matrix = values
}

// Matrix Возвращение матрицы
func (m *Matrix[T]) Matrix() [][]T {
	return m.matrix
}

// SetDeterminant Определение детерминанта
func (m *Matrix[T]) SetDeterminant() {
	m.determinant = dopf.RecursSearchDet(m.matrix)
}

// Determinant Возвращение детерминанта
func (m *Matrix[T]) Determinant() T {
	return m.determinant
}

// SetRow Определение кол-во строк
func (m *Matrix[T]) SetRow(x int) {
	m.row = x
}

// Row Возвращение кол-во строк
func (m *Matrix[T]) Row() int {
	return m.row
}

// SetColumn Определение кол-во столбцов
func (m *Matrix[T]) SetColumn(x int) {
	m.column = x
}

// Column Возвращение кол-во столбцов
func (m *Matrix[T]) Column() int {
	return m.column
}

// SetMax Определение максимального значения матрицы
func (m *Matrix[T]) SetMax(val ...T) {
	if len(val) == 0 {
		maxVal := T(math.MinInt)
		for i := 0; i < m.row; i++ {
			for j := 0; j < m.column; j++ {
				if m.matrix[i][j] > maxVal {
					maxVal = m.matrix[i][j]
				}
			}
		}
		m.max = maxVal
	} else {
		m.max = val[0]
	}
}

// Max Возвращение максимального значения матрицы
func (m *Matrix[T]) Max() T {
	return m.max
}

// SetMin Определение минимального значения матрицы
func (m *Matrix[T]) SetMin(val ...T) {
	if len(val) == 0 {
		minVal := T(math.MaxInt)
		for i := 0; i < m.row; i++ {
			for j := 0; j < m.column; j++ {
				if m.matrix[i][j] < minVal {
					minVal = m.matrix[i][j]
				}
			}
		}
		m.min = minVal
	} else {
		m.min = val[0]
	}
}

// Min Возвращение минимального значения матрицы
func (m *Matrix[T]) Min() T {
	return m.min
}

// Печать структуры
func (m *Matrix[T]) String() string {
	var s string = fmt.Sprintf("Матрица %dx%d", m.row, m.column)
	precision := 0
	width := int(math.Max(float64(len(fmt.Sprintf("%.0f", float64(m.Max())))), float64(len(fmt.Sprintf("%.0f", float64(m.Min())))))) + 2
	if _, ok := any(m.Matrix()[0][0]).(float64); ok {
		precision = 2
	}
	for i := 0; i < m.Row(); i++ {
		s += "\n["
		for j := 0; j < m.Column(); j++ {
			s += fmt.Sprintf(" %*.*f ", width+precision, precision, float64(m.Matrix()[i][j]))
		}
		s += " ]"
	}
	s += fmt.Sprintf("\nОпределитель: %0.3f\n", float64(m.determinant))
	return s
}
