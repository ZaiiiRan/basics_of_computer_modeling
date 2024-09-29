package area

import (
    "math"
    "math/rand"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
	"image/color"
)

type AreaCalculator struct {
    n    float64
    N    int
    a    float64
    b    float64
    inside []point
    outside []point
}

type point struct {
    x float64
    y float64
}

func NewAreaCalculator(n float64, N int) *AreaCalculator {
    return &AreaCalculator{n: n, N: N}
}

func (ac *AreaCalculator) f(phi float64) float64 {
    return math.Sqrt((10) * math.Cos(phi)*math.Cos(phi) + (ac.n - 10) * math.Sin(phi)*math.Sin(phi))
}

func (ac *AreaCalculator) GeneratePoints() {
    var X, Y []float64
    for phi := 0.0; phi < 2*math.Pi; phi += 0.0001 {
        r := ac.f(phi)
        X = append(X, r*math.Cos(phi))
        Y = append(Y, r*math.Sin(phi))
    }
    ac.a = math.Abs(max(X) - min(X))
    ac.b = math.Abs(max(Y) - min(Y))

    for i := 0; i < ac.N; i++ {
        x := rand.Float64()*(2*ac.a) - ac.a
        y := rand.Float64()*(2*ac.b) - ac.b
        pp := math.Sqrt(x*x + y*y)

        // Используем atan2 для лучшей проверки попадания
        var ff float64
        if pp != 0 {
            ff = math.Atan2(y, x)
        }

        if pp < ac.f(ff) {
            ac.inside = append(ac.inside, point{x, y})
        } else {
            ac.outside = append(ac.outside, point{x, y})
        }
    }
}

func (ac *AreaCalculator) CalculateArea() (float64, float64) {
    m := float64(len(ac.inside))
    s := (m / float64(ac.N)) * ac.a * ac.b * 4
    return s, 7 * math.Pi
}

func (ac *AreaCalculator) BuildPlot(path string) {
	p := plot.New()

    // Рисуем фигуру
    figurePoints := make(plotter.XYs, 0)
    for phi := 0.0; phi <= 2*math.Pi; phi += 0.01 {
        r := ac.f(phi)
        x := r * math.Cos(phi)
        y := r * math.Sin(phi)
        figurePoints = append(figurePoints, plotter.XY{X: x, Y: y})
    }
    
    figureLine, err := plotter.NewLine(figurePoints)
    if err != nil {
        panic(err)
    }
    figureLine.LineStyle.Color = color.RGBA{B: 0, A: 255}
    figureLine.LineStyle.Width = vg.Points(2)
    p.Add(figureLine)

    // Рисуем точки внутри
    insidePoints := make(plotter.XYs, len(ac.inside))
    for i, pt := range ac.inside {
        insidePoints[i].X = pt.x
        insidePoints[i].Y = pt.y
    }
    insideScatter, err := plotter.NewScatter(insidePoints)
    if err != nil {
        panic(err)
    }
    insideScatter.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
    p.Add(insideScatter)

    // Рисуем точки снаружи
    outsidePoints := make(plotter.XYs, len(ac.outside))
    for i, pt := range ac.outside {
        outsidePoints[i].X = pt.x
        outsidePoints[i].Y = pt.y
    }
    outsideScatter, err := plotter.NewScatter(outsidePoints)
    if err != nil {
        panic(err)
    }
    outsideScatter.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
    p.Add(outsideScatter)

    p.Title.Text = "Площадь"
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"

    // Сохранение графика
    if err := p.Save(8*vg.Inch, 8*vg.Inch, path); err != nil {
        panic(err)
    }
}

func min(arr []float64) float64 {
    minVal := arr[0]
    for _, v := range arr {
        if v < minVal {
            minVal = v
        }
    }
    return minVal
}

func max(arr []float64) float64 {
    maxVal := arr[0]
    for _, v := range arr {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}