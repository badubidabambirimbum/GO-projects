package main

import (
	matrixInterface "matrix/interface"
	matrixInt "matrix/matrixInt"
)

func main() {
	m := matrixInt.MatrixInt{}
	m.Create()
	m2 := matrixInt.MatrixInt{}
	matrixInterface.MultiplicationByNumber(&m, 5, &m)
	matrixInterface.PrintMatrix(&m2)
	matrixInterface.PrintMatrix(&m)
}
