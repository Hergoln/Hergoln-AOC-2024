package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

var INPUT_FILENAME string = "input"

func readInput(filename string) string {
	readFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(readFile)
}

func multiply(mul string) int {
	left := strings.Index(mul, "(")
	right := strings.Index(mul, ")")
	numbers := strings.Split(mul[left+1:right], ",")

	prod := 1
	for i := range(numbers) {
		val, _ := strconv.Atoi(numbers[i])
		prod *= val
	}
	return prod
}

// find mul(x,y)
func day1() {
	input := readInput(INPUT_FILENAME)
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	muls := r.FindAllString(input, -1)	
	sum := 0
	for mul := range(muls) {
		sum += multiply(muls[mul])
	}
	fmt.Println(sum)
}

func day2() {
	input := readInput(INPUT_FILENAME)
	r := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|do\(\)|don't\(\)`)
	instructions := r.FindAllString(input, -1)

	sum := 0
	enable := true
	for idx := range(instructions) {
		if instructions[idx] == "do()" {
			enable = true
		}
		if instructions[idx] == "don't()" {
			enable = false
		}
		if enable {
			sum += multiply(instructions[idx])
		}
	}

	fmt.Println(sum)
}

func main() {
	day2()
}