package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type NeutronState int
type Medium int

const (
	scatter NeutronState = iota
	absorb
	fission
	nothing
)

const (
	u235 Medium = iota
	water
	ZR2
	empty
)

const (
	fastScatterLambdaU235 = 5.107798
	fastAbsorbLambdaU235  = 227.013246
	fastFissionLambdaU235 = 20.43119214

	// Meters (m) - thermal
	//scatterLambdaU235 float64 = 0.02043501314551705
	//absorbLambdaU235  float64 = 0.0020641427419714192
	//fissionLambdaU235 float64 = 0.00035051480524042965

	// Meters (m) - Fast
	scatterLambdaU235 float64 = 0.05108753286379263
	absorbLambdaU235  float64 = 2.2705570161685613
	fissionLambdaU235 float64 = 0.2043501314551705

	scatterLambdaWater float64 = 1
	absorbLambdaWater  float64 = 10

	scatterLambdaZR2 float64 = 1
	absorbLambdaZR2  float64 = 1
)

//var fissionLambdaU235 float64 = 5.8220509150435375e-56  // meters
//var absorbLambdaU235 float64 = 3.428541094414528e-55  // m
//var scatterLambdaU235 float64 = 3.3942556834703824e-54 // m

func randomPointInSphere(radius float64) Vector {
	x := ((rand.Float64() * 2) - 1) * radius
	y := ((rand.Float64() * 2) - 1) * radius
	z := ((rand.Float64() * 2) - 1) * radius

	r2 := math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2)
	r := math.Sqrt(r2)
	if r <= radius {
		return Vector{x, y, z}
	}
	return randomPointInSphere(radius)
}

type reactor struct {
	spheres []sphere
}

func newReactor() *reactor {
	r := &reactor{[]sphere{
		newSphere(2, 0, Vector{0, 0, 0}),
		newSphere(2, 0, Vector{4, 0, 0}),
	}}
	return r
}

func main() {
	startTime := time.Now()

	rand.Seed(time.Now().UnixNano())

	//kForR()

	run()

	endTime := time.Now()
	duration := float64(endTime.Sub(startTime).Nanoseconds()) * math.Pow(10, -9)
	fmt.Println("Duration =", duration, "Seconds")
}

func debug() {
	sPositionVector := Vector{0,0,0}
	s := newSphere(1, 1, sPositionVector)

	r := sPositionVector
	r = Vector{0,0,0}
	nPositionVector := r.addVector(sPositionVector)
	n := newNeutron(nPositionVector, &s)

	for i := 0; i < 10; i++ {
		fmt.Print("Is in sphere: ")
		fmt.Println(s.isNeutronInU235(n))

		fmt.Print("Step Vector: ")
		n.stepVector.print()

		fmt.Printf("Step magnitude: %v\n", n.stepVector.magnitude())

		fmt.Print("Position: ")
		n.positionVector.print()

		p := n.probabilityForDistance(n.stepVector.magnitude(), fissionLambdaU235)
		fmt.Printf("Propability of fission: %v\n", p)
		fmt.Println("----------------")

		lam  := n.getLambda()
		fmt.Printf("lambda: %v\n", lam)

		n.move()
	}
}

func kForR() {
	r := 0.7
	s := newSphere(r, 0, Vector{0, 0, 0})

	n := 1000

	a, aSD := averageNeutronsProduced(s, n)

	k := a / float64(n)
	kUncert :=  k * aSD / a

	fmt.Printf("r = %v, a = %v ± %v, k = %v ± %v\n", r, a, aSD, k, kUncert)

}

func run() {
	var radius float64 = 0.08       // meters
	var shieldingThickness float64 = 0 // meters
	var numberNeutronsStart = 10000	// Number of neutrons starting

	s := newSphere(radius, shieldingThickness, Vector{0, 0, 0})

	var k float64
	var averageProduced float64
	var sd float64
	var kUncert float64


	averageProduced, sd = averageNeutronsProduced(s, numberNeutronsStart)
	k = averageProduced / float64(numberNeutronsStart)
	kUncert = k * sd / averageProduced

	fmt.Printf("Finished:k = %v ± %v, produced = %v ± %v\n", k, kUncert, averageProduced, sd)


	fmt.Println("Writing data to file")
	writeColumnsToFile("data/infinite-fast.txt", []float64{averageProduced}, []float64{sd}, []float64{k}, []float64{kUncert})


	fmt.Println("Finished")

	//
	//a := averageNeutronsProduced(s, numberNeutronsStart)
	//fmt.Printf("Average Neutrons produced: %v\n", a)
	//fmt.Printf("Ratio: %v\n", a / float64(numberNeutronsStart))

	//var tolerance float64 = 10
	//var step float64 = 0.01
	//
	//radiusHistory := make([]float64, 0)
	//averageHistory := make([]float64, 0)
	//
	//for {
	//	averageProduced := averageNeutronsProduced(s, numberNeutronsStart)
	//
	//	fmt.Printf("Radius = %f, Average Produced = %f\n", radius, averageProduced)
	//
	//	radiusHistory = append(radiusHistory, radius)
	//	averageHistory = append(averageHistory, averageProduced)
	//
	//	delta := float64(numberNeutronsStart) - averageProduced
	//
	//	if delta < -tolerance {
	//		radius -= step
	//	} else if delta > tolerance {
	//		radius += step
	//	} else {
	//		step *= 0.1
	//		tolerance *= 0.1
	//		if step < math.Pow(10, -5) {
	//			break
	//		}
	//	}
	//}

	//a := averageNeutronsProduced(numberNeutronsStart)
	//fmt.Println("Average Neutrons Produced = ", a)
}

// Returns # average produced, uncertainty (one sd)
func averageNeutronsProduced(s sphere, nStartNeutrons int) (float64, float64) {
	n := 1000 // Number of iterations
	neutronsProduced := make([]int, n)
	for i := 0; i < n; i++ {
		neutronsProduced[i] = performIteration(&s, nStartNeutrons)
	}
	averageProduced := averageOfArray(neutronsProduced)
	sd := sdOfArray(neutronsProduced, averageProduced)
	return averageProduced, sd
}

func averageOfArray(arr []int) float64 {
	var total = 0
	for i := range arr {
		total += arr[i]
	}
	return float64(total) / float64(len(arr))
}

func sdOfArray(arr []int, mean float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += math.Pow(float64(arr[i]) - mean, 2)
	}
	sd := math.Sqrt(sum / float64(len(arr)))
	return sd
}

func performIteration(s *sphere, nStartNeutrons int) int {
	ns := make([]Neutron, nStartNeutrons)
	for i := 0; i < nStartNeutrons; i++ {
		position := randomPointInSphere(s.radius)
		ns[i] = newNeutron(position, s)
	}

	var numberNeutronsProduced = 0
	var numberNeutronsFree int = 0
	for {
		for i := 0; i < len(ns); i++ {
			if ns[i].isFree {
				numberNeutronsFree += 1
				ns[i].move()

				if ns[i].state == fission {
					x := rand.Float64()
					if x <= 0.7 {
						numberNeutronsProduced += 2
					} else {
						numberNeutronsProduced += 3
					}
				}
			}
		}

		if numberNeutronsFree == 0 {
			break
		}
		numberNeutronsFree = 0
	}

	return numberNeutronsProduced
}
