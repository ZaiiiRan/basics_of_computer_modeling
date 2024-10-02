package generator

import (
	"fmt"
	"strconv"
	"strings"
)

func MiddleSquareMethod(seed float64, N int) []float64 {
	result := make([]float64, N + 1)
	result[0] = seed

	current := seed

	for i := 1; i < N + 1; i++ {
		squared := current * current

		squaredString := fmt.Sprintf("%0.8f", squared)
		squaredString = strings.Replace(squaredString, "0.", "", -1)

		middle := squaredString[2:6]

		next, _ := strconv.ParseFloat("0."+middle, 64)

		result[i] = next
		current = next
	}

	return result
}

