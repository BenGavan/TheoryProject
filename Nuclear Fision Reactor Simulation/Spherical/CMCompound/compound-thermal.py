import matplotlib.pyplot as plt
import numpy as np

f = open("data/sphere-critical-mass-plot-fast-u235-238-compound.txt")

u235s = np.array([])
rs = np.array([])

for line in f:
    l = line.split()

    u = float(l[0])
    r = float(l[1]) * 100

    u235s = np.append(u235s, u)
    rs = np.append(rs, r)

f.close()

r_uncerts = np.array([0.01] * len(rs))

plt.grid(color='0.5', linestyle='-', linewidth=1, alpha=0.2)
plt.errorbar(x=u235s, y=rs, yerr=r_uncerts, fmt='k.', elinewidth=1, ms=3.5)
# plt.plot(u235s, rs, 'k.')
plt.xlabel('Proportion of U-235')
plt.ylabel('Critical Radius (cm)')
plt.savefig('plots/sphere-critical-mass-plot-fast-u235-238-compound.png')
plt.show()

if __name__ == '__main__':
    pass
