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
 - number of neutrons produced by fission
*/

type Medium int

const (
	u235 Medium = iota
	u238
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

func randomPointInCylinder(radius, height float64) Vector {
	x := ((rand.Float64() * 2) - 1) * radius
	y := ((rand.Float64() * 2) - 1) * radius
	z := ((rand.Float64() * 2) - 1) * height / 2

	r2 := math.Pow(x, 2) + math.Pow(y, 2)
	r := math.Sqrt(r2)

	dz := math.Abs(z)

	if (r < radius) && (dz < height/2) {
		return Vector{x, y, z}
	}
	return randomPointInCylinder(radius, height)
}

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

	cmPlot()
	//kForR()

	endTime := time.Now()
	duration := float64(endTime.Sub(startTime).Nanoseconds()) * math.Pow(10, -9)
	fmt.Println("Duration =", duration, "Seconds")
}

//func debug() {
//	sPositionVector := Vector{0,0,0}
//	c := newSphere(1, 1, sPositionVector)
//
//	r := sPositionVector
//	r = Vector{0,0,0}
//	nPositionVector := r.addVector(sPositionVector)
//	n := newNeutron(nPositionVector, &c)
//
//	for i := 0; i < 10; i++ {
//		fmt.Print("Is in sphere: ")
//		fmt.Println(c.isNeutronInU235(n))
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

	var u235Percentage = 0.90
	var u238Percentage = 1 - u235Percentage

	r := 0.0001
	height := 2.0
	c := newCylinder(r, height, 0, Vector{0, 0, 0}, u235Percentage, u238Percentage)

	n := 1000

	a, aSD := averageNeutronsProduced(c, n)

	k := a / float64(n)
	kUncert := k * aSD / a

	fmt.Printf("r = %v, a = %v ± %v, k = %v ± %v\n", r, a, aSD, k, kUncert)

}

//func findStartRadius(u235Percentage, u238Percentage float64) float64 {
//	r := 0.0005
//	c := newSphere(r, 0, Vector{0, 0, 0}, u235Percentage, u238Percentage)
//
//	n := 1000
//
//	for i := 0;
//	a, aSD := averageNeutronsProduced(c, n)
//
//	k := a / float64(n)
//	kUncert := k * aSD / a
//}

