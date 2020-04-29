package main

type sphere struct {
	radius          float64
	casingThickness float64

	position Vector

	u235Percentage float64
	u238Percentage float64
}

func newSphere(radius, casingThickness float64, position Vector, u235Percentage, u238Percentage float64) sphere {
	s := sphere{
		radius:          radius,
		casingThickness: casingThickness,
		position:        position,
		u235Percentage:  u235Percentage,
		u238Percentage:  u238Percentage,
	}
	return s
}

func getU235Lambda() float64 {
	return -1
}

func getShieldingLambda() float64 {
	return -1
}

func (s *sphere) isNeutronInU235(n Neutron) bool {
	deltaV := s.position.subVector(n.positionVector)
	magnitude := deltaV.magnitude()

	if magnitude < s.radius {
		return true
	}
	return false
}
