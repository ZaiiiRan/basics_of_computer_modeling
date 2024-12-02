package gasStation

import (
	"fmt"
	"lab6/internal/queue"
	"lab6/internal/request"
	"math/rand"
	"time"
	"math"
)

type GasStation struct {
	queue1                 *queue.Queue
	queue2                 *queue.Queue
	currentRequest1        *request.Request
	currentRequest2        *request.Request
	currentTime            float64
	nextRequestArrivalTime float64
	lostRequests           int
	servedRequests         []*request.Request
	totalRequests          []*request.Request
	departureIntervals     []float64
	random                 *rand.Rand
	queue1Lengths          []float64
	queue2Lengths          []float64
	queueTimestamps        []float64
}

func NewGasStation() *GasStation {
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	gs := &GasStation{
		queue1:             queue.NewQueue(),
		queue2:             queue.NewQueue(),
		random:             rnd,
		servedRequests:     make([]*request.Request, 0),
		totalRequests:      make([]*request.Request, 0),
		departureIntervals: make([]float64, 0),
		queue1Lengths:      make([]float64, 0),
		queue2Lengths:      make([]float64, 0),
		queueTimestamps:    make([]float64, 0),
	}

	gs.nextRequestArrivalTime = gs.generateNextRequestArrivalTime()

	return gs
}

func (gs *GasStation) Run(numRequests int) {
	gs.nextRequestArrivalTime = gs.generateNextRequestArrivalTime()

	for len(gs.totalRequests) < numRequests {
		// Новая заявка
		if gs.currentTime >= gs.nextRequestArrivalTime {
			newRequest := request.NewRequest(gs.nextRequestArrivalTime)
			fmt.Println("Клиент подъехал")
			if gs.queue1.Size() < 5 && (gs.queue1.Size() <= gs.queue2.Size()) {
				fmt.Println("Клиент добавлен в 1 очередь")
				gs.queue1.Enqueue(newRequest)
				gs.updateQueueLengths()
			} else if gs.queue2.Size() < 5 {
				fmt.Println("Клиент добавлен во 2 очередь")
				gs.queue2.Enqueue(newRequest)
				gs.updateQueueLengths()
			} else {
				fmt.Println("Клиент уехал, так как очереди переполнены")
				gs.lostRequests++
			}
			gs.totalRequests = append(gs.totalRequests, newRequest)
			gs.nextRequestArrivalTime = gs.generateNextRequestArrivalTime()
		}

		// Обработка в первой очереди
		if gs.currentRequest1 == nil && gs.queue1.Size() > 0 {
			gs.currentRequest1 = gs.queue1.Dequeue().(*request.Request)
			gs.currentRequest1.StartTime = gs.currentTime
			gs.currentRequest1.EndTime = gs.currentRequest1.StartTime + gs.generateServiceTime()
		}

		// Обработка во второй очереди
		if gs.currentRequest2 == nil && gs.queue2.Size() > 0 {
			gs.currentRequest2 = gs.queue2.Dequeue().(*request.Request)
			gs.currentRequest2.StartTime = gs.currentTime
			gs.currentRequest2.EndTime = gs.currentRequest2.StartTime + gs.generateServiceTime()
		}

		// Завершение обслуживания в первой очереди
		if gs.currentRequest1 != nil && gs.currentTime >= gs.currentRequest1.EndTime {
			gs.servedRequests = append(gs.servedRequests, gs.currentRequest1)
			if len(gs.servedRequests) > 1 {
				prevEndTime := gs.servedRequests[len(gs.servedRequests)-2].EndTime
				gs.departureIntervals = append(gs.departureIntervals, gs.currentTime-prevEndTime)
			}
			gs.currentRequest1 = nil
			gs.updateQueueLengths()
			fmt.Println("Клиент уехал, обработка завершена (из первой очереди)")
		}

		// Завершение обслуживания во второй очереди
		if gs.currentRequest2 != nil && gs.currentTime >= gs.currentRequest2.EndTime {
			gs.servedRequests = append(gs.servedRequests, gs.currentRequest2)
			if len(gs.servedRequests) > 1 {
				prevEndTime := gs.servedRequests[len(gs.servedRequests)-2].EndTime
				gs.departureIntervals = append(gs.departureIntervals, gs.currentTime-prevEndTime)
			}
			gs.currentRequest2 = nil
			gs.updateQueueLengths()
			fmt.Println("Клиент уехал, обработка завершена (из второй очереди)")
		}

		// Увеличение времени
		gs.currentTime += 0.01
	}
}

