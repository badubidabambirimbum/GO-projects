// Функции для функционала main функции
package matrix

import (
	"fmt"
	"log"
	matrix "matrix/matrixStruct"
	"reflect"
	"strconv"
)

// Печать матриц
func PrintMatrix[T matrix.Number](matrixList map[int]matrix.Matrix[T]) {
	for key, value := range matrixList {
		fmt.Println(key)
		fmt.Println(&value)
	}
}

// Добавление матрицы в словарь
func AddMatrix[T matrix.Number](matrixList *map[int]matrix.Matrix[T]) {
	m := matrix.Matrix[T]{}
	m.Create()
	n := len(*matrixList)
	(*matrixList)[n+1] = m
}

// Операции между матрицами
func OperationBetweenMatrices[T matrix.Number](
	input_split []string,
	matrixList *map[int]matrix.Matrix[T],
	operation func(matrix.Matrix[T], matrix.Matrix[T]) matrix.Matrix[T]) {

	num1, err := strconv.Atoi(input_split[1])
	if err != nil {
		log.Fatalf("%s не переводится в int", input_split[1])
	}
	num2, err := strconv.Atoi(input_split[2])
	if err != nil {
		log.Fatalf("%s не переводится в int", input_split[2])
	}
	resultMatrix := operation((*matrixList)[num1], (*matrixList)[num2])
	n := len(*matrixList)
	(*matrixList)[n+1] = resultMatrix
	fmt.Println("\nРезультат:")
	fmt.Println(&resultMatrix)
}

// Операция умножения матрицы на число
func OperationMatrixByNubmer[T matrix.Number](
	input_split []string,
	matrixList *map[int]matrix.Matrix[T],
	operation func(matrix.Matrix[T], T) matrix.Matrix[T]) {
	var resultMatrix matrix.Matrix[T]
	num, err := strconv.Atoi(input_split[1])
	if err != nil {
		log.Fatalf("%s не переводится в int", input_split[1])
	}

	switch reflect.TypeOf(*new(T)).Kind() {
	case reflect.Int:
		val, err := strconv.Atoi(input_split[2])
		if err != nil {
			log.Fatalf("%s не переводится в int", input_split[2])
		} else {
			resultMatrix = operation((*matrixList)[num], T(any(val).(T)))
		}
	case reflect.Float64:
		val, err := strconv.ParseFloat(input_split[2], 64)
		if err != nil {
			log.Fatalf("%s не переводится в float64", input_split[2])
		} else {
			resultMatrix = operation((*matrixList)[num], T(any(val).(T)))
		}
	default:
		log.Fatalf("Тип %T не поддерживается", *new(T))
	}
	n := len(*matrixList)
	(*matrixList)[n+1] = resultMatrix
	fmt.Println("\nРезультат:")
	fmt.Println(&resultMatrix)
}

// Операции над матрицей
func OperationMatrix[T matrix.Number](
	input_split []string,
	matrixList *map[int]matrix.Matrix[T],
	operation func(matrix.Matrix[T]) matrix.Matrix[T]) {
	num, err := strconv.Atoi(input_split[1])
	if err != nil {
		log.Fatalf("%s не переводится в int", input_split[1])
	}
	resultMatrix := operation((*matrixList)[num])
	n := len(*matrixList)
	(*matrixList)[n+1] = resultMatrix
	fmt.Println("\nРезультат:")
	fmt.Println(&resultMatrix)
}
