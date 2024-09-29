package integral

import (
	"math/rand"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

type Integral struct {
	F func(float64) float64
	A float64
	B float64
	N int
	MaxF float64
}

func (i *Integral) Init(a, b float64, N int, f func(float64) float64) {
	i.A = a
	i.B = b
	i.F = f
	i.N = N
}

func (integral *Integral) GeneratePoints() (inside plotter.XYs, outside plotter.XYs) {
	// Поиск максимального значения функции для корректного генератора точек по оси Y
	maxF := 0.0
	for x := integral.A; x <= integral.B; x += 0.001 {
		val := integral.F(x)
		if val > maxF {
			maxF = val
		}
	}

	for i := 0; i < integral.N; i++ {
		px := (integral.B - integral.A) * rand.Float64() + integral.A // x генерируется в диапазоне [A, B]
		py := rand.Float64() * maxF                                    // y генерируется в диапазоне [0, maxF]
		if py < integral.F(px) {
			inside = append(inside, plotter.XY{X: px, Y: py})
		} else {
			outside = append(outside, plotter.XY{X: px, Y: py})
		}
	}
	integral.MaxF = maxF
	return inside, outside
}

func (i *Integral) CalculateArea(inside plotter.XYs) float64 {
	M := len(inside)
	if M == 0 {
		return 0.0
	}
	return float64(M) / float64(i.N) * (i.B - i.A) * i.MaxF
}

func (i *Integral) BuildPlot(inside, outside plotter.XYs, path string) {
	p := plot.New()

	p.Title.Text = "Интеграл"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// p.X.Min = i.A - 1  // Можно уменьшить нижний предел для увеличения области
	// p.X.Max = i.B + 1  // Можно увеличить верхний предел
	
	// p.Y.Min = 0        // Задать минимальное значение для оси Y
	// p.Y.Max = maxF + 1 // Увеличить диапазон по оси Y для лучшей видимости


	insideScatter, err := plotter.NewScatter(inside)
	if err != nil {
		panic(err)
	}
	insideScatter.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}

	outsideScatter, err := plotter.NewScatter(outside)
	if err != nil {
		panic(err)
	}
	outsideScatter.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	p.Add(insideScatter, outsideScatter)

	// Построение графика самой функции
	linePoints := plotter.XYs{}
	for x := i.A; x <= i.B; x += 0.001 { // Исправляем диапазон от A до B
		linePoints = append(linePoints, plotter.XY{X: x, Y: i.F(x)})
	}
	line, err := plotter.NewLine(linePoints)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Width = vg.Points(2.7)
	line.LineStyle.Color = plotter.DefaultLineStyle.Color

	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, path); err != nil {
		panic(err)
	}
}