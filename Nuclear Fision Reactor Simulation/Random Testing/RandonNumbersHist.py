import matplotlib.pyplot as plt
import numpy as np

f = open("data/random-numbers.txt")

xs = np.array([])
for v in f:
    x = float(v)
    xs = np.append(xs, x)

f.close()

# xs = [0, 0.999, 0.1, 0.2, 0.3, .4, .5, .6, .7, .8, .9]

x_label = "Random values generated"
y_label = "Frequency"
#
# print("plotting 2")
#
# plt.hist(xs, bins=2, color="grey", ec='black')
# plt.xlabel(x_label)
# plt.ylabel(y_label)
# plt.savefig("plots/randomNumberHist2.png", ppi=300)
# plt.show()
#
# print("plotting 10")
#
# plt.hist(xs, bins=10, color="grey", ec='black')
# plt.xlabel(x_label)
# plt.ylabel(y_label)
# plt.savefig("plots/randomNumberHist10.png", ppi=300)
# plt.show()
#
# print("plotting 20")
#
# plt.hist(xs, bins=20, color="grey", ec='black')
# plt.xlabel(x_label)
# plt.ylabel(y_label)
# plt.savefig("plots/randomNumberHist20.png", ppi=300)
# plt.show()

print("plotting 100")

x = plt.hist(xs, bins=100, color="grey", ec='black')[0]
plt.savefig("plots/randomNumberHist100Second.png", ppi=300)
plt.show()

print(x)

print('Plotting 100')

plt.hist(x, bins=100, color="grey", ec='grey')
plt.xlabel('Quantity in each bin')
plt.ylabel('Frequency')
plt.savefig("plots/randomNumberHistOfHist100.png")
plt.show()

print('Plotting 10')

plt.hist(x, bins=10, color="grey", ec='grey')
plt.xlabel('Quantity in each bin')
plt.ylabel('Frequency')
plt.savefig("plots/randomNumberHistOfHist10.png")
plt.show()

print('Plotting 20')

plt.hist(x, bins=20, color="grey", ec='grey')
plt.xlabel('Quantity in each bin')
plt.ylabel('Frequency')
plt.savefig("plots/randomNumberHistOfHist20.png")
plt.show()

print('Plotting 50')

plt.hist(x, bins=50, color="grey", ec='grey')
plt.xlabel('Quantity in each bin')
plt.ylabel('Frequency')
plt.savefig("plots/randomNumberHistOfHist50.png")
plt.show()

print('Plotting 200')

plt.hist(x, bins=200, color="grey", ec='grey')
plt.xlabel('Quantity in each bin')
plt.ylabel('Frequency')
plt.savefig("plots/randomNumberHistOfHist200.png")
plt.show()

# plt.show()
#
# print("plotting 1000")
#
# plt.hist(xs, bins=1000, color="grey", ec='grey')
# plt.xlabel(x_label)
# plt.ylabel(y_label)
# plt.savefig("plots/randomNumberHist1000.png", ppi=300)
# plt.show()
