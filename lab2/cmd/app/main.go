package main

import (
	"fmt"
	"time"
	"math/rand"
	"flag"
	"os"
	"os/exec"
	"bufio"

	
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

	fmt.Printf("Реальная площадь треугольника: %.4f\n", triangle.TriangleArea())
	fmt.Printf("Вычисленная площадь треугольника: %.4f\n", triangle.FindTriangleArea(points))

	triangle.BuildPlot(points, "./task1.png")
}

func interactive() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1 - Площадь теругольника")
		fmt.Println("0 - Площадь теругольника")

		fmt.Print("\n\n\nВведите номер действия: ")
		choice, _ := reader.ReadByte()

		switch choice {
		case '1':
			clearCmd()
			reader.ReadString('\n')
			triangle()
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
		default:
			fmt.Println("Неизвестное действие")
		}
	}
}