// Сбора статистики по очередям
func (gs *GasStation) updateQueueLengths() {
	gs.queue1Lengths = append(gs.queue1Lengths, float64(gs.queue1.Size()))
	gs.queue2Lengths = append(gs.queue2Lengths, float64(gs.queue2.Size()))
	gs.queueTimestamps = append(gs.queueTimestamps, gs.currentTime)
}

// Генерация времени следующей заявки
func (gs *GasStation) generateNextRequestArrivalTime() float64 {
	// Из предыдущих наблюдений известно, что интервалы времени между прибытием клиентов в час пик распределены экспоненциально с математическим ожиданием, равным 0.1 единицы времени.
	return gs.currentTime + gs.generateExponential(0.1)
}

// Генерация времени обработки
func (gs *GasStation) generateServiceTime() float64 {
	// Продолжительность обслуживания всех колонок одинакова и распределена экспоненциально с математическим ожиданием, равным 0.5 единицы времени.
	return gs.generateExponential(0.5)
}

// Генерация экспоненциального распределения
func (gs *GasStation) generateExponential(lambda float64) float64 {
	return generateExponential(gs.random.ExpFloat64(), lambda)
}

// Средняя длина очереди 1
func (gs *GasStation) AverageQueueLength1() float64 {
	if len(gs.queueTimestamps) < 2 {
		return 0
	}
	totalLength := 0.0
	for i := 1; i < len(gs.queueTimestamps); i++ {
		duration := gs.queueTimestamps[i] - gs.queueTimestamps[i-1]
		totalLength += gs.queue1Lengths[i-1] * duration
	}
	totalTime := gs.queueTimestamps[len(gs.queueTimestamps)-1] - gs.queueTimestamps[0]
	if totalTime == 0 {
		return 0
	}
	return totalLength / totalTime
}

// Средняя длина очереди 2
func (gs *GasStation) AverageQueueLength2() float64 {
	if len(gs.queueTimestamps) < 2 {
		return 0
	}
	totalLength := 0.0
	for i := 1; i < len(gs.queueTimestamps); i++ {
		duration := gs.queueTimestamps[i] - gs.queueTimestamps[i-1]
		totalLength += gs.queue2Lengths[i-1] * duration
	}
	totalTime := gs.queueTimestamps[len(gs.queueTimestamps)-1] - gs.queueTimestamps[0]
	if totalTime == 0 {
		return 0
	}
	return totalLength / totalTime
}

// Процент потерянных заявок
func (gs *GasStation) LostRequestPercentage() float64 {
	totalRequests := len(gs.servedRequests) + gs.lostRequests
	if totalRequests == 0 {
		return 0
	}
	return float64(gs.lostRequests) / (float64(totalRequests) * 2)
}

// Средний интервал между отъездами
func (gs *GasStation) AverageDepartureInterval() float64 {
	if len(gs.departureIntervals) == 0 {
		return 0
	}
	var total float64
	for _, interval := range gs.departureIntervals {
		total += interval
	}
	return total / float64(len(gs.departureIntervals))
}

// Среднее время пребывания клиента на заправке
func (gs *GasStation) AverageTimeInSystem() float64 {
	if len(gs.servedRequests) == 0 {
		return 0
	}
	var total float64
	for _, req := range gs.servedRequests {
		total += req.EndTime - req.ReceiptTime
	}
	return total / float64(len(gs.servedRequests))
}

func generateExponential(random float64, lambda float64) float64 { 
	return random / math.Pow(lambda, -1)
}