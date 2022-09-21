package main

import (
	"ann/linearsystems"
)

//func f(x float64) float64{ return math.Exp(x)-2*math.Pow(x,2)+x-1.5}

func main() {
	input := [linearsystems.SIZE1 * linearsystems.SIZE2]float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
		10, 11, 12,
	}

	matrix := linearsystems.CreateMatrix(input)

	op1 := linearsystems.MultiplyLineElements(matrix, 3, 3.0/4.0)
	op2 := linearsystems.ChangeLines(op1,3,4)
	op3 := linearsystems.ChangeLines(op2,1,3)
	op4 := linearsystems.AddLineWeightedElements(op3, 2, 1, -9.0/8.0)
	op5 := linearsystems.AddLineWeightedElements(op4, 4, 1, 4.0)
	op6 := linearsystems.MultiplyLineElements(op5, 1, 5.0)

	linearsystems.FormatMatrixResult(op6)
}
