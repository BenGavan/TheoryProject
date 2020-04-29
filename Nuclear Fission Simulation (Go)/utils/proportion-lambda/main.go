package main

import (
	"fmt"
	"math"
)

func  main() {
	l()
}

func total() {
	nU235 := nU235(1)
	nU238 := nU238(0)

	var sigma235 float64 = (10 + 99 + 583) * math.Pow(10,  -24)
	var sigma238 float64 = (9 + 2 + 0.00002) * math.Pow(10,  -24)

	t := (nU235 * sigma235) + (nU238 * sigma238)

	fmt.Printf("SigmaT = %v\n", t)
	l := 1 / t
	fmt.Println(l / 100)
}

func lambdaU235(percentage, sigma float64) float64 {
	n := nU235(percentage)
	l := lambda(n, sigma) / 100
	fmt.Printf("l = %v\n", l)
	return l
}

func lambdaU238(percentage, sigma float64) float64 {
	n := nU238(percentage)
	l := lambda(n, sigma) / 100
	fmt.Printf("l = %v\n", l)
	return l
}

func nU235(percentage float64) float64 {
	NA := 6.022 * math.Pow(10, 23) // Avogadro's number

	M := 235.0439299 // Atomic Mass

	var p float64 = 19.1 // Density (g/cm^3)

	n := (NA / M) * p
	return n * percentage
}

func nU238(percentage float64) float64 {
	NA := 6.022 * math.Pow(10, 23) // Avogadro's number

	M := 238.05078826 // Atomic Mass

	var p float64 = 19.1 // Density (g/cm^3)

	n := (NA / M) * p
	return n * percentage
}


func lambda(n, sigma float64) float64 {
	lambda := 1 / (n * sigma)
	return lambda
}
