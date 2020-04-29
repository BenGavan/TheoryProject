package main

import (
	"fmt"
	"math"
)

func main() {
	n := nU238()
	fmt.Println(n)

	var sigma float64 = e-24

	fmt.Println("Sigma = ", n * sigma)

	lambda := lambda(n, sigma)
	fmt.Print("U235 lambda: ")
	fmt.Println(lambda / 100)
}

func nU235() float64 {
	NA := 6.022 * math.Pow(10, 23)

	M := 235.0439299

	var p float64 = 19.1

	n := (NA / M) * p
	return n
}

func nU238() float64 {
	NA := 6.022 * math.Pow(10, 23)

	M := 235.0439299

	var p float64 = 19.1

	n := (NA / M) * p
	return n
}

func lambda(n, sigma float64) float64 {
	lambda := 1 / (n * sigma)
	return lambda
}


//func crossSection(u235, u238 float64) float64 {
//	var sigma235 float64 =  10 + 99 + 583
//	var sigma238 float64 = 9 + 2 + 0.00002
//	c :=  (u235 * sigma235) + (u238 * sigma238)
//
//
//}

//func test() {
//	x := math.Pow(10, -100)
//	fmt.Println(x)
//}

//func random() {
//	r := rand.Float64()
//}