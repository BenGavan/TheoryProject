package main

import (
	"fmt"
	"math"
	"math/rand"
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
			currentMin  = options[i]
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
	c                     *cylinder
}

func newNeutron(position Vector, s *cylinder) Neutron {
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
		c:                     s,
	}
	return n
}

func randomDistanceFor(lambda float64) float64 {
	u := rand.Float64()
	return - lambda * math.Log(u)
}

func (n *Neutron) move() {
	options := []StepOption{
		{
			lambda:   fissionLambdaU235,
			distance: randomDistanceFor(fissionLambdaU235),
			state:    fission,
		},
		{
			lambda:   absorbLambdaU235,
			distance: randomDistanceFor(absorbLambdaU235),
			state:    absorb,
		},
		{
			lambda:   scatterLambdaU235,
			distance: randomDistanceFor(scatterLambdaU235),
			state:    scatter,
		},
	}

	minOption := minOptionOf(options)

	deltaDistance := minOption.distance
	n.state = minOption.state

	deltaVector := n.directionVector.multiplyBy(deltaDistance)
	n.stepVector = n.stepVector.addVector(deltaVector)
	n.positionVector = n.positionVector.addVector(deltaVector)

	//m := n.positionVector.magnitude()
	//fmt.Printf("m = %v\n", m)

	if n.c.isNeutronInU235(*n) == false {
		n.medium = empty
		n.isFree = false
		return
	}

	n.medium = u235

	//n.printState()

	if (n.state == absorb) || (n.state == fission) || (n.medium == empty) {
		n.isFree = false
	}
}

func (n *Neutron) probabilityForDistanceNew(d, lambda, u float64) float64 {
	return (math.Exp(d/lambda) - 1) * u
}

//
//func (n *Neutron) getTotalLambda() float64 {
//	if n.medium == u235 {
//		inverseLambda := 1/scatterLambdaU235 + 1/absorbLambdaU235 + 1/fissionLambdaU235
//		return math.Pow(inverseLambda, -1)
//	}
//	//} else if n.medium == water {
//	//	inverseLambda := 1/scatterLambdaWater + 1/absorbLambdaWater
//	//	return math.Pow(inverseLambda, -1)
//	//} else if n.medium == ZR2 {
//	//	inverseLambda := 1/scatterLambdaZR2 + 1/absorbLambdaZR2
//	//	return math.Pow(inverseLambda, -1)
//	//}
//	panic("Neutron is somehow in a material that does not exist/is supported (could be in nothing: " + fmt.Sprint(n.medium == empty) + ")")
//}

func (n *Neutron) getCurrentMedium(c *cylinder) Medium {

	isInU235 := c.isNeutronInU235(*n)
	if isInU235 {
		return u235
	}
	return empty
}

//func (n *Neutron) randomDistance() float64 {
//	u := rand.Float64()
//	mag := - n.getTotalLambda() * math.Log(u)
//	//fmt.Printf("Total Labda: %v\n", n.getTotalLambda()) //0.0002953036581722117
//	return mag
//}

////
//func (n *Neutron) probabilityForDistance(d float64, lambda float64) float64 {
//	return 1 - math.Exp(- d/lambda)
//}

////

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
