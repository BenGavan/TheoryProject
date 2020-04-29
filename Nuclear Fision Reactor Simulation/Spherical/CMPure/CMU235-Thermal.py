import matplotlib.pyplot as plt
import numpy as np
from scipy.interpolate import interp1d

rs = np.array([])
neutrons_produced = np.array([])
neutrons_produced_uncert = np.array([])
ks = np.array([])
k_uncerts = np.array([])

f = open("data/sphere-critical-mass-plot-thermal-depends-on-previous-300.txt")
for line in f:
    l = line.split()

    r = float(l[0]) * 100
    produced = float(l[3])
    producedUncert = float(l[4])
    k = float(l[1])
    kUncert = float(l[2])

    rs = np.append(rs, r)
    neutrons_produced = np.append(neutrons_produced, produced)
    neutrons_produced_uncert = np.append(neutrons_produced_uncert, producedUncert)
    ks = np.append(ks, k)
    k_uncerts = np.append(k_uncerts, kUncert)

f.close()

# print(rs)
# print(neutrons_produced)
# print(neutrons_produced_uncert)
# print(ks)
# print(k_uncerts)

plt.grid(color='0.5', linestyle='-', linewidth=1, alpha=0.2)
plt.errorbar(x=rs, y=ks, yerr=k_uncerts, fmt="k.", ecolor="grey", elinewidth=1, ms=1, marker=".", label='Data')
# plt.plot([min(rs), max(rs)], [1,1], color = '0.75', label='Critical mass k value')
plt.plot([min(rs), max(rs)], [2.077199113499998, 2.077199113499998], color='0.3', label='Asymptote (k$= 2.077 \pm 0.004)$')
plt.xlabel('Sphere radius (cm)')
plt.ylabel('k value')
plt.legend()
plt.savefig("spherical-critical-mass-u235-thermal.png")
plt.show()

# f = interp1d(ks, rs, kind='quadratic')
# r = f(1)
#
# ys = np.linspace(0.1,1.48, 100)
# xs = f(ys)
# plt.title('quadratic')
# plt.plot(xs, ys)
# plt.show()
# print("r = ", r)


f = interp1d(ks, rs, kind='cubic')
r = f(1)

ys = np.linspace(min(ks), max(ks), 100)
xs = f(ys)
plt.title('cubic')
plt.plot(xs, ys)
plt.show()
print("(cubic) r = ", r)

if __name__ == '__main__':
    pass


