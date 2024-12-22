package simulation

import (
	"fmt"
	"lab8/internal/queue"
	"math"
	"math/rand"
	"time"
)

// Статусы светофора
const (
	Green1 = "g1"
	Green2 = "g2"
	Red    = "r"
)

// Класс светофора
type TrafficLight struct {
	g1Dur   int // Длительность зеленого сигнала для 1-го направления
	g2Dur   int // Длительность зеленого сигнала для 2-го направления
	rDur    int // Длительность красного сигнала
	curTime int // Текущее время
	cycle   int // Полный цикл работы светофора
}

// Создание экземпляра светофора
func NewTrafficLight(g1Dur, g2Dur, rDur int) *TrafficLight {
	return &TrafficLight{
		g1Dur: g1Dur,
		g2Dur: g2Dur,
		rDur:  rDur,
		cycle: g1Dur + g2Dur + 2*rDur,
	}
}

// Текущее состояние светофора
func (t *TrafficLight) Status() string {
	pos := t.curTime % t.cycle
	if pos < t.g1Dur {
		return Green1
	} else if pos < t.g1Dur+t.rDur {
		return Red
	} else if pos < t.g1Dur+t.rDur+t.g2Dur {
		return Green2
	}
	return Red
}

// Увеличение времени
func (t *TrafficLight) Advance(step int) {
	t.curTime += step
}

// Симуляция движения в туннеле
type TunnelSim struct {
	nCars      int           // Количество машин
	rate1      float64       // Интенсивность потока 1
	rate2      float64       // Интенсивность потока 2
	light      *TrafficLight // Светофор
	q1         *queue.Queue  // Очередь потока 1
	q2         *queue.Queue  // Очередь потока 2
	curTime    int           // текущее время
	carsPassed int           // Количество проехавших машин
	q1Lens     []int         // История длин первой очереди
	q2Lens     []int         // История длин второй очереди
	times      []int         // История времени
	wait1      []int         // История ожиданий 1
	wait2      []int         // История ожиданий 2
	randSrc    *rand.Rand
}

// Создание симуляции
func NewTunnelSim(nCars int, rate1, rate2 float64, g1Dur, g2Dur, rDur int) *TunnelSim {
	return &TunnelSim{
		nCars:   nCars,
		rate1:   rate1,
		rate2:   rate2,
		light:   NewTrafficLight(g1Dur, g2Dur, rDur),
		q1:      queue.NewQueue(),
		q2:      queue.NewQueue(),
		randSrc: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// генерация времени прибытия автомобиля на основе интенсивности (экспоненциальное распределение)
func (sim *TunnelSim) genArrivalTime(rate float64) float64 {
	return -math.Log(1.0-sim.randSrc.Float64()) / rate
}

// Запуск симуляции
func (sim *TunnelSim) Run() {
	next1 := sim.genArrivalTime(sim.rate1)
	next2 := sim.genArrivalTime(sim.rate2)

	for sim.carsPassed < sim.nCars {
		sim.light.Advance(1)
		sim.curTime++

		// обработка прибытия машин
		if float64(sim.curTime) >= next1 {
			sim.q1.Enqueue(sim.curTime)
			next1 += sim.genArrivalTime(sim.rate1)
		}

		if float64(sim.curTime) >= next2 {
			sim.q2.Enqueue(sim.curTime)
			next2 += sim.genArrivalTime(sim.rate2)
		}

		// пропуск машин через туннель в зависимости от статуса светофора
		status := sim.light.Status()
		if status == Green1 && !sim.q1.IsEmpty() {
			arrival := sim.q1.Dequeue().(int)
			sim.wait1 = append(sim.wait1, sim.curTime-arrival)
			sim.carsPassed++
		} else if status == Green2 && !sim.q2.IsEmpty() {
			arrival := sim.q2.Dequeue().(int)
			sim.wait2 = append(sim.wait2, sim.curTime-arrival)
			sim.carsPassed++
		}

		// сохранение длин очередей
		sim.q1Lens = append(sim.q1Lens, sim.q1.Length())
		sim.q2Lens = append(sim.q2Lens, sim.q2.Length())
		sim.times = append(sim.times, sim.curTime)
	}
}

// вычисление средней длины очереди
func avgQueueLen(lens, times []int) float64 {
	totalLen := 0.0
	totalTime := times[len(times)-1] - times[0]
	for i := 1; i < len(times); i++ {
		duration := times[i] - times[i-1]
		totalLen += float64(lens[i-1]) * float64(duration)
	}
	if totalTime > 0 {
		return totalLen / float64(totalTime)
	}
	return 0.0
}

// Анализ результатов
func (sim *TunnelSim) Analyze() {
	avgQ1 := avgQueueLen(sim.q1Lens, sim.times)
	avgQ2 := avgQueueLen(sim.q2Lens, sim.times)

	var avgW1, avgW2 float64
	if len(sim.wait1) > 0 {
		sum := 0
		for _, w := range sim.wait1 {
			sum += w
		}
		avgW1 = float64(sum) / float64(len(sim.wait1))
	}
	if len(sim.wait2) > 0 {
		sum := 0
		for _, w := range sim.wait2 {
			sum += w
		}
		avgW2 = float64(sum) / float64(len(sim.wait2))
	}

	fmt.Printf("Средняя длина очереди 1-го направления: %.2f\n", avgQ1)
	fmt.Printf("Средняя длина очереди 2-го направления: %.2f\n", avgQ2)
	fmt.Printf("Среднее время ожидания 1-го направления: %.2f секунд\n", avgW1)
	fmt.Printf("Среднее время ожидания 2-го направления: %.2f секунд\n", avgW2)
	fmt.Printf("Обслужено автомобилей: %d\n", sim.carsPassed)
}

// Нахождение оптимальных длительностей зеленых сигналов
func (sim *TunnelSim) OptimizeDurations() {
	bestWait := math.Inf(1)
	var bestDurs [2]int
	step := 5
	var avgTimes [2]float64

	for g1 := 30; g1 <= 90; g1 += step {
		for g2 := 30; g2 <= 90; g2 += step {
			tempSim := NewTunnelSim(sim.nCars, sim.rate1, sim.rate2, g1, g2, sim.light.rDur)
			tempSim.Run()

			avgW1, avgW2 := 0.0, 0.0
			if len(tempSim.wait1) > 0 {
				sum := 0
				for _, w := range tempSim.wait1 {
					sum += w
				}
				avgW1 = float64(sum) / float64(len(tempSim.wait1))
			}
			if len(tempSim.wait2) > 0 {
				sum := 0
				for _, w := range tempSim.wait2 {
					sum += w
				}
				avgW2 = float64(sum) / float64(len(tempSim.wait2))
			}

			avgWait := (avgW1 + avgW2) / 2
			if avgWait < bestWait {
				bestWait = avgWait
				bestDurs = [2]int{g1, g2}
				avgTimes = [2]float64{avgW1, avgW2}
			}
		}
	}
	fmt.Printf("Оптимальное решение: %d для 1-го направления и %d для 2-го направления.\n", bestDurs[0], bestDurs[1])
	fmt.Printf("Среднее время ожидания: %f для 1-го направления и %f для второго направления\n", avgTimes[0], avgTimes[1])
}
