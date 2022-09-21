package linearsystems

import (
	"fmt"
)

const SIZE1 = 3
const SIZE2 = 3

type Matrix struct {
	Matrix [SIZE1][SIZE2]float64
}

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

