import numpy as np
import matplotlib.pyplot as plt
import matplotlib as mpl
from mpl_toolkits.mplot3d import Axes3D


def random_point_in_sphere(radius):
    rands = np.random.uniform(-1, 1, 3)
    x = rands[0] * radius
    y = rands[1] * radius
    z = rands[2] * radius

    r2 = pow(x, 2) + pow(y, 2) + pow(z, 2)
    r = np.sqrt(r2)
    if r < radius:
        return x, y, z
    else:
        return random_point_in_sphere(radius)


number_neutrons = 10000

xs = np.zeros(number_neutrons)
ys = np.zeros(number_neutrons)
zs = np.zeros(number_neutrons)

for i in range(number_neutrons):
    xs[i], ys[i], zs[i] = random_point_in_sphere(1)

plt.plot(xs, ys, ".k")
plt.xlabel("x")
plt.ylabel("y")
plt.savefig("plots/RandomPointInSphere/xy.png")
plt.show()

plt.plot(ys, zs, ".k")
plt.xlabel("y")
plt.ylabel("z")
plt.savefig("plots/RandomPointInSphere/yz.png")
plt.show()

plt.plot(xs, zs, ".k")
plt.xlabel("x")
plt.ylabel("z")
plt.savefig("plots/RandomPointInSphere/xz.png")
plt.show()


# import pickle


mpl.rcParams['legend.fontsize'] = 8

fig = plt.figure()
ax = fig.gca(projection='3d')

ax.plot(xs, ys, zs, "k.")

# pickle.dump(fig, open('FigureObject.fig.pickle', 'wb'))
#
# figx = pickle.load(open('FigureObject.fig.pickle', 'rb'))

# figx.show()
plt.show()
