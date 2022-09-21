package main

import (
	"ann/linearsystems"
)

//func f(x float64) float64{ return math.Exp(x)-2*math.Pow(x,2)+x-1.5}

func main(){
	in := [linearsystems.SIZE1*linearsystems.SIZE2]float64{
		1,2,3,
		4,5,6,
		7,8,9,
		10,11,12,
	}

	linearsystems.FormatMatrixResult(linearsystems.CreateMatrix(in))
}

