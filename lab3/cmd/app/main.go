package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"lab3/internal/generator"
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

func middleSquareMethod() {
	N := 6
	R0 := 0.583
	fmt.Println("Начальные данные: ")
	fmt.Println("N = ", N)
	fmt.Println("R0 = ", R0)

	result := generator.MiddleSquareMethod(R0, N)
	fmt.Println("\nПолученная последовательность: ")
	for i, num := range result {
		fmt.Printf("R%d =  %.4f\n", i, num)
	}
}

func modifiedMiddleSquareMethod() {
	N := 6
	R0 := 0.5836
	R1 := 0.2176
	fmt.Println("Начальные данные: ")
	fmt.Println("N = ", N)
	fmt.Println("R0 = ", R0)
	fmt.Println("R1 = ", R1)


	result := generator.ModifiedMiddleSquareMethod(R0, R1, N)
	fmt.Println("\nПолученная последовательность: ")
	for i, num := range result {
		fmt.Printf("R%d =  %.4f\n", i, num)
	}
}

func lemerMethod() {
	N := 5
	R0 := 0.585
	g := 927
	fmt.Println("Начальные данные: ")
	fmt.Println("N = ", N)
	fmt.Println("R0 = ", R0)
	fmt.Println("g = ", g)


	result := generator.LemerMethod(R0, float64(g), N)
	fmt.Println("\nПолученная последовательность: ")
	for i, num := range result {
		fmt.Printf("R%d =  %.4f\n", i, num)
	}
}

func lemerMethod2() {
	N := 6
	a := 265
	m := 129
	seed := 122
	fmt.Println("Начальные данные: ")
	fmt.Println("N = ", N)
	fmt.Println("a = ", a)
	fmt.Println("m = ", m)
	fmt.Println("X0 = ", seed)

	result := generator.MultiplicativeCongruentialMethod(seed, a, m, N)
	fmt.Println("\nПолученная последовательность: ")
	for i, num := range result {
		fmt.Printf("R%d =  %.4f\n", i + 1, num)
	}
}

func interactive() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1 - N случайных чисел методом Неймана (срединных квадратов)")
		fmt.Println("2 - N случайных чисел модифицированным методом Неймана (срединных квадратов)")
		fmt.Println("3 - N случайных чисел при помощи алгоритма Лемера")
		fmt.Println("4 - N случайных чисел мультипликативным конгруэнтным методом")
		fmt.Println("0 - Выход")

		fmt.Print("\n\n\nВведите номер действия: ")
		choice, _ := reader.ReadByte()

		switch choice {
		case '1':
			clearCmd()
			reader.ReadString('\n')
			middleSquareMethod()
			reader.ReadByte()
			clearCmd()
		case '2':
			clearCmd()
			reader.ReadString('\n')
			modifiedMiddleSquareMethod()
			reader.ReadByte()
			clearCmd()
		case '3':
			clearCmd()
			reader.ReadString('\n')
			lemerMethod()
			reader.ReadByte()
			clearCmd()
		case '4':
			clearCmd()
			reader.ReadString('\n')
			lemerMethod2()
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
	clearCmd()
	interactive()
}