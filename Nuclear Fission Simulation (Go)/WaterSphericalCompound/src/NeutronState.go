package main

import "fmt"

type NeutronState int

const (
	scatter NeutronState = iota
	absorb
	fissionU235
	fissionU238
	nothing
)

func (state *NeutronState) toString() string {
	var s string
	switch *state {
	case scatter:
		s = "Scatter"
	case absorb:
		s = "Absorb"
	case fissionU235:
		s = "Fission U-235"
	case fissionU238:
		s = "Fission U-238"
	case nothing:
		s = "Nothing"
	}
	return s
}

func (state *NeutronState) print() {
	s := state.toString()
	fmt.Printf("State = %v\n", s)
}
