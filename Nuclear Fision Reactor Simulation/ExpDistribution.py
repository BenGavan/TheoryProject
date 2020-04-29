# Plots the histogram of the length of the random walks for a specific mean free path

import numpy as np
import matplotlib.pyplot as plt

lam = 0.1

print("ExpDist")
#
# def p(x):
#     return pow(np.e, - x / lam)
#
#
# xs = np.linspace(0, 50, 1000)
# ys = 1 - p(xs) / lam
#
# print(xs, ys)
#
# plt.plot(xs, ys)
# plt.show()

def x():
    u = np.random.uniform(0, 1, 1)[0]
    return - lam * np.log(u)

n = 10000
xs = np.zeros(n)
for i in range(n):
    xs[i] = x()

axes = plt.gca()
plt.hist(xs, bins=100, ec='black')
axes.set_xlim([0, 1])
plt.xlabel("distances values")
plt.ylabel("Frequency")
plt.title("Histogram of lengths of individual walks for $\lambda$ = {}, $n={}$".format(lam, n))
plt.show()

def f(x):
    return pow(np.e, - x / lam)

xs = np.linspace(0, 5, 100)
plt.plot(xs, f(xs))
plt.show()

