package main

import (
	"fmt"
	"math"
	"math/rand"
)

const ( // All  in barns (10^-24 cm^-2)
	// *** Thermal *** ///
	//u235ThermalScatterSigma float64 = 10
	//u235ThermalCaptureSigma float64 = 99
	//u235ThermalFissionSigma float64 = 583
	//u235ThermalTotalSigma           = u235ThermalScatterSigma + u235ThermalCaptureSigma + u235ThermalFissionSigma
	//
	//u238ThermalScatterSigma float64 = 9
	//u238ThermalCaptureSigma float64 = 2
	//u238ThermalFissionSigma float64 = 0.00002
	//u238ThermalTotalSigma           = u238ThermalScatterSigma + u238ThermalCaptureSigma + u238ThermalFissionSigma

	//waterThermalScatterSigma float64 = 44
	//waterThermalCaptureSigma float64 = 0.4001

	//heavyWaterThermalScatterSigma float64 = 12
	//heavyWaterThermalCaptureSigma float64 = 0.0007

	/// *** FAST *** ///

	u235FastScatterSigma float64 = 4
	u235FastCaptureSigma float64 = 0.09
	u235FastFissionSigma float64 = 1
	u235FastTotalSigma           = u235FastScatterSigma + u235FastCaptureSigma + u235FastFissionSigma

	u238FastScatterSigma float64 = 5
	u238FastCaptureSigma float64 = 0.07
	u238FastFissionSigma float64 = 0.3
	u238FastTotalSigma           = u238FastScatterSigma + u238FastCaptureSigma + u238FastFissionSigma

	//waterFastScatterSigma float64 = 11
	//waterFastCaptureSigma float64 = 8.003e-5

	heavyWaterFastScatterSigma float64 = 9
	heavyWaterFastCaptureSigma float64 = 1.403e-5
)

type StepOption struct {
	lambda   float64
	distance float64
	state    NeutronState
}

func (o *StepOption) print() {
	fmt.Printf("lambda = %v, distance = %v, state = %v\n", o.lambda, o.distance, o.state.toString())
}

func minOptionOf(options []StepOption) StepOption {
	var currentMin StepOption
	for i := 0; i < len(options); i++ {
		if i == 0 {
			currentMin = options[i]
			continue
		}
		if options[i].distance < currentMin.distance {
			currentMin = options[i]
		}
	}
	return currentMin
}

type Neutron struct {
	isFree                bool
	state                 NeutronState
	generation            int
	scattersInWater       int
	directionVector       Vector
	positionVector        Vector
	positionVectorHistory []Vector
	stepVector            Vector
	medium                Medium
	energy                float64 // MeV
	positionHistory       []Vector
	s                     *sphere
	fissionedWith         Medium
}

func generateNeutron(s *sphere) Neutron {
	startPosition := randomPointInSphere(1)
	n := newNeutron(startPosition, s)
	return n
}

func newNeutron(position Vector, s *sphere) Neutron {
	startPosition := position
	startDirection := newRandomUnitVector()
	positionVectorHistory := make([]Vector, 1)
	positionVectorHistory[0] = startPosition
	stepVector := Vector{0, 0, 0}
	n := Neutron{
		isFree:                true,
		state:                 nothing,
		generation:            0,
		scattersInWater:       0,
		directionVector:       startDirection,
		positionVector:        startPosition,
		positionVectorHistory: positionVectorHistory,
		stepVector:            stepVector,
		medium:                u235,
		s:                     s,
		//fissionedWith:         nil,
	}
	return n
}

func (n *Neutron) randomDistanceFor(lambda float64) float64 {
	u := rand.Float64()
	return - lambda * math.Log(u)
}

func (n *Neutron) stepInFissionable() StepOption {
	u235ThermalFissionLambda := lambda(nU235(n.s.u235Percentage), u235FastFissionSigma)
	u235ThermalCaptureLambda := lambda(nU235(n.s.u235Percentage), u235FastCaptureSigma)
	u235ThermalScatterLambda := lambda(nU235(n.s.u235Percentage), u235FastScatterSigma)

	u238ThermalFissionLambda := lambda(nU238(n.s.u238Percentage), u238FastFissionSigma)
	u238ThermalCaptureLambda := lambda(nU238(n.s.u238Percentage), u238FastCaptureSigma)
	u238ThermalScatterLambda := lambda(nU238(n.s.u238Percentage), u238FastScatterSigma)

	options := []StepOption{
		{
			lambda:   u235ThermalFissionLambda,
			distance: n.randomDistanceFor(u235ThermalFissionLambda),
			state:    fissionU235,
		},
		{
			lambda:   u235ThermalCaptureLambda,
			distance: n.randomDistanceFor(u235ThermalCaptureLambda),
			state:    absorb,
		},
		{
			lambda:   u235ThermalScatterLambda,
			distance: n.randomDistanceFor(u235ThermalScatterLambda),
			state:    scatter,
		},
		{
			lambda:   u238ThermalFissionLambda,
			distance: n.randomDistanceFor(u238ThermalFissionLambda),
			state:    fissionU238,
		},
		{
			lambda:   u238ThermalCaptureLambda,
			distance: n.randomDistanceFor(u238ThermalCaptureLambda),
			state:    absorb,
		},
		{
			lambda:   u238ThermalScatterLambda,
			distance: n.randomDistanceFor(u238ThermalScatterLambda),
			state:    scatter,
		},
	}

	return minOptionOf(options)
}

