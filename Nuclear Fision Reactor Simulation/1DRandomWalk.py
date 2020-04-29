import numpy as np
import matplotlib.pyplot as plt

print("1D Random Walk")

n = 100000000
randoms = np.random.uniform(-1, 1, n)

moves = np.zeros(n, dtype=bool)
for i in range(n):
    moves[i] = randoms[i] < 0

position = 0
positions = np.zeros(n, dtype=int)
for i in range(n):
    if moves[i]:
        position += 1
    else:
        position -= 1
    positions[i] = position

plt.plot(np.linspace(0, n, n), positions)
plt.plot([0, n], [0, 0])
plt.show()