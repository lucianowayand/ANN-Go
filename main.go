package main

import (
	"ann/zeroes"
	"math"
	
)

func f(x float64) float64{ return math.Exp(x)-2*math.Pow(x,2)+x-1.5}

func main(){
	a := 3.497683
	b := 4.230358
	tol := 2.40849e-06

	zeroes.Bisection(f, a, b , 100, tol)

}

