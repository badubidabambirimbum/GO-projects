// Матричный калькулятор
package main

import (
	"bufio"
	"fmt"
	mainfunct "matrix/mainFunctions"
	funct "matrix/matrixFunctions"
	matrix "matrix/matrixStruct"
	"os"
	"strings"
)

const text = `
		***Указатель***
+-----------------------------------------------------+
|       | type - int/float  (int по умолчанию)        |
| info  | num1/num2 - Номер матрицы в списке          |
|       | val - Множитель матрицы                     |
+-----------------------------------------------------+
|   Ввод              | Описание                      |
+---------------------+-------------------------------+
|   0                 | Вывод указателя               |
|   1 [    type    ]  | Выбор типа матрицы            |
|   2 [            ]  | Список доступных матриц       |
|   3 [            ]  | Добавить матрицу              |
|   4 [  num1 num2 ]  | Сумма матриц                  |
|   5 [  num1 num2 ]  | Разность матриц               |
|   6 [  num1 num2 ]  | Произведение матриц           |
|   7 [  num1 val  ]  | Произведение матрицы на число |
|   8 [    num1    ]  | Транспонирование матрицы      |
|   9 [    num1    ]  | Обратная матрица              |
|  10                 | Выход                         |
+-----------------------------------------------------+
`

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	matrixIntMap := map[int]matrix.Matrix[int]{}
	matrixFloatMap := map[int]matrix.Matrix[float64]{}
	matrixType := "int"
	fmt.Print(text)
	for {
		fmt.Print("\nВведите число: ")
		if scanner.Scan() {
			input := scanner.Text()
			switch {

			case input == "0":
				fmt.Print(text)

			case input == "1 int":
				matrixType = "int"
			case input == "1 float":
				matrixType = "float"

			case input == "2":
				if matrixType == "int" {
					mainfunct.PrintMatrix(matrixIntMap)
				} else {
					mainfunct.PrintMatrix(matrixFloatMap)
				}

			case input == "3":
				if matrixType == "int" {
					mainfunct.AddMatrix(&matrixIntMap)
				} else {
					mainfunct.AddMatrix(&matrixFloatMap)
				}
			case strings.HasPrefix(input, "4"):
				if matrixType == "int" {
					mainfunct.OperationBetweenMatrices(strings.Split(input, " "), &matrixIntMap, funct.Sum)
				} else {
					mainfunct.OperationBetweenMatrices(strings.Split(input, " "), &matrixFloatMap, funct.Sum)
				}

			case strings.HasPrefix(input, "5"):
				if matrixType == "int" {
					mainfunct.OperationBetweenMatrices(strings.Split(input, " "), &matrixIntMap, funct.Difference)
				} else {
					mainfunct.OperationBetweenMatrices(strings.Split(input, " "), &matrixFloatMap, funct.Difference)
				}

			case strings.HasPrefix(input, "6"):
				if matrixType == "int" {
					mainfunct.OperationBetweenMatrices(strings.Split(input, " "), &matrixIntMap, funct.Product)
				} else {
					mainfunct.OperationBetweenMatrices(strings.Split(input, " "), &matrixFloatMap, funct.Product)
				}

			case strings.HasPrefix(input, "7"):
				if matrixType == "int" {
					mainfunct.OperationMatrixByNubmer(strings.Split(input, " "), &matrixIntMap, funct.MultiplicationByNumber)
				} else {
					mainfunct.OperationMatrixByNubmer(strings.Split(input, " "), &matrixFloatMap, funct.MultiplicationByNumber)
				}

			case strings.HasPrefix(input, "8"):
				if matrixType == "int" {
					mainfunct.OperationMatrix(strings.Split(input, " "), &matrixIntMap, funct.Transposition)
				} else {
					mainfunct.OperationMatrix(strings.Split(input, " "), &matrixFloatMap, funct.Transposition)
				}

			case strings.HasPrefix(input, "9"):
				if matrixType == "int" {
					mainfunct.OperationMatrix(strings.Split(input, " "), &matrixIntMap, funct.Inverse)
				} else {
					mainfunct.OperationMatrix(strings.Split(input, " "), &matrixFloatMap, funct.Inverse)
				}

			case input == "10":
				println("Выход из программы!\n")
				os.Exit(0)

			default:
				fmt.Println("Некорректный ввод!")
			}
		}
		fmt.Print("+-----------------------------------------------------+")
	}
}
