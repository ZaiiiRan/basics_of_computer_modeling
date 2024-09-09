package approx

import (
	"lab1/internal/funcs"
	"lab1/internal/matrix"
	"math"
)

func Round(x float64, n int) float64 {
	pow := math.Pow(10, float64(n))
	return math.Round(x*pow) / pow
}

// y = a * x + b
func Linear(x, y []float64) (float64, float64) {
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

	det := Round(D.Determinant(), 4)

	a := Round(Da.Determinant()/det, 4)
	b := Round(Db.Determinant()/det, 4)

	return Round(a, 2), Round(b, 2)
}

// y = a * x^b
func Power(x, y []float64) (float64, float64) {
	logX, logY := make([]float64, len(x)), make([]float64, len(y))
	n := float64(len(x))
	sumLogX, sumLogY, sumLogX2, sumLogXY := 0.0, 0.0, 0.0, 0.0

	for i := 0; i < len(x); i++ {
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

	det := Round(D.Determinant(), 4)

	lnA := Round(DlnA.Determinant()/det, 4)
	b := Round(Db.Determinant()/det, 4)

	return Round(math.Exp(lnA), 2), Round(b, 2)
}

// y = a * e^(bx)
func Exponential(x, y []float64) (float64, float64) {
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

	det := Round(D.Determinant(), 4)

	lnA := Round(DlnA.Determinant()/det, 4)
	b := Round(Db.Determinant()/det, 4)

	return Round(math.Exp(lnA), 2), Round(b, 2)
}

// y = ax^2 + bx + c
func Quadratic(x, y []float64) (float64, float64, float64) {
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

	det := Round(D.Determinant(), 4)

	a := Round(Da.Determinant()/det, 4)
	b := Round(Db.Determinant()/det, 4)
	c := Round(Dc.Determinant()/det, 4)

	return Round(a, 2), Round(b, 2), Round(c, 2)
}

func LinearErrorCalc(x, y []float64, aLin, bLin float64) float64 {
	err := 0.0
	for i := range x {
		err += math.Pow(y[i] - funcs.LinearFunc(aLin, bLin, x[i]), 2)
	}
	return Round(err, 2)
}

func PowerErrorCalc(x, y []float64, aPow, bPow float64) float64 {
	err := 0.0
	for i := range x {
		err += math.Pow(y[i] - funcs.PowerFunc(aPow, bPow, x[i]), 2)
	}
	return Round(err, 2)
}

func ExpErrorCalc(x, y []float64, aExp, bExp float64) float64 {
	err := 0.0
	for i := range x {
		err += math.Pow(y[i] - funcs.ExpFunc(aExp, bExp, x[i]), 2)
	}
	return Round(err, 2)
}

func QuadraticErrorCalc(x, y []float64, aQuad, bQuad, cQuad float64) float64 {
	err := 0.0
	for i := range x {
		err += math.Pow(y[i] - funcs.QuadraticFunc(aQuad, bQuad, cQuad, x[i]), 2)
	}
	return Round(err, 2)
}
