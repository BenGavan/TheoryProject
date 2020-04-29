import matplotlib as mpl
from mpl_toolkits.mplot3d import Axes3D
import numpy as np
import matplotlib.pyplot as plt
from enum import Enum


print("Infinite 2D Reactor")


def random_point_in_sphere(radius):
    rands = np.random.uniform(-1, 1, 3)
    x = rands[0] * radius
    y = rands[1] * radius
    z = rands[2] * radius

    r2 = pow(x, 2) + pow(y, 2) + pow(z, 2)
    r = np.sqrt(r2)
    if r < radius:
        return x, y, z
    else:
        return random_point_in_sphere(radius)


# TODO - Test the histogram of the theta and phi values calculated
class NeutronState(Enum):
    scatter = 0
    absorb = 1
    fission = 2
    nothing = 3


class Medium(Enum):
    u235 = 0
    water = 1


class RandomUnitVector(object):

    def __init__(self):
        self.x, self.y, self.z = random_point_in_sphere(1)

        magnitude = self.magnitude()
        self.x = self.x / magnitude
        self.y = self.y / magnitude
        self.z = self.z / magnitude

    def magnitude(self):
        r2 = pow(self.x, 2) + pow(self.y, 2) + pow(self.z, 2)
        return np.sqrt(r2)


class Vector(object):

    def __init__(self, unit_vector, magnitude):
        self.unit_vector = unit_vector
        self.x = unit_vector.x * magnitude
        self.y = unit_vector.y * magnitude
        self.z = unit_vector.z * magnitude

    def add(self, magnitude):
        self.x += magnitude * self.unit_vector.x
        self.y += magnitude * self.unit_vector.y
        self.z += magnitude * self.unit_vector.z

    def magnitude(self):
        r2 = pow(self.x, 2) + pow(self.y, 2) + pow(self.z, 2)
        return np.sqrt(r2)

    def print_magnitude(self):
        print("Magnitude = {}".format(self.magnitude()))

    def to_string(self):
        return "({}, {}, {})".format(self.x, self.y, self.z)

    def print(self):
        print(self.to_string())


class Neutron(object):

    fast_scatter_lambda_u235 = 5.107798
    fast_absorb_lambda_u235 = 227.013246
    fast_fission_lambda_u235 = 20.43119214

    scatter_lambda_u235 = 204.31192
    absorb_lambda_u235 = 20.63756782
    fission_lambda_u235 = 3.504492648

    scatter_lambda_water = 1
    absorb_lambda_water = 10

    is_free = True

    state = NeutronState.nothing
    generation = 0

    scatters_in_water = 0

    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

        self.start_x = x
        self.start_y = y
        self.start_z = z

        self.x_history = np.array([x])
        self.y_history = np.array([y])
        self.z_history = np.array([z])

        self.lam = pow(1/self.scatter_lambda_u235 + 1/self.absorb_lambda_u235 + 1/self.fission_lambda_u235, -1)

        self.vector = Vector(RandomUnitVector(), 1)

        self.medium = Medium.u235

    def random_distance(self):
        u = np.random.uniform(0, 1, 1)[0]
        magnitude = - self.lam * np.log(1 - u)
        return magnitude

    def move_distance(self):

        delta_distance = self.random_distance()

        if self.state == NeutronState.nothing:
            self.vector.add(delta_distance)
        else:
            v = RandomUnitVector()
            self.vector = Vector(v, delta_distance)

        self.x += self.vector.x
        self.y += self.vector.y
        self.z += self.vector.z

        self.x_history = np.append(self.x_history, self.x)
        self.y_history = np.append(self.y_history, self.y)
        self.z_history = np.append(self.z_history, self.z)

    def is_in_sphere(self):
        r2 = pow(self.x, 2) + pow(self.y, 2) + pow(self.z, 2)
        r = np.sqrt(r2)
        if r < radius:
            return True
        return False

    def move(self):
        if not self.is_free:
            return

        if self.scatters_in_water > 100:
            self.is_free = False
            print("Neutrons scattered more than 100 times in water")
            return

        self.generation += 1

        self.move_distance()

        magnitude = self.vector.magnitude()

        if self.is_in_sphere():
            if self.medium == Medium.water:
                self.scatters_in_water = 0
            self.medium = Medium.u235
            p_scatter = 1 - pow(np.e, - magnitude / self.scatter_lambda_u235)
            p_absorb = 1 - pow(np.e, - magnitude / self.absorb_lambda_u235)
            p_fission = 1 - pow(np.e, -magnitude / self.fission_lambda_u235)

            p_total = p_scatter + p_absorb + p_fission

            if p_total > 1:
                p_scatter = p_scatter / p_total
                p_absorb = p_absorb / p_total
                p_fission = p_fission / p_total
                p_total = p_scatter + p_absorb + p_fission

            a = np.random.uniform(0, 1, 1)[0]
            if a <= p_scatter:
                self.state = NeutronState.scatter
                return NeutronState.scatter
            elif p_scatter < a <= (p_scatter + p_absorb):
                self.is_free = False
                self.state = NeutronState.absorb
                return NeutronState.absorb
            elif (p_scatter + p_absorb) < a <= (p_scatter + p_absorb + p_fission):
                self.is_free = False
                self.state = NeutronState.fission
                return NeutronState.fission
            else:
                self.state = NeutronState.nothing
                return NeutronState.nothing

        else:
            self.scatters_in_water += 1
            self.medium = Medium.water
            p_scatter = 1 - pow(np.e, - magnitude / self.scatter_lambda_water)
            p_absorb = 1 - pow(np.e, - magnitude / self.absorb_lambda_water)

            p_total = p_scatter + p_absorb

            if p_total > 1:
                p_scatter = p_scatter / p_total
                p_absorb = p_absorb / p_total
                p_total = p_scatter + p_absorb

            a = np.random.uniform(0, 1, 1)[0]
            if a <= p_scatter:
                self.state = NeutronState.scatter
                return NeutronState.scatter
            elif p_scatter < a <= (p_scatter + p_absorb):
                self.is_free = False
                self.state = NeutronState.absorb
                return NeutronState.absorb
            else:
                self.state = NeutronState.nothing
                return NeutronState.nothing




    def get_last_move_distance(self):
        dx = self.x_history[-1] - self.x_history[-2]
        dy = self.y_history[-1] - self.y_history[-2]
        dz = self.z_history[-1] - self.z_history[-2]

        r2 = pow(dx, 2) + pow(dy, 2) + pow(dz, 2)
        r = np.sqrt(r2)
        return r

    def get_position(self):
        return self.x, self.y, self.z

    def plot(self):
        mpl.rcParams['legend.fontsize'] = 8

        fig = plt.figure()
        ax = fig.gca(projection='3d')

        ax.plot(self.x_history, self.y_history, self.z_history, label='Neutron Path')
        ax.plot([self.start_x], [self.start_y], [self.start_z], "rx", label="Start Position")
        ax.plot([self.x_history[-1]], [self.y_history[-1]], [self.z_history[-1]], "kx", label="End Position")
        ax.legend()

        plt.show()


