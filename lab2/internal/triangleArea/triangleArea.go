package triangleArea

import (
	"math"
	"math/rand"

	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Triangle struct {
	yMin float64
	yMax float64
	xMin float64
	xMax float64

	N int
	f1 func(float64) float64
	f2 func(float64) float64

	expectedMeanX float64
	expectedMeanY float64
	expectedVarianceX float64
	expectedVarianceY float64
	tolerance float64
}

// инициализация треугольника
func (t *Triangle) Init(xMin, xMax, yMin, yMax float64, N int, f1 func(float64) float64, f2 func(float64) float64) {
	t.xMin = xMin
	t.xMax = xMax
	t.yMin = yMin
	t.yMax = yMax
	t.N = N
	t.f1 = f1
	t.f2 = f2
	t.expectedMeanX = (xMin + xMax) / 2
	t.expectedMeanY = (yMin + yMax) / 2
	t.expectedVarianceX = math.Pow(xMax-xMin, 2) / 12
	t.expectedVarianceY = math.Pow(yMax-yMin, 2) / 12
	t.tolerance = 0.01
}

// поиск границ треугольника
func (t *Triangle) boundariesSearching(x float64) float64 {
	y := t.f1(x)
	if y < t.yMax {
		return y
	}
	return t.f2(x)
}

// генерация случайных точек
func (t *Triangle) generatePoints() ([]float64, []float64) {
	x := make([]float64, t.N)
	y := make([]float64, t.N)
	for i := 0; i < t.N; i++ {
		x[i] = rand.Float64()*(t.xMax-t.xMin) + t.xMin
		y[i] = rand.Float64()*(t.yMax-t.yMin) + t.yMin
	}
	return x, y
}



func (t *Triangle) checkConvergence(meanX, varX, meanY, varY float64) bool {
	meanXDiff := math.Abs(meanX-t.expectedMeanX) / t.expectedMeanX
	varXDiff := math.Abs(varX-t.expectedVarianceX) / t.expectedVarianceX
	meanYDiff := math.Abs(meanY-t.expectedMeanY) / t.expectedMeanY
	varYDiff := math.Abs(varY-t.expectedVarianceY) / t.expectedVarianceY

	return meanXDiff < t.tolerance && varXDiff < t.tolerance && meanYDiff < t.tolerance && varYDiff < t.tolerance
}

func (t *Triangle) meanAndVariance(arr []float64) (float64, float64) {
	mean := 0.0
	for _, val := range arr {
		mean += val
	}
	mean /= float64(len(arr))

	variance := 0.0
	for _, val := range arr {
		variance += math.Pow(val-mean, 2)
	}
	variance /= float64(len(arr))
	return mean, variance
}

// основная функция по генерации точек
func (t *Triangle) GenerateNormalPoints() ([][2]float64) {
	converged := false
	var x, y []float64

	for !converged {
		x, y = t.generatePoints()

		meanX, varX := t.meanAndVariance(x)
		meanY, varY := t.meanAndVariance(y)

		converged = t.checkConvergence(meanX, varX, meanY, varY)
	}

	points := make([][2]float64, t.N)
	for i := range points {
		points[i] = [2]float64{x[i], y[i]}
	}
	return points
}

// Функция для вычисления доли точек внутри треугольника
func (t *Triangle) findPointsProportion(points [][2]float64) float64 {
	n := len(points)
	m := 0
	for _, point := range points {
		if point[1] <= t.boundariesSearching(point[0]) {
			m++
		}
	}
	return float64(m) / float64(n)
}

// Функция для нахождения площади треугольника
func (t *Triangle) FindTriangleArea(points [][2]float64) float64 {
	proportion := t.findPointsProportion(points)
	return proportion * (t.xMax - t.xMin) * (t.yMax - t.yMin)
}

// Нахождение площади по формуле
func (t *Triangle) TriangleArea() float64 {
	return (t.yMax / 2) * t.xMax
}







// Функция для построения графика
func (t *Triangle) BuildPlot(points [][2]float64, path string) {
	p := plot.New()

	p.Title.Text = "Площадь треугольника"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	var insidePoints, outsidePoints plotter.XYs
	for _, point := range points {
		if point[1] <= t.boundariesSearching(point[0]) {
			insidePoints = append(insidePoints, plotter.XY{X: point[0], Y: point[1]})
		} else {
			outsidePoints = append(outsidePoints, plotter.XY{X: point[0], Y: point[1]})
		}
	}
	// Точки внутри треугольника
	inside, err := plotter.NewScatter(insidePoints)
	if err != nil {
		panic(err)
	}
	inside.GlyphStyle.Color = color.RGBA{R:0, G:255, B:255} // Красные точки 

	// Точки вне треугольника
	outside, err := plotter.NewScatter(outsidePoints)
	if err != nil {
		panic(err)
	}
	outside.GlyphStyle.Color = color.RGBA{R:255, G:255, B:0} // Синие точки

	p.Add(inside, outside)

	linePoints := plotter.XYs{}
	for x := t.xMin; x <= t.xMax; x += 0.1 {
		linePoints = append(linePoints, plotter.XY{X: x, Y: t.boundariesSearching(x)})
	}
	line, err := plotter.NewLine(linePoints)
	if err != nil {
		panic(err)
	}
	line.Color = plotter.DefaultLineStyle.Color

	p.Add(line)

	if err := p.Save(8*vg.Inch, 6*vg.Inch, path); err != nil {
		panic(err)
	}
}