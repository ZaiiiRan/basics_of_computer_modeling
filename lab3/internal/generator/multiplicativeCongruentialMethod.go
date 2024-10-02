package generator


func MultiplicativeCongruentialMethod(seed int, a int, m int, N int) []float64 {
	results := make([]float64, N)
	x := seed

	for i := 0; i < N; i++ {
		x = (a * x) % m
		results[i] = float64(x) / float64(m)
	}

	return results
}