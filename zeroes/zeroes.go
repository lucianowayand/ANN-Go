package zeroes

import (
	"fmt"
	"math"
)

type FunctionFloat64 func(float64) float64

func Bisection(f FunctionFloat64, a float64, b float64, n int, tol float64) float64{
	fa := f(a)
	fb := f(b)

	if fa * fb >= 0{
		fmt.Println("Intervalo inválido")
		return 0
	} else {
		for i := 0; i < n; i++ {
			m := 0.5 * (a+b)
			fm := f(m)
			if fm == 0{
				fmt.Printf("Voce encontrou uma raiz r = %.16f\n", m)
				return m
			}
			fmt.Printf("x_%d = %.16f\n", i+1, m)
			if math.Abs(fm) < tol{
				fmt.Printf("atingiu a tolerancia => x_%d = %.16f\n", i+1, m)
				return m
			}
			if fa * fm < 0{
				b = m
			} else {
				a = m
				fa = fm
			}
		}
	}
	return 0
}

func BissectionByInterval(a float64, b float64, tol float64) float64{
	return math.Ceil(math.Log2(math.Abs(b-a)/math.Abs(tol)))+1
}