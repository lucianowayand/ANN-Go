package zeroes

import (
	"fmt"
	"math"
)

type FunctionFloat64 func(float64) float64
const SIZE = 10

func Bisection(f FunctionFloat64, a float64, b float64, n int, tol float64) float64 {
	fa := f(a)
	fb := f(b)

	if fa*fb >= 0 {
		fmt.Println("Intervalo inválido")
		return 0
	} else {
		for i := 0; i < n; i++ {
			m := 0.5 * (a + b)
			fm := f(m)
			if fm == 0 {
				fmt.Printf("Voce encontrou uma raiz r = %.16f\n", m)
				return m
			}
			fmt.Printf("x_%d = %.16f\n", i+1, m)
			if math.Abs(fm) < tol {
				fmt.Printf("atingiu a tolerancia => x_%d = %.16f\n", i+1, m)
				return m
			}
			if fa*fm < 0 {
				b = m
			} else {
				a = m
				fa = fm
			}
		}
	}
	return 0
}

func BissectionByInterval(a float64, b float64, tol float64) float64 {
	return math.Ceil(math.Log2(math.Abs(b-a)/math.Abs(tol))) + 1
}

func Newton(f FunctionFloat64, df FunctionFloat64, x0 float64, n int, p [SIZE]int) {
	iteracao := 1
	i2 := 0
	for i := 0; i < n; i++ {
		dfx0 := df(x0)
		if dfx0 == 0 {
			break
		}
		xi := x0 - f(x0)/dfx0
		if p[i2] == iteracao {
			fmt.Printf("%.16f,\n", xi)
			i2++
		}
		iteracao++
		x0 = xi
	}
}

func Secant(f FunctionFloat64, x0 float64, x1 float64, n int) {
	for i := 0; i < n; i++ {
		fx0 := f(x0)
		fx1 := f(x1)

		if fx1 == fx0 {
			fmt.Printf("Divisão por zero na iteração %d", i+2)
			return
		}

		x2 := (x0*fx1 - x1*fx0) / (fx1 - fx0)
		fmt.Printf("x_%d = %.16f\n", i+2, x2)
		x0 = x1
		x1 = x2
	}
}

func FalsePosition(f FunctionFloat64, a float64, b float64, n int, p [SIZE]int) {
	iteracao := 1
	i2 := 0
	
	fa := f(a)
	fb := f(b)

	if fa*fb >= 0 {
		fmt.Printf("O teorema de bolzano nao sabe dizer se existe raiz para f no intervalo [%.16f, %.16f]", a, b)
		return
	} else {
		for i := 0; i < n; i++ {
			x := (a*fb - b*fa) / (fb - fa)
			if p[i2] == iteracao {
				fmt.Printf("%.16f, \n", x)
				i2++
			}
			fx := f(x)
			if fx == 0 {
				fmt.Printf("Encontramos uma raiz para f, x = %.16f", x)
				return
			}
			if fa*fx < 0 {
				b = x
				fb = fx
			} else {
				a = x
				fa = fx
			}
			iteracao++
		}

	}
}

func FixedPoint(f FunctionFloat64, x0 float64, n int, p [SIZE]int){
	iteracao := 1
	i2 := 0

	for i := 0; i < n; i++{
        x1 := f(x0)
        if iteracao == p[i2]{
            fmt.Printf("%.16f,\n", x1);
            i2++;
        }
        iteracao++;
        x0 = x1;
    }
}