package nonlinearsystems

import (
	"fmt"
)

type FunctionFloat64 func(float64, float64) float64

func Newton(
	f FunctionFloat64, fx FunctionFloat64, fy FunctionFloat64,
	g FunctionFloat64, gx FunctionFloat64, gy FunctionFloat64,
	n int, x float64, y float64,
) {
	x0 := x
	y0 := y

	for i := 0; i < n; i++ {
		a := fx(x0, y0)
		b := fy(x0, y0)
		c := gx(x0, y0)
		d := gy(x0, y0)

		det := a*d - b*c

		xk := x0 - (gy(x0, y0)*f(x0, y0)-fy(x0, y0)*g(x0, y0))/det
		yk := y0 - (-gx(x0, y0)*f(x0, y0)+fx(x0, y0)*g(x0, y0))/det

		x0 = xk
		y0 = yk
		fmt.Printf("%.12f,\n %.12f,\n", x0, y0)
	}
}
