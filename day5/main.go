package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"slices"
)

func readInput(filename string) (map[int][]int,[][]int) {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	rules := make(map[int][]int)
	var pages_sets [][]int
	text := "gibberish"

	// rules
	for fileScanner.Scan() {
		text = fileScanner.Text()
		if text == "" {
			break
		}
		parts := strings.Split(text, "|")
		dependant, _ := strconv.Atoi(parts[0])
		if rules[dependant] == nil {
			rules[dependant] = make([]int, 0)
		}
		rule, _ := strconv.Atoi(parts[1])
		rules[dependant] = append(rules[dependant], rule)
	}

	// pages
	pages_sets = make([][]int, 0)
	for i := 0; fileScanner.Scan(); i++ {
		text = fileScanner.Text()
		pages := strings.Split(text, ",")
		pages_sets = append(pages_sets, make([]int, 0))
		for p := range(pages) {
			val, _ := strconv.Atoi(pages[p])
			pages_sets[i] = append(pages_sets[i], val)
		}
	}

	readFile.Close()
	return rules, pages_sets
}

func validatePage(pageI int, rules map[int][]int, pages []int) bool {
	var rule []int
	if rules[pages[pageI]] == nil {
		rule = make([]int, 0)
	} else {
		rule = rules[pages[pageI]]
	}
	
	for i := pageI; i >= 0; i-- {
		if slices.Contains(rule, pages[i]) {
			return false
		}
	}

	return true
}

func checkPagesPart1(rules map[int][]int, pages []int) int {
	for i := 0; i < len(pages); i++ {
		if !validatePage(i, rules, pages) {
			return 0
		}
	}
	return pages[len(pages)/2]
}

func checkPagesPart2(rules map[int][]int, pages []int) int {
	return 0
}

func puzzle(checker func (map[int][]int, []int)int) {
	rules, pages_sets := readInput("input")
	count := 0
	for p := range(pages_sets) {
		count += checker(rules, pages_sets[p])
	}

	fmt.Println(count)
}

func main() {
	part1(checkPagesPart1)
}
