package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
)


func readInput(filename string) ([]int, []int) {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var result []string
	var l1 []int
	var l2 []int
	for fileScanner.Scan() {
		result = strings.Split(fileScanner.Text(), " ")
		left, _ := strconv.Atoi(result[0])
		l1 = append(l1, left)
		right, _ := strconv.Atoi(result[3])
		l2 = append(l2, right)
	}
	readFile.Close()
	return l1, l2
}

func part1() {
	l1, l2 := readInput("input.txt")
	sort.Ints(l1)
	sort.Ints(l2)
	sum := 0

	for i := range(l1) {
		val := l2[i] - l1[i]
		if val < 0 {
			val = -val
		}
		sum += val
	}
	fmt.Println(sum)
}

func findFactor(val int, l2 []int) int {
	count := 0
	for j := range(l2) {
		if val == l2[j] {
			count++
		}
	}
	return val * count
}

func part2() {
	l1, l2 := readInput("input.txt")
	l1Map := map[int]int{}
	sum := 0
	var exists bool
	var factor int

	for i := range(l1) {
		currentLeft := l1[i]
		factor, exists = l1Map[currentLeft]
		if !exists {
			factor = findFactor(currentLeft, l2)
			l1Map[currentLeft] = factor
		}
		sum += factor
	}
	fmt.Println(sum)
}

func main() {
	part2()
}