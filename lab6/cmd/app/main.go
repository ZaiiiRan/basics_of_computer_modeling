package main

import (
	"fmt"
	"lab6/internal/gasStation"
	"lab6/internal/machine"
)

func simulateMachine() {
	machine := machine.NewMachine()
	machine.Run(500)

	completedRequests := machine.CompletedRequests()
	var totalTime float64
	for _, request := range completedRequests {
		totalTime += (request.EndTime - request.StartTime)
	}
	avgTime := totalTime / float64(len(completedRequests))
	util := totalTime / completedRequests[len(completedRequests)-1].EndTime

	fmt.Printf("Загрузка станка:  %d%%\n", int(util*100))
	fmt.Printf("Время выполнения заявки: %.5f ч.\n", avgTime)
}

func simulateGasStation() {
	gs := gasStation.NewGasStation()
	gs.Run(400)

	avgQueueLength1 := gs.AverageQueueLength1()
	avgQueueLength2 := gs.AverageQueueLength2()
	lostRequestPercentage := gs.LostRequestPercentage()
	avgDepartureInterval := gs.AverageDepartureInterval()
	avgTimeInSystem := gs.AverageTimeInSystem()

	fmt.Printf("Среднее число клиентов в первой очереди: %.3f\n", avgQueueLength1)
	fmt.Printf("Среднее число клиентов во второй очереди: %.3f\n", avgQueueLength2)
	fmt.Printf("Процент клиентов, которые отказались от обслуживания: %f%%\n", lostRequestPercentage*100)
	fmt.Printf("Средний интервал между отъездами клиентов: %.3f\n", avgDepartureInterval)
	fmt.Printf("Среднее время пребывания клиента на запраке: %.3f\n", avgTimeInSystem)
}

func main() {
	fmt.Println("Станок")
	simulateMachine()
	fmt.Print("\n\n\n")

	fmt.Println("Заправка")
	simulateGasStation()
}
