package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
)

var INPUT_LENGTH int = 1
var INPUT_FILENAME string = "test_input.txt"

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func readInput(filename string) [][]int {
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	reports := make([][]int, INPUT_LENGTH)
	for i := 0; fileScanner.Scan() && i < INPUT_LENGTH; i++ {
		result := strings.Split(fileScanner.Text(), " ")
		report := make([]int, len(result))
		for j := range(result) {
			level, _ := strconv.Atoi(result[j])
			report[j] = level
		}
		reports[i] = report
	}

	return reports
}

func checkReport_pt1(report []int) bool {
	isRising := report[0] - report[1] > 0 
	for i := 0; i < len(report) -1; i++  {
		diff := report[i] - report[i+1]
		if (diff > 0) != isRising {
			return false
		}
		diff = abs(diff)
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func checkPair(left, right int, isRising bool) bool{
	diff := left - right
	if (diff > 0) != isRising {
		return false
	}
	diff = abs(diff)
	if diff < 1 || diff > 3 {
		return false
	}
	return true
}

func checkReport_pt2(report []int) bool {
	firstUnsafe := 0
	isRising := report[0] - report[1] > 0 
	for i := 0; i < len(report) -1; i++ {
		if !checkPair(report[i], report[i+1], isRising) && firstUnsafe < 1 {
			firstUnsafe++
			fmt.Println("I'm here")
			if (i != len(report) -2) && !checkPair(report[i], report[i+2], isRising) {
				return false
			}
			if i != 0 && !checkPair(report[i-1], report[i+1], isRising) {
				return false
			}
		}
	}
	fmt.Println(report)
	return true
}

func puzzle(predicate func([]int)bool) {
	reports := readInput(INPUT_FILENAME)
	sum := 0
	for idx := range(reports) {
		report := reports[idx]
		if predicate(report) {
			sum++
		}
	}

	fmt.Println(sum)
}

func main() {
	puzzle(checkReport_pt2)
}