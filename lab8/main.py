import math
import random
from queue import Queue


class TrafficLight:
    """Класс управления светофором."""
    def __init__(self, green1_duration, green2_duration, red_duration):
        self.green1_duration = green1_duration
        self.green2_duration = green2_duration
        self.red_duration = red_duration
        self.current_time = 0
        self.cycle_time = green1_duration + green2_duration + 2 * red_duration

    def get_light_status(self):
        """Возвращает текущий статус светофора."""
        cycle_position = self.current_time % self.cycle_time
        if cycle_position < self.green1_duration:
            return "green1"
        elif cycle_position < self.green1_duration + self.red_duration:
            return "red"
        elif cycle_position < self.green1_duration + self.red_duration + self.green2_duration:
            return "green2"
        else:
            return "red"

    def advance_time(self, time_step):
        """Увеличивает текущее время."""
        self.current_time += time_step


class TunnelSimulation:
    """Класс для моделирования движения через перекрёсток и туннель."""
    def __init__(self, num_cars, arrival_rate1, arrival_rate2, green1_duration, green2_duration, red_duration):
        self.num_cars = num_cars
        self.arrival_rate1 = arrival_rate1
        self.arrival_rate2 = arrival_rate2

        self.light = TrafficLight(green1_duration, green2_duration, red_duration)
        self.queue1 = Queue()
        self.queue2 = Queue()
        self.current_time = 0
        self.cars_passed = 0

        self.queue1_lengths = []
        self.queue2_lengths = []
        self.queue_timestamps = []
        self.wait_times1 = []  # Время ожидания для 1-го направления
        self.wait_times2 = []  # Время ожидания для 2-го направления

        self.random = random.Random()
        self.random.seed()

    def generate_arrival_time(self, rate):
        """Генерация времени прибытия машины."""
        return -math.log(1 - self.random.random()) / rate

    def run(self):
        """Основной цикл симуляции."""
        next_arrival1 = self.generate_arrival_time(self.arrival_rate1)
        next_arrival2 = self.generate_arrival_time(self.arrival_rate2)

        while self.cars_passed < self.num_cars:
            self.light.advance_time(1)
            self.current_time += 1

            # Прибытие машин
            if self.current_time >= next_arrival1:
                self.queue1.put(self.current_time)
                next_arrival1 += self.generate_arrival_time(self.arrival_rate1)

            if self.current_time >= next_arrival2:
                self.queue2.put(self.current_time)
                next_arrival2 += self.generate_arrival_time(self.arrival_rate2)

            # Движение через туннель
            light_status = self.light.get_light_status()
            if light_status == "green1" and not self.queue1.empty():
                arrival_time = self.queue1.get()
                self.wait_times1.append(self.current_time - arrival_time)
                self.cars_passed += 1
            elif light_status == "green2" and not self.queue2.empty():
                arrival_time = self.queue2.get()
                self.wait_times2.append(self.current_time - arrival_time)
                self.cars_passed += 1

            # Сохранение длины очередей
            self.queue1_lengths.append(self.queue1.qsize())
            self.queue2_lengths.append(self.queue2.qsize())
            self.queue_timestamps.append(self.current_time)

    def average_queue_length(self, queue_lengths):
        """Средняя длина очереди."""
        total_length = 0.0
        for i in range(1, len(self.queue_timestamps)):
            duration = self.queue_timestamps[i] - self.queue_timestamps[i - 1]
            total_length += queue_lengths[i - 1] * duration
        total_time = self.queue_timestamps[-1] - self.queue_timestamps[0]
        return total_length / total_time if total_time > 0 else 0
    

    def analyze_results(self):
        """Анализ результатов симуляции."""
        avg_queue1 = self.average_queue_length(self.queue1_lengths)
        avg_queue2 = self.average_queue_length(self.queue2_lengths)
        avg_wait_time1 = sum(self.wait_times1) / len(self.wait_times1) if self.wait_times1 else 0
        avg_wait_time2 = sum(self.wait_times2) / len(self.wait_times2) if self.wait_times2 else 0

        print(f"Средняя длина очереди 1-го направления: {avg_queue1:.2f}")
        print(f"Средняя длина очереди 2-го направления: {avg_queue2:.2f}")
        print(f"Среднее время ожидания автомобилей в 1-м направлении: {avg_wait_time1:.2f} секунд")
        print(f"Среднее время ожидания автомобилей во 2-м направлении: {avg_wait_time2:.2f} секунд")
        print(f"Обслужено автомобилей: {self.cars_passed}")

    def optimize_green_durations(self):
        """Оптимизация длительности зелёных сигналов."""
        best_avg_wait_time = float('inf')
        best_durations = None
        step = 5  # Шаг изменения длительности зелёного сигнала
        avg_times = ()
        for green1 in range(30, 91, step):
            for green2 in range(30, 91, step):
                temp_simulation = TunnelSimulation(
                    self.num_cars, 
                    self.arrival_rate1, 
                    self.arrival_rate2, 
                    green1, 
                    green2, 
                    self.light.red_duration
                )
                temp_simulation.run()
                avg_wait_time1 = sum(temp_simulation.wait_times1) / len(temp_simulation.wait_times1) if temp_simulation.wait_times1 else float('inf')
                avg_wait_time2 = sum(temp_simulation.wait_times2) / len(temp_simulation.wait_times2) if temp_simulation.wait_times2 else float('inf')
                avg_wait_time = (avg_wait_time1 + avg_wait_time2) / 2

                if avg_wait_time < best_avg_wait_time:
                    best_avg_wait_time = avg_wait_time
                    best_durations = (green1, green2)
                    avg_times = (avg_wait_time1, avg_wait_time2)

        print(f"Оптимальные зелёные интервалы: {best_durations[0]} с для 1-го направления, {best_durations[1]} с для 2-го направления.")
        print("Среднее время ожидания по двум направлениям с оптимизированными зелеными интервалами:")
        print("1 направление: ", avg_times[0])
        print("2 направление: ", avg_times[1])
        


# Пример использования
if __name__ == "__main__":
    simulation = TunnelSimulation(
        num_cars=1000,
        arrival_rate1=1/12,  # Среднее время прибытия 12 с
        arrival_rate2=1/9,   # Среднее время прибытия 9 с
        green1_duration=30,
        green2_duration=30,
        red_duration=60
    )
    simulation.run()
    simulation.analyze_results()
    simulation.optimize_green_durations()