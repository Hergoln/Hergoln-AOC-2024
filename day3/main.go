package main

import (
	"fmt"
	"os"
)

var INPUT_FILENAME string = "example"

func readInput(filename string) string {
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return string(readFile)
}

func day1() {
	fmt.Println(readInput(INPUT_FILENAME))
}

func main() {
	day1()
}