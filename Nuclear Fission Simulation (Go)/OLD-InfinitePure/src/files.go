package main

import (
	"fmt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
	"os"
)

func saveFloatArray(results []float64, path string) {
	var saveString string
	for i := 0; i < len(results); i++ {
		saveString += fmt.Sprintf("%f", results[i]) + "\n"
	}
	bytes := []byte(saveString)

	writeToFile(path, bytes)
}

func writeColumnsToFile(path string, cols ...[]float64) {
	if !isArraysOfEqualLength(cols) {
		fmt.Println("ERROR: Arrays not of equal length")
		return
	}
	var outString string
	for i := 0; i < len(cols[0]); i++ {
		for j := range cols {
			outString += fmt.Sprint(cols[j][i], "\t")
		}
		outString += "\n"
	}
	outBytes := []byte(outString)
	writeToFile(path, outBytes)
}

func writeToFile(path string, bytes []byte) {
	//fmt.Println("path:", path, "bytes:", bytes)
	createFile(path)
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = file.Sync()
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("File successfully uploaded")
}

func createFile(path string) {
	if fileDoesExist(path) {
		fmt.Println("File", path, "already exists")
		return
	}

	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Println("file has not been created")
		return
	}
	fmt.Println("file created")
}

// Check if file exists
func fileDoesExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func isArraysOfEqualLength(arrays [][]float64) bool {
	for i := 0; i < len(arrays)-1; i++ {
		if len(arrays[i]) != len(arrays[i+1]) {
			return false
		}
	}
	return true
}
