from math import * 
from matplotlib import pyplot as plt
from numpy import linspace

# Генерация последовательности
def random_sequence(N, a = 36261, m = 312500, X = 0.123):
    element_sequence = []
    for i in range (N):
        X_new = (a * X) % m
        R = round(X_new / m, 4)
        X = X_new
        element_sequence.append(R)
    return element_sequence

# Параметры распределений
a = 5
b = 12
lambd = 3
mu = 2
sigma = 2

def uniform_distribution(R):
    return a + R * (b - a)

def exponential_distribution(R):
    return -log(R)/lambd

# Метод Бокс-Мюллера
def normal_distribution(R1, R2):
    z0 = mu + sigma * (cos(2 * pi * R2) * sqrt(-2 * log(R1)))
    z1 = mu + sigma * (sin(2 * pi * R2) * sqrt(-2 * log(R1)))
    return z0, z1

# Оценка мат. ожидания
def math_expectation(values):
    return sum(values)/len(values)

# Оценка дисперсии
def dispersion(values):
    return (sum(list(map(lambda x: x * x, values))) - sum(values)**2/len(values))/(len(values) - 1)

# Плотность вероятности
def density_uniform(x):
    return 1 / (b - a) if a <= x <= b else 0

def density_exponential(x):
    return lambd * exp(-lambd * x) if x >= 0 else 0

def density_normal(x):
    return exp(-0.5 * ((x - mu) / sigma) ** 2) / (sigma * sqrt(2 * pi))

# Теоретическое мат.ожидание
math_expectation_uniform = (a + b)/2 
math_expectation_exponential = -1/lambd
math_expectation_normal = mu

# Теоретическая дисперсия
dispersion_uniform = (b - a)**2/12
dispersion_exponential = 1/lambd**2
dispersion_normal = sigma**2

# Выборки
N1 = 10
N2 = 20
N3 = 50
N4 = 100
N = 1000

random_values = random_sequence(N)

uniform = [uniform_distribution(R) for R in random_values]
exponential = [exponential_distribution(R) for R in random_values]
normal = []
for R in range(1, N, 2):
    normal.extend(normal_distribution(random_values[R - 1], random_values[R]))

