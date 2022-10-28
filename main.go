package main

import (
	"ann/linearsystems"

)

func main() {
	matrix := linearsystems.CreateMatrix([]float64{
		7, -2, 2, -5,
		8, 6, -8, 6, 
		5, 8, 6, 6, 
		-2, -3, 4, 4,
	}, 4, 4)
	
	linearsystems.Gauss(*matrix)
}