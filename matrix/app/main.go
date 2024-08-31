package main

import (
	"bufio"
	"fmt"
	funct "matrix/matrixFunctions"
	matrix "matrix/matrixStruct"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	matrixIntList := map[int]matrix.Matrix[int]{}
	matrixFloatList := map[int]matrix.Matrix[float64]{}
	text := fmt.Sprint(
		"Указатель\n",
		"0. Вывод указателя\n",
		"1. Список доступных матриц (int или float)\n",
		"2. Добавить матрицу (int или float)\n",
		"3. Сложить матрицы")
	fmt.Println(text)
	for {
		fmt.Print("Введите число: ")
		if scanner.Scan() {
			input := scanner.Text()
			switch input {
			case "0":
				fmt.Println(text)
			case "1 int":
				for key, value := range matrixIntList {
					fmt.Println(key)
					fmt.Println(&value)
				}
			case "1 float":
				for key, value := range matrixFloatList {
					fmt.Println(key)
					fmt.Println(&value)
				}
			case "2 int":
				m := matrix.Matrix[int]{}
				m.Create()
				n := len(matrixIntList)
				matrixIntList[n+1] = m
			case "2 float":
				m := matrix.Matrix[float64]{}
				m.Create()
				n := len(matrixFloatList)
				matrixFloatList[n+1] = m
			}
		}
	}
	m := matrix.Matrix[int]{}
	m.Create()
	fmt.Println(&m)
	m1 := matrix.Matrix[int]{}
	m1.Create()
	fmt.Println(&m1)
	m2 := funct.ProductMatrix(m, m1)
	fmt.Println(&m2)
}
