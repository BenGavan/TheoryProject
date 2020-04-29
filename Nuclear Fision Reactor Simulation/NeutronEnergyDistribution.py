import numpy as np
import matplotlib.pyplot as plt
import math

print("Neutron Energy Distribution")


def func1(x):
    return 0.4865 * np.sqrt(2 * 6 * pow(10, 16) / 939.9) * np.sqrt(x) * np.sinh(np.sqrt(2 * x)) * pow(np.e, -x)


def func(x):
    return 0.4865 * np.sinh(np.sqrt(2 * x)) * pow(np.e, -x)


def integral(x):
    """
    cumulative probability distribution for the energy spectrum of u235 + n -> u236
    Source: https://indico.cern.ch/event/145296/contributions/1381141/attachments/136909/194258/lecture24.pdf
    :param x:
    :return:
    """
    first = -pow(np.e, - np.sqrt(2) * np.sqrt(x) - x) * (-1 + pow(np.e, 2 * np.sqrt(2) * np.sqrt(x)))
    second = np.sqrt(np.e * np.pi / 2) * math.erf(1 / np.sqrt(2) - np.sqrt(x))
    third = np.sqrt(np.e * np.pi / 2) * math.erf(1 / np.sqrt(2) + np.sqrt(x))
    return 0.24325 * (first - second + third)


xs = np.linspace(0, 10, 1000)
ys = []
for x in xs:
    ys.append(integral(x))
# ys = integral(xs)

plt.title("cumulative probability distribution")
plt.plot(xs, ys)

plt.show()
