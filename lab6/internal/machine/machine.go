package machine

import (
	"lab6/internal/queue"
	"lab6/internal/request"
	"math"
	"math/rand"
	"time"
)

type Machine struct {
	queue                  *queue.Queue
	currentRequest         *request.Request
	currentTime            float64
	nextBreakdownTime      float64
	nextRequestArrivalTime float64
	completedRequests      []*request.Request
	random                 *rand.Rand
}

func NewMachine() *Machine {
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	m := &Machine{
		completedRequests:      make([]*request.Request, 0),
		nextBreakdownTime:      0,
		nextRequestArrivalTime: 0,
		queue:                  queue.NewQueue(),
		random:                 rnd,
	}
	m.nextBreakdownTime = m.generateNextBreakdownTime()
	m.nextRequestArrivalTime = m.generateNextRequestArrivalTime()
	return m
}

func (m *Machine) CompletedRequests() []*request.Request {
	return m.completedRequests
}

func (m *Machine) Run(requestsNumber int) {
	for len(m.completedRequests) < requestsNumber {
		// Новая заявка
		if m.currentTime >= m.nextRequestArrivalTime {
			newRequest := request.NewRequest(m.nextRequestArrivalTime)
			m.queue.Enqueue(newRequest)
			m.nextRequestArrivalTime = m.generateNextRequestArrivalTime()
		}

		// Установка новой заявки в обработку
		if m.currentRequest == nil && m.queue.Size() > 0 {
			element := m.queue.Dequeue()
			if element != nil {
				m.currentRequest = element.(*request.Request)
				m.currentRequest.StartTime = m.currentTime + m.generateSetupTime()
			}
		}

		// Обработка текущей заявки
		if m.currentRequest != nil && m.currentTime >= m.currentRequest.StartTime {
			processingTime := m.generateProcessingTime()
			m.currentRequest.EndTime = m.currentRequest.StartTime + processingTime

			if m.currentTime >= m.nextBreakdownTime {
				// Поломка станка
				m.onBreakDown()
			} else if m.currentTime >= m.currentRequest.EndTime {
				// Завершение заявки
				m.completedRequests = append(m.completedRequests, m.currentRequest)
				m.currentRequest = nil
			}
		}
		m.currentTime += 0.01
	}
}

func (m *Machine) onBreakDown() {
	repairTime := m.random.Float64()*(0.5-0.1) + 0.1
	m.currentTime += repairTime
	m.nextBreakdownTime = m.generateNextBreakdownTime()

	if m.currentRequest != nil {
		m.queue.Enqueue(m.currentRequest)
		m.currentRequest = nil
	}
}

func (m *Machine) generateNextRequestArrivalTime() float64 {
	// Задания поступают на станок в среднем один раз в час. Распределение величины интервала между ними экспоненциально.
	return m.currentTime + m.generateExponential(1.0)
}

func (m *Machine) generateSetupTime() float64 {
	// Перед выполнением задания производится наладка станка, время осуществления которой распределено равномерно на интервале 0.2-0.5 ч.
	return m.random.Float64()*(0.5-0.2) + 0.2
}

func (m *Machine) generateProcessingTime() float64 {
	// Время выполнения задания нормально распределено с математическим ожиданием 0.5 ч и среднеквадратичным отклонением 0.1 ч
	return m.generateNormal(0.5, 0.1)
}

func (m *Machine) generateNextBreakdownTime() float64 {
	// Интервалы между поломками распределены нормально с математическим ожиданием 20 ч и среднеквадратичным отклонением 2 ч.
	return m.currentTime + m.generateNormal(20, 2)
}

func (m *Machine) generateExponential(lambda float64) float64 {
	return -math.Log(1-m.random.Float64()) / lambda
}

func (m *Machine) generateNormal(mean, stdDev float64) float64 {
	u1 := 1.0 - m.random.Float64()
	u2 := 1.0 - m.random.Float64()
	randStdNormal := math.Sqrt(-2.0*math.Log(u1)) * math.Sin(2.0*math.Pi*u2)
	return mean + stdDev*randStdNormal
}
