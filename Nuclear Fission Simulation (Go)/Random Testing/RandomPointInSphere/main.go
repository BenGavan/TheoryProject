package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Random Point in sphere")
	filepath := "Random Testing/RandomPointInSphere/random-points-sphere.txt"
	fmt.Printf("Opening file %v", filepath)
	f, err := os.OpenFile(filepath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer f.Close()

	for i := 0; i < 2000; i++ {
		r := randomPointInSphere(1)
		r.print()

		_, err = f.WriteString(fmt.Sprintf("%v\t%v\t%v\n", r.x, r.y, r.z))
		if err != nil {
			fmt.Println("ERROR: error appending to file: " + filepath)
			panic(err)
		}
	}
}

func randomPointInSphere(radius float64) Vector {
	var a float64 = 1
	var b float64 = 2
	x := ((rand.Float64() * b) - a) * radius
	y := ((rand.Float64() * b) - a) * radius
	z := ((rand.Float64() * b) - a) * radius

	r2 := math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2)
	r := math.Sqrt(r2)
	if r <= radius {
		return Vector{x, y, z}
	}
	return randomPointInSphere(radius)
}

type Vector struct {
	x float64
	y float64
	z float64
}

func (v *Vector) magnitude() float64 {
	r2 := math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2)
	return math.Sqrt(r2)
}

func (v *Vector) makeUnit() {
	m := v.magnitude()

	v.x = v.x / m
	v.y = v.y / m
	v.z = v.z / m
}

func (v *Vector) multiplyBy(factor float64) Vector {
	x := v.x * factor
	y := v.y * factor
	z := v.z * factor
	return Vector{x, y, z}
}

func (v *Vector) addVector(delta Vector) Vector {
	x := v.x + delta.x
	y := v.y + delta.y
	z := v.z + delta.z
	return Vector{x, y, z}
}

func (v *Vector) subVector(delta Vector) Vector {
	x := v.x - delta.x
	y := v.y - delta.y
	z := v.z - delta.z
	return Vector{x, y, z}
}

func (v *Vector) print() {
	fmt.Println(v.toString())
}

func (v *Vector) toString() string {
	return fmt.Sprintf("(%v, %v, %v)\n", v.x, v.y, v.z)
}
