import numpy as np
import matplotlib.pyplot as plt
from scipy.interpolate import interp1d

rs = np.array([])
ks = np.array([])
k_uncerts = np.array([])

f = open('sphere-cm-u235-vac-fast.txt')
for line in f:
    l = line.split()
    print(l)

    r = float(l[3].replace(',', '')) * 100
    k = float(l[6].replace(',', ''))
    k_uncert = float(l[8].replace(',', ''))

    print(r, k, k_uncert)

    rs = np.append(rs, r)
    ks = np.append(ks, k)
    k_uncerts = np.append(k_uncerts, k_uncert)

f.close()

plt.grid(color='0.5', linestyle='-', linewidth=1, alpha=0.2)
plt.errorbar(x=rs, y=ks, yerr=k_uncerts, fmt="kx", ecolor="grey", elinewidth=1, ms=1, marker=".", label='Data')
plt.plot([min(rs), max(rs)], [2.0810221, 2.0810221], color='0.3', label='Asymptote (k$= 2.081 \pm 0.008)$')
plt.xlabel('Sphere radius (cm)')
plt.ylabel('k value')
plt.legend()
plt.savefig("spherical-critical-mass-u235-fast-vac.png")
plt.show()


#
# rs = np.array([])
# neutrons_produced = np.array([])
# neutrons_produced_uncert = np.array([])
# ks = np.array([])
# k_uncerts = np.array([])
#
# f = open("sphere-critical-mass-plot-250.txt")
# for line in f:
#     l = line.split()
#
#     r = float(l[0]) * 100
#     produced = float(l[1])
#     producedUncert = float(l[2])
#     k = float(l[3])
#     kUncert = float(l[4])
#
#     rs = np.append(rs, r)
#     neutrons_produced = np.append(neutrons_produced, produced)
#     neutrons_produced_uncert = np.append(neutrons_produced_uncert, producedUncert)
#     ks = np.append(ks, k)
#     k_uncerts = np.append(k_uncerts, kUncert)
#
# f.close()
#
# plt.errorbar(x=rs, y=ks, yerr=k_uncerts, fmt="k.", ecolor="grey", elinewidth=1, ms=1, marker=".", label='Data')
# plt.savefig("spherical-critical-mass-u235-fast-and-thermal-vac.png")
# plt.show()


# f = interp1d(ks, rs, kind='quadratic')
# r = f(1)
#
# ys = np.linspace(min(ks), max(ks), 100)
# xs = f(ys)
# plt.title('cubic')
# plt.plot(xs, ys)
# plt.show()
# print("(cubic) r = ", r)