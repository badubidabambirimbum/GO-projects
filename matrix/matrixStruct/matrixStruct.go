// Тип Matrix и его методы
package matrix

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Ограничения на тип
type Number interface {
	int | float64
}

// Структура Matrix, которая работает с типами из интерфейса Number
type Matrix[T Number] struct {
	matrix      [][]T // матрица
	determinant T     // определитель
	row         int   // ряд
	column      int   // столбец
}

// Создание матрицы
func (m *Matrix[T]) Create() {
	fmt.Print("Введите количество строк и столбцов: ")
	fmt.Scan(&(m.row), &(m.column))
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

// Создание матрицы из нулей
func (m *Matrix[T]) CreateEmptyMatrix(row, col int) {
	if len(m.matrix) != 0 {
		log.Fatal("Матрица уже создана!")
	}
	for i := 0; i < row; i++ {
		m.matrix = append(m.matrix, []T{})
		for j := 0; j < col; j++ {
			m.matrix[i] = append(m.matrix[i], 0)
		}
	}
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
	m.determinant = recursSearchDet(m.matrix)
}

// Вовращение детерминанта
func (m *Matrix[T]) Determinant() T {
	return m.determinant
}

// Вспомогательная функция для поиска определителя
func recursSearchDet[T Number](m [][]T) T {
	var det T
	det = 0

	if len(m) == 2 {
		det = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	} else {
		for i := 0; i < len(m); i++ {
			mRec := [][]T{}
			for j := 0; j < len(m); j++ {
				if i != j {
					mRec = append(mRec, m[j][1:])
				}
			}
			det += T(math.Pow(-1, float64(i))) * m[i][0] * recursSearchDet(mRec)
		}
	}
	return det
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
