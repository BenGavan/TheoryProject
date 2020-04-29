import numpy as np
import matplotlib.pyplot as plt

lam = 1


def f(x):
    return 1 - pow(np.e, -x / lam)


def x(u):
    return -lam * np.log(1 - u)

xs = np.linspace(0, 10, 100)
plt.plot(xs, f(xs))
plt.show()
