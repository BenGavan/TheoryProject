package main

import (
	"fmt"
	"math"
)

func main() {
	testSwitch()
}

func testSwitch() {
	switch 2 {
	case 1:
		fmt.Println("first")
	case 2:
		fmt.Println("second")
	case 3:
		fmt.Println("third")
	}
}

func old2() {
	var produced int = 1000
	var need int = 1999
	var r float64 = float64(need) / float64(produced)
	//x := need % produced
	fmt.Printf("Produced = %v, Need = %v, r = %v\n", produced, need, r)

	if produced < need {
		n1 := produced -1
		v1 := int(math.Floor(r))
		one := v1 * n1

		v2 := need - one

		sum := one + v2

		fmt.Printf("%v of %v = %v\n", n1, v1, n1 * v1)
		fmt.Printf("%v of %v = %v\n", 1, v2, v2)
		fmt.Printf("Sum = %v\n", sum)
		fmt.Printf("Used = %v / %v\n", n1 + 1, produced)
	} else if produced == need {
		fmt.Printf("%v of %v = %v\n", produced, 1, produced)
		fmt.Printf("Used = %v / %v spaces\n", produced, produced)
	} else {
		fmt.Printf("%v of %v = %v\n", need, 1, need)
		fmt.Printf("%v of %v = %v\n", produced - need, 0, 0)
		fmt.Printf("Sum = %v\n", need)
	}


	//run()
}

// The one which works
func option() {
	var nPreviousFissionlocations int = 6
	var neutronsRequired int = 12

	previousLocations := make([]int, nPreviousFissionlocations)
	for i := 0; i < neutronsRequired; i++ {
		previousLocations[i % nPreviousFissionlocations] += 1
	}

	nextLocations := make([]int, neutronsRequired)
	for i := 0; i < nPreviousFissionlocations; i++ {
		nextLocations[i % neutronsRequired] += 1
	}


	fmt.Println(previousLocations)
	fmt.Println(nextLocations)
}

//func test() {
//	in := 100
//	possiblePlaces := 10
//
//	places := make([]int, possiblePlaces)
//
//	for i := 0; i < in; i++ {
//
//	}
//}

func run() {
	var produced float64 = 551
	var need float64 = 500
	var r float64 = need / produced

	fmt.Println(r)

	floor := math.Floor(r)
	ceil := math.Ceil(r)

	p := r - floor

	y := p * 2
	z := (1-p) * 1

	fmt.Println(y * need, z * need)

	fmt.Printf("floor = %v\nceil= %v\n", floor, ceil)

	x := (need - produced * ceil) / (floor - ceil)
	fmt.Printf("x = %v of %v\n",  x, floor)
}

func old() {
	var n float64 = 501
	var r float64 = n / 500

	floor := math.Floor(r)
	ceil := math.Ceil(r)

	fmt.Printf("floor = %v\nceil= %v\n", floor, ceil)

	proportionLow := (r + ceil) / (floor + ceil)
	proportionHigh := 1 - proportionLow

	//proportionHigh := r - floor
	fmt.Printf("proportionHigh = %v\n", proportionHigh)

	//proportionLow := 1 - proportionHigh
	fmt.Printf("Proportion low = %v\n", proportionLow)

	numberLow := math.Round(proportionLow * n)
	numberHigh := math.Round(proportionHigh * n)

	fmt.Printf("will need\n%v of %v\t and \t %v of %v\n", numberLow, floor,  numberHigh, ceil)

	//proportionLow := (r + ceil) / (floor + ceil)
	//proportionHigh := 1 - proportionLow
	//fmt.Printf("Proportion low = %v\n", proportionLow)
}