package matrix

import (
	"bufio"
	"fmt"
	"log"
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
					num, err := strconv.Atoi(parts[i*m.row+j])
					if err != nil {
						fmt.Println("Ошибка преобразования:", err)
						return
					}
					m.matrix[i] = append(m.matrix[i], T(num))
				case float64:
					num, err := strconv.ParseFloat(parts[i*m.row+j], 64)
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

func (m *Matrix[T]) SetMatrix(i, j int, val T) {
	if i < 0 || i >= m.row || j < 0 || j >= m.column {
		log.Fatalf("Размер матрицы %d x %d, а вы хотите изменить значение [%d,%d]", m.row, m.column, i, j)
	}
	valT, ok := any(val).(T)
	if !ok {
		log.Fatalf("Значение типа %T не подходит!", val)
	}
	m.matrix[i][j] = valT
}

func (m *Matrix[T]) Matrix() [][]T {
	return (*m).matrix
}

func (m *Matrix[T]) SetDeterminant() {

}

func (m *Matrix[T]) Determinant() T {
	return (*m).determinant
}

func (m *Matrix[T]) SetRow(x int) {
	m.row = x
}

func (m *Matrix[T]) Row() int {
	return (*m).row
}

func (m *Matrix[T]) SetColumn(x int) {
	m.column = x
}

func (m *Matrix[T]) Column() int {
	return (*m).column
}