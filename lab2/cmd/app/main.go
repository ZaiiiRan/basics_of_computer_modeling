package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"lab2/internal/area"
	"lab2/internal/integral"
	"lab2/internal/pi"
	"lab2/internal/triangleArea"
)

func clearCmd() {
	var cmd *exec.Cmd

    if os.Getenv("OS") == "Windows_NT" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }

    cmd.Stdout = os.Stdout
    cmd.Run()
}

func triangle() {
	var triangle triangleArea.Triangle
	triangle.Init(0.0, 32.0, 0.0, 16.0, 1000, func(x float64) float64 { return 10 * x / 14}, func(x float64) float64 { return 10 * (x - 20) / (-6) + 20 })

	points := triangle.GenerateNormalPoints()

	exact_area := triangle.TriangleArea()
	approx_area := triangle.FindTriangleArea(points)
	fmt.Printf("Реальная площадь треугольника: %.4f\n", exact_area)
	fmt.Printf("Вычисленная площадь треугольника: %.4f\n", approx_area)

	abs_error := math.Abs(exact_area - approx_area)
	fmt.Printf("Абсолютная погрешность: %.4f\n", abs_error)
	fmt.Printf("Относительная погрешность: %4f\n", (abs_error / exact_area) * 100)

	triangle.BuildPlot(points, "./task1.png")
}

func integrate() {
	var integral integral.Integral
	integral.Init(0, 5, 1000, func(x float64) float64 { return math.Sqrt(29 - 14 * math.Pow(math.Cos(x), 2)) })

	inside, outside := integral.GeneratePoints()

	approxArea := integral.CalculateArea(inside)
	fmt.Printf("Приблизительная площадь: %4f\n", approxArea)

	exactValue := 23.49836573097829
	fmt.Printf("Точная площадь: %4f\n", exactValue)
	absoluteError := math.Abs(approxArea - exactValue)
	fmt.Printf("Абсолютная погрешность: %.4f\n", absoluteError)
	fmt.Printf("Относительная погрешность: %.4f\n", (absoluteError/exactValue) * 100)

	integral.BuildPlot(inside, outside, "./task2.png")
}

func findPi() {
	radius := 14.0
	N := 1000

	simulator := pi.PiSimulator{
		N: N,
		R: radius,
	}

	simulator.GeneratePoints()
	piEstimate, _:= simulator.CalculatePi()

	fmt.Printf("Вычисленное значение pi = %.6f\n", piEstimate)
	fmt.Printf("Значение числа pi: %.6f\n", math.Pi)

	// Вычисление абсолютной и относительной погрешности
	absoluteError := math.Abs(piEstimate - math.Pi)
	relativeError := (absoluteError / math.Pi) * 100
	fmt.Printf("Абсолютная погрешность: %.6f\n", absoluteError)
	fmt.Printf("Относительная погрешность: %.6f\n", relativeError)

	simulator.BuildPlot("./task3.png")
}

func findArea() {
	n := 14.0
	N := 10000
	ac := area.NewAreaCalculator(n, N)
	ac.GeneratePoints()
    approxArea, exactValue := ac.CalculateArea()
    
    fmt.Printf("Приблизительная площадь: %.4f\n", approxArea)
    fmt.Printf("Точное значение площади: %.4f\n", exactValue)
    absoluteError := math.Abs(approxArea - exactValue)
    fmt.Printf("Абсолютная погрешность: %.4f\n", absoluteError)
    fmt.Printf("Относительная погрешность: %.4f\n", (absoluteError/exactValue) * 100)

    ac.BuildPlot("./task4.png")
}

func interactive() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1 - Площадь теругольника")
		fmt.Println("2 - Интеграл")
		fmt.Println("3 - π")
		fmt.Println("4 - Площадь фигуры")
		fmt.Println("0 - Выход")

		fmt.Print("\n\n\nВведите номер действия: ")
		choice, _ := reader.ReadByte()

		switch choice {
		case '1':
			clearCmd()
			reader.ReadString('\n')
			triangle()
			reader.ReadByte()
			clearCmd()
		case '2':
			clearCmd()
			reader.ReadString('\n')
			integrate()
			reader.ReadByte()
			clearCmd()
		case '3':
			clearCmd()
			reader.ReadString('\n')
			findPi()
			reader.ReadByte()
			clearCmd()
		case '4':
			clearCmd()
			reader.ReadString('\n')
			findArea()
			reader.ReadByte()
			clearCmd()
		case '0':
			return
		default:
			clearCmd()
		}
	}
}

func main() {
	action := flag.Int("action", -1, "Выберите действие")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if *action == -1 {
		clearCmd()
		interactive()
	} else {

		switch *action  {
		case 1:
			triangle()
		case 2:
			integrate()
		case 3:
			findPi()
		case 4:
			findArea()
		default:
			fmt.Println("Неизвестное действие")
		}
	}
}
