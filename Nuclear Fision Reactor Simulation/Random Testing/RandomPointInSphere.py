import matplotlib as mpl
from mpl_toolkits.mplot3d import Axes3D
import numpy as np
import matplotlib.pyplot as plt

f = open("data/random-points-sphere.txt")

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


fig = plt.figure()
ax = fig.gca(projection='3d')

ax.plot(xs, ys, zs, "k.")

ax.set_xlabel('x')
ax.set_ylabel('y')
ax.set_zlabel('z')

ax.set_xlim([-1, 1])
ax.set_ylim([-1, 1])
ax.set_zlim([-1, 1])


plt.savefig("plots/randomPointInSphere2000.png")
plt.show()
