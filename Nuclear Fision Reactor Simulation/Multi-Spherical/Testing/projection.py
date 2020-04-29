import numpy as np
import matplotlib.pyplot as plt

f = open("data.txt")

xs = np.array([])
ys = np.array([])
zs = np.array([])

for line in f:
    l = line.split()

    x = float(l[0])
    y = float(l[1])
    z = float(l[2])

    xs = np.append(xs, x)
    ys = np.append(ys, y)
    zs = np.append(zs, z)

f.close()

plt.plot(xs, ys, "kx")
plt.show()

if __name__ == '__main__':
    pass