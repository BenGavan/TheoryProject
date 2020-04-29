package main

type sphere struct {
	radius float64
	casingThickness float64

	position Vector
}

func newSphere(radius, casingThickness float64, position Vector) sphere {
	s := sphere{
		radius:          radius,
		casingThickness: casingThickness,
		position:        position,
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
	//deltaV := s.position.subVector(n.positionVector)
	//magnitude := deltaV.magnitude()

	return true

	//if magnitude < s.radius {
	//	return true
	//}
	//return false
}