func (n *Neutron) stepInWater() StepOption {
	waterScatterLambda := lambda(nWater(1), heavyWaterFastScatterSigma)
	waterCaptureLambda := lambda(nWater(1), heavyWaterFastCaptureSigma)

	options := []StepOption{
		{
			lambda:   waterScatterLambda,
			distance: n.randomDistanceFor(waterScatterLambda),
			state:    scatter,
		},
		{
			lambda:   waterCaptureLambda,
			distance: n.randomDistanceFor(waterCaptureLambda),
			state:    absorb,
		},
	}

	return minOptionOf(options)
}

func (n *Neutron) move() {
	n.directionVector = newRandomUnitVector()

	var minOption StepOption

	switch n.medium {
	case u235:
		minOption = n.stepInFissionable()
	case water:
		minOption = n.stepInWater()
	}

	deltaDistance := minOption.distance
	n.state = minOption.state


	deltaVector := n.directionVector.multiplyBy(deltaDistance)
	n.positionVector = n.positionVector.addVector(deltaVector)

	n.medium = n.getCurrentMedium(n.s)

	switch n.medium {
	case u235:
		if (n.state == absorb) || (n.state == fissionU235) || (n.state == fissionU238) {
			n.isFree = false
		}
	case water:
		if n.state == absorb {
			n.isFree = false
		}
		if (n.state == fissionU235) || (n.state == fissionU238) {
			n.state = scatter
		}
	case empty:
		n.isFree = false
	}


	//fmt.Println(n.s.isNeutronInU235(*n))

	if (n.state == absorb) || (n.state == fissionU235) || (n.state == fissionU238) || (n.medium == empty) {
		n.isFree = false
	}
}

//func (n *Neutron) getTotalLambda() float64 {
//	if n.medium == u235 {
//		inverseLambda := (nU235(n.s.u235Percentage) * u235ThermalTotalSigma) + (nU238(n.s.u238Percentage) * u238ThermalTotalSigma)
//		inverseLambda *= math.Pow(10,  -24)
//		lambda := math.Pow(inverseLambda, -1) / 100
//		return lambda
//	}
//	//else if n.medium == water {
//	//	inverseLambda := 1/scatterLambdaWater + 1/absorbLambdaWater
//	//	return math.Pow(inverseLambda, -1)
//	//} else if n.medium == ZR2 {
//	//	inverseLambda := 1/scatterLambdaZR2 + 1/absorbLambdaZR2
//	//	return math.Pow(inverseLambda, -1)
//	//}
//	panic("Neutron is somehow in a material that does not exist/is supported (could be in nothing: " + fmt.Sprint(n.medium == empty) + ")")
//}

func (n *Neutron) getCurrentMedium(s *sphere) Medium {
	isInU235 := s.isNeutronInU235(*n)
	if isInU235 {
		return u235
	}
	isInWater := s.isNeutronInWater(*n)
	if isInWater {
		return water
	}
	return empty
}

//func (n *Neutron) randomDistance() float64 {
//	u := rand.Float64()
//	mag := - n.getTotalLambda() * math.Log(u)
//	//fmt.Printf("Total Labda: %v\n", n.getTotalLambda())
//	return mag
//}

func (n *Neutron) probabilityForDistance(d float64, lambda float64) float64 {
	return 1 - math.Exp(- d/lambda)
}

func (n *Neutron) probabilityNotOccurringForDistance(d, lambda float64) float64 {
	return math.Exp(- d / lambda)
}

func (n *Neutron) printState() {
	n.state.print()
}

func (n *Neutron) printCurrentMedium() {
	s := n.getCurrentMediumString()
	fmt.Printf("Current Medium: %v\n", s)
}

func (n *Neutron) getCurrentMediumString() string {
	var s string
	switch n.medium {
	case u235:
		s = "U235"
	case water:
		s = "Water"
	case ZR2:
		s = "Zr2"
	case empty:
		s = "Empty"
	}
	return s
}