print(f'''Равномерное распределение:
    Математическое ожидание:
        Выборка из {N1} чисел: {math_expectation(uniform[:N1])}
        Выборка из {N2} чисел: {math_expectation(uniform[:N2])}
        Выборка из {N3} чисел: {math_expectation(uniform[:N3])}
        Выборка из {N4} чисел: {math_expectation(uniform[:N4])}
        Выборка из {N} чисел: {math_expectation(uniform)}
        Теоретическое: {math_expectation_uniform}
        Абсолютная погрешность для выборки из {N} чисел: {abs(math_expectation(uniform) - math_expectation_uniform)}
        Относительная погрешность для выборки из {N} чисел: {abs(math_expectation(uniform) - math_expectation_uniform) / math_expectation_uniform * 100}
    Дисперсия:
        Выборка из {N1} чисел: {dispersion(uniform[:N1])}
        Выборка из {N2} чисел: {dispersion(uniform[:N2])}
        Выборка из {N3} чисел: {dispersion(uniform[:N3])}
        Выборка из {N4} чисел: {dispersion(uniform[:N4])}
        Выборка из {N} чисел: {dispersion(uniform)}
        Теоретическая: {dispersion_uniform}
        Абсолютная погрешность для выборки из {N} чисел: {abs(dispersion(uniform) - dispersion_uniform)}
        Относительная погрешность для выборки из {N} чисел: {abs(dispersion(uniform) - dispersion_uniform) / dispersion_uniform * 100}
Экспоненциальное распределение:
    Математическое ожидание:
        Выборка из {N1} чисел: {math_expectation(exponential[:N1])}
        Выборка из {N2} чисел: {math_expectation(exponential[:N2])}
        Выборка из {N3} чисел: {math_expectation(exponential[:N3])}
        Выборка из {N4} чисел: {math_expectation(exponential[:N4])}
        Выборка из {N} чисел: {math_expectation(exponential)}
        Теоретическое: {math_expectation_exponential}
        Абсолютная погрешность для выборки из {N} чисел: {abs(math_expectation(exponential) - math_expectation_exponential)}
        Относительная погрешность для выборки из {N} чисел: {abs(math_expectation(exponential) - math_expectation_exponential) / math_expectation_exponential * 100}
    Дисперсия:
        Выборка из {N1} чисел: {dispersion(exponential[:N1])}
        Выборка из {N2} чисел: {dispersion(exponential[:N2])}
        Выборка из {N3} чисел: {dispersion(exponential[:N3])}
        Выборка из {N4} чисел: {dispersion(exponential[:N4])}
        Выборка из {N} чисел: {dispersion(exponential)}
        Теоретическая: {dispersion_exponential}
        Абсолютная погрешность для выборки из {N} чисел: {abs(dispersion(exponential) - dispersion_exponential)}
        Относительная погрешность для выборки из {N} чисел: {abs(dispersion(exponential) - dispersion_exponential) / dispersion_exponential * 100}
Нормальное распределение:
    Математическое ожидание:
        Выборка из {N1} чисел: {math_expectation(normal[:N1])}
        Выборка из {N2} чисел: {math_expectation(normal[:N2])}
        Выборка из {N3} чисел: {math_expectation(normal[:N3])}
        Выборка из {N4} чисел: {math_expectation(normal[:N4])}
        Выборка из {N} чисел: {math_expectation(normal)}
        Теоретическое: {math_expectation_normal}
        Абсолютная погрешность для выборки из {N} чисел: {abs(math_expectation(normal) - math_expectation_normal)}
        Относительная погрешность для выборки из {N} чисел: {abs(math_expectation(normal) - math_expectation_normal) / math_expectation_normal * 100}
    Дисперсия:
        Выборка из {N1} чисел: {dispersion(normal[:N1])}
        Выборка из {N2} чисел: {dispersion(normal[:N2])}
        Выборка из {N3} чисел: {dispersion(normal[:N3])}
        Выборка из {N4} чисел: {dispersion(normal[:N4])}
        Выборка из {N} чисел: {dispersion(normal)}
        Теоретическая: {dispersion_normal}
        Абсолютная погрешность для выборки из {N} чисел: {abs(dispersion(normal) - dispersion_normal)}
        Относительная погрешность для выборки из {N} чисел: {abs(dispersion(normal) - dispersion_normal) / dispersion_normal * 100}''')

# Построение графиков
bins = 10
rwidth = 0.75
fontsize = 8

figure, axis = plt.subplots(3, 6)

X = linspace(a, b, 100)
for i in (axis[0, 0], axis[0, 1], axis[0, 2], axis[0, 3], axis[0, 4]):
    i.fill_between(X, [density_uniform(x) for x in X], color='pink')
for i in ((0, uniform[:N1]), (1, uniform[:N2]), (2, uniform[:N3]), (3, uniform[:N4]), (4, uniform)):
    axis[0, i[0]].hist(i[1], bins=bins, rwidth=rwidth, density=True, range=(a, b))
axis[0, 5].ecdf(uniform)
for i in ((0, N1), (1, N2), (2, N3), (3, N4), (4, N)):
    axis[0, i[0]].set_title(f'Равномерное\nN = {i[1]}', fontsize=fontsize)
axis[0, 5].set_title('Равномерное\nДиаграммы накопленных частот', fontsize=fontsize)

X = linspace(min(exponential), max(exponential), 100)
for i in (axis[1, 0], axis[1, 1], axis[1, 2], axis[1, 3], axis[1, 4]):
    i.fill_between(X, [density_exponential(x) for x in X], color='pink')
for i in ((0, exponential[:N1]), (1, exponential[:N2]), (2, exponential[:N3]), (3, exponential[:N4]), (4, exponential)):
    axis[1, i[0]].hist(i[1], bins=bins, rwidth=rwidth, density=True, range=(min(exponential), max(exponential)))
