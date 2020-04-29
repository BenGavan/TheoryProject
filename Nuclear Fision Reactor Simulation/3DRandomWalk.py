import matplotlib as mpl
from mpl_toolkits.mplot3d import Axes3D
import numpy as np
import matplotlib.pyplot as plt


class Neutron(object):

    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

        self.start_x = x
        self.start_y = y
        self.start_z = z

        self.x_history = np.array([x])
        self.y_history = np.array([y])
        self.z_history = np.array([z])

    def move(self):
        self.x += np.random.uniform(-1, 1, 1)[0]
        self.y += np.random.uniform(-1, 1, 1)[0]
        self.z += np.random.uniform(-1, 1, 1)[0]
        self.x_history = np.append(self.x_history, self.x)
        self.y_history = np.append(self.y_history, self.y)
        self.z_history = np.append(self.z_history, self.z)

    def plot(self):
        mpl.rcParams['legend.fontsize'] = 8

        fig = plt.figure()
        ax = fig.gca(projection='3d')

        ax.plot(self.x_history, self.y_history, self.z_history, label='Neutron Path')
        ax.plot([self.start_x], [self.start_y], [self.start_z], "rx", label="Start Position")
        ax.plot([self.x_history[-1]], [self.y_history[-1]], [self.z_history[-1]], "kx", label="End Position")
        ax.legend()

        plt.show()


n = Neutron(0, 0, 0)
for _ in range(1000):
    n.move()

n.plot()
