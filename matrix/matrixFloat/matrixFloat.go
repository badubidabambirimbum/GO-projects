package matrix

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MatrixFloat struct {
	matrix      [][]float64 // матрица
	determinant float64     // определитель
	row         int         // ряд
	column      int         // столбец
}

// Создание матрицы
func (m *MatrixFloat) Create() {
	fmt.Print("Введите количество строк и столбцов: ")
	fmt.Scan(&(m.row), &(m.column))
	fmt.Print("Введите матрицу: ")

	// Создаем новый сканер
	scanner := bufio.NewScanner(os.Stdin)

	// Считываем строку с консоли
	scanner.Scan()

	if scanner.Scan() {

		// Получаем строку и разбиваем её на части
		input := scanner.Text()
		parts := strings.Split(input, " ")

		if len(parts) != m.row*m.column {
			log.Fatalf("Вы задали размер матрицы %d x %d, а ввели %d", m.row, m.column, len(parts))
		}

		for i := 0; i < m.row; i++ {
			m.matrix = append(m.matrix, []float64{})
			for j := 0; j < m.column; j++ {
				num, err := strconv.ParseFloat(parts[i*m.row+j], 64)
				if err != nil {
					fmt.Println("Ошибка преобразования:", err)
					return
				}
				m.matrix[i] = append(m.matrix[i], num)
			}
		}
	}
}

func (m *MatrixFloat) SetMatrix(i, j int, val any) {
	if i < 0 || i >= m.row || j < 0 || j >= m.column {
		log.Fatalf("Размер матрицы %d x %d, а вы хотите изменить значение [%d,%d]", m.row, m.column, i, j)
	}
	valFloat, ok := val.(float64)
	if !ok {
		log.Fatalf("Значение типа %T не подходит!", val)
	}
	m.matrix[i][j] = float64(valFloat)
}

func (m *MatrixFloat) Matrix() [][]float64 {
	return (*m).matrix
}

func (m *MatrixFloat) SetDeterminant() {

}

func (m *MatrixFloat) Determinant() any {
	return (*m).determinant
}

func (m *MatrixFloat) Row() int {
	return (*m).row
}

func (m *MatrixFloat) Column() int {
	return (*m).column
}
