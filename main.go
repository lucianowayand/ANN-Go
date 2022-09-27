package main

import (
	"ann/nonlinearsystems"
)

func f(x float64, y float64) float64  { return x*x + y*y - 5 }
func fx(x float64, y float64) float64 { return 2 * x }
func fy(x float64, y float64) float64 { return 2 * y }

func g(x float64, y float64) float64  { return x*x + x*y*y*y - 3 }
func gx(x float64, y float64) float64 { return 2 * x + y*y*y }
func gy(x float64, y float64) float64 { return 3 * x*y*y }

func main() {
	nonlinearsystems.Newton(f, fx, fy, g, gx, gy, 5, -2.1985,0.598)
}
