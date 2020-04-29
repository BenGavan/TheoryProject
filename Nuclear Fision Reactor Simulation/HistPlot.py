import matplotlib.pyplot as plt
import numpy as np

f = open("data2/magnitudes-of-steps.txt")

xs = np.array([])
for v in f:
    x = float(v)
    xs = np.append(xs, x)

f.close()

plt.hist(xs, bins=100, ec='black')
plt.savefig("plots/stepsizehist.png")
plt.show()


