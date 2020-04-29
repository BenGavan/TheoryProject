package main

import "math"

type cylinder struct {
	radius          float64
	length          float64
	casingThickness float64

	position Vector

	u235Percentage float64
	u238Percentage float64
}

func newCylinder(radius, length, casingThickness float64, position Vector, u235Percentage, u238Percentage float64) cylinder {
	c := cylinder{
		radius:          radius,
		length:          length,
		casingThickness: casingThickness,
		position:        position,
		u235Percentage:  u235Percentage,
		u238Percentage:  u238Percentage,
	}
	return c
}

func getU235Lambda() float64 {
	return -1
}

func getShieldingLambda() float64 {
	return -1
}

func (c *cylinder) isNeutronInU235(n Neutron) bool {
	deltaV := c.position.subVector(n.positionVector)

	r2 := math.Pow(deltaV.x, 2) + math.Pow(deltaV.y, 2)
	r := math.Sqrt(r2)

	dz := math.Abs(deltaV.z)

	if (r < c.radius) && (dz < c.length/2) {
		return true
	}

	return false
}