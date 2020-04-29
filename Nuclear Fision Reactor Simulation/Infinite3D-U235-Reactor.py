import matplotlib as mpl
from mpl_toolkits.mplot3d import Axes3D
import numpy as np
import matplotlib.pyplot as plt
from enum import Enum

print("Infinite 2D Reactor")

#TODO - Test the histog]ram of the theta and phi valueas calculated
class NeutronState(Enum):
    scatter = 0
    absorb = 1
    fission = 2
    nothing = 3


class RandomUnitVector(object):

    def __init__(self):
        phi = self.random_phi()
        z = self.random_z()
        self.x = self.random_x(phi, z)
        self.y = self.random_y(phi, z)
        self.z = z

    def magnitude(self):
        r2 = pow(self.x, 2) + pow(self.y, 2) + pow(self.z, 2)
        return np.sqrt(r2)

    def random_phi(self):
        p = np.random.uniform(0, 1, 1)[0]
        phi = 2 * np.pi * p
        return phi

    def random_z(self):
        return np.random.uniform(-1, 1, 1)[0]

    def random_x(self, phi, z):
        return np.sin(np.arccos(z)) * np.cos(phi)

    def random_y(self, phi, z):
        return np.sin(np.arccos(z)) * np.sin(phi)


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

    fast_scatter_lambda = 5.107798
    fast_absorb_lambda = 227.013246
    fast_fission_lambda = 20.43119214

    scatter_lambda = 204.31192
    absorb_lambda = 20.63756782
    fission_lambda = 3.504492648

    is_free = True

    state = NeutronState.nothing
    generation = 0

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

        self.lam = pow(1/self.scatter_lambda + 1/self.absorb_lambda + 1/self.fission_lambda, -1)

        self.vector = Vector(RandomUnitVector(), 1)

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

    def move(self):
        if not self.is_free:
            return

        self.generation += 1

        self.move_distance()

        magnitude = self.vector.magnitude()

        p_scatter = 1 - pow(np.e, - magnitude / self.scatter_lambda)
        p_absorb = 1 - pow(np.e, - magnitude / self.absorb_lambda)
        p_fission = 1 - pow(np.e, -magnitude / self.fission_lambda)

        p_total = p_scatter + p_absorb + p_fission

        if p_total > 1:
            p_scatter = p_scatter / p_total
            p_absorb = p_absorb / p_total
            p_fission = p_fission / p_total
            p_total = p_scatter + p_absorb + p_fission

        a = np.random.uniform(0, 1, 1)[0]
        if a <= p_scatter:
            print("{} - Scatter".format(self.generation))
            self.state = NeutronState.scatter
            return NeutronState.scatter
        elif p_scatter < a <= (p_scatter + p_absorb):
            print("{} - Absorb".format(self.generation))
            self.is_free = False
            self.state = NeutronState.absorb
            return NeutronState.absorb
        elif (p_scatter + p_absorb) < a <= (p_scatter + p_absorb + p_fission):
            print("{} - Fission".format(self.generation))
            self.is_free = False
            self.state = NeutronState.fission
            return NeutronState.fission
        else:
            print("{} - nothing".format(self.generation))
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


number_neutrons = 1000
neutrons = np.zeros(number_neutrons, dtype=Neutron)

for i in range(number_neutrons):
    neutrons[i] = Neutron(0, 0, 0)

generations = 1000

number_neutrons_history = np.zeros(generations, dtype=int)
current_number_neutrons = 0

for g in range(generations):


    for n in neutrons:
        n.move()
        print(n.state)

        if n.state == NeutronState.fission:
            rand = np.random.uniform(0, 1, 1)[0]
            if rand <= 0.7:
                # 2 neutrons
                current_number_neutrons += 2
            else:
                # 3 neutrons
                current_number_neutrons += 3


print("Number of Neutrons = {}".format(current_number_neutrons))




