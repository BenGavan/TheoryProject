package main

import (
	"fmt"
	"math"
	"math/rand"
)

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
	}
	return n
}

func randomDistanceFor(lambda float64) float64 {
	u := rand.Float64()
	return - lambda * math.Log(u)
}

func (n *Neutron) moveEGS() {
	dF := randomDistanceFor(fissionLambdaU235)
	dA := randomDistanceFor(absorbLambdaU235)
	dS := randomDistanceFor(scatterLambdaU235)

	d1 := math.Min(dF, dA)
	deltaDistance := math.Min(d1, dS)

	deltaVector := n.directionVector.multiplyBy(deltaDistance)
	n.stepVector = n.stepVector.addVector(deltaVector)
	n.positionVector = n.positionVector.addVector(deltaVector)

	switch deltaDistance {
	case dF:
		n.state = fission
	case dA:
		n.state = absorb
	case dS:
		n.state = scatter
	}

	n.printState()
	fmt.Printf("f: %v, a: %v, s: %v\n", dF,  dA, dS)

	if (n.state == absorb) || (n.state == fission) || (n.medium == empty){
		n.isFree = false
	}
}

func (n *Neutron) move() {
	deltaDistance := n.randomDistance()

	//fmt.Printf("Delta Step distance: %v\n", deltaDistance)

	deltaVector := n.directionVector.multiplyBy(deltaDistance)
	n.stepVector = n.stepVector.addVector(deltaVector)
	n.positionVector = n.positionVector.addVector(deltaVector)

	distance := n.stepVector.magnitude()

	n.medium = u235
	//n.printCurrentMedium()

	//fmt.Print("Step Vector")
	//n.stepVector.print()
	//
	//fmt.Printf("Step Magnitude: %v\n", n.stepVector.magnitude())



	pScatter := n.probabilityForDistance(distance, scatterLambdaU235)
	pAbsorb := n.probabilityForDistance(distance, absorbLambdaU235)
	pFission := n.probabilityForDistance(distance, fissionLambdaU235)

	pTotal := pScatter + pAbsorb + pFission

	if pTotal > 1 {
		pScatter = pScatter / pTotal
		pAbsorb = pAbsorb / pTotal
		pFission = pFission / pTotal
	}

	//fmt.Println(pScatter, pAbsorb, pFission)
	//fmt.Println(scatterLambdaU235, absorbLambdaU235, fissionLambdaU235)

	a := rand.Float64()
	if a <= pScatter {
		n.state = scatter
	} else if a <= (pScatter + pAbsorb) {
		n.state = absorb
	} else if a <= (pScatter + pAbsorb + pFission) {
		n.state = fission
	} else {
		n.state = nothing
	}


	if n.state != nothing {
		n.stepVector = Vector{0,0,0}
		n.directionVector = newRandomUnitVector()
	}

	if (n.state == absorb) || (n.state == fission) || (n.medium == empty){
		n.isFree = false
	}

	//n.printState()
	//fmt.Println("Medium: ", n.medium)
}

func (n *Neutron) getLambda() float64 {
	if n.medium == u235 {
		inverseLambda := 1/scatterLambdaU235 + 1/absorbLambdaU235 + 1/fissionLambdaU235
		return math.Pow(inverseLambda, -1)
	} else if n.medium == water {
		inverseLambda := 1/scatterLambdaWater + 1/absorbLambdaWater
		return math.Pow(inverseLambda, -1)
	} else if n.medium == ZR2 {
		inverseLambda := 1/scatterLambdaZR2 + 1/absorbLambdaZR2
		return math.Pow(inverseLambda, -1)
	}
	panic("Neutron is somehow in a material that does not exist/is supported (could be in nothing: " + fmt.Sprint(n.medium == empty) + ")")
}

func (n *Neutron) getCurrentMedium(s *sphere) Medium {

	isInU235 := s.isNeutronInU235(*n)
	if isInU235 {
		return u235
	}
	return empty
}

func (n *Neutron) randomDistance() float64 {
	u := rand.Float64()
	mag := - n.getLambda() * math.Log(u)
	return mag
}

func (n *Neutron) probabilityForDistance(d float64, lambda float64) float64 {
	return 1 - math.Exp(- d/lambda)
}

func (n *Neutron) printState() {
	var s string
	switch n.state {
	case scatter:
		s = "Scatter"
	case absorb:
		s = "Absorb"
	case fission:
		s = "Fission"
	case nothing:
		s = "Nothing"
	}
	fmt.Printf("State = %v\n", s)
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