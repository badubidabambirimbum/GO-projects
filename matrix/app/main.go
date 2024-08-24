package main

import (
	funct "matrix/matrixFunctions"
	matrix "matrix/matrixStruct"
)

func main() {
	m := matrix.Matrix[int]{}
	m.Create()
	m = funct.MultiplicationByNumber(m, 5)
	funct.PrintMatrix(m)
}