func findCR(c cylinder) float64 {
	n := 200

	rLower := 0.00000001
	rUpper := 1.0

	for {
		c.radius = rUpper
		a, _ := averageNeutronsProduced(c, n)
		kNew := a / float64(n)
		if kNew <= 1 {
			rUpper += 1
		} else {
			break
		}
	}

	var r float64

	for {
		rNew := (rUpper + rLower) / 2

		fmt.Println(rUpper, rLower)

		c.radius = rNew
		a, sd := averageNeutronsProduced(c, n)
		kNew := a / float64(n)
		kNewUncert := kNew * sd / a

		fmt.Printf("Finished: radius = %v,  k = %v ± %v, produced = %v ± %v\n", rNew, kNew, kNewUncert, a, sd)

		r = rNew

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
	return r
}

func cmPlot() {
	var u235Percentage float64
	var u238Percentage float64

	var s cylinder
	height := 2.0

	crs := make([]float64, 100)
	u235s := make([]float64, 100)

	f := openFile()
	defer f.Close()

	for i := 1; i < 100; i++ {
		u235Percentage = 0.01 * float64(i)
		u238Percentage = 1 - u235Percentage

		s = newCylinder(1, height, 0, Vector{0, 0, 0}, u235Percentage, u238Percentage)

		cr := findCR(s)

		crs[i] = cr
		u235s[i] = u235Percentage

		appendLineToFile(f, u235Percentage, cr)

		n := 100

		s.radius = cr

		a, sd := averageNeutronsProduced(s, n)
		kLower := a / float64(n)
		kLowerUncert := kLower * sd / a

		fmt.Printf("Final: u235 = %v, radius = %v,  k = %v ± %v, produced = %v ± %v\n", u235Percentage, cr, kLower, kLowerUncert, a, sd)
	}
}

func openFile() *os.File {
	filepath := "data/cylinder-critical-mass-plot-thermal-u235-238-compound.txt"
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

//func run() {
//	var u235Percentage = 0.01
//	var u238Percentage = 0.99
//
//	var radius = 0.08                  // meters
//
//	var shieldingThickness float64 = 0 // meters
//	var numberNeutronsStart = 5000     // Number of neutrons starting
//
//	c := newCylinder(radius, shieldingThickness, Vector{0, 0, 0}, u235Percentage, u238Percentage)
//
//	var nRadii = 250
//	radii := make([]float64, nRadii)
//	produced := make([]float64, nRadii)
//	producedSD := make([]float64, nRadii)
//	kValues := make([]float64, nRadii)
//	kValueUncerts := make([]float64, nRadii)
//
//	var r float64
//	var k float64
//	var averageProduced float64
//	var sd float64
//	var kUncert float64
//	for i := 0; i < nRadii; i++ {
//		r = 0.01*float64(i) + 0.00249 + 0.005
//		c.radius = r
//		averageProduced, sd = averageNeutronsProduced(c, numberNeutronsStart)
//		k = averageProduced / float64(numberNeutronsStart)
//		kUncert = k * sd / averageProduced
//
//		radii[i] = r
//		kValues[i] = k
//		produced[i] = averageProduced
//		producedSD[i] = sd
//		kValueUncerts[i] = kUncert
//
//		fmt.Printf("Finished: radius = %v,  k = %v ± %v, produced = %v ± %v\n", r, k, kUncert, averageProduced, sd)
//	}
//
//	fmt.Println("Writing data to file")
//	writeColumnsToFile("../data/sphere-critical-mass-plot-thermal-0.01-intervals.txt", radii, produced, producedSD, kValues, kValueUncerts)
//
//	fmt.Println("Finished")
//
//	//
//	//a := averageNeutronsProduced(c, numberNeutronsStart)
//	//fmt.Printf("Average Neutrons produced: %v\n", a)
//	//fmt.Printf("Ratio: %v\n", a / float64(numberNeutronsStart))
//
//	//var tolerance float64 = 10
//	//var step float64 = 0.01
//	//
//	//radiusHistory := make([]float64, 0)
//	//averageHistory := make([]float64, 0)
//	//
//	//for {
//	//	averageProduced := averageNeutronsProduced(c, numberNeutronsStart)
//	//
//	//	fmt.Printf("Radius = %f, Average Produced = %f\n", radius, averageProduced)
//	//
//	//	radiusHistory = append(radiusHistory, radius)
//	//	averageHistory = append(averageHistory, averageProduced)
//	//
//	//	delta := float64(numberNeutronsStart) - averageProduced
//	//
//	//	if delta < -tolerance {
//	//		radius -= step
//	//	} else if delta > tolerance {
//	//		radius += step
//	//	} else {
//	//		step *= 0.1
//	//		tolerance *= 0.1
//	//		if step < math.Pow(10, -5) {
//	//			break
//	//		}
//	//	}
//	//}
//
//	//a := averageNeutronsProduced(numberNeutronsStart)
//	//fmt.Println("Average Neutrons Produced = ", a)
//}

// Returns # average produced, uncertainty (one sd)
func averageNeutronsProduced(c cylinder, nStartNeutrons int) (float64, float64) {
	n := 1000 // Number of iterations
	neutronsProduced := make([]float64, n)

	fissionLocations := make([]FissionLocation, nStartNeutrons)
	for i := 0; i < nStartNeutrons; i++ {
		fissionLocations[i].location = randomPointInCylinder(c.radius, c.length)
		fissionLocations[i].numberNeutronsProduced = 1
	}

	for i := 0; i < n; i++ {
		previousFissionLocations, nNeutronsProduced := performIteration(&c, nStartNeutrons, fissionLocations)
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
func performIteration(c *cylinder, nStartNeutrons int, previousFissionLocations []FissionLocation) ([]Vector, float64) {
	//fmt.Println(previousFissionLocations)
	ns := make([]Neutron, nStartNeutrons)

	index := 0
	var location Vector
	for i := 0; i < len(previousFissionLocations); i++ {
		for x := 0; x < previousFissionLocations[i].numberNeutronsProduced; x++ {
			location = previousFissionLocations[i].location
			ns[index] = newNeutron(location, c)
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
					numberNeutronsProduced += 2.43                                    // TODO: Change when doing thermal (2.43) / fast (2.50)
					fissionLocations = append(fissionLocations, ns[i].positionVector) // TODO: Change when doing fast
				}
				if ns[i].state == fissionU238 {
					numberNeutronsProduced += 0 // TODO: Change when doing thermal (0) / fast (2.46)
					//fissionLocations = append(fissionLocations, ns[i].positionVector) // TODO: Change when doing thermal
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
