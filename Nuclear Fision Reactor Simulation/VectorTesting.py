import numpy as np


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
        return np.random.uniform(0, 1, 1)[0]

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


v = Vector(RandomUnitVector(), 1)
v.print()
v.print_magnitude()
v.add(1)
v.print()
v.print_magnitude()
v.add(2)

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

    def print_magnitude(self):
        print("Magnitude = {}".format(self.magnitude()))

    def to_string(self):
        return "({}, {}, {})".format(self.x, self.y, self.z)

    def print(self):
        print(self.to_string())



r = RandomUnitVector()
r.print()
r.print_magnitude()