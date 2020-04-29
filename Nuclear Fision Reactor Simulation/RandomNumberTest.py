import numpy as np
import matplotlib.pyplot as plt

random_numbers = np.random.uniform(0, 1, 1000000)
x, y, z = plt.hist(random_numbers, bins=100, ec='black')
print(x, y)
plt.show()

mean = np.std(x)
print("std = {}".format(mean))