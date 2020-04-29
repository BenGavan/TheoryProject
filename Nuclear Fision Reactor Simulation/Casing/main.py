import numpy as np
import matplotlib.pyplot as plt

f = open("data/fast-pure-u235-casing.txt")

cs = np.array([])
rs = np.array([])
r_uncerts = np.array([])

for line in f:
    l = line.split()

    print(l)

    c = float(l[0]) * 100
    r = float(l[1]) * 100

    cs = np.append(cs, c)
    rs = np.append(rs, r)
    r_uncerts = np.append(r_uncerts, 0.00001)


plt.grid(color='0.5', linestyle='-', linewidth=1, alpha=0.2)

plt.errorbar(cs, rs, r_uncerts, fmt="kx")
# plt.ylim(0.0275, 0.0283)
plt.ylabel("Critical radius (cm)")
plt.xlabel("Shielding Thickness (cm)")
plt.tight_layout()
plt.savefig("plots/casing-fast.png")
plt.show()


if __name__ == "__main__":
    pass