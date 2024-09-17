// Тип Matrix и его методы
package matrix

import (
	"bufio"
	"fmt"
	"log"
	dopf "matrix/dopFunctions"
	inter "matrix/numInterface"
	"os"
	"strconv"
	"strings"
)

// Структура Matrix, которая работает с типами из интерфейса Number
type Matrix[T inter.Number] struct {
	matrix      [][]T // матрица
	determinant T     // определитель
	row         int   // ряд
	column      int   // столбец
}

// Создание матрицы
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
				case float64:
					num, err := strconv.ParseFloat(parts[i*m.column+j], 64)
					if err != nil {
						fmt.Println("Ошибка преобразования:", err)
						return
					}
					m.matrix[i] = append(m.matrix[i], T(num))
				default:
					fmt.Println("Ошибка типа")
					return
				}
			}
		}
	}
	m.SetDeterminant()
}

// Длина матрицы (row * col)
func (m *Matrix[T]) Len() int {
	return m.row * m.column
}

//  Изменение значения матрицы ?
func (m *Matrix[T]) SetMatrix(values [][]T) {
	m.matrix = values
}

// Возвращение матрицы
func (m *Matrix[T]) Matrix() [][]T {
	return m.matrix
}

// Определение детерминанта
func (m *Matrix[T]) SetDeterminant() {
	m.determinant = dopf.RecursSearchDet(m.matrix)
}

// Вовращение детерминанта
func (m *Matrix[T]) Determinant() T {
	return m.determinant
}

func (m *Matrix[T]) SetRow(x int) {
	m.row = x
}

func (m *Matrix[T]) Row() int {
	return m.row
}

func (m *Matrix[T]) SetColumn(x int) {
	m.column = x
}

func (m *Matrix[T]) Column() int {
	return m.column
}

func (m *Matrix[T]) String() string {
	var s string = fmt.Sprintf("Матрица %dx%d", m.row, m.column)
	for i := 0; i < m.Row(); i++ {
		s += fmt.Sprint("\n", m.Matrix()[i])
	}
	s += fmt.Sprint("\nОпределитель: ", m.determinant, "\n")
	return s
}
