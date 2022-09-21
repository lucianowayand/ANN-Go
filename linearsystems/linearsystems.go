package main

import (
	"fmt"
)

const SIZE1 = 3
const SIZE2 = 3

func FormatMatrixResult(matrix [SIZE1][SIZE2]float64) {
	for position := range matrix {
		for _, element := range matrix[position] {
			fmt.Printf("%.8f, ", element)
		}
	}

}

func CreateMatrix(a float64, b float64, c float64, d float64, e float64, f float64, g float64, h float64, i float64) [SIZE1][SIZE2]float64 {
	return [SIZE1][SIZE2]float64{
		{a, b, c},
		{d, e, f},
		{g, h, i},
	}
}

func ChangeLines(matrix [SIZE1][SIZE2]float64, line1 int, line2 int) [SIZE1][SIZE2]float64 {
	aux := matrix

	matrix[line1-1] = aux[line2-1]
	matrix[line2-1] = aux[line1-1]

	return matrix
}

func MultiplyLineElements(matrix [SIZE1][SIZE2]float64, line int, constant float64) [SIZE1][SIZE2]float64 {
	for position, element := range matrix[line-1] {
		matrix[line-1][position] = element * constant
	}
	return matrix

}

func AddLineElements(matrix [SIZE1][SIZE2]float64, line1 int, line2 int) [SIZE1][SIZE2]float64 {
	aux := matrix

	for position := range matrix {
		aux[line1-1][position] = matrix[line1-1][position] + matrix[line2-1][position]
	}

	return aux
}

func main() {
	matrix := CreateMatrix(
		9.0/7.0, 9.0/2.0, 8.0/9.0, 
		-9.0/7.0, 5.0/3.0, -7.0/5.0, 
		-9.0/7.0, 4.0/7.0, -8.0/7.0,
	)

	op1 := ChangeLines(matrix, 1, 2)

	op2 := MultiplyLineElements(op1, 1, -3.0/2.0)

	op3 := ChangeLines(op2, 2, 3)

	op4 := MultiplyLineElements(op3, 1, -4.0/7.0)

	op5 := AddLineElements(op4, 3, 1)

	op6 := MultiplyLineElements(op5, 1, -7.0/4.0)

	FormatMatrixResult(op6)

}
