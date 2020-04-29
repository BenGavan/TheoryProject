import numpy as np
import matplotlib.pyplot as plt

f = open("data/sphere-critical-mass-plot-fast-pure-u235-varying-heavy-water.txt")

ws_h = np.array([])
crs_h = np.array([])
crs_uncert_h = np.array([])

for line in f:
    l = line.split()

    w = float(l[0]) * 100
    cr = float(l[1]) * 100

    ws_h = np.append(ws_h, w)
    crs_h = np.append(crs_h, cr)
    crs_uncert_h = np.append(crs_uncert_h, 0.01 / 100)

f.close()

f = open("data/sphere-critical-mass-plot-fast-pure-u235-varying-water.txt")

ws = np.array([])
crs = np.array([])
crs_uncert = np.array([])

for line in f:
    l = line.split()

    w = float(l[0]) * 100
    cr = float(l[1]) * 100

    ws = np.append(ws, w)
    crs = np.append(crs, cr)
    crs_uncert = np.append(crs_uncert, 0.01 / 100)

f.close()

plt.errorbar(ws, crs, crs_uncert, fmt="C1x", label="H$_2$O")
plt.errorbar(ws_h, crs_h, crs_uncert_h, fmt="bx", label="D$_2$O")

plt.grid(color='0.5', linestyle='-', linewidth=1, alpha=0.2)
plt.xlabel('Water shell thickness (cm)')
plt.ylabel('Critical Radius (cm)')
# plt.ylim(0.025,  0.03)
plt.legend()
plt.savefig("plots/sphere-heavy-water-and-water-fast.png")
plt.show()

if __name__ == '__main__':
    pass