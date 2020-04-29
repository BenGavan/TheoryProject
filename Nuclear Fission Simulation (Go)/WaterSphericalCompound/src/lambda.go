package main

import (
	"math"
)

//func lambdaU235(percentage, sigma float64) float64 {
//	n := nU235(percentage)
//	sigma *= math.Pow(10, -24)
//	l := lambda(n, sigma)
//	//fmt.Printf("l = %v\n", l)
//	l /= 100
//	return l
//}
//
//func lambdaU238(percentage, sigma float64) float64 {
//	n := nU238(percentage)
//	sigma *= math.Pow(10, -24)
//	l := lambda(n, sigma) / 100
//	//fmt.Printf("l = %v\n", l)
//	return l
//}

// MARK: - Number Densities //

func nU235(percentage float64) float64 {
	NA := 6.022 * math.Pow(10, 23) // Avogadro's number

	M := 235.0439299 // Atomic Mass

	var p = 19.1 // Density (g/cm^3)

	n := (NA / M) * p
	return n * percentage
}

func nU238(percentage float64) float64 {
	NA := 6.022 * math.Pow(10, 23) // Avogadro's number

	M := 238.05078826 // Atomic Mass (g/mol)

	var p = 19.1 // Density (g/cm^3)

	n := (NA / M) * p
	return n * percentage
}

func nWater(percentage float64) float64 {
	NA := 6.022 * math.Pow(10, 23) // Avogadro's number

	M := 18.015 // Atomic Mass (g/mol)

	var p = 0.997 // Density (g/cm^3)

	n := (NA / M) * p
	return n * percentage
}

// MARK: - Utils

func lambda(n, sigma float64) float64 {
	sigma *= math.Pow(10, -24)
	lambda := 1 / (n * sigma)
	return lambda / 100
}
