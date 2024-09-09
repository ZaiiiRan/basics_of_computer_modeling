package funcs

import (
	"math"
)

func LinearFunc(a, b, x float64) float64 {
	return a * x + b
}

func PowerFunc(a, b, x float64) float64 {
	return a * math.Pow(x, b)
}

func ExpFunc(a, b, x float64) float64 {
	return a * math.Exp(x * b)
}

func QuadraticFunc(a, b, c, x float64) float64 {
	return a * x * x + b * x + c
}