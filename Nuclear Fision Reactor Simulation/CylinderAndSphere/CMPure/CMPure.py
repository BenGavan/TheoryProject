import matplotlib.pyplot as plt
import numpy as np
from scipy.interpolate import interp1d

rs_c = np.array([])
neutrons_produced_c = np.array([])
neutrons_produced_uncert_c = np.array([])
ks_c = np.array([])
k_uncerts_c = np.array([])

f = open("data/cylindrical-critical-mass-plot-thermal-depends-on-previous.txt")
for line in f:
    l = line.split()

    r = float(l[0]) * 100
    produced = float(l[3])
    producedUncert = float(l[4])
    k = float(l[1])
    kUncert = float(l[2])

    rs_c = np.append(rs_c, r)
    neutrons_produced_c = np.append(neutrons_produced_c, produced)
    neutrons_produced_uncert_c = np.append(neutrons_produced_uncert_c, producedUncert)
    ks_c = np.append(ks_c, k)
    k_uncerts_c = np.append(k_uncerts_c, kUncert)

f.close()

rs_s = np.array([])
neutrons_produced_s = np.array([])
neutrons_produced_uncert_s = np.array([])
ks_s = np.array([])
k_uncerts_s = np.array([])

f = open("data/sphere-critical-mass-plot-thermal-depends-on-previous-300.txt")
for line in f:
    l = line.split()

    r = float(l[0]) * 100
    produced = float(l[3])
    producedUncert = float(l[4])
    k = float(l[1])
    kUncert = float(l[2])

    # if r > 0.1:
    #     continue

    rs_s = np.append(rs_s, r)
    neutrons_produced_s = np.append(neutrons_produced_s, produced)
    neutrons_produced_uncert_s = np.append(neutrons_produced_uncert_s, producedUncert)
    ks_s = np.append(ks_s, k)
    k_uncerts_s = np.append(k_uncerts_s, kUncert)

f.close()

# print(rs)
# print(neutrons_produced)
# print(neutrons_produced_uncert)
# print(ks)
# print(k_uncerts)

print(rs_c)

plt.grid(color='0.5', linestyle='-', linewidth=1, alpha=0.2)
plt.errorbar(x=rs_c, y=ks_c, yerr=k_uncerts_c, fmt="k.", ecolor="blue", elinewidth=1, ms=1, marker=".", label='Cylinder (length = 200 cm)')
# plt.plot([min(rs), max(rs)], [1,1], color = '0.75', label='Critical mass k value')
plt.plot([min(rs_c), max(rs_c)], [2.077199113499998, 2.077199113499998], color='0.3', label='Asymptote (k$= 2.077 \pm 0.004)$')

plt.errorbar(x=rs_s, y=ks_s, yerr=k_uncerts_s, fmt="k.", ecolor="red", elinewidth=1, ms=1, marker=".", label='Sphere')
# plt.plot([min(rs), max(rs)], [1,1], color = '0.75', label='Critical mass k value')
# plt.plot([min(rs_s), max(rs_s)], [2.077199113499998, 2.077199113499998], color='0.3', label='Asymptote (k$= 2.077 \pm 0.004)$')

plt.xlabel('Cylinder radius (cm)')
plt.ylabel('k value')
plt.legend()
plt.savefig("plots/cylinder-sphere-critical-mass-u235-fast.png")
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


f = interp1d(ks_c, rs_c, kind='cubic')
r = f(1)

ys = np.linspace(min(ks_c), max(ks_c), 100)
xs = f(ys)
plt.title('cubic')
plt.plot(xs, ys)
plt.show()
print(" cylinder(cubic) r = ", r)

if __name__ == '__main__':
    pass


