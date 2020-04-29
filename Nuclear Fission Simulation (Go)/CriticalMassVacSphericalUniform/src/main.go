package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

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
	scatterLambdaU235 float64 = 0.02043501314551705
	absorbLambdaU235  float64 = 0.0020641427419714192
	fissionLambdaU235 float64 = 0.00035051480524042965

	// Meters (m) - Fast
	//scatterLambdaU235 float64 = 0.05108753286379263
	//absorbLambdaU235  float64 = 2.2705570161685613
	//fissionLambdaU235 float64 = 0.2043501314551705
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

//type reactor struct {
//	spheres []sphere
//}
//
//func newReactor() *reactor {
//	r := &reactor{[]sphere{
//		newSphere(2, 0, Vector{0, 0, 0}),
//		newSphere(2, 0, Vector{4, 0, 0}),
//	}}
//	return r
//}

func main() {
	startTime := time.Now()

	rand.Seed(time.Now().UnixNano())

	//kForR()

	run()

	endTime := time.Now()
	duration := float64(endTime.Sub(startTime).Nanoseconds()) * math.Pow(10, -9)
	fmt.Println("Duration =", duration, "Seconds")
}

//func debug() {
//	sPositionVector := Vector{0, 0, 0}
//	s := newSphere(1, 1, sPositionVector)
//
//	r := sPositionVector
//	r = Vector{0, 0, 0}
//	nPositionVector := r.addVector(sPositionVector)
//	n := newNeutron(nPositionVector, &s)
//
//	for i := 0; i < 10; i++ {
//		fmt.Print("Is in sphere: ")
//		fmt.Println(s.isNeutronInU235(n))
//
//		fmt.Print("Step Vector: ")
//		n.stepVector.print()
//
//		fmt.Printf("Step magnitude: %v\n", n.stepVector.magnitude())
//
//		fmt.Print("Position: ")
//		n.positionVector.print()
//
//		p := n.probabilityForDistance(n.stepVector.magnitude(), fissionLambdaU235)
//		fmt.Printf("Propability of fission: %v\n", p)
//		fmt.Println("----------------")
//
//		n.move()
//	}
//}

func kForR() {
	r := 0.0003
	s := newSphere(r, 0, Vector{0, 0, 0})

	//n := 1000
	n := 1000

	a, aSD := averageNeutronsProduced(s, n)

	k := a / float64(n)
	kUncert := k * aSD / a

	fmt.Printf("r = %v, a = %v ± %v, k = %v ± %v\n", r, a, aSD, k, kUncert)

}

func appendLineToFile(f *os.File, xs ...float64) {
	var s string
	for i := 0; i < len(xs); i++ {
		s += fmt.Sprintf("%v\t", xs[i])
	}
	s += "\n"

	_, err := f.WriteString(s)
	if err != nil {
		fmt.Println("ERROR: error appending to file: " + f.Name())
		panic(err)
	}
}

func openFile() *os.File {
	filepath := "data/sphere-critical-mass-plot-thermal-depends-on-previous-300.txt"
	fmt.Printf("Opening file %v\n", filepath)
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return f
}

func run() {
	var radius = 0.000005              // meters
	var shieldingThickness float64 = 0 // meters
	var numberNeutronsStart = 10000     // Number of neutrons starting

	s := newSphere(radius, shieldingThickness, Vector{0, 0, 0})

	var nRadii = 300
	radii := make([]float64, nRadii)
	produced := make([]float64, nRadii)
	producedSD := make([]float64, nRadii)
	kValues := make([]float64, nRadii)
	kValueUncerts := make([]float64, nRadii)

	var r float64
	var k float64
	var averageProduced float64
	var sd float64
	var kUncert float64

	f := openFile()
	defer f.Close()

	for i := 1; i < nRadii; i++ {
		r = 0.000005 * float64(i) + (0.1 / 100)
		s.radius = r
		averageProduced, sd = averageNeutronsProduced(s, numberNeutronsStart)
		k = averageProduced / float64(numberNeutronsStart)
		kUncert = k * sd / averageProduced

		radii[i] = r
		kValues[i] = k
		produced[i] = averageProduced
		producedSD[i] = sd
		kValueUncerts[i] = kUncert

		appendLineToFile(f, r, k, kUncert, averageProduced, sd)

		fmt.Printf("Finished: radius = %v,  k = %v ± %v, produced = %v ± %v\n", r, k, kUncert, averageProduced, sd)
	}

	//fmt.Println("Writing data to file")
	//writeColumnsToFile("../data/sphere-critical-mass-plot-thermal-independent-300.txt", radii, produced, producedSD, kValues, kValueUncerts)

	fmt.Println("Finished")
}

