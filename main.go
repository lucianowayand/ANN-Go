package main

import (
	"ann/linearsystems"
)

//func f(x float64) float64{ return x }

func main() {
	input := [linearsystems.SIZE1 * linearsystems.SIZE2]float64{
		6.51, -0.77, 4.66,
		3.46, 4.97, 0.44,
		-2.26, -4.39, 7.73,
	}

	matrix := linearsystems.CreateMatrix(input)

	results := [linearsystems.SIZE2]float64 {2.25, 0.65, 4.41}

	estimation := [linearsystems.SIZE2]float64 {0.43, 0.39, 2.54}

	linearsystems.Jacobi(matrix, results, estimation, 19)
}