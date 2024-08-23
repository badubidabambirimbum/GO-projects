package matrix

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MatrixInt struct {
	matrix      [][]int // матрица
	determinant int     // определитель
	row         int     // ряд
	column      int     // столбец
}

// Создание матрицы
func (m *MatrixInt) Create() {
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
			m.matrix = append(m.matrix, []int{})
			for j := 0; j < m.column; j++ {
				num, err := strconv.Atoi(parts[i*m.row+j])
				if err != nil {
					fmt.Println("Ошибка преобразования:", err)
					return
				}
				m.matrix[i] = append(m.matrix[i], num)
			}
		}
	}
}

func (m *MatrixInt) SetMatrix(i, j int, val any) {
	if i < 0 || i >= m.row || j < 0 || j >= m.column {
		log.Fatalf("Размер матрицы %d x %d, а вы хотите изменить значение [%d,%d]", m.row, m.column, i, j)
	}
	valInt, ok := val.(int)
	if !ok {
		log.Fatalf("Значение типа %T не подходит!", val)
	}
	m.matrix[i][j] = int(valInt)
}

func (m *MatrixInt) Matrix() [][]int {
	return (*m).matrix
}

func (m *MatrixInt) SetDeterminant() {

}

func (m *MatrixInt) Determinant() any {
	return (*m).determinant
}

func (m *MatrixInt) Row() int {
	return (*m).row
}

func (m *MatrixInt) Column() int {
	return (*m).column
}