axis[1, 5].ecdf(exponential)
for i in ((0, N1), (1, N2), (2, N3), (3, N4), (4, N)):
    axis[1, i[0]].set_title(f'Экспоненциальное\nN = {i[1]}', fontsize=fontsize)
axis[1, 5].set_title('Экспоненциальное\nДиаграммы накопленных частот', fontsize=fontsize)

X = linspace(min(normal), max(normal), 100)
for i in (axis[2, 0], axis[2, 1], axis[2, 2], axis[2, 3], axis[2, 4]):
    i.fill_between(X, [density_normal(x) for x in X], color='pink')
for i in ((0, normal[:N1]), (1, normal[:N2]), (2, normal[:N3]), (3, normal[:N4]), (4, normal)):
    axis[2, i[0]].hist(i[1], bins=bins, rwidth=rwidth, density=True, range=(min(normal), max(normal)))
axis[2, 5].ecdf(normal)
for i in ((0, N1), (1, N2), (2, N3), (3, N4), (4, N)):
    axis[2, i[0]].set_title(f'Нормальное\nN = {i[1]}', fontsize=fontsize)
axis[2, 5].set_title('Нормальное\nДиаграммы накопленных частот', fontsize=fontsize)

figure.tight_layout(pad=0.5, w_pad=0, h_pad=0)
plt.show()


figure, axis = plt.subplots(3, 2)
axis[0, 0].set_title('Зависимость оценки мат. ожидания\nравномерного распределения\nот объёма выборки', fontsize=fontsize)
axis[0, 0].bar([str(N1), str(N2), str(N3), str(N4), str(N)], [math_expectation(uniform[:N1]), math_expectation(uniform[:N2]), math_expectation(uniform[:N3]), math_expectation(uniform[:N4]), math_expectation(uniform[:N])])
axis[0, 1].set_title('Зависимость оценки дисперсии\nравномерного распределения\nот объёма выборки', fontsize=fontsize)
axis[0, 1].bar([str(N1), str(N2), str(N3), str(N4), str(N)], [dispersion(uniform[:N1]), dispersion(uniform[:N2]), dispersion(uniform[:N3]), dispersion(uniform[:N4]), dispersion(uniform[:N])])
axis[1, 0].set_title('Зависимость оценки мат. ожидания\nэкспоненциального распределения\nот объёма выборки', fontsize=fontsize)
axis[1, 0].bar([str(N1), str(N2), str(N3), str(N4), str(N)], [math_expectation(exponential[:N1]), math_expectation(exponential[:N2]), math_expectation(exponential[:N3]), math_expectation(exponential[:N4]), math_expectation(exponential[:N])])
axis[1, 1].set_title('Зависимость оценки дисперсии\nэкспоненциального распределения\nот объёма выборки', fontsize=fontsize)
axis[1, 1].bar([str(N1), str(N2), str(N3), str(N4), str(N)], [dispersion(exponential[:N1]), dispersion(exponential[:N2]), dispersion(exponential[:N3]), dispersion(exponential[:N4]), dispersion(exponential[:N])])
axis[2, 0].set_title('Зависимость оценки мат. ожидания\nнормального распределения\nот объёма выборки', fontsize=fontsize)
axis[2, 0].bar([str(N1), str(N2), str(N3), str(N4), str(N)], [math_expectation(normal[:N1]), math_expectation(normal[:N2]), math_expectation(normal[:N3]), math_expectation(normal[:N4]), math_expectation(normal[:N])])
axis[2, 1].set_title('Зависимость оценки дисперсии\nнормального распределения\nот объёма выборки', fontsize=fontsize)
axis[2, 1].bar([str(N1), str(N2), str(N3), str(N4), str(N)], [dispersion(normal[:N1]), dispersion(normal[:N2]), dispersion(normal[:N3]), dispersion(normal[:N4]), dispersion(normal[:N])])
figure.tight_layout()
plt.show()