package main

import (
	"fmt"
	"os"
	"bufio"
	"reflect"
)

type Coo struct {
	x int
	y int
}
type Dir struct {
	dirId int
	dir Coo
}
var UP Dir = Dir{dirId:0, dir:Coo{x:0,y:-1}}
var RT Dir = Dir{dirId:1, dir:Coo{x:1,y:0}}
var DN Dir = Dir{dirId:2, dir:Coo{x:0,y:1}}
var LT Dir = Dir{dirId:3, dir:Coo{x:-1,y:0}}
var DIRS []Dir = []Dir{UP,RT,DN,LT}

func readInput(filename string) ([][]rune,Coo,Dir) {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var startPos Coo
	dir := UP
	placeSchema := make([][]rune, 0)
	for y := 0; fileScanner.Scan(); y++ {
		nextLine := []rune(fileScanner.Text())
		placeSchema = append(placeSchema, nextLine)
		for x := 0; startPos.x == 0 && startPos.y == 0 && x < len(nextLine); x++ {
			if nextLine[x] == '^' {
				startPos = Coo{x:x,y:y}
			}
		}
	}

	return placeSchema, startPos, dir
}

func guardOutside(schema [][]rune, pos Coo) bool {
	if pos.x < 0 || pos.x >= len(schema[0]) {
		return true
	}
	if pos.y < 0 || pos.y >= len(schema) {
		return true
	}
	return false
}

func rotate(dir Dir) Dir {
	return DIRS[(dir.dirId + 1) % len(DIRS)]
}

func dirControl(schema [][]rune, pos Coo, dir Dir) Dir {
	checkY := pos.y + dir.dir.y
	checkX := pos.x + dir.dir.x
	if checkY >= len(schema) || checkX >= len(schema[0]) {
		return dir
	}

	if checkY < 0 || checkX < 0 {
		return dir
	}

	if schema[checkY][checkX] == '#' || schema[checkY][checkX] == 'O' {
		return rotate(dir)
	}
	return dir
}

func move(pos Coo, dir Dir) Coo {
	pos.x += dir.dir.x
	pos.y += dir.dir.y
	return pos
}

func printSchema(schema [][]rune) {
	for y := range(schema) {
		fmt.Println(string(schema[y]))
	}
}

func part1() {
	placeSchema, pos, dir := readInput("input")

	for i:= 0; !guardOutside(placeSchema, pos); i++ {
		placeSchema[pos.y][pos.x] = 'X'
		dir = dirControl(placeSchema, pos, dir)
		pos = move(pos, dir)
	}

	count := 0
	for y := range(placeSchema) {
		for x := range(placeSchema[y]) {
			if placeSchema[y][x] == 'X' {
				count++
			}
		}
	}

	fmt.Println(count)
}

func canPutObstacle(schema [][]rune, pos Coo, dir Dir) bool {
	checkY := pos.y + dir.dir.y
	checkX := pos.x + dir.dir.x
	if checkY >= len(schema) || checkX >= len(schema[0]) {
		return false
	}

	if checkY < 0 || checkX < 0 {
		return false
	}

	if schema[checkY][checkX] == '#' {
		return false
	}
	return true
}

func matchAtEnd(visitedBefore []Coo, toMatch []Coo) bool {
	if len(visitedBefore)-len(toMatch) < 0 {
		return false
	}
	return reflect.DeepEqual(visitedBefore[len(visitedBefore)-len(toMatch):], toMatch)
}

func isLoopIn(visited []Coo) bool {
	for i := len(visited)-2; i > 0; i-- {
		if matchAtEnd(visited[:i], visited[i:]) {
			return true
		}
	}
	return false
}

func samePosDir(pos Coo, dir Dir, otherPos Coo, otherDir Dir) bool {
	return pos.x == otherPos.x && pos.y == otherPos.y && dir.dirId == otherDir.dirId
}

func createsALoop(baseSchema [][]rune, initialPos Coo, initialDir Dir) bool {
	pos := Coo{x:initialPos.x, y:initialPos.y}
	dir := initialDir
	schema := baseSchema
	if !canPutObstacle(schema, pos, dir) {
		return false
	}
	// put obstacle
	schema[initialPos.y + initialDir.dir.y][initialPos.x + initialDir.dir.x] = 'O'
	visitedPos := make([]Coo, 0)

	for i:= 0; !guardOutside(schema, pos); i++ {
		dir = dirControl(schema, pos, dir)
		pos = move(pos, dir)
		visitedPos = append(visitedPos, pos)
		// fmt.Println(len(visitedPos))
		if samePosDir(pos, dir, initialPos, initialDir) || isLoopIn(visitedPos) {
			schema[initialPos.y + initialDir.dir.y][initialPos.x + initialDir.dir.x] = '.'
			return true
		}
	}
	
	schema[initialPos.y + initialDir.dir.y][initialPos.x + initialDir.dir.x] = '.'
	return false
}

func part2() {
	placeSchema, pos, dir := readInput("input")

	count := 0
	for i := 0; !guardOutside(placeSchema, pos); i++ {
		// fmt.Println("Iteracja #", i)
		if createsALoop(placeSchema, pos, dir) {
			count++
		}
		dir = dirControl(placeSchema, pos, dir)
		pos = move(pos, dir)
	}

	fmt.Println(count)
}

func main() {
	part2()
}