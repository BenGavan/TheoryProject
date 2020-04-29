package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

/*
Changes from Thermal vs Fast Neutrons:
 - cross sections
	- Comment out constants
	- Use the correct constant in distance calculations
 - number of neutrons produced by fission
 - File name being written to
*/

type Medium int

const (
	u235 Medium = iota
	//u238
	water
	ZR2
	empty
)

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
	spheres     []sphere
	waterRadius float64
}

func newReactor(reactorWaterRadius, radius, casingThickness, u235Percentage, u238Percentage float64) *reactor {
	r := &reactor{
		spheres: []sphere{
			newSphere(radius, casingThickness, 0, Vector{0, 0, 0}, u235Percentage, u238Percentage),
			newSphere(radius, casingThickness, 0, Vector{2 * radius, 0, 0}, u235Percentage, u238Percentage),
			newSphere(radius, casingThickness, 0, Vector{-2 * radius, 0, 0}, u235Percentage, u238Percentage),

			newSphere(radius, casingThickness, 0, Vector{2 * radius, 2 * radius, 0}, u235Percentage, u238Percentage),
			newSphere(radius, casingThickness, 0, Vector{0, 2 * radius, 0}, u235Percentage, u238Percentage),
			newSphere(radius, casingThickness, 0, Vector{-2 * radius, 2 * radius, 0}, u235Percentage, u238Percentage),

			newSphere(radius, casingThickness, 0, Vector{2 * radius, -2 * radius, 0}, u235Percentage, u238Percentage),
			newSphere(radius, casingThickness, 0, Vector{0, -2 * radius, 0}, u235Percentage, u238Percentage),
			newSphere(radius, casingThickness, 0, Vector{-2 * radius, -2 * radius, 0}, u235Percentage, u238Percentage),
		},
		waterRadius: reactorWaterRadius,
	}
	return r
}

func main() {
	startTime := time.Now()

	rand.Seed(time.Now().UnixNano())

	//cmPlot()
	kForR()

	endTime := time.Now()
	duration := float64(endTime.Sub(startTime).Nanoseconds()) * math.Pow(10, -9)
	fmt.Println("Duration =", duration, "Seconds")
}

//func debug() {
//	sPositionVector := Vector{0,0,0}
//	s := newSphere(1, 1, sPositionVector)
//
//	r := sPositionVector
//	r = Vector{0,0,0}
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
//		p := n.probabilityForDistance(n.stepVector.magnitude(), lambda(nU235(1), u235ThermalFission))
//		fmt.Printf("Propability of fission: %v\n", p)
//		fmt.Println("----------------")
//
//		lam  := n.getTotalLambda()
//		fmt.Printf("lambda: %v\n", lam)
//
//		n.move()
//	}
//}

func kForR() {

	var u235Percentage = 1.0
	var u238Percentage = 1 - u235Percentage

	radius := 0.1

	r := newReactor(1, radius, 0, u235Percentage, u238Percentage)

	n := 100

	a, aSD := averageNeutronsProduced(r, n*len(r.spheres))
	k := a / float64(n)
	kUncert := k * aSD / a

	fmt.Printf("radius = %v, a = %v ± %v, k = %v ± %v\n", radius, a, aSD, k, kUncert)

}

func howKVariesWithR() {
	u235Percentage := 1.0
	u238Percentage := 0.0

	casingThickness := 0.0
	waterThickness := 0.0

	radius := 0.0

	var nStartNeutrons = 1000

	f := openFile()
	defer f.Close()

	for i := 0; i < 100; i++ {
		radius = 1 / 100 * float64(i)
		r := newReactor(radius, casingThickness, waterThickness, u235Percentage, u238Percentage)

		a, sd := averageNeutronsProduced(r, nStartNeutrons)
		k := a / float64(nStartNeutrons)
		kUncert := k * sd / a

		appendLineToFile(f, radius, a, sd, k, kUncert)
		fmt.Println("----------------")
	}

	fmt.Println("Finished")
}

