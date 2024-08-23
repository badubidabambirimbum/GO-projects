package main

import (
	matrixInterface "matrix/interface"
	matrixInt "matrix/matrixInt"
)

func main() {
	m := matrixInt.MatrixInt{}
	m.Create()

	m.SetMatrix(5, 5, 5)
	matrixInterface.MultiplicationByNumber(&m, "5")
	matrixInterface.PrintMatrix(&m)
}
