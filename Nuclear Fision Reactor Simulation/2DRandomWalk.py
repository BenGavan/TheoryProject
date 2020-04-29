import numpy as np
import matplotlib.pyplot as plt


class Walker(object):

    x = 0
    y = 0

    x_history = np.array([0])
    y_history = np.array([0])

    def move(self):
        self.x += np.random.uniform(-1, 1, 1)[0]
        self.y += np.random.uniform(-1, 1, 1)[0]
        self.x_history = np.append(self.x_history, self.x)
        self.y_history = np.append(self.y_history, self.y)

    def plot(self):
        plt.plot(self.x_history, self.y_history)
        plt.plot(0, 0, "kx")
        plt.show()


w = Walker()
for _ in range(100):
    w.move()


w.plot()