func findCR(r *reactor) float64 {
	n := 300

	rLower := 0.00000001
	rUpper := 0.5

	for {
		for i := 0; i < len(r.spheres); i++ {
			r.spheres[i].radius = rUpper
		}
		a, _ := averageNeutronsProduced(r, n)
		kNew := a / float64(n)
		if kNew <= 1 {
			rUpper += 1
		} else {
			break
		}
	}

	var radius float64

	for {
		rNew := (rUpper + rLower) / 2

		for i := 0; i < len(r.spheres); i++ {
			r.spheres[i].radius = rNew
		}

		a, sd := averageNeutronsProduced(r, n)
		kNew := a / float64(n)
		kNewUncert := kNew * sd / a

		fmt.Printf("Finished: radius = %v,  k = %v ± %v, produced = %v ± %v\n", rNew, kNew, kNewUncert, a, sd)

		radius = rNew

		delta := math.Abs(kNew - 1)
		if delta < 0.001 {
			break
		}

		if kNew < 1 {
			rLower = rNew
		} else if kNew > 1 {
			rUpper = rNew
		} else {
			break
		}
	}
	return radius
}

func cmPlot() {
	var u235Percentage float64 = 1
	var u238Percentage float64 = 0

	var waterThickness float64 = 0
	var casingThickness float64 = 0

	var radius float64 = 1

	crs := make([]float64, 100)
	casingThicknesses := make([]float64, 100)

	f := openFile()
	defer f.Close()

	for i := 0; i < 40; i++ {
		casingThickness = 1.0 / 40.0 * float64(i)

		fmt.Printf("casing thickness = %v\n", casingThickness)

		r := newReactor(radius, casingThickness, waterThickness, u235Percentage, u238Percentage)

		cr := findCR(r)

		crs[i] = cr
		casingThicknesses[i] = casingThickness

		appendLineToFile(f, casingThickness, cr)

		fmt.Println("----------------")
	}
}

func openFile() *os.File {
	filepath := "data/sphere-critical-mass-plot-fast-pure-u235-varying-zr-casing.txt"
	fmt.Printf("Opening file %v\n", filepath)
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return f
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

// Returns # average produced, uncertainty (one sd)
func averageNeutronsProduced(r *reactor, nStartNeutrons int) (float64, float64) {
	n := 1000 // Number of iterations
	neutronsProduced := make([]float64, n)

	neutronsInEach := nStartNeutrons / len(r.spheres)

	var fissionLocations []FissionLocation
	for x := 0; x < len(r.spheres); x++ {
		for i := 0; i < neutronsInEach; i++ {
			fl := FissionLocation{
				location:               addVectors(randomPointInSphere(r.spheres[x].radius), r.spheres[x].position),
				numberNeutronsProduced: 1,
			}
			fissionLocations = append(fissionLocations, fl)
			fmt.Printf("%v\t%v\t%v\n", fl.location.x, fl.location.y, fl.location.z)
		}
	}



	for i := 0; i < n; i++ {
		previousFissionLocations, nNeutronsProduced := performIteration(r, nStartNeutrons, fissionLocations)
		neutronsProduced[i] = nNeutronsProduced

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
	return total / float64(len(arr))
}

func sdOfArray(arr []float64, mean float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += math.Pow(float64(arr[i])-mean, 2)
	}
	sd := math.Sqrt(sum / float64(len(arr)))
	return sd
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
		locations[j%len(locations)].numberNeutronsProduced += 1
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

// Returns the positions of the fission locations of the previous generation
func performIteration(r *reactor, nStartNeutrons int, previousFissionLocations []FissionLocation) ([]Vector, float64) {
	//fmt.Println(previousFissionLocations)
	ns := make([]Neutron, nStartNeutrons)

	index := 0
	var location Vector
	for i := 0; i < len(previousFissionLocations); i++ {
		for x := 0; x < previousFissionLocations[i].numberNeutronsProduced; x++ {
			location = previousFissionLocations[i].location
			ns[index] = newNeutron(location, r)
			index += 1
		}
	}

	var numberNeutronsFree = 0
	var fissionLocations []Vector
	var numberNeutronsProduced float64 = 0
	for {
		for i := 0; i < len(ns); i++ {
			if ns[i].isFree {
				numberNeutronsFree += 1
				ns[i].move()
				if ns[i].medium == empty {
					continue
				}

				if ns[i].state == fissionU235 {
					numberNeutronsProduced += 2.5                                     // TODO: Change when doing thermal (2.43) / fast (2.50)
					fissionLocations = append(fissionLocations, ns[i].positionVector) // TODO: Change when doing fast
				}
				if ns[i].state == fissionU238 {
					numberNeutronsProduced += 2.46                                    // TODO: Change when doing thermal (0) / fast (2.46)
					fissionLocations = append(fissionLocations, ns[i].positionVector) // TODO: Change when doing thermal
				}
			}
		}

		if numberNeutronsFree == 0 {
			break
		}
		numberNeutronsFree = 0
	}
	return fissionLocations, numberNeutronsProduced
}