def perform_generation(start_neutrons):
    neutrons = start_neutrons
    current_number_neutrons = 0

    next_generation_neutrons = np.array([])

    iterations = 0

    while True:
        iterations += 1
        number_free_neutrons = 0

        xs = np.zeros(len(neutrons), dtype=float)
        ys = np.zeros(len(neutrons), dtype=float)
        zs = np.zeros(len(neutrons), dtype=float)

        i = 0

        for n in neutrons:
            if not n.is_free:
                continue

            number_free_neutrons += 1

            n.move()

            xs[i] = n.x
            ys[i] = n.y
            zs[i] = n.z

            if n.state == NeutronState.fission:
                rand = np.random.uniform(0, 1, 1)[0]
                if rand <= 0.7:
                    # 2 neutrons
                    current_number_neutrons += 2
                else:
                    # 3 neutrons
                    current_number_neutrons += 3
            i += 1

        # fig = plt.figure()
        # ax = fig.gca(projection='3d')
        #
        # ax.plot(xs, ys, zs, "k.")
        # plt.show()

        if number_free_neutrons == 0:
            # print("iterations = {}".format(iterations))
            return current_number_neutrons, iterations


def average_produced(radius):
    number_history = np.zeros(generations)
    iteration_history = np.zeros(generations)

    for i in range(generations):
        neutrons = np.zeros(number_neutrons, dtype=Neutron)

        for j in range(number_neutrons):
            x, y, z = random_point_in_sphere(radius)
            neutrons[j] = Neutron(x, y, z)

        neutrons_produced, iterations = perform_generation(neutrons)
        number_history[i] = neutrons_produced
        iteration_history[i] = iterations

    average_produced = np.average(number_history)
    print("Average Produced = {}".format(average_produced))

    xs = np.linspace(0, generations, generations)
    plt.plot(xs, number_history)
    plt.plot([0, generations], [average_produced, average_produced])
    plt.title("Generation in water, radius = {}".format(radius))
    plt.xlabel("iteration")
    plt.ylabel("Neutrons produced")
    plt.show()

    return average_produced


radius = 6.664
number_neutrons = 1000
generations = 200

tolerance = 1
step = 0.001

radius_history = np.array([])
average_history = np.array([])

while True:

    average = average_produced(radius)

    print("Radius = {}".format(radius))

    radius_history = np.append(radius_history, radius)
    average_history = np.append(average_history, average)

    delta = number_neutrons - average

    if delta < -tolerance:
        radius -= step
    elif delta > tolerance:
        radius += step
    else:
        step = step * 0.1
        tolerance = tolerance * 0.1
        if step < pow(10, -4):
            break



print(radius)

plt.plot(radius_history, average_history, "kx")
plt.title("Spherical U-235 in an water (n = {})".format(number_neutrons))
plt.xlabel("Radius")
plt.ylabel("Average Neutrons produced")
plt.savefig("plots/sphere-u235-water-cm.png")
plt.show()

# Radius = 6.662200000000004







