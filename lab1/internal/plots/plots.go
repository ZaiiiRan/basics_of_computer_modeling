package plots

import (
	"log"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"lab1/internal/funcs"
)

func CreatePlot(x []float64, y []float64, aLin, bLin, aPow, bPow, aExp, bExp, aQuad, bQuad, cQuad float64) {
	p := plot.New()
	p.Title.Text = "Графики аппроксимаций"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	var X []float64
	for i := 0.1; i <= 10.0; i += 0.01 {
		X = append(X, i)
	}

	linearData := make(plotter.XYs, len(X))
	powerData := make(plotter.XYs, len(X))
	expData := make(plotter.XYs, len(X))
	quadData := make(plotter.XYs, len(X))
	points := make(plotter.XYs, len(x))

	for i := range X {
		linearData[i].X = X[i]
		linearData[i].Y = funcs.LinearFunc(aLin, bLin, X[i])

		powerData[i].X = X[i]
		powerData[i].Y = funcs.PowerFunc(aPow, bPow, X[i])

		expData[i].X = X[i]
		expData[i].Y = funcs.ExpFunc(aExp, bExp, X[i])

		quadData[i].X = X[i]
		quadData[i].Y = funcs.QuadraticFunc(aQuad, bQuad, cQuad, X[i])
	}

	for i := range x {
		points[i].X = x[i]
		points[i].Y = y[i]
	}

	err := plotutil.AddLinePoints(p, "Точки", points)
	if err != nil {
		log.Fatalf("Произошла ошибка при заполнении точек: %v", err)
	}

	linearLine, err := plotter.NewLine(linearData)
	if err != nil {
		log.Fatalf("Произошла ошибка при рисовании линейной аппроксимации: %v", err)
	}
	linearLine.Color = plotutil.Color(0)

	powerLine, err := plotter.NewLine(powerData)
	if err != nil {
		log.Fatalf("Произошла ошибка при рисовании степенной аппроксимации: %v", err)
	}
	powerLine.Color = plotutil.Color(1)

	expLine, err := plotter.NewLine(expData)
	if err != nil {
		log.Fatalf("Произошла ошибка при рисовании показательной аппроксимации: %v", err)
	}
	expLine.Color = plotutil.Color(2)

	quadLine, err := plotter.NewLine(quadData)
	if err != nil {
		log.Fatalf("Произошла ошибка при рисовании показательной квадратичной: %v", err)
	}
	quadLine.Color = plotutil.Color(3)

	p.Add(linearLine, powerLine, expLine, quadLine)

	p.Legend.Add("Линейная", linearLine)
	p.Legend.Add("Степенная", powerLine)
	p.Legend.Add("Показательная", expLine)
	p.Legend.Add("Квадратичная", quadLine)

	if err := p.Save(8*vg.Inch, 6*vg.Inch, "plot.png"); err != nil {
		log.Fatalf("Произошла ошибка при сохранении графика: %v", err)
	}

	log.Println("График сохранен как plot.png")
}