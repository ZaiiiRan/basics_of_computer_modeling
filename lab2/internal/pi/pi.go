package pi

import (
	"math/rand"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
)

type Point struct {
	x, y float64
}

type PiSimulator struct {
	N      int      // Количество точек
	points []Point  // Массив точек
	R      float64  // Радиус круга
}

// Метод для генерации точек
func (sim *PiSimulator) GeneratePoints() {
	sim.points = make([]Point, sim.N)

	for i := 0; i < sim.N; i++ {
		sim.points[i] = Point{
			x: rand.Float64() * 2 * sim.R,
			y: rand.Float64() * 2 * sim.R,
		}
	}
}

// Метод для вычисления числа π
func (sim *PiSimulator) CalculatePi() (float64, int) {
	insideCount := 0

	for _, point := range sim.points {
		// Проверка, попала ли точка внутрь круга
		if (point.x-sim.R)*(point.x-sim.R)+(point.y-sim.R)*(point.y-sim.R) < sim.R*sim.R {
			insideCount++
		}
	}

	// Вычисление приближенного значения π
	piEstimate := 4.0 * float64(insideCount) / float64(sim.N)
	return piEstimate, insideCount
}

// Метод для построения графика
func (sim *PiSimulator) BuildPlot(path string) {
	// Создаем новый график
	p:= plot.New()

	// Заголовок и подписи
	p.Title.Text = "PI"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Разделяем точки на внутри и снаружи круга
	var inside, outside plotter.XYs
	for _, point := range sim.points {
		if (point.x-sim.R)*(point.x-sim.R)+(point.y-sim.R)*(point.y-sim.R) < sim.R*sim.R {
			inside = append(inside, plotter.XY{X: point.x, Y: point.y})
		} else {
			outside = append(outside, plotter.XY{X: point.x, Y: point.y})
		}
	}

	// Добавляем точки внутри круга (зеленые)
	insideScatter, err := plotter.NewScatter(inside)
	if err != nil {
		panic(err)
	}
	insideScatter.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	p.Add(insideScatter)

	// Добавляем точки вне круга (красные)
	outsideScatter, err := plotter.NewScatter(outside)
	if err != nil {
		panic(err)
	}
	outsideScatter.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	p.Add(outsideScatter)

	// Рисуем квадрат
	square := plotter.NewFunction(func(x float64) float64 {
		return 0 // Нижняя сторона квадрата
	})
	p.Add(square)

	p.Add(plotter.NewFunction(func(x float64) float64 {
		return 2 * sim.R // Верхняя сторона квадрата
	}))

	// Левая сторона квадрата
	p.Add(plotter.NewFunction(func(y float64) float64 {
		return 0 // Левая сторона (x = 0)
	}))

	// Правая сторона квадрата
	p.Add(plotter.NewFunction(func(y float64) float64 {
		return 2 * sim.R // Правая сторона (x = 2R)
	}))

	// Добавляем круг
	circleTop := plotter.NewFunction(func(x float64) float64 {
		return sim.R + math.Sqrt(sim.R*sim.R-(x-sim.R)*(x-sim.R)) // Верхняя полусфера
	})
	circleTop.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	p.Add(circleTop)

	// Добавляем нижнюю полусферу
	circleBottom := plotter.NewFunction(func(x float64) float64 {
		return sim.R - math.Sqrt(sim.R*sim.R-(x-sim.R)*(x-sim.R)) // Нижняя полусфера
	})
	circleBottom.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	p.Add(circleBottom)

	// Устанавливаем пределы графика
	p.X.Min = 0
	p.X.Max = 2 * sim.R
	p.Y.Min = 0
	p.Y.Max = 2 * sim.R

	// Сохраняем график в файл
	if err := p.Save(8*vg.Inch, 8*vg.Inch, path); err != nil {
		panic(err)
	}
}