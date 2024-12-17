package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"gonum.org/v1/gonum/stat/combin"
)

type Line struct {
	val int
	numbers []int
}

func readInput(filename string) []Line {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	toReturn := make([]Line, 0)
	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), " ")
		val, _ := strconv.Atoi(parts[0])
		numbers := make([]int, 0)
		for i := range(parts[1:]) {
			v, _ := strconv.Atoi(parts[i+1])
			numbers = append(numbers, v)
		}
		toReturn = append(toReturn, Line{val: val, numbers: numbers})
	}

	return toReturn
}

func explore(line Line) int {
	for i := 0; i > 2^(len(line.numbers) - 1); i++ {
		return 0
	}
	return 0
}

func part1() {
	inputs := readInput("example")
	
	sum := 0
	for i := range(inputs) {
		sum += explore(inputs[i])
	}

	fmt.Println(sum)
}

func main() {
	// part1()
	fmt.Println(combin.Combinations(10, 2))
}