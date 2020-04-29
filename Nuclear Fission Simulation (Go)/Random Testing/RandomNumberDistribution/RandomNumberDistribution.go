package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
)

func main() {
	filepath := "Random Testing/data/random-numbers.txt"
	fmt.Printf("Opening file %v", filepath)
	f, err := os.OpenFile(filepath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer f.Close()

	var x float64
	for i := 0; i < int(math.Pow(10, 6)); i++ {
		x = rand.Float64()
		_, err = f.WriteString(fmt.Sprint(x) + "\n")
		if err != nil {
			fmt.Println("ERROR: error appending to file: " + filepath)
			panic(err)
		}
	}
}