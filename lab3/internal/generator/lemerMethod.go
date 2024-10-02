package generator

import (
	"math"
)

func fractionalPart(x float64) float64 {
	return x - math.Floor(x)
}

func LemerMethod(seed float64, g float64, N int) []float64 {
	results := make([]float64, N + 1)
	results[0] = seed

	for i := 1; i <= N; i++ {
		results[i] = fractionalPart(g * results[i-1])
	}

	return results
}