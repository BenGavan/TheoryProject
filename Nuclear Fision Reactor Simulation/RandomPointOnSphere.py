import numpy as np
import matplotlib.pyplot as plt


def magnitude(x, y, z):
    r2 = pow(x, 2) + pow(y, 2) + pow(z, 2)
    return np.sqrt(r2)


def random_phi():
    p = np.random.uniform(0, 1, 1)[0]
    phi = 2 * np.pi * p
    return phi


def random_z():
    return np.random.uniform(0, 1, 1)[0]


def random_x(phi, z):
    return np.sin(np.arccos(z)) * np.cos(phi)


def random_y(phi, z):
    return np.sin(np.arccos(z)) * np.sin(phi)


def random_point():
    phi = random_phi()
    z = random_z()
    x = random_x(phi, z)
    y = random_y(phi, z)
    return x, y, x


n = 1000000
xs = np.zeros(n, dtype=float)
ys = np.zeros(n, dtype=float)
zs = np.zeros(n, dtype=float)

for i in range(n):
    p = random_point()
    xs[i] = p[0]
    ys[i] = p[1]
    zs[i] = p[2]

plt.hist(xs, bins=50, ec='black')
plt.title("Distribution of x values on a sphere")
plt.xlabel("x values")
plt.ylabel("frequency")
plt.savefig("plots/x_sphere_hist_n={}.png".format(n))
plt.show()

plt.hist(ys, bins=50, ec='black')
plt.title("Distribution of y values on a sphere")
plt.xlabel("y values")
plt.ylabel("frequency")
plt.savefig("plots/y_sphere_hist_n={}.png".format(n))
plt.show()

plt.hist(zs, bins=50, ec='black')
plt.title("Distribution of z values on a sphere")
plt.xlabel("x values")
plt.ylabel("frequency")
plt.savefig("plots/z_sphere_hist_n={}.png".format(n))
plt.show()
