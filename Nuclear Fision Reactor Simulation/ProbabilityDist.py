import numpy as np
import matplotlib.pyplot as plt

ps = np.linspace(0, 0.9999999999999999, 10000)


def func(p):
    lam = 1
    return - lam * np.log(1- p)


def inverse(x):
    lam = 1
    return 1 - np.power(np.e, - x / lam)


xs = func(ps)

plt.plot(xs, ps)
plt.show()



xs = np.linspace(0, 10, 100)

ys = inverse(xs)

plt.plot(ys, xs)
plt.show()
