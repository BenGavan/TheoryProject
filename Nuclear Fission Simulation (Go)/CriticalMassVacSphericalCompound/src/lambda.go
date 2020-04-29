package main

import (
	"math"
)

const ( // All  in barns (10^-24 cm^-2)
	//u235ThermalScatterSigma float64 = 10
	//u235ThermalCaptureSigma float64 = 99
	//u235ThermalFissionSigma float64 = 583
	//u235ThermalTotalSigma           = u235ThermalScatterSigma + u235ThermalCaptureSigma + u235ThermalFissionSigma
	//
	//u238ThermalScatterSigma float64 = 9
	//u238ThermalCaptureSigma float64 = 2
	//u238ThermalFissionSigma float64 = 0.00002
	//u238ThermalTotalSigma           = u238ThermalScatterSigma + u238ThermalCaptureSigma + u238ThermalFissionSigma

	u235FastScatterSigma float64 = 4
	u235FastCaptureSigma float64 = 0.09
	u235FastFissionSigma float64 = 1
	u235FastTotalSigma           = u235FastScatterSigma + u235FastCaptureSigma + u235FastFissionSigma

	u238FastScatterSigma float64 = 5
	u238FastCaptureSigma float64 = 0.07
	u238FastFissionSigma float64 = 0.3
	u238FastTotalSigma           = u238FastScatterSigma + u238FastCaptureSigma + u238FastFissionSigma
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

	M := 238.05078826 // Atomic Mass

	var p = 19.1 // Density (g/cm^3)

	n := (NA / M) * p
	return n * percentage
}

// MARK: - Utils

func lambda(n, sigma float64) float64 {
	sigma *= math.Pow(10, -24)
	lambda := 1 / (n * sigma)
	return lambda / 100
}
