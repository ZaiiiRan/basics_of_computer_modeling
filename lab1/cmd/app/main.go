package main

import (
	"fmt"
	"sync"

	"lab1/internal/approx"
	"lab1/internal/plots"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	x := []float64{1, 2, 3, 4, 5, 6}
	y := []float64{2.0, 0.68, 0.44, 0.24, 0.12, 0.14}

	// коэффициенты
	var aLin, bLin, aPow, bPow, aExp, bExp, aQuad, bQuad, cQuad float64

	wg.Add(4)

	go func() {
		defer wg.Done()
		mu.Lock()
		aLin, bLin = approx.Linear(x, y)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		aPow, bPow = approx.Power(x, y)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		aExp, bExp = approx.Exponential(x, y)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		aQuad, bQuad, cQuad = approx.Quadratic(x, y)
		mu.Unlock()
	}()

	wg.Wait()

	fmt.Printf("Линейная аппроксимация: y = %.2f * x + %.2f\n", aLin, bLin)
	fmt.Printf("Степенная аппроксимация: y = %.2f * x^%.2f\n", aPow, bPow)
	fmt.Printf("Показательная аппроксимация: y = %.2f * e^%.2fx\n", aExp, bExp)
	fmt.Printf("Квадратичная аппроксимация: y = %.2fx^2 + %.2fx + %.2f\n", aQuad, bQuad, cQuad)

	fmt.Print("\n\n")


	// погрешности и график
	var linErr, powErr, expErr, quadErr float64

	wg.Add(5)

	go func() {
		defer wg.Done()
		mu.Lock()
		linErr = approx.LinearErrorCalc(x, y, aLin, bLin)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		powErr = approx.PowerErrorCalc(x, y, aPow, bPow)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		expErr = approx.ExpErrorCalc(x, y, aExp, bExp)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		quadErr = approx.QuadraticErrorCalc(x, y, aQuad, bQuad, cQuad)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		plots.CreatePlot(x, y, aLin, bLin, aPow, bPow, aExp, bExp, aQuad, bQuad, cQuad)
		mu.Unlock()
	}()

	wg.Wait()

	fmt.Print("\n\n")

	fmt.Printf("Погрешность линейной аппроксимации: %.2f\n", linErr)
	fmt.Printf("Погрешность степенной аппроксимации: %.2f\n", powErr)
	fmt.Printf("Погрешность показательной аппроксимации: %.2f\n", expErr)
	fmt.Printf("Погрешность квадратичной аппроксимации: %.2f\n", quadErr)
}