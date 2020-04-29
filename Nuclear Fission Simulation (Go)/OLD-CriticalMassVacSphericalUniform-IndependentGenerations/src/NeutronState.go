package main

import "fmt"

type NeutronState int

func (state *NeutronState) toString() string {
	var s string
	switch *state {
	case scatter:
		s = "Scatter"
	case absorb:
		s = "Absorb"
	case fission:
		s = "Fission"
	case nothing:
		s = "Nothing"
	}
	return s
}

func (state *NeutronState) print() {
	s := state.toString()
	fmt.Printf("State = %v\n", s)
}
