import numpy as np
import matplotlib.pyplot as plt

f = open("data/results.txt")

u235s = np.array([])

rs = np.array([])
r_uncerts = np.array([])

for line in f:
    l = line.split()

    u235 = float(l[0]) * 100
    u238 = 1 - u235
    r = float(l[2]) * 100

    r_uncert = float(l[3]) * 100

    u235s = np.append(u235s, u235)
    rs = np.append(rs, r)
    r_uncerts = np.append(r_uncerts, r_uncert)

f.close()

plt.xlabel('Percentage of U-235')
plt.ylabel('Critical Radius (cm)')
plt.errorbar(x=u235s, y=rs, yerr=r_uncerts, fmt='k.')
plt.savefig('plots/u235-u238-cm-plot.png')
plt.show()