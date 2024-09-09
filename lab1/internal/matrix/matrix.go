package matrix

type Matrix struct {
	data [][]float64
}

func InitMatrix(data [][]float64) Matrix {
	return Matrix{data: data}
}

func (m Matrix) Determinant() float64 {
	size := len(m.data)

	if size == 1 {
		return m.data[0][0]
	}
	if size == 2 {
		return m.data[0][0] * m.data[1][1] - m.data[0][1] * m.data[1][0]
	}

	det := 0.0
	for i := 0; i < size; i++ {
		minor := make([][]float64, size - 1)
		for j := range minor {
			minor[j] = make([]float64, size - 1)
		}

		for row := 1; row < size; row++ {
			columnIndex := 0
			for col := 0; col < size; col++ {
				if col == i {
					continue
				}
				minor[row-1][columnIndex] = m.data[row][col]
				columnIndex++
			}
		}
		if i % 2 == 1 {
			det -= m.data[0][i] * InitMatrix(minor).Determinant()
		} else {
			det += m.data[0][i] * InitMatrix(minor).Determinant()
		}
	}

	return det
}