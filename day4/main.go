package main

import (
	"fmt"
	"os"
	"bufio"
)

type Pred func(int,int,[][]rune)bool
type Trans func(int,int,int,[][]rune)rune

var INPUT_FILENAME string = "input"
var SEARCHED []rune = []rune("XMAS")
var INPUT_LENGTH int = 140 // 10

func readInput(filename string) [][]rune {
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lines := make([][]rune, INPUT_LENGTH)
	for i := 0; fileScanner.Scan(); i++ {
		line := []rune(fileScanner.Text())
		lines[i] = line
	}
	return lines
}

func genericCheck(input [][]rune, h, v int, predicate Pred, transform Trans ) int {
	if predicate(h,v,input) {
		return 0
	}
	for s := range(SEARCHED) {
		if transform(h, v, s, input) != SEARCHED[s] {
			return 0
		}
	}
	return 1
}

func FVP(h, v int, input [][]rune) bool{
	return v + (len(SEARCHED) -1) >= len(input[0])
}
func FVT(h, v, s int, input[][]rune) rune {
	return input[h][v+s]
}

func BVP(h, v int, input [][]rune) bool{
	return v - (len(SEARCHED) -1) < 0
}
func BVT(h, v, s int, input [][]rune) rune {
	return input[h][v-s]
}


func FHP(h, v int, input [][]rune) bool{
	return h + (len(SEARCHED) -1) >= len(input)
}
func FHT(h, v, s int, input [][]rune) rune {
	return input[h+s][v]
}

func BHP(h, v int, input [][]rune) bool{
	return h - (len(SEARCHED) -1) < 0
}
func BHT(h, v, s int, input [][]rune) rune {
	return input[h-s][v]
}

func FFDP(h, v int, input [][]rune) bool{
	return v + (len(SEARCHED) -1) >= len(input[0]) || h + (len(SEARCHED) -1) >= len(input)
}
func FFDT(h, v, s int, input [][]rune) rune {
	return input[h+s][v+s]
}

func BFDP(h, v int, input [][]rune) bool{
	return v - (len(SEARCHED) -1) < 0 || h - (len(SEARCHED) -1) < 0
}
func BFDT(h, v, s int, input [][]rune) rune {
	return input[h-s][v-s]
}

func FBDP(h, v int, input [][]rune) bool{
	return v + (len(SEARCHED) -1) >= len(input[0]) || h - (len(SEARCHED) -1) < 0
}
func FBDT(h, v, s int, input [][]rune) rune {
	return input[h-s][v+s]
}

func BBDP(h, v int, input [][]rune) bool{
	return v - (len(SEARCHED) -1) < 0 || h + (len(SEARCHED) -1) >= len(input)
}
func BBDT(h, v, s int, input [][]rune) rune {
	return input[h+s][v-s]
}

var PREDICATES []Pred  =  []Pred{FVP, BVP, FHP, BHP, FFDP, BFDP, FBDP, BBDP}
var TRANSFORMS []Trans = []Trans{FVT, BVT, FHT, BHT, FFDT, BFDT, FBDT, BBDT}

func allRoundCheck(input [][]rune, h, v int) int {
	count := 0
	for i := range(PREDICATES) {
		count += genericCheck(input, h, v, PREDICATES[i], TRANSFORMS[i])
	}
	return count
}

func part1() {
	input := readInput(INPUT_FILENAME)

	count := 0
	for horI := range(input) {
		for verI := range(input[0]) {
			count += allRoundCheck(input, horI, verI)
		}
	}

	fmt.Println(count)
}

// Left Upper corner going Downwards
func LUDCheck(input [][]rune, h, v int) bool {
	return input[h-1][v-1] == 'M' && input[h+1][v+1] == 'S'
}

// Left Down corner going Uppwards
func LDUCheck(input [][]rune, h, v int) bool {
	return input[h+1][v-1] == 'M' && input[h-1][v+1] == 'S'
}

// Right Down corner going Uppwords
func RDUCheck(input [][]rune, h, v int) bool {
	return input[h+1][v+1] == 'M' && input[h-1][v-1] == 'S'
}

// Right Upper corner going Downwards
func RUDCheck(input [][]rune, h, v int) bool {
	return input[h-1][v+1] == 'M' && input[h+1][v-1] == 'S'
}

func xCheck(input [][]rune, h, v int) int {
	if input[h][v] != 'A' {
		return 0
	}

	LUD := LUDCheck(input, h, v)
	LDU := LDUCheck(input, h, v)
	RDU := RDUCheck(input, h, v)
	RUD := RUDCheck(input, h, v)

	if LUD && LDU || LUD && RUD || RDU && LDU || RDU && RUD {
		return 1
	}

	return 0
}

func part2() {
	input := readInput(INPUT_FILENAME)

	count := 0
	for horI := 1; horI < len(input) -1; horI++ {
		for verI := 1; verI < len(input) -1; verI++ {
			count += xCheck(input, horI, verI)
		}
	}

	fmt.Println(count)
}

func main() {
	part2()
}