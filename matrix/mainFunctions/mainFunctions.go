// Package matrix Функции для функционала main функции
package matrix

import (
	"fmt"
	matrix "matrix/matrixStruct"
	inter "matrix/numInterface"
	"reflect"
	"strconv"
)

// PrintMatrix Печать матриц
func PrintMatrix[T inter.Number](matrixList map[int]matrix.Matrix[T]) {
	for key := 1; key <= len(matrixList); key++ {
		fmt.Println(key)
		m := matrixList[key]
		fmt.Println(&m)
	}
}

// AddMatrix Добавление матрицы в словарь
func AddMatrix[T inter.Number](matrixList *map[int]matrix.Matrix[T]) error {
	m := matrix.Matrix[T]{}
	err := m.Create()
	if err != nil {
		return err
	}
	n := len(*matrixList)
	(*matrixList)[n+1] = m

	return nil
}

// OperationBetweenMatrices Операции между матрицами
func OperationBetweenMatrices[T inter.Number](
	inputSplit []string,
	matrixList *map[int]matrix.Matrix[T],
	operation func(matrix.Matrix[T], matrix.Matrix[T]) (matrix.Matrix[T], error)) error {
	var resultMatrix matrix.Matrix[T]
	num1, err := strconv.Atoi(inputSplit[1])
	if err != nil {
		return fmt.Errorf("%s не переводится в int", inputSplit[1])
	}
	num2, err := strconv.Atoi(inputSplit[2])
	if err != nil {
		return fmt.Errorf("%s не переводится в int", inputSplit[2])
	}
	resultMatrix, err = operation((*matrixList)[num1], (*matrixList)[num2])
	if err != nil {
		return err
	}
	n := len(*matrixList)
	(*matrixList)[n+1] = resultMatrix
	fmt.Println("\nРезультат:")
	fmt.Println(&resultMatrix)

	return nil
}

// OperationMatrixByNumber Операция умножения матрицы на число
func OperationMatrixByNumber[T inter.Number](
	inputSplit []string,
	matrixList *map[int]matrix.Matrix[T],
	operation func(matrix.Matrix[T], T) (matrix.Matrix[T], error)) error {
	var resultMatrix matrix.Matrix[T]
	num, err := strconv.Atoi(inputSplit[1])
	if err != nil {

		return fmt.Errorf("%s не переводится в int", inputSplit[1])
	}

	switch reflect.TypeOf(*new(T)).Kind() {
	case reflect.Int:
		val, err := strconv.Atoi(inputSplit[2])
		if err != nil {
			return fmt.Errorf("%s не переводится в int", inputSplit[2])
		} else {
			resultMatrix, err = operation((*matrixList)[num], T(any(val).(T)))
			if err != nil {
				return err
			}
		}
	case reflect.Float64:
		val, err := strconv.ParseFloat(inputSplit[2], 64)
		if err != nil {
			return fmt.Errorf("%s не переводится в float64", inputSplit[2])
		} else {
			resultMatrix, err = operation((*matrixList)[num], T(any(val).(T)))
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("тип %T не поддерживается", *new(T))
	}
	n := len(*matrixList)
	(*matrixList)[n+1] = resultMatrix
	fmt.Println("\nРезультат:")
	fmt.Println(&resultMatrix)

	return nil
}

// OperationMatrix Операции над матрицей
func OperationMatrix[T inter.Number](
	inputSplit []string,
	matrixList *map[int]matrix.Matrix[T],
	operation func(matrix.Matrix[T]) (matrix.Matrix[T], error)) error {
	num, err := strconv.Atoi(inputSplit[1])
	if err != nil {
		return fmt.Errorf("%s не переводится в int", inputSplit[1])
	}
	resultMatrix, err := operation((*matrixList)[num])
	if err != nil {
		return err
	}
	n := len(*matrixList)
	(*matrixList)[n+1] = resultMatrix
	fmt.Println("\nРезультат:")
	fmt.Println(&resultMatrix)

	return nil
}
