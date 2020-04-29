package main

import (
	"fmt"
	"math"
)

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
	fmt.Printf("(%v, %v, %v)\n", v.x, v.y, v.z)
}

func newRandomUnitVector() Vector {
	v := randomPointInSphere(1)
	v.makeUnit()
	return v
}
