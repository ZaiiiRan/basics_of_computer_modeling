package generator

import (
	"fmt"
	"strconv"
	"strings"
)

func ModifiedMiddleSquareMethod(R0, R1 float64, N int) []float64 {
	result := make([]float64, N + 2)
	result[0] = R0
	result[1] = R1

	for i := 2; i < N+2; i++ {
		squared := (result[i-1] + result[i-2]) * (result[i-1] + result[i-2])
		
		squaredString := fmt.Sprintf("%0.8f", squared)
		squaredString = strings.Replace(squaredString, "0.", "", -1)

		middle := squaredString[2:6]

		next, _ := strconv.ParseFloat("0."+middle, 64)

		result[i] = next
	}

	return result
}