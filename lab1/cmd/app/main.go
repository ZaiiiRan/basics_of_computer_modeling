package main

import(
	"fmt"
	"math"
	"lab1/internal/matrix"
)

func round(x float64, n int) float64 {
	pow := math.Pow(10, float64(n))
	return math.Round(x * pow) / pow
}

// y = a * x + b
func linear(x, y []float64) (float64, float64) {
	n := float64(len(x))
	sumX, sumY, sumX2, sumXY := 0.0, 0.0, 0.0, 0.0

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumX2 += x[i] * x[i]
		sumXY += x[i] * y[i]
	}

	D := matrix.InitMatrix([][]float64{
		{sumX2, sumX},
		{sumX, n},
	})

	Da := matrix.InitMatrix([][]float64{
		{sumXY, sumX},
		{sumY, n},
	})

	Db := matrix.InitMatrix([][]float64{
		{sumX2, sumXY},
		{sumX, sumY},
	})

	det := round(D.Determinant(), 4)

	a := round(Da.Determinant() / det, 4)
	b := round(Db.Determinant() / det, 4)

	return round(a, 2), round(b, 2)
}

// y = a * x^b
func power(x, y []float64) (float64, float64) {
	logX, logY := make([]float64, len(x)), make([]float64, len(y))
	n := float64(len(x))
	sumLogX, sumLogY, sumLogX2, sumLogXY := 0.0, 0.0, 0.0, 0.0

	for i := 0; i< len(x); i++ {
		logX[i] = math.Log(x[i])
		logY[i] = math.Log(y[i])
		sumLogX += logX[i]
		sumLogY += logY[i]
		sumLogX2 += logX[i] * logX[i]
		sumLogXY += logX[i] * logY[i]
	}

	D := matrix.InitMatrix([][]float64{
		{sumLogX2, sumLogX},
		{sumLogX, n},
	})

	Db := matrix.InitMatrix([][]float64{
		{sumLogXY, sumLogX},
		{sumLogY, n},
	})

	DlnA := matrix.InitMatrix([][]float64{
		{sumLogX2, sumLogXY},
		{sumLogX, sumLogY},
	})

	det := round(D.Determinant(), 4)

	lnA := round(DlnA.Determinant() / det, 4)
	b := round(Db.Determinant() / det, 4)

	return round(math.Exp(lnA), 2), round(b, 2)
}

// y = a * e^(bx)
func exponential(x, y []float64) (float64, float64) {
	logY := make([]float64, len(y))
	n := float64(len(x))
	sumX, sumLogY, sumX2, sumXLogY := 0.0, 0.0, 0.0, 0.0

	for i := 0; i < len(x); i++ {
		logY[i] = math.Log(y[i])
		sumX += x[i]
		sumLogY += logY[i]
		sumX2 += x[i] * x[i]
		sumXLogY += x[i] * logY[i]
	}

	D := matrix.InitMatrix([][]float64{
		{sumX2, sumX},
		{sumX, n},
	})

	Db := matrix.InitMatrix([][]float64{
		{sumXLogY, sumX},
		{sumLogY, n},
	})

	DlnA := matrix.InitMatrix([][]float64{
		{sumX2, sumXLogY},
		{sumX, sumLogY},
	})

	det := round(D.Determinant(), 4)

	lnA := round(DlnA.Determinant() / det, 4)
	b := round(Db.Determinant() / det, 4)

	return round(math.Exp(lnA), 2), round(b, 2)
}

// y = ax^2 + bx + c
func quadratic(x, y []float64) (float64, float64, float64) {
	n := float64(len(x))
	sumX, sumX2, sumX3, sumX4 := 0.0, 0.0, 0.0, 0.0
	sumY, sumXY, sumX2Y := 0.0, 0.0, 0.0

	for i := 0; i < len(x); i++ {
		x2 := x[i] * x[i]
		x3 := x2 * x[i]
		x4 := x3 * x[i]
		sumX += x[i]
		sumX2 += x2
		sumX3 += x3
		sumX4 += x4
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2Y += x2 * y[i]
	}

	D := matrix.InitMatrix([][]float64{
		{sumX4, sumX3, sumX2},
		{sumX3, sumX2, sumX},
		{sumX2, sumX, n},
	})

	Da := matrix.InitMatrix([][]float64{
		{sumX2Y, sumX3, sumX2},
		{sumXY, sumX2, sumX},
		{sumY, sumX, n},
	})

	Db := matrix.InitMatrix([][]float64{
		{sumX4, sumX2Y, sumX2},
		{sumX3, sumXY, sumX},
		{sumX2, sumY, n},
	})

	Dc := matrix.InitMatrix([][]float64{
		{sumX4, sumX3, sumX2Y},
		{sumX3, sumX2, sumXY},
		{sumX2, sumX, sumY},
	})

	det := round(D.Determinant(), 4)

	a := round(Da.Determinant() / det, 4)
	b := round(Db.Determinant() / det, 4)
	c := round(Dc.Determinant() / det, 4)

	return round(a, 2), round(b, 2), round(c, 2)
}

func main() {
	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{1.0, 1.5, 3.0, 4.5, 7.0, 8.5}

	aLin, bLin := linear(x, y)
	fmt.Printf("Линейная аппроксимация: y = %.2f * x + %.2f\n", aLin, bLin)

	aPow, bPow := power(x, y)
	fmt.Printf("Степенная аппроксимация: y = %.2f * x^%.2f\n", aPow, bPow)

	aExp, bExp := exponential(x, y)
	fmt.Printf("Показательная аппроксимация: y = %.2f * e^%.2fx\n", aExp, bExp)

	aQuad, bQuad, cQuad := quadratic(x, y)
	fmt.Printf("Квадратичная аппроксимация: y = %.2fx^2 + %.2fx + %.2f\n", aQuad, bQuad, cQuad)
}