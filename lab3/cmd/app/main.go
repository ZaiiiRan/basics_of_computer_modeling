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

func interactive() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1 - N случайных чисел методом Неймана (срединных квадратов)")
		fmt.Println("2 - N случайных чисел модифицированным методом Неймана (срединных квадратов)")
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