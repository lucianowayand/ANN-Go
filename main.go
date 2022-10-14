package main

import (
	"ann/linearsystems"

)

func main() {
	input := []float64{
		7, -2, 2, -5,
		8, 6, -8, 6, 
		5, 8, 6, 6, 
		-2, -3, 4, 4,
	}

	matrix := linearsystems.CreateMatrix(input, 4, 4)
	linearsystems.FormatMatrixResult(*matrix)

	op1 := linearsystems.ChangeLines(*matrix, 1, 2)
	linearsystems.FormatMatrixResult(op1)

	op2 := linearsystems.MultiplyLineElements(op1, 3, 10)
	linearsystems.FormatMatrixResult(op2)
}