type FissionLocation struct {
	location               Vector
	numberNeutronsProduced int
}

func generateFissionLocations(previousFissionLocations []Vector, nStartNeutrons int) []FissionLocation {
	locations := make([]FissionLocation, len(previousFissionLocations))

	for x := 0; x < len(locations); x++ {
		locations[x].location = previousFissionLocations[x]
		locations[x].numberNeutronsProduced = 0
	}

	for j := 0; j < nStartNeutrons; j++ {
		locations[j % len(locations)].numberNeutronsProduced += 1
	}

	var newLocations []FissionLocation

	if len(previousFissionLocations) > nStartNeutrons {
		for x := 0; x < len(locations); x++ {
			if locations[0].numberNeutronsProduced != 0 {
				newLocations = append(newLocations, locations[0])
			}
		}
	} else {
		newLocations = locations
	}


	return newLocations
}

// Returns # average produced, uncertainty (one sd)
func averageNeutronsProduced(s sphere, nStartNeutrons int) (float64, float64) {
	n := 1000 // Number of iterations
	neutronsProduced := make([]float64, n)

	fissionLocations := make([]FissionLocation, nStartNeutrons)
	for i := 0; i < nStartNeutrons; i++ {
		fissionLocations[i].location = randomPointInSphere(s.radius)
		fissionLocations[i].numberNeutronsProduced = 1
	}

	for i := 0; i < n; i++ {
		previousFissionLocations := performIteration(&s, nStartNeutrons, fissionLocations)
		neutronsProduced[i] = 2.43 * float64(len(previousFissionLocations))

		if neutronsProduced[i] == 0 {
			continue
		}

		fissionLocations = generateFissionLocations(previousFissionLocations, nStartNeutrons)
	}

	averageProduced := averageOfArray(neutronsProduced)
	sd := sdOfArray(neutronsProduced, averageProduced)
	return averageProduced, sd
}

func averageOfArray(arr []float64) float64 {
	var total float64 = 0
	for i := range arr {
		total += arr[i]
	}
	return float64(total) / float64(len(arr))
}

func sdOfArray(arr []float64, mean float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += math.Pow(float64(arr[i])-mean, 2)
	}
	sd := math.Sqrt(sum / float64(len(arr)))
	return sd
}

// Returns the positions of the fission locations of the previous generation
func performIteration(s *sphere, nStartNeutrons int, previousFissionLocations []FissionLocation) []Vector {
	//fmt.Println(previousFissionLocations)
	ns := make([]Neutron, nStartNeutrons)

	index := 0
	var location Vector
	for i := 0; i < len(previousFissionLocations); i++ {
		for x := 0; x < previousFissionLocations[i].numberNeutronsProduced; x++ {
			location = previousFissionLocations[i].location
			ns[index] = newNeutron(location, s)
			index += 1
		}
	}

	//for i := 0; i < nStartNeutrons; i++ {
	//	position := randomPointInSphere(s.radius)
	//	ns[i] = newNeutron(position, s)
	//}

	var numberNeutronsFree = 0
	var fissionLocations []Vector
	for {
		for i := 0; i < len(ns); i++ {
			if ns[i].isFree {
				numberNeutronsFree += 1
				ns[i].move()

				if ns[i].state == fission && ns[i].medium == u235 {
					//numberNeutronsProduced += 2.43 // TODO: Check  average (find teh correct distribution)
					fissionLocations = append(fissionLocations, ns[i].positionVector)
				}
			}
		}

		if numberNeutronsFree == 0 {
			break
		}
		numberNeutronsFree = 0
	}
	return fissionLocations
}
