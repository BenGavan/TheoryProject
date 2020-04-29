package main

type sphere struct {
	radius          float64
	casingThickness float64
	waterThickness  float64

	position Vector

	u235Percentage float64
	u238Percentage float64
}

func newSphere(radius, casingThickness, waterThickness float64, position Vector, u235Percentage, u238Percentage float64) sphere {
	s := sphere{
		radius:          radius,
		casingThickness: casingThickness,
		waterThickness:  waterThickness,
		position:        position,
		u235Percentage:  u235Percentage,
		u238Percentage:  u238Percentage,
	}
	return s
}

func (s *sphere) isNeutronInU235(n Neutron) bool {
	deltaV := s.position.subVector(n.positionVector)
	magnitude := deltaV.magnitude()

	if magnitude < s.radius {
		return true
	}
	return false
}

func (s * sphere) isNeutronInWater(n Neutron) bool {
	deltaV := s.position.subVector(n.positionVector)
	magnitude := deltaV.magnitude()

	if (magnitude > (s.radius + s.casingThickness)) && (magnitude < (s.radius + s.casingThickness + s.waterThickness)) {
		return true
	}
	return false
}

func (s *sphere) isNeutronInZr(n Neutron) bool {
	deltaV := s.position.subVector(n.positionVector)
	magnitude := deltaV.magnitude()

	if (magnitude > s.radius) && (magnitude < (s.radius + s.casingThickness)) {
		return true
	}
	return false
}